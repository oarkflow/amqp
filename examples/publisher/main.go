package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/oarkflow/amqp/amqp091"

	grabbit "github.com/oarkflow/amqp"
)

func OnPubReattempting(name string, retry int) bool {
	log.Printf("callback_redo: {%s} retry count {%d}", name, retry)
	return true // want continuing
}

// OnNotifyPublish CallbackNotifyPublish
func OnNotifyPublish(confirm amqp.Confirmation, ch *grabbit.Channel) {
	log.Printf("callback: publish confirmed status [%v] from queue [%s]\n", confirm.Ack, ch.Queue())
}

// OnNotifyReturn CallbackNotifyReturn
func OnNotifyReturn(_ amqp.Return, ch *grabbit.Channel) {
	log.Printf("callback: publish returned from queue [%s]\n", ch.Queue())
}

func PublishMsg(publisher *grabbit.Publisher, start, end int) {
	message := amqp.Publishing{}
	data := make([]byte, 0, 64)
	buff := bytes.NewBuffer(data)

	for i := start; i < end; i++ {
		<-time.After(1 * time.Second)
		buff.Reset()
		buff.WriteString(fmt.Sprintf("test number %04d", i))
		message.Body = buff.Bytes()
		log.Println("going to send:", buff.String())

		if err := publisher.Publish(message); err != nil {
			log.Println("publishing failed with: ", err)
		}
	}
}

func main() {
	ctxMaster, ctxCancel := context.WithCancel(context.TODO())

	conn := grabbit.NewConnection(
		"amqp://guest:guest@localhost:5672", amqp.Config{},
		grabbit.WithConnectionCtx(ctxMaster),
		grabbit.WithConnectionName("conn.main"),
	)

	pubOpt := grabbit.DefaultPublisherOptions()
	pubOpt.WithKey("workload").WithContext(ctxMaster).WithConfirmationsCount(20)

	topos := make([]*grabbit.TopologyOptions, 0, 8)
	topos = append(topos, &grabbit.TopologyOptions{
		Name:          "workload",
		IsDestination: true,
		Durable:       true,
		Declare:       true,
	})

	publisher := grabbit.NewPublisher(conn, pubOpt,
		grabbit.WithChannelCtx(ctxMaster),
		grabbit.WithChannelName("chan.publisher.example"),
		grabbit.WithChannelTopology(topos),
		grabbit.OnChannelRecovering(OnPubReattempting),
		grabbit.OnPublishSuccess(OnNotifyPublish),
		grabbit.OnPublishFailure(OnNotifyReturn),
	)

	if !publisher.AwaitAvailable(30*time.Second, 1*time.Second) {
		log.Println("publisher not ready yet")
		ctxCancel()
		return
	}

	PublishMsg(publisher, 0, 5)
}
