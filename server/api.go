package main

import (
	// "github.com/LenzEducation/lenz-server/config"
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	// config := p.getConfiguration()

	// if err := config.isValidConfig(); err != nil {
	// 	http.Error(w, "This plugin is not properly configured.", http.StatusNotImplemented)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path

	switch path {
	case "/api/v1/employees":
		p.getEmployeesDirectory(w, r)
	default:
		http.NotFound(w, r)
	}
}

type APIErrorMessage struct {
	message string
}

func writeError(w http.ResponseWriter, errMessage *APIErrorMessage) {
	errBytes, _ := json.Marshal(errMessage)
	w.Write(errBytes)
}

func (p *Plugin) getEmployeesDirectory(w http.ResponseWriter, r *http.Request) {
	pluginConfig := p.getConfiguration()
	bambooClient := NewClient(nil, pluginConfig.BambooDomain)
	directory, _, err := bambooClient.buildEmployeeDirectory(pluginConfig.BambooAPIKey)
	if err != nil {
		writeError(w, err)
		return
	}
	employees, _ := json.Marshal(directory.Employees)
	w.Write(employees)
}
