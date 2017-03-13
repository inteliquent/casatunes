package casatunes

type Client struct {
  config struct {
    endpoint string
  }
}

// Returns new Casatunes API client *Client
func New(endpoint string) *Client {
  s := &Client{}
  s.config.endpoint = endpoint
  return s
}
