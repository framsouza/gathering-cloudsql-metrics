package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func List(client *pubsub.Client) ([]*pubsub.Topic, error) {
	ctx := context.Background()
	var topics []*pubsub.Topic
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
func Create(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	t, err := client.CreateTopic(ctx, topic)
	if err != nil {
		return err
	}

	fmt.Printf("Topic created: %v\n", t)
	return nil
}

func Publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}
