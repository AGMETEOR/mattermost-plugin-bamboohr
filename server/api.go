package main

import (
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, world!")
	config := p.getConfiguration()

	if err := config.isValidConfig(); err != nil {
		http.Error(w, "This plugin is not properly configured.", http.StatusNotImplemented)
		return
	}

	// Pick the API key
	// TODO: Authenticate request with key
	key := config.BambooSubdomainAPIKey

	w.Header().Set("Content-Type", "application/json")

	subdomain = p.bambooSubdomain

	p.getEmployeesDirectory(w)

}

func writeError(w http.ResponseWriter, errMessage string) {
	w.Write(errMessage)
}

func (p *Plugin) getEmployeesDirectory(w http.ResponseWriter) {
	config := p.getConfiguration()
	bambooClient := NewClient(nil, p.bambooSubdomain)
	err := bambooClient.buildEmployeeDirectory(config.BambooSubdomainAPIKey)
	if err != nil {
		return
	}
	w.Write(bambooClient.Directory)
}
