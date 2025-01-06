package dnd5e

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	apiUrl string
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can not be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			// Transport: &loggingRoundTripper{
			// 	logger: os.Stdout,
			// 	next:   http.DefaultTransport,
			// },
		},
		apiUrl: "https://www.dnd5eapi.co/api/monsters",
	}, nil
}

func (c Client) GetMonsters() ([]MonsterShot, error) {
	res, err := c.client.Get(c.apiUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r monstersResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Monsters, nil
}

func (c Client) GetMonster(index string) (MonsterFullResponse, error) {
	url := fmt.Sprintf("%s/%s", c.apiUrl, index)
	res, err := c.client.Get(url)
	if err != nil {
		return MonsterFullResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return MonsterFullResponse{}, err
	}

	var r MonsterFullResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return MonsterFullResponse{}, err
	}

	return r, nil
}
