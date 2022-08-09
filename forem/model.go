package forem

import (
	"errors"
	"time"
)

var (
	ErrTitleIsRequired = errors.New("title is required")
)

type Article struct {
	Title          string   `json:"title" required:"true"`
	Published      bool     `json:"published"`
	BodyMarkdown   string   `json:"body_markdown"`
	Tags           []string `json:"tags"`
	Series         string   `json:"series"`
	OrganizationID int      `json:"organization_id"`
}

type ArticleInput struct {
	Article Article `json:"article"`
}

type UserResponse struct {
	ID              int    `json:"id"`
	TypeOf          string `json:"type_of"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	Summary         string `json:"summary"`
	Location        string `json:"location"`
	WebsiteUrl      string `json:"website_url"`
	JoinedAt        string `json:"joined_at"`
	ProfileImage    string `json:"profile_image"`
}

type ArticleResponse struct {
	TypeOf                 string      `json:"type_of"`
	ID                     int         `json:"id"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	CoverImage             string      `json:"cover_image"`
	ReadablePublishDate    string      `json:"readable_publish_date"`
	SocialImage            string      `json:"social_image"`
	TagList                string      `json:"tag_list"`
	Tags                   []string    `json:"tags"`
	Slug                   string      `json:"slug"`
	Path                   string      `json:"path"`
	URL                    string      `json:"url"`
	CanonicalURL           string      `json:"canonical_url"`
	CommentsCount          int         `json:"comments_count"`
	PositiveReactionsCount int         `json:"positive_reactions_count"`
	PublicReactionsCount   int         `json:"public_reactions_count"`
	CollectionID           int         `json:"collection_id"`
	CreatedAt              time.Time   `json:"created_at"`
	EditedAt               interface{} `json:"edited_at"`
	CrosspostedAt          interface{} `json:"crossposted_at"`
	PublishedAt            time.Time   `json:"published_at"`
	LastCommentAt          time.Time   `json:"last_comment_at"`
	BodyHTML               string      `json:"body_html"`
	BodyMarkdown           string      `json:"body_markdown"`
	User                   struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
	Organization struct {
		Name           string `json:"name"`
		Username       string `json:"username"`
		Slug           string `json:"slug"`
		ProfileImage   string `json:"profile_image"`
		ProfileImage90 string `json:"profile_image_90"`
	} `json:"organization"`
}

func (i Article) validate() error {
	if i.Title == "" {
		return errors.New("title is required")
	}

	return nil
}
