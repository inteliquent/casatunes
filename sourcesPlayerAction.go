package casatunes

import (
  "net/url"
  "net/http"
  "errors"
  "fmt"
  "strings"
)

func (client *Client) SourcesPlayerAction(source string, playerAction string) (error) {
  action := strings.ToLower(playerAction)
  endpoint := client.config.endpoint + "/sources/" +
    url.PathEscape(source) + "/player/" +
    url.PathEscape(action)
  // Validate URL
  _, err := url.ParseRequestURI(endpoint)

  if err != nil {
    return err
  }

  // Validate player action
  switch action {
  case "play", "pause", "next", "previous":
    // Do nothing, because these actions are all valid
  default:
    // return an error if a player action isn't supported.
    // valid player actions that aren't supported include:
    //   shuffle, repeat, favorite, position
    err := errors.New(fmt.Sprintf("Unsupported player action [%s]", action))
    return err
  }

  resp, err := http.Get(endpoint)

  if err != nil {
    return err
  }

  if resp.StatusCode != http.StatusOK {
    err := errors.New(fmt.Sprintf("Received HTTP status code [%d]", resp.StatusCode))
    return err
  }

  return nil
}
