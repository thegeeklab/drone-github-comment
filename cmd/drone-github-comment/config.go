package main

import (
	"github.com/thegeeklab/drone-github-comment/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings, category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "api-key",
			EnvVars:     []string{"PLUGIN_API_KEY", "GITHUB_COMMENT_API_KEY"},
			Usage:       "personal access token access github api",
			Destination: &settings.APIKey,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "base-url",
			EnvVars:     []string{"PLUGIN_BASE_URL", "GITHUB_COMMENT_BASE_URL"},
			Usage:       "api url",
			Value:       "https://api.github.com/",
			Destination: &settings.BaseURL,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "key",
			EnvVars:     []string{"PLUGIN_KEY", "GITHUB_COMMENT_KEY"},
			Usage:       "unique identifier to assign to comment",
			Destination: &settings.Key,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "message",
			EnvVars:     []string{"PLUGIN_MESSAGE", "GITHUB_COMMENT_MESSAGE"},
			Usage:       "file or string that holds the comment text",
			Destination: &settings.Message,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "update",
			EnvVars:     []string{"PLUGIN_UPDATE", "GITHUB_COMMENT_UPDATE"},
			Usage:       "update existing comment that matches the key",
			Value:       false,
			Destination: &settings.Update,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "skip-missing",
			EnvVars:     []string{"PLUGIN_SKIP_MISSING", "GITHUB_COMMENT_SKIP_MISSING"},
			Usage:       "skip creation of comment if message file does not exist",
			Value:       false,
			Destination: &settings.SkipMissing,
			Category:    category,
		},
	}
}
