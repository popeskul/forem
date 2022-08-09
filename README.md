Forem client.
=======================

You need API key from forem.com.

See it in action:

```go
package main

import (
	"context"
	"fmt"
	"github.com/popeskul/forem"
	"log"
	"os"
	"time"
)

func main() {
	client, err := forem.NewClient(forem.Client{
		ApiKey:  os.Getenv("FOREM_API_KEY"),
		Timeout: time.Second * 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get information about the current user.
	user, err := client.User(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", user)

	// Create a new topic.
	newArticle := forem.Article{
		Title:     "Article was created at " + strconv.FormatInt(date, 10),
		Published: false,
		Tags:      []string{"golang"},
	}
	article, err := client.CreateArticle(ctx, newArticle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Article created: %+v \n", article)
}
```

![Go Report](https://goreportcard.com/badge/github.com/popeskul/forem)
![Repository Top Language](https://img.shields.io/github/languages/top/popeskul/forem)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/popeskul/forem)
![Github Repository Size](https://img.shields.io/github/repo-size/popeskul/forem)
![Github Open Issues](https://img.shields.io/github/issues/popeskul/forem)
![Lines of code](https://img.shields.io/tokei/lines/github/popeskul/forem)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/popeskul/forem)