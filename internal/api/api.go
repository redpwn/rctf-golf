package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type APIResponse struct {
	Kind    string          `json:"kind"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type ClientConfig struct {
	StartTime int64 `json:"startTime"`
	//EndTime   int64 `json:"endTime"`
}

type ChallengeSolve struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
}

type ChallengeSolvesData struct {
	Solves []ChallengeSolve `json:"solves"`
}

type GetChallengeSolvesParams struct {
	Limit  int
	Offset int
}

type ResponseError = APIResponse

func (e *ResponseError) Error() string {
	return e.Kind + ": " + e.Message
}

func parseResponse(resp *http.Response, expectedKind string, data interface{}) error {
	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}

	if apiResp.Kind != expectedKind {
		err := ResponseError(apiResp)
		return &err
	}

	if data != nil {
		if err := json.Unmarshal(apiResp.Data, data); err != nil {
			return err
		}
	}

	return nil
}

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *APIClient {
	return NewClientWithHTTPClient(baseURL, http.DefaultClient)
}

func NewClientWithHTTPClient(baseURL string, httpClient *http.Client) *APIClient {
	return &APIClient{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}

func (c *APIClient) GetChallengeSolves(challID string, params GetChallengeSolvesParams) ([]ChallengeSolve, error) {
	qs := url.Values{}
	qs.Add("limit", fmt.Sprint(params.Limit))
	qs.Add("offset", fmt.Sprint(params.Offset))
	resp, err := c.httpClient.Get(c.baseURL + fmt.Sprintf("/api/v1/challs/%s/solves?%s", url.PathEscape(challID), qs.Encode()))
	if err != nil {
		return nil, err
	}
	var data ChallengeSolvesData
	if err := parseResponse(resp, "goodChallengeSolves", &data); err != nil {
		return nil, err
	}
	return data.Solves, nil
}

func (c *APIClient) GetClientConfig() (*ClientConfig, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/v1/integrations/client/config")
	if err != nil {
		return nil, err
	}
	var data ClientConfig
	if err := parseResponse(resp, "goodClientConfig", &data); err != nil {
		return nil, err
	}
	return &data, nil
}
