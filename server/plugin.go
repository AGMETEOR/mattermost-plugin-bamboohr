package main

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *Configuration

	BotUserID string
}

func (p *Plugin) OnActivate() error {
	p.API.RegisterCommand(getCommand())

	profileImage := filepath.Join("assets", "bamboo.png")

	botId, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    "bamboo",
		DisplayName: "Bamboo",
		Description: "Created by the BambooHR plugin.",
	}, plugin.ProfileImagePath(profileImage))

	if err != nil {
		return errors.Wrap(err, "failed to ensure bamboo bot")
	}
	p.BotUserID = botId

	return nil
}

func (p *Plugin) isUserAuthorized(id string) bool {
	pluginConfig := p.getConfiguration()
	adminsList := pluginConfig.BambooAdmins
	allowedBambooAdmins := strings.Split(adminsList, ",")
	userAllowed := contains(allowedBambooAdmins, id)
	return userAllowed
}

func contains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
