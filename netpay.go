package netpay

// Client type.
type Client struct {
}

// New returns a new client.
func New(apiKey string) (*Client, error) {
	return &Client{}, nil
}
