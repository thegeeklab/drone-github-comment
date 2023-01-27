package plugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v50/github"
)

// Release holds ties the drone env data and github client together.
type commentClient struct {
	Message  string
	Update   bool
	Key      string
	Repo     string
	Owner    string
	IssueNum int

	*github.Client
	context.Context
}

func (cc *commentClient) issueComment() error {
	var err error
	var comment *github.IssueComment
	var resp *github.Response

	ic := &github.IssueComment{
		Body: &cc.Message,
	}

	if cc.Update {
		// Append plugin comment ID to comment message so we can search for it later
		message := fmt.Sprintf("%s\n<!-- id: %s -->\n", cc.Message, cc.Key)
		ic.Body = &message

		comment, err = cc.comment()

		if err == nil && comment != nil {
			_, resp, err = cc.Client.Issues.EditComment(cc.Context, cc.Owner, cc.Repo, int64(*comment.ID), ic)
		}
	}

	if err == nil && resp == nil {
		_, _, err = cc.Client.Issues.CreateComment(cc.Context, cc.Owner, cc.Repo, cc.IssueNum, ic)
	}

	if err != nil {
		return err
	}

	return nil
}

func (cc *commentClient) comment() (*github.IssueComment, error) {
	var allComments []*github.IssueComment

	opts := &github.IssueListCommentsOptions{}

	for {
		comments, resp, err := cc.Client.Issues.ListComments(cc.Context, cc.Owner, cc.Repo, cc.IssueNum, opts)
		if err != nil {
			return nil, err
		}

		allComments = append(allComments, comments...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	for _, comment := range allComments {
		if strings.Contains(*comment.Body, fmt.Sprintf("<!-- id: %s -->", cc.Key)) {
			return comment, nil
		}
	}

	return nil, nil
}
