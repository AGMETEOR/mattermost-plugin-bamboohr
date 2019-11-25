package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, world!")
	config := p.getConfiguration()

	// Pick the API key
	// TODO: Authenticate request with key
	key := config.BambooSubdomainAPIKey

	if err := config.isValidConfig(); err != nil {
		http.Error(w, "This plugin is not properly configured.", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	subdomain = p.bambooSubdomain

	p.getEmployeesDirectory(w)

}

func (p *Plugin) getEmployeesDirectory(w http.ResponseWriter) {
	config := p.getConfiguration()
	// ctx := context.Background()
	response, err := http.Get(buildBambooURL(p.bambooSubdomain))
	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	w.Write(body)
}

func buildBambooURL(sd string) string {
	baseUrl := "https://api.bamboohr.com/api/gateway.php/%s/v1/employees/directory"
	return fmt.Sprintf(baseUrl, bar)
}
