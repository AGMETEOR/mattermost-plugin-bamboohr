package main

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/plugin/plugintest/mock"

	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
)

func TestPlugin(t *testing.T) {

	t.Run("test proper setup of plugin on activate", func(t *testing.T) {
		p := Plugin{}
		api := &plugintest.API{}
		helpers := &plugintest.Helpers{}

		api.On("RegisterCommand", mock.AnythingOfType("*model.Command")).Return(nil)

		helpers.On("EnsureBot", mock.AnythingOfType("*model.Bot"), mock.AnythingOfType("EnsureBotOption")).Return("botId", nil)

		p.SetAPI(api)
		p.SetHelpers(helpers)

		p.OnActivate()
		helpers.AssertExpectations(t)

		// We assert that all API calls required to set up plugin were called
		api.AssertExpectations(t)
	})
}
