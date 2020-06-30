package manager

type managerClient struct {
	eventChan chan string
}

func newManagerClient() *managerClient {
	return &managerClient{
		eventChan: make(chan string),
	}
}

func (c *managerClient) Send(v string) {
	c.eventChan <- v
}
