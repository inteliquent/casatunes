package casatunes

import (
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
  "fmt"
)

func (client *Client) MediaSourcesPlayMedia(source, mediaId, addToQueue string) (*RESTResultInteger, error) {
  // addToQueue
  // playNow (Replace Queue Items and Play Now)
  // add (Add Items to Queue)
  // addplay (Add Items to Queue and start playing the first newly added item)
  switch addToQueue {
  case "playnow", "add", "addplay":
    // Do nothing, because these actions are valid
  default:
    addToQueue = "addplay"
  }
  endpoint := client.config.endpoint + "/media/sources/" +
    url.PathEscape(source) + "/play/" +
    url.PathEscape(mediaId) + "?addToQueue=" + addToQueue
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
