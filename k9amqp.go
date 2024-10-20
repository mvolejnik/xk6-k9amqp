package k9amqp

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/k9amqp", new(K9amqp))
}

type K9amqp struct{}

func (k9amqp *K9amqp) Test() string {
	println("K9 AMQP")
	return "test"
}
