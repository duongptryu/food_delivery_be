package subscriber

import (
	"context"
	"fmt"
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/asyncjob"
	"food_delivery_be/pubsub"
	"food_delivery_be/skio"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx   component.AppContext
	rtEngine skio.RealtimeEngine
}

func NewEngine(appCtx component.AppContext, realtimeEngine skio.RealtimeEngine) *consumerEngine {
	return &consumerEngine{
		appCtx:   appCtx,
		rtEngine: realtimeEngine,
	}
}

func (e *consumerEngine) Start() error {
	e.startSubTopic(common.TopicUserLikeRestaurant, true, RunIncreaseLikeCountAfterUserLikeRestaurant(e.appCtx), EmitRealtimeAfterUserLikeRestaurant(e.appCtx, e.rtEngine))

	e.startSubTopic(common.TopicUserDislikeRestaurant, true, RunDecreaseLikeCountAfterUserUnlikeRestaurant(e.appCtx, e.rtEngine))
	return nil
}

//type GroupJob interface {
//	Run(ctx context.Context) error
//}

func (e *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...consumerJob) error {
	c, _ := e.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range jobs {
		fmt.Println("Setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			fmt.Println("Running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := c

			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&jobs[i], <-msg))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				fmt.Println(err)
			}
		}
	}()

	return nil
}
