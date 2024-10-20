package amqp

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/amqp", new(K9Amqp))
}

type K9Amqp struct{}

func (k9Amqp *K9Amqp) Test() string {
	println("K9 AMQP")
	return "test"
}
