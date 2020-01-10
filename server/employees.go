package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	employeeDirectoryLink = "/v1/employees/directory"
	createEmployeeLink    = "/v1/employees/"
)

type EmployeeService service

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

func (eS *EmployeeService) AddNewEmployee(ctx context.Context, key string, createURL string, employee *Employee) (*Employee, *APIErrorMessage) {
	employeeData := new(Employee)

	jsonData, _ := json.Marshal(employee)

	req, err := http.NewRequest("POST", createURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, &APIErrorMessage{message: "Error making a new request"}
	}

	bytes, _, err := eS.client.Do(ctx, key, req)

	if err != nil {
		return nil, &APIErrorMessage{message: "Error making the request"}
	}

	e := json.Unmarshal(bytes, employeeData)

	if e != nil {
		return nil, &APIErrorMessage{message: "Error unmarshaling the request"}
	}

	return employeeData, nil
}

func (eS *EmployeeService) BuildEmployeeDirectory(ctx context.Context, key string, directoryUrl string) (*EmployeeDirectoryResult, int, *APIErrorMessage) {
	directory := new(EmployeeDirectoryResult)

	req, err := http.NewRequest("GET", directoryUrl, nil)

	if err != nil {
		return nil, 0, &APIErrorMessage{message: "Error making a new request"}
	}
	bytes, statusCode, err := eS.client.Do(ctx, key, req)

	if err != nil {
		return nil, statusCode, &APIErrorMessage{message: "Error making the request"}
	}

	e := json.Unmarshal(bytes, directory)

	if e != nil {
		return nil, statusCode, &APIErrorMessage{message: "Error unmarshaling the request"}
	}

	return directory, statusCode, nil
}
