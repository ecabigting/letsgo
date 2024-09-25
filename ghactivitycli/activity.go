package main

// Define the structs to match the JSON structure
type Actor struct {
	AvatarURL    string `json:"avatar_url"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	ID           int    `json:"id"`
	Login        string `json:"login"`
	URL          string `json:"url"`
}

type CommitAuthor struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Commit struct {
	Author   CommitAuthor `json:"author"`
	Distinct bool         `json:"distinct"`
	Message  string       `json:"message"`
	SHA      string       `json:"sha"`
	URL      string       `json:"url"`
}

type Payload struct {
	Before       string   `json:"before"`
	Commits      []Commit `json:"commits"`
	DistinctSize int      `json:"distinct_size"`
	Head         string   `json:"head"`
	PushID       int      `json:"push_id"`
	Ref          string   `json:"ref"`
	RepositoryID int      `json:"repository_id"`
	Size         int      `json:"size"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Event struct {
	Actor     Actor   `json:"actor"`
	CreatedAt string  `json:"created_at"`
	ID        string  `json:"id"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	Repo      Repo    `json:"repo"`
	Type      string  `json:"type"`
}

type Activity struct {
	Count    int
	RepoName string
	Type     string
}
