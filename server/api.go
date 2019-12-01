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

	if path == "/api/v1/employees" {
		p.getEmployeesDirectory(w)
		return
	}
	http.NotFound(w, r)
}

type APIErrorMessage struct {
	message string
}

func writeError(w http.ResponseWriter, errMessage *APIErrorMessage) {
	errBytes, _ := json.Marshal(errMessage)
	w.Write(errBytes)
}

func (p *Plugin) getEmployeesDirectory(w http.ResponseWriter) {
	// config := p.getConfiguration()
	bambooClient := NewClient(nil, p.bambooSubdomain)
	// directory, err := bambooClient.buildEmployeeDirectory(config.BambooSubdomainAPIKey)
	directory, err := bambooClient.buildEmployeeDirectory("101915139c16b11e950816e743ae4b1fa96b93e6")
	if err != nil {
		writeError(w, err)
		return
	}
	employees, _ := json.Marshal(directory.Employees)
	w.Write(employees)
}
