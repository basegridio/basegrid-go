package basegrid

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// BaseURL is the default API endpoint
const BaseURL = "https://basegrid-production.up.railway.app/v1"

// Client is the BaseGrid API client
type Client struct {
    APIKey     string
    BaseURL    string
    HTTPClient *http.Client
}

// Memory represents a stored memory
type Memory struct {
    ID         string                 `json:"id,omitempty"`
    AgentID    string                 `json:"agentId"`
    Content    string                 `json:"content"`
    Metadata   map[string]interface{} `json:"metadata,omitempty"`
    Importance float64                `json:"importance,omitempty"`
    CreatedAt  time.Time              `json:"createdAt,omitempty"`
}

// SearchParams represents search query parameters
type SearchParams struct {
    AgentID   string                 `json:"agentId"`
    Query     string                 `json:"query"`
    Limit     int                    `json:"limit,omitempty"`
    Threshold float64                `json:"threshold,omitempty"`
    Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// SearchResult represents a single search result
type SearchResult struct {
    ID         string                 `json:"id"`
    Content    string                 `json:"content"`
    Similarity float64                `json:"similarity"`
    Metadata   map[string]interface{} `json:"metadata"`
    CreatedAt  time.Time              `json:"createdAt"`
}

// searchResponse is the internal API response structure
type searchResponse struct {
    Success bool           `json:"success"`
    Results []SearchResult `json:"results"`
}

// New creates a new BaseGrid client
func New(apiKey string) *Client {
    return &Client{
        APIKey:     apiKey,
        BaseURL:    BaseURL,
        HTTPClient: &http.Client{Timeout: 10 * time.Second},
    }
}

// Add stores a new memory
func (c *Client) Add(mem Memory) (*Memory, error) {
    body, err := json.Marshal(mem)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", c.BaseURL+"/memories", bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    c.addHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("API error: %s", resp.Status)
    }

    var result struct {
        Success bool   `json:"success"`
        Data    Memory `json:"data"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return &result.Data, nil
}

// Search retrieves relevant memories
func (c *Client) Search(params SearchParams) ([]SearchResult, error) {
    body, err := json.Marshal(params)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", c.BaseURL+"/memories/search", bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    c.addHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("API error: %s", resp.Status)
    }

    var result searchResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Results, nil
}

func (c *Client) addHeaders(req *http.Request) {
    req.Header.Set("Authorization", "Bearer "+c.APIKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "basegrid-go/1.0.0")
}
