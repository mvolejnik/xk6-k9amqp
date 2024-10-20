package main

import (
	"fmt"

	"github.com/mvolejnik/xk6-k9-amqp/amqp"
)

var k9amqp = new(amqp.K9Amqp)

func main() {
	fmt.Println(k9amqp)
	k9amqp.Test()
}
