package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestConfiguration struct {
	BambooDomain string
	BambooAPIKey string
	BambooAdmins string
}

func TestExecuteCommand(t *testing.T) {
	t.Run("test for test command", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		defer svr.Close()

		p := Plugin{
			BambooBaseURL: svr.URL + "/%s",
		}

		p.setConfiguration(&Configuration{
			BambooDomain: "test_domain",
			BambooAPIKey: "APIKEY",
			BambooAdmins: "id1, id2",
		})

		commandArgs := &model.CommandArgs{
			Command: "/bamboo test",
			UserId:  "id1",
		}

		isSendEphemeralPostCalled := false

		siteURL := "http://test.com"

		api := &plugintest.API{}
		api.On("GetConfig").Return(&model.Config{ServiceSettings: model.ServiceSettings{SiteURL: &siteURL}})

		api.On("SendEphemeralPost", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Post")).Run(func(args mock.Arguments) {
			isSendEphemeralPostCalled = true
			post := args.Get(1).(*model.Post)
			assert.Equal(t, "Congratulations! Bamboo is correctly configured on your server.", post.Message)
		}).Once().Return(&model.Post{})
		p.SetAPI(api)
		p.ExecuteCommand(&plugin.Context{}, commandArgs)
		assert.Equal(t, true, isSendEphemeralPostCalled)
	})
}
