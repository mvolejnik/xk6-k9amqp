package k9amqp

import (
	"fmt"
	"go.k6.io/k6/js/modules"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpConnection struct {
	Server string
	Port int
	Username string
	Password string
}

type K9amqp struct {
	Pool map[AmqpConnection]*amqp.Connection
}

func init() {
	modules.Register("k6/x/k9amqp", new(K9amqp))
}

func (k9amqp *K9amqp)  InitPool() {
	k9amqp.Pool = make(map[AmqpConnection]*amqp.Connection)
}

func (k9amqp *K9amqp) Connect(server string, port int, vhost string, username string, password string) string {
	/*config := amqp.Config{
		Vhost:      vhost,
		Properties: amqp.NewConnectionProperties(),
	}*/
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", username, password, server, port))
	println(err)
	k9amqp.Pool[AmqpConnection{Server: server, Port: port, Username: username, Password: password}] = conn
	channel, err := conn.Channel()
	println(err)
	println(channel)
	/*err = channel.Publish(
		"amq.topic",
		"test",
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte("{}"),
			DeliveryMode:    amqp.Transient,
			Priority: 0,
		},
	)*/
	return "test"
}

func (k9amqp *K9amqp) Test() string {
	return "test"
}
