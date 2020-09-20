package main

import (
	"github.com/urfave/cli/v2"
	"github.com/thegeeklab/drone-github-comment/plugin"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "api-key",
			Usage:       "api key to access github api",
			EnvVars:     []string{"PLUGIN_API_KEY", "GITHUB_COMMENT_API_KEY"},
			Destination: &settings.APIKey,
		},
		&cli.StringFlag{
			Name:        "base-url",
			Value:       "https://api.github.com/",
			Usage:       "api url, needs to be changed for ghe",
			EnvVars:     []string{"PLUGIN_BASE_URL", "GITHUB_COMMENT_BASE_URL"},
			Destination: &settings.BaseURL,
		},
		&cli.StringFlag{
			Name:        "key",
			Usage:       "key to assign comment",
			EnvVars:     []string{"PLUGIN_KEY", "GITHUB_COMMENT_KEY"},
			Destination: &settings.Key,
		},
		&cli.StringFlag{
			Name:        "message",
			Usage:       "file or string with comment message",
			EnvVars:     []string{"PLUGIN_MESSAGE", "GITHUB_COMMENT_MESSAGE"},
			Destination: &settings.Message,
		},
		&cli.BoolFlag{
			Name:        "update",
			Usage:       "update an existing comment that matches the key",
			EnvVars:     []string{"PLUGIN_UPDATE", "GITHUB_COMMENT_UPDATE"},
			Destination: &settings.Update,
		},
	}
}
