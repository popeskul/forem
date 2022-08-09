package forem

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"runtime"
	"testing"
	"time"
)

var c *Client

// import files from root directory
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "./..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	c, err = NewClient(Client{
		ApiKey:  os.Getenv("FOREM_API_KEY"),
		Timeout: time.Second * 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func TestClient_User(t *testing.T) {
	ctx := context.Background()
	me, err := c.User(ctx)
	if err != nil {
		t.Error(err)
	}

	if me.ID == 0 {
		t.Errorf("User() error = %v, wantErr %v", err, "0")
	}
}

func TestClient_CreateArticle(t *testing.T) {
	title := "Article was created at " + time.Now().UTC().String()

	type want struct {
		Article *Article
		Err     bool
	}

	tests := []struct {
		name string
		args Article
		want want
	}{
		{
			name: "create article",
			args: Article{
				Title:     title,
				Published: false,
				Tags:      []string{"golang"},
			},
			want: want{
				Article: &Article{
					Title:     title,
					Published: false,
					Tags:      []string{"golang"},
				},
				Err: false,
			},
		},
		{
			name: "without title",
			args: Article{
				Published: false,
				Tags:      []string{"golang"},
			},
			want: want{
				Article: nil,
				Err:     true,
			},
		},
		{
			name: "bad request with the same title",
			args: Article{
				Title:     "111",
				Published: false,
				Tags:      []string{"golang"},
			},
			want: want{
				Article: nil,
				Err:     true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := c.CreateArticle(ctx, tt.args)
			if (err != nil) != tt.want.Err {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.Err)
			}

			if tt.want.Article != nil {
				if tt.want.Article.Title != tt.args.Title {
					t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.Err)
				}
			}
		})
	}
}
