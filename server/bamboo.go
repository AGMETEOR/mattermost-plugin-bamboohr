package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl               = "https://api.bamboohr.com/api/gateway.php/%s"
	apiVersion            = "v1"
	employeeDirectoryLink = "/v1/employees/directory"
)

type Employee struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	JobTitle    string `json:"jobTitle"`
	Location    string `json:"location"`
}

type EmployeeField struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type EmployeeDirectoryResult struct {
	Fields    []EmployeeField `json:"fields"`
	Employees []Employee      `json:"employees"`
}

type Client struct {
	client  *http.Client
	BaseUrl string
}

// Takes in your company's bamboo subdomain
func NewClient(httpClient *http.Client, subdomain string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	builtBambooUrl := buildBambooURL(subdomain, baseUrl)

	c := &Client{
		client:  httpClient,
		BaseUrl: builtBambooUrl,
	}

	return c
}

func (c *Client) do(key string, req *http.Request) ([]byte, int, error) {
	req.SetBasicAuth(key, "")
	req.Header.Set("Accept", "application/json")

	// Might not be necessary
	client := &http.Client{}
	resp, err := client.Do(req)

	statusCode := resp.StatusCode

	if err != nil {
		return nil, statusCode, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	if 200 != statusCode {
		return nil, statusCode, fmt.Errorf("%s", body)
	}

	return body, statusCode, nil
}

func (c *Client) buildEmployeeDirectory(key string) (*EmployeeDirectoryResult, int, *APIErrorMessage) {
	directory := new(EmployeeDirectoryResult)
	directoryUrl := buildUrlToDirectory(c.BaseUrl, employeeDirectoryLink)
	req, err := http.NewRequest("GET", directoryUrl, nil)

	if err != nil {
		return nil, 0, &APIErrorMessage{message: "Error making a new request"}
	}
	bytes, statusCode, err := c.do(key, req)

	if err != nil {
		return nil, 0, &APIErrorMessage{message: "Error making the request"}
	}

	e := json.Unmarshal(bytes, directory)

	if e != nil {
		return nil, 0, &APIErrorMessage{message: "Error unmarshaling the request"}
	}

	return directory, statusCode, nil
}

func buildBambooURL(subdomain string, baseUrl string) string {
	return fmt.Sprintf(baseUrl, subdomain)
}

func buildUrlToDirectory(b string, d string) string {
	return b + d
}
