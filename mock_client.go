package versions

// MockClient is a mock of Client
type MockClient struct {
	Client
	MockList func() ([]string, error)
}

// List is a mock of Client.List()
func (c *MockClient) List() ([]string, error) {
	return c.MockList()
}
