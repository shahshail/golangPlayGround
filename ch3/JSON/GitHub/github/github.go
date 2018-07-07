// This program provides a GO API for the GitHub Issue Tracker.

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

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

//SearchIssue quaries the GitHub issue tracker
func SearchIssue(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ""))

	resp, err := http.Get(IssueURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed %s :", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
