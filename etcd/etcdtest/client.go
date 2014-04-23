// Copyright (c) 2013 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package etcdtest

import (
	"errors"
	"github.com/coreos/go-etcd/etcd"
	"strings"
	"time"
)

// Client represents a fake etcd client. Used for testing.
type Client struct {
	Responses map[string]*etcd.Response
}

// Get mimics the etcd.Client.Get() method.
func (c *Client) Get(key string, sort, recurse bool) (*etcd.Response, error) {
	return c.Responses[key], nil
}

// Wait mimics waiting for changes
func (c *Client) Watch(prefix string, waitIndex uint64, recursive bool,
	receiver chan *etcd.Response, stop chan bool) (*etcd.Response, error) {
	for key, value := range c.Responses {
		if strings.HasPrefix(key, prefix) {
			receiver <- value
		}
	}
	select {
	case <-stop:
		return nil, errors.New("Etcd watch signalled to stop!")
	case <-time.After(time.Duration(1 * time.Millisecond)):
		return nil, nil
	}
}

// AddResponses adds or updates the Client.Responses map.
func (c *Client) AddResponse(key string, response *etcd.Response) {
	c.Responses[key] = response
}

// NewClient returns a fake etcd client.
func NewClient() *Client {
	responses := make(map[string]*etcd.Response)
	return &Client{responses}
}
