package cache

import "net/http"

import "fmt"

import "errors"

const vapourPort int = 3009

//Client defines the getter and setter cache methods
type Client struct {
	Port        int
	BaseURL     string
	StatusCheck string
	GetKeyURL   string
	SetKeyURL   string
}

//Entry defines a key value struct
type Entry struct {
	Key   string
	Value interface{}
}

//Create checks for a connection with the vapour server
func (client *Client) Create() error {
	client.createURLs()
	res, err := http.Get(client.StatusCheck)
	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return errors.New("Connection failed")
	}
	return nil
}

func (client *Client) createURLs() {
	client.StatusCheck = fmt.Sprintf("%s:%d/status", client.BaseURL, client.Port)
	client.GetKeyURL = fmt.Sprintf("%s:%d/get/", client.BaseURL, client.Port)
	client.SetKeyURL = fmt.Sprintf("%s:%d/set", client.BaseURL, client.Port)
}

func (client *Client) createGetURL(key string) string {
	return client.BaseURL + key
}

//Get returns a key value from the cache
func (client *Client) Get(key string) (interface{}, error) {
	res, err := http.Get(client.createGetURL(key))
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		{
			return nil, errors.New("Cache: GET failed")
		}
	}
	return nil, nil
}

//Set puts a key value to the cache
//TODO
func (client *Client) Set(key string) error {
	res, err := http.Get(client.createGetURL(key))
	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		{
			return errors.New("Cache: GET failed")
		}
	}
	return nil
}