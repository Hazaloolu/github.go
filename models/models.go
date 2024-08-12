package models

// repo owner

type Owner struct {
	Login string
}

// item = repo data structure

type Item struct {
	ID              int
	Name            string
	FullName        string `json:"full_name`
	Owner           Owner
	Description     string
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:stargazers_count`
}

// api response
type JSONData struct {
	Count int `json:"total_count`
	Items []Item
}
