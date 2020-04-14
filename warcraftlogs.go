package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"time"
)

const APIDomain = "classic.warcraftlogs.com"

type Client struct {
	httpClient *http.Client
	APIKey     string
}

func NewClient(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		httpClient: httpClient,
		APIKey:     apiKey,
	}
}

func (c *Client) doGet(endpoint string, v interface{}) error {
	urlO := "https://" + path.Join(APIDomain, endpoint)
	urlP, _ := url.Parse(urlO)
	q := urlP.Query()
	q.Set("api_key", c.APIKey)
	urlP.RawQuery = q.Encode()
	response, err := c.httpClient.Get(urlP.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(contents, &v)
}

func (c *Client) GetParses(query ParsesQuery) (ParsesResponse, error) {
	var result ParsesResponse
	endpoint := fmt.Sprintf("/v1/parses/character/%s/%s/%s", query.CharacterName, query.Server, query.Region)
	if err := c.doGet(endpoint, &result); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	sort.Sort(sort.Reverse(ByDate(result)))

	p := result[0]
	if p.Spec == "Healer" {
		return c.GetHealingParses(query)
	}
	return result, nil
}

func (c *Client) GetHealingParses(query ParsesQuery) (ParsesResponse, error) {
	var result ParsesResponse
	endpoint := fmt.Sprintf("/v1/parses/character/%s/%s/%s?metric=hps", query.CharacterName, query.Server, query.Region)
	if err := c.doGet(endpoint, &result); err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(ByDate(result)))

	return result, nil
}

type ByDate ParsesResponse

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDate) Less(i, j int) bool { return d[i].StartTime < d[j].StartTime }
