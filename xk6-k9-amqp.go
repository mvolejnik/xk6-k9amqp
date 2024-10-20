package k9amqp

import (
	"github.com/mvolejnik/xk6-k9-amqp/amqp"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/k9amqp", new(amqp.K9Amqp))
}
