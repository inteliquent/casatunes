package casatunes

import (
  "testing"
  "net/http"
  "fmt"
)

type sourcesPlayerAction struct {
  response string
  status int
}

func newSourcesPlayerAction() *sourcesPlayerAction {
  return &sourcesPlayerAction{
    response: `{"SourceID": 0,"QueueCount": 82,"QueueSongIndex": 7,"ChangeQueueID": 60,"ChangeSlideShowID": 1,"ChangeMetaDataID": 3333,"Status": 2,"Controls": 50594303,"RepeatMode": 1,"ShuffleMode": true,"CurrSong": {  "Flags": 12289,  "ID": "c0566cf3e6bd94868d152fc6086d7c6a",  "Title": "Reaching Out",  "Type": 0,  "ArtworkURI": "http://i.scdn.co/image/4b1e40eff8d4ab6f380a712ecef89c8945fee104",  "ArtworkRatio": 0,  "Album": "Unlimited",  "Artists": "Bassnectar",  "Duration": 294,  "Track": "1",  "BitRate": 0,  "ListenerCount": 0,  "Rating": 62,  "TotalItems": 0,  "DisplayInfo": [    "Reaching Out",    "Bassnectar",    "Unlimited"  ],  "ContextMenuItems": [    {      "Type": 1,      "Title": "Go to Album",      "Value": "spotify4:album:846df04b-59c4-4289-83fc-28267ed59cfb:2n9RwIM1CdRV4GZzC7sfWa"    },    {      "Type": 1,      "Title": "Go to Artist",      "Value": "spotify4:artist:846df04b-59c4-4289-83fc-28267ed59cfb:1JPy5PsJtkhftfdr6saN2i"    },    {      "Type": 1,      "Title": "More Like This",      "Value": "spotify4:rec:846df04b-59c4-4289-83fc-28267ed59cfb:3anyoDE1gcNsRtLmkE55bU"    }  ]},"NextSong": {  "Flags": 12289,  "ID": "ec1d6940d001b34e76e5e5467b45b39d",  "Title": "Fardration",  "Type": 0,  "ArtworkURI": "http://i.scdn.co/image/a292d6bdab4369cb4f5ce8d1a3e60b5d5fd608a5",  "ArtworkRatio": 0,  "Album": "Getting Along",  "Artists": "LRKR",  "Duration": 216,  "Track": "2",  "BitRate": 0,  "ListenerCount": 0,  "Rating": 54,  "TotalItems": 0,  "DisplayInfo": [    "Fardration",    "LRKR - Getting Along"  ],  "ContextMenuItems": [    {      "Type": 1,      "Title": "Go to Album",      "Value": "spotify4:album:846df04b-59c4-4289-83fc-28267ed59cfb:6Wva0hYPq5hNnMAzM5jG15"    },    {      "Type": 1,      "Title": "Go to Artist",      "Value": "spotify4:artist:846df04b-59c4-4289-83fc-28267ed59cfb:0yTK74zLEsMyrdVPjw3Zqi"    },    {      "Type": 1,      "Title": "More Like This",      "Value": "spotify4:rec:846df04b-59c4-4289-83fc-28267ed59cfb:6f2O2P8QWAui2uAkXwWojT"    }  ]},"CurrProgress": 207,"SlideShowAvailable": false,"SearchPromptText": "Search for music by artist, album, or title","SourceLockedByZoneID": -1,"ServiceName": "Spotify","ServiceLogoURI": "4615DECB-5CBC-44CA-AE4C-12E002F3385E"}`,
    status: http.StatusOK,
  }
}

func (spah *sourcesPlayerAction) handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(spah.status)
  w.Write([]byte(spah.response))
}

// casatunes.NowPlaying should return object casatunes.RESTNowPlayingMediaItem
// from casaplayerEndpoint
func TestSourcesPlayerAction(t *testing.T) {
  once.Do(startServer)
  action := "play"
  casaplayerEndpoint := "http://" + serverAddr + "/"
  casaSource := "0"
  spa := newSourcesPlayerAction()

  http.HandleFunc("/sources/" + casaSource + "/player/" + "play", func(w http.ResponseWriter, r *http.Request) { spa.handler(w, r) })

  t.Run("play music", func(t *testing.T) {
    client := New(casaplayerEndpoint)
    err := client.SourcesPlayerAction(casaSource, action)

    if err != nil {
      t.Fatal(fmt.Sprintf("Could not play music! error: [%s]", err))
    } else {
      t.Log("Successfully played music!")
    }

  })

  t.Run("Nonexistant endpoint", func(t *testing.T) {
    client := New("http://unresolvable:456")
    err := client.SourcesPlayerAction(casaSource, action)

    if err != nil {
      t.Log("Received error on nonexistant endpoint")
    } else {
      t.Fatal("No error received on nonexistant endpoint !")
    }
  })

  t.Run("Invalid URI", func(t *testing.T) {
    client := New("GarbageText")
    err := client.SourcesPlayerAction(casaSource, action)

    if err != nil {
      t.Log("Invalid URI rejected.")
    } else {
      t.Fatal("NowPlaying accepted casatunes.Client with bad URI !")
    }
  })

  t.Run("HTTP 503 Response Code", func(t *testing.T) {
    client := New(casaplayerEndpoint)
    spa.status = http.StatusServiceUnavailable

    err := client.SourcesPlayerAction(casaSource, action)

    if err != nil {
      t.Log("[503] Error received.")
    } else {
      t.Fatal(err)
    }
  })

  t.Run("Return error on unsupported player action 'ignite'", func(t *testing.T) {
    client := New(casaplayerEndpoint)

    err := client.SourcesPlayerAction(casaSource, "ignite")

    if err != nil {
      t.Log(fmt.Sprintf("Error received! [%s]", err))
    } else {
      t.Fatal("No error received!")
    }
  })
}
