package amqpClient

type Feature interface {
	Shutdown() error
	SetDisconnectChannel(ch chan error) *Consumer
}


type Parameters struct {
	ExchangeType string
	RoutingKey string
	PrefetchCount int
}

type Credentials struct {
	User string
	Password string
	Host string
	Port string
}