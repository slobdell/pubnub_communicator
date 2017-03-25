package pubnub_communicator

import (
	"fmt"
	"github.com/pubnub/go/messaging"
)

func publishOnEvent(pubnubChannelId string, messagesToPublish chan string) {
	pubnub := messaging.NewPubnub(
		PUBLISH_KEY,
		SUBSCRIBE_KEY,
		SECRET_KEY,
		CIPHER_KEY,
		USE_SSL,
		"",  // custom UUID
		nil, // optional logger
	)
	successChannel := make(chan []byte)
	errorChannel := make(chan []byte)
	go handlePublishCallbacks(successChannel, errorChannel)

	var payload string
	for {
		payload = <-messagesToPublish
		pubnub.Publish(
			pubnubChannelId,
			payload,
			successChannel,
			errorChannel,
		)
	}
}

func handlePublishCallbacks(successChannel, errorChannel chan []byte) {
	select {
	case <-successChannel:
		//fmt.Println(string(response))
	case err := <-errorChannel:
		fmt.Println(string(err))
	case <-messaging.Timeout():
		fmt.Println("Publish() timeout")
	}
}
