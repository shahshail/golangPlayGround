// This program provides a GO API for the GitHub Issue Tracker.

package github

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	User    *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
