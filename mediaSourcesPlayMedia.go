package casatunes

import (
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
  "fmt"
)

func (client *Client) MediaSourcesPlayMedia(source string, mediaId string) (*RESTResultInteger, error) {
  endpoint := client.config.endpoint + "/media/sources/" + source + "/play/" + mediaId
  // Validate URL
  _, err := url.ParseRequestURI(endpoint)

  if err != nil {
    return nil, err
  }

  resultInteger := &RESTResultInteger{}
  resp, err := http.Get(endpoint)

  if err != nil {
    return nil, err
  }

  if resp.StatusCode != http.StatusOK {
    err := errors.New(fmt.Sprintf("Received HTTP status code [%d]", resp.StatusCode))
    return nil, err
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(data, resultInteger)

  if err != nil {
    return nil, err
  }

  return resultInteger, nil
}
