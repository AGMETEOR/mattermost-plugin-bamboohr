package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	employeeDirectoryLink = "/v1/employees/directory"
	createEmployeeLink    = "/v1/employees/"
)

type Employee struct {
	ID             string `json:"id"`
	EmployeeNumber string `json:"employeeNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DisplayName    string `json:"displayName"`
	JobTitle       string `json:"jobTitle"`
	Location       string `json:"location"`
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

func reqWithContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

func (c *Client) do(ctx context.Context, key string, r *http.Request) ([]byte, int, error) {
	req := reqWithContext(ctx, r)
	req.SetBasicAuth(key, "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	statusCode := resp.StatusCode

	if err != nil {
		select {
		case <- ctx.Done():
			// we want the ctx error to indicate cancellation
			return nil, statusCode, ctx.Err()
		default:
		}
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

func (c *Client) buildEmployeeDirectory(ctx context.Context, key string, directoryUrl string) (*EmployeeDirectoryResult, int, *APIErrorMessage) {
	directory := new(EmployeeDirectoryResult)

	req, err := http.NewRequest("GET", directoryUrl, nil)

	if err != nil {
		return nil, 0, &APIErrorMessage{message: "Error making a new request"}
	}
	bytes, statusCode, err := c.do(ctx,key, req)

	if err != nil {
		return nil, statusCode, &APIErrorMessage{message: "Error making the request"}
	}

	e := json.Unmarshal(bytes, directory)

	if e != nil {
		return nil, statusCode, &APIErrorMessage{message: "Error unmarshaling the request"}
	}

	return directory, statusCode, nil
}

func (c *Client) addNewEmployee(ctx context.Context, key string, createURL string, employee *Employee) (*Employee, *APIErrorMessage) {
	employeeData := new(Employee)

	jsonData, _ := json.Marshal(employee)

	req, err := http.NewRequest("POST", createURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, &APIErrorMessage{message: "Error making a new request"}
	}

	bytes, _, err := c.do(ctx,key, req)

	if err != nil {
		return nil, &APIErrorMessage{message: "Error making the request"}
	}

	e := json.Unmarshal(bytes, employeeData)

	if e != nil {
		return nil, &APIErrorMessage{message: "Error unmarshaling the request"}
	}

	return employeeData, nil

}

func buildBambooURL(subdomain, baseURL string) string {
	return fmt.Sprintf(baseURL, subdomain)
}

func buildUrlToEndpoint(b, d string) string {
	return b + d
}
