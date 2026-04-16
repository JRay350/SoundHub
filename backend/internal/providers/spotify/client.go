package spotify

type Client struct {
	// later: add token, http client, etc.
}

func NewClient() *Client {
	return &Client{}
}