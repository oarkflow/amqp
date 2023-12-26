package main

import (
	"context"
	"log"
	"syscall"
	"time"

	amqp "github.com/oarkflow/amqp/amqp091"

	grabbit "github.com/oarkflow/amqp"
)

func MsgHandler(props *grabbit.DeliveriesProperties, messages []grabbit.DeliveryData, mustAck bool, ch *grabbit.Channel) {
	for _, msg := range messages {
		log.Printf("  [%s][%s] got message: %s;\n", ch.Name(), props.ConsumerTag, string(msg.Body))
		ch.Ack(msg.DeliveryTag, false)
	}
}

func main() {
	ctxMaster, ctxCancel := context.WithCancel(context.TODO())
	conn := grabbit.NewConnection("amqp://guest:guest@localhost:5672", amqp.Config{}, grabbit.WithConnectionCtx(ctxMaster))
	topos := []*grabbit.TopologyOptions{
		{Name: "workload", IsDestination: true, Durable: true, Declare: true},
	}

	optConsumerOne := grabbit.DefaultConsumerOptions()
	optConsumerOne.WithName("consumer.one").WithQosGlobal(true).WithPrefetchTimeout(1 * time.Second)
	consumer := grabbit.NewConsumer(conn, optConsumerOne,
		grabbit.WithChannelName("chan.consumer.one"),
		grabbit.WithChannelTopology(topos),
		grabbit.OnChannelDown(func(name string, err grabbit.OptionalError) bool {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			log.Printf("callback: {%s} went down with {%s}", name, err)
			return true
		}),
		grabbit.OnChannelUp(func(name string) {
			log.Printf("callback: {%s} went up", name)
		}),
		grabbit.OnChannelRecovering(func(name string, retry int) bool {
			log.Printf("callback: connection establish {%s} retry count {%d}", name, retry)
			return true
		}),
		grabbit.WithChannelProcessor(MsgHandler),
	)
	consumer.Wait(ctxCancel)
}
