package casatunes

import (
  "testing"
  "net/http"
  "reflect"
)

type mediaSourcesPlayMediaHandler struct {
  response string
  status int
}

func newMediaSourcesPlayMediaHandler() *mediaSourcesPlayMediaHandler {
  return &mediaSourcesPlayMediaHandler{
    response: `{"Result": 1}`,
    status: http.StatusOK,
  }
}

func (msh *mediaSourcesPlayMediaHandler) handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(msh.status)
  w.Write([]byte(msh.response))
}

// casatunes.NowPlaying should return object casatunes.RESTNowPlayingMediaItem
// from casaplayerEndpoint
func TestMediaSourcesPlayMedia(t *testing.T) {
  once.Do(startServer)
  casaplayerEndpoint := "http://" + serverAddr + "/"
  mediaId := "0123abasdf4567890" // Spotify internet radio
  source := "0"
  mspm := newMediaSourcesPlayMediaHandler()

  http.HandleFunc(
    "/media/sources/" + source + "/play/" + mediaId,
    func(w http.ResponseWriter, r *http.Request) { mspm.handler(w, r) },
  )

  t.Run("Valid object returned", func(t *testing.T) {
    client := New(casaplayerEndpoint)
    resp, err := client.MediaSourcesPlayMedia(source, mediaId)

    if err != nil {
      t.Fatal(err)
    }

    if reflect.TypeOf(resp) != reflect.TypeOf(&RESTResultInteger{}) {
      t.Fatal("NowPlaying did not return object of type casatunes.RESTResultInteger !")
    } else {
      t.Log("NowPlaying returned casatunes.RESTNowPlayingMediaItem object")
    }

  })
}
