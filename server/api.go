package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
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
	bambooClient := p.getClient(pluginConfig.BambooDomain)
	dURL := buildUrlToEndpoint(bambooClient.BaseUrl, employeeDirectoryLink)

	ctx := context.Background()
	directory, _, err := bambooClient.EmployeeService.BuildEmployeeDirectory(ctx, pluginConfig.BambooAPIKey, dURL)
	if err != nil {
		writeError(w, err)
		return
	}
	employees, _ := json.Marshal(directory.Employees)
	w.Write(employees)
}
