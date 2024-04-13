package gorest

// Client any struct that implements the HttpClient interface is considered an httpClient , when creating a client you
// need to create an instance of the struct and not the interface
type Client struct{}

func New() HttpClient {
	client := &Client{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func (c *Client) Get() {

}

func (c *Client) Post() {

}

func (c *Client) Put() {

}

func (c *Client) Patch() {

}

func (c *Client) Delete() {

}
