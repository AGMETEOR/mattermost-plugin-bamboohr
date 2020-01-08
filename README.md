# (WIP) Mattermost BambooHR Plugin
[![CircleCI](https://circleci.com/gh/AGMETEOR/mattermost-plugin-bamboohr.svg?style=svg)](https://circleci.com/gh/AGMETEOR/mattermost-plugin-bamboohr)

[Bamboo HR](https://www.bamboohr.com/) is a software platform for companies to manage their human resources (HR).

This plugin is meant to be used by users/companies using both [Mattermost](https://www.mattermost.org/) for their chat and BambooHR for human resource related issues.

Contributions are welcome!

## Table of Contents
- [1. Features](#1-features)
- [2. Configuration](#2-configuration)
- [3. Development](#3-development)

## 1. Features
### 1.1 Create an employee on BambooHR using an existing Mattermost user
If a system admin has added a new Mattermost account for an employee, an authorized user can use ```/bamboo add``` from the DM channel to add that user to BambooHR.

## 2. Configuration
Configure the plugin in Mattermost by going to ```System Console > Plugins > Bamboohr```. Enable the plugin if it's not enabled. Set your BambooHR domain name, your BambooHR API Key and a list of comma separated Mattermost user IDs that'll be allowed to run some crucial bamboo plugin commands.

## 3. Development
- Fork this repo
- Clone your fork and make changes on your branch
- Run ```$ make``` at the root of this project
- Install the generated tar on your server to see your changes

