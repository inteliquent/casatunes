package casatunes

import (
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
  "fmt"
)

func (client *Client) MediaSearchMC(mcId string, searchText string) (*RESTMediaCollectionItem, error) {
  endpoint := client.config.endpoint + "/media/search/" + mcId + "/" + searchText
  // Validate URL
  _, err := url.ParseRequestURI(endpoint)

  if err != nil {
    return nil, err
  }

  MediaCollectionItem := &RESTMediaCollectionItem{}
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

  err = json.Unmarshal(data, MediaCollectionItem)

  if err != nil {
    return nil, err
  }
  return MediaCollectionItem, nil
}
