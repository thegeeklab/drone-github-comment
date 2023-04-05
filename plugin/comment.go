package plugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v51/github"
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
}

func (cc *commentClient) issueComment(ctx context.Context) error {
	var (
		err     error
		comment *github.IssueComment
		resp    *github.Response
	)

	issueComment := &github.IssueComment{
		Body: &cc.Message,
	}

	if cc.Update {
		// Append plugin comment ID to comment message so we can search for it later
		message := fmt.Sprintf("%s\n<!-- id: %s -->\n", cc.Message, cc.Key)
		issueComment.Body = &message

		comment, err = cc.comment(ctx)

		if err == nil && comment != nil {
			_, resp, err = cc.Client.Issues.EditComment(ctx, cc.Owner, cc.Repo, *comment.ID, issueComment)
		}
	}

	if err == nil && resp == nil {
		_, _, err = cc.Client.Issues.CreateComment(ctx, cc.Owner, cc.Repo, cc.IssueNum, issueComment)
	}

	if err != nil {
		return err
	}

	return nil
}

func (cc *commentClient) comment(ctx context.Context) (*github.IssueComment, error) {
	var allComments []*github.IssueComment

	opts := &github.IssueListCommentsOptions{}

	for {
		comments, resp, err := cc.Client.Issues.ListComments(ctx, cc.Owner, cc.Repo, cc.IssueNum, opts)
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

	//nolint:nilnil
	return nil, nil
}
