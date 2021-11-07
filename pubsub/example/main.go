package main

import (
	"context"
	"fmt"
	"food_delivery_be/pubsub"
	"food_delivery_be/pubsub/pblocal"
	"time"
)

func main() {
	localPb := pblocal.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, _ := localPb.Subscribe(context.Background(), topic)
	sub2, _ := localPb.Subscribe(context.Background(), topic)

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		//defer common.AppRecovery()
		for {
			fmt.Println("Con1: ", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		//defer common.AppRecovery()
		for {
			fmt.Println("Con2: ", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	//time.Sleep(time.Second * 3)
	//close1()
	//close2()

	//localPb.Public(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 2)
}
