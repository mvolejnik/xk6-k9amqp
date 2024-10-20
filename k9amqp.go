package k9amqp

import (
	"go.k6.io/k6/js/modules"
)

type AmqpConnection struct {
	Url string
	Username string
	Password string
}

type K9amqp struct {
	Pool map[AmqpConnection]string
}

func init() {
	modules.Register("k6/x/k9amqp", new(K9amqp))
}

func (k9amqp *K9amqp)  InitPool() {
	k9amqp.Pool = make(map[AmqpConnection]string)
}

func (k9amqp *K9amqp) Connect(url string, username string, password string) string {
	println("K9amqp connect")
	k9amqp.Pool[AmqpConnection{Url: url, Username: username, Password: password}] = "test"
	return "test"
}

func (k9amqp *K9amqp) Test() string {
	println("K9 AMQP test()")
	return "test"
}
