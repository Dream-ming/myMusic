package search

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type OpenSearchClient struct {
    endpoint   string 
    index      string 
    httpClient *http.Client
}

func NewOpenSearchClient(endpoint, index string) *OpenSearchClient {
    return &OpenSearchClient{
        endpoint:   endpoint,
        index:      index,
        httpClient: &http.Client{Timeout: 10 * time.Second},
    }
}

func (c *OpenSearchClient) IndexDoc(ctx context.Context, id string, doc interface{}) error {
    url := fmt.Sprintf("%s/%s/_doc/%s", c.endpoint, c.index, id)
    body, err := json.Marshal(doc)
    if err != nil {
        return err
    }
    req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(body))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode >= 300 {
        respBody, _ := ioutil.ReadAll(resp.Body)
        return fmt.Errorf("index doc failed: %s", string(respBody))
    }
    return nil
}

func (c *OpenSearchClient) GetDoc(ctx context.Context, id string) ([]byte, error) {
    url := fmt.Sprintf("%s/%s/_doc/%s", c.endpoint, c.index, id)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode == 404 {
        return nil, nil
    }
    if resp.StatusCode >= 300 {
        respBody, _ := ioutil.ReadAll(resp.Body)
        return nil, fmt.Errorf("get doc failed: %s", string(respBody))
    }
    return ioutil.ReadAll(resp.Body)
}

func (c *OpenSearchClient) Search(ctx context.Context, query map[string]interface{}) ([]byte, error) {
    url := fmt.Sprintf("%s/%s/_search", c.endpoint, c.index)
    body, err := json.Marshal(query)
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewReader(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode >= 300 {
        respBody, _ := ioutil.ReadAll(resp.Body)
        return nil, fmt.Errorf("search failed: %s", string(respBody))
    }
    return ioutil.ReadAll(resp.Body)
}

func (c *OpenSearchClient) DeleteDoc(ctx context.Context, id string) error {
    url := fmt.Sprintf("%s/%s/_doc/%s", c.endpoint, c.index, id)
    req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
    if err != nil {
        return err
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 && resp.StatusCode != 404 {
        respBody, _ := ioutil.ReadAll(resp.Body)
        return fmt.Errorf("delete doc failed: %s", string(respBody))
    }
    return nil
}