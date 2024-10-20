package main

import (
	"fmt"

	"github.com/mvolejnik/xk6-k9-amqp/amqp"
)

var k9amqp = new(amqp.K9Amqp)

func main() {
	fmt.Println("Hello, World!")
	k9amqp.Test()
}
