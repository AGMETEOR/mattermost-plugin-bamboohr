package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const COMMAND_HELP = `* |/bamboo test| - Run 'test' to see if you're configured to run bamboo commands
* |/bamboo help| - Run 'help' to see a list of commands available for you
`

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "bamboo",
		DisplayName:      "Bamboo",
		Description:      "Integration with BambooHR.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: test",
		AutoCompleteHint: "[command]",
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, message string) {
	post := &model.Post{
		UserId:    p.BotUserID,
		ChannelId: args.ChannelId,
		Message:   message,
	}

	_ = p.API.SendEphemeralPost(args.UserId, post)
}

func (p *Plugin) responsef(commandArgs *model.CommandArgs, format string, args ...interface{}) *model.CommandResponse {
	p.postCommandResponse(commandArgs, fmt.Sprintf(format, args...))
	return &model.CommandResponse{}
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	action := ""
	if len(split) > 1 {
		action = split[1]
	}

	if command != "/bamboo" {
		return &model.CommandResponse{}, nil
	}

	switch action {
	case "test":
		return p.testCommandFunc(args)
	case "add":
		return p.createEmployeeCommandFunc(args)
	case "help":
		return p.helpCommandFunc(args)
	default:
		return p.responsef(args, fmt.Sprintf("Unknown action %v", action)), nil
	}
}

func (p *Plugin) testCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	serverConfig := p.API.GetConfig()
	pluginConfig := p.getConfiguration()

	if serverConfig.ServiceSettings.SiteURL == nil {
		return p.responsef(args, "SiteURL not set. Encountered an error testing integration to BambooHR."), nil
	}

	token := pluginConfig.BambooAPIKey
	domain := pluginConfig.BambooDomain

	if token == "" || domain == "" {
		return p.responsef(args, "Bamboo configuration for your server is not set."), nil
	}

	if !p.isUserAuthorized(args.UserId) {
		return p.responsef(args, "You will not be authorized to run Bamboo commands."), nil
	}

	// Verify that configured Bamboo token works
	bambooClient := p.getClient(domain)
	dURL := buildUrlToEndpoint(bambooClient.BaseUrl, employeeDirectoryLink)
	_, statusCode, _ := bambooClient.buildEmployeeDirectory(token, dURL)

	if statusCode == 200 {
		return p.responsef(args, "Congratulations! Bamboo is correctly configured on your server."), nil
	}

	return p.responsef(args, "Ooops! Test was unsuccessful."), nil

}

func (p *Plugin) createEmployeeCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	serverConfig := p.API.GetConfig()
	pluginConfig := p.getConfiguration()

	if serverConfig.ServiceSettings.SiteURL == nil {
		return p.responsef(args, "SiteURL not set. Encountered an error testing integration with BambooHR."), nil
	}

	token := pluginConfig.BambooAPIKey
	domain := pluginConfig.BambooDomain

	channelID := args.ChannelId
	userID := args.UserId

	// Only users allowed by system admins can run bamboo commands
	// For example, an HR Manager is an allowed user to run bamboo commands
	if !p.isUserAuthorized(userID) {
		return p.responsef(args, "You will not be authorized to run Bamboo commands."), nil
	}

	channel, _ := p.API.GetChannel(channelID)
	channelName := channel.Name

	if channel.Type != "D" {
		return p.responsef(args, "Run add command against a DM channel"), nil
	}

	splitChannelName := strings.Split(channelName, "__")

	employeeId := ""

	employeeId = splitChannelName[1]

	if employeeId == userID {
		employeeId = splitChannelName[0]
	}

	employee, _ := p.API.GetUser(employeeId)

	if token == "" || domain == "" {
		return p.responsef(args, "Bamboo configuration for your server is not properly set."), nil
	}

	bambooClient := p.getClient(domain)
	empURL := buildUrlToEndpoint(bambooClient.BaseUrl, createEmployeeLink)

	data := &Employee{
		EmployeeNumber: employee.Id,
		FirstName:      employee.FirstName,
		LastName:       employee.LastName,
	}

	_, err := bambooClient.addNewEmployee(token, empURL, data)

	if err != nil {
		return p.responsef(args, "Failed to add this employee"), nil
	}

	return p.responsef(args, "New employee added"), nil

}

func (p *Plugin) helpCommandFunc(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := "###### Mattermost Bamboo Plugin - Slash Command Help\n" + strings.Replace(COMMAND_HELP, "|", "`", -1)
	return p.responsef(args, text), nil
}
