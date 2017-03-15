package casatunes

import (
  "net/url"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "errors"
  "fmt"
)

type RESTNowPlayingMediaItem struct {
  SourceID int
  QueueCount int
  QueueSongIndex int
  ChangeQueueID int
  ChangeSlideShowID int
  ChangeMetaDataID int
  Status int
  Controls int
  RepeatMode int
  ShuffleMode bool
  MessageID string
  CurrSong RESTMediaItem
  NextSong RESTMediaItem
  CurrProgress int
  SlideShowAvailable bool
  SearchPromptText string
  SourceLockedByZoneID int
  ServiceName string
  ServiceLogoURI string
  PlaylistID string
  PlaylistName string
}

// Access the /sources/{id}/nowplaying API and return *RESTNowPlayingMediaItem
func (client *Client) NowPlaying(source string) (*RESTNowPlayingMediaItem, error) {
  endpoint := client.config.endpoint + "/sources/" + source + "/nowplaying"
  // Validate URL
  _, err := url.ParseRequestURI(endpoint)

  if err != nil {
    return nil, err
  }

  nowPlayingMediaItem := &RESTNowPlayingMediaItem{}
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

  json.Unmarshal(data, nowPlayingMediaItem)
  return nowPlayingMediaItem, nil
}
