package publisher

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

func Publish(w io.Writer, projectId, topicID, msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
