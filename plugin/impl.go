package plugin

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v54/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// Settings for the Plugin.
type Settings struct {
	BaseURL     string
	IssueNum    int
	Key         string
	Message     string
	Update      bool
	APIKey      string
	SkipMissing bool
	IsFile      bool

	baseURL *url.URL
}

var ErrPluginEventNotSupported = errors.New("event not supported")

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	var err error

	if p.pipeline.Build.Event != "pull_request" {
		return fmt.Errorf("%w: %s", ErrPluginEventNotSupported, p.pipeline.Build.Event)
	}

	if p.settings.Message != "" {
		if p.settings.Message, p.settings.IsFile, err = readStringOrFile(p.settings.Message); err != nil {
			return fmt.Errorf("error while reading %s: %w", p.settings.Message, err)
		}
	}

	if !strings.HasSuffix(p.settings.BaseURL, "/") {
		p.settings.BaseURL += "/"
	}

	p.settings.baseURL, err = url.Parse(p.settings.BaseURL)
	if err != nil {
		return fmt.Errorf("failed to parse base url: %w", err)
	}

	if p.settings.Key == "" {
		key := fmt.Sprintf("%s/%s/%d", p.pipeline.Repo.Owner, p.pipeline.Repo.Name, p.settings.IssueNum)
		hash := sha256.Sum256([]byte(key))
		p.settings.Key = fmt.Sprintf("%x", hash)
	}

	if p.settings.Key, _, err = readStringOrFile(p.settings.Key); err != nil {
		return fmt.Errorf("error while reading %s: %w", p.settings.Key, err)
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: p.settings.APIKey})
	tc := oauth2.NewClient(
		context.WithValue(context.Background(), oauth2.HTTPClient, p.network.Client),
		ts,
	)

	client := github.NewClient(tc)
	client.BaseURL = p.settings.baseURL

	commentClient := commentClient{
		Client:   client,
		Repo:     p.pipeline.Repo.Name,
		Owner:    p.pipeline.Repo.Owner,
		Message:  p.settings.Message,
		Update:   p.settings.Update,
		Key:      p.settings.Key,
		IssueNum: p.pipeline.Build.PullRequest,
	}

	if p.settings.SkipMissing && !p.settings.IsFile {
		logrus.Infof("comment skipped: 'message' is not a valid path or file does not exist while 'skip-missing' is enabled")

		return nil
	}

	err := commentClient.issueComment(p.network.Context)
	if err != nil {
		return fmt.Errorf("failed to create or update comment: %w", err)
	}

	return nil
}
