package api

type Client interface {
	Call(model string, method string, props any, res any) error
}

type Model interface {
	Call(method string, props any, res any) error
}

type model struct {
	client Client
	name   string
}

func NewModel(client Client, name string) Model {
	return &model{client: client, name: name}
}

func (m *model) Call(method string, props any, res any) error {
	return m.client.Call(m.name, method, props, res)
}
