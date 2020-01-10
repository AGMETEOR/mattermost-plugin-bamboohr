package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	BaseUrl string

	common service

	// Services to talk to different groups of Bamboo API
	EmployeeService *EmployeeService
}

type service struct {
	client *Client
}

func reqWithContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

func (c *Client) Do(ctx context.Context, key string, r *http.Request) ([]byte, int, error) {
	req := reqWithContext(ctx, r)
	req.SetBasicAuth(key, "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	statusCode := resp.StatusCode

	if err != nil {
		select {
		case <-ctx.Done():
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

func buildBambooURL(subdomain, baseURL string) string {
	return fmt.Sprintf(baseURL, subdomain)
}

func buildUrlToEndpoint(b, d string) string {
	return b + d
}
