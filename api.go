package pubnub_communicator

import (
	"github.com/slobdell/observerPattern"
)

type PubnubCommunicator struct {
	observerPattern.ConcreteObservable
	channelId      string
	publishChannel chan string
}

func NewPubnubCommunicator(channelId string) *PubnubCommunicator {
	pubnubCommunicator := &PubnubCommunicator{
		ConcreteObservable: *observerPattern.NewConcreteObservable(),
		channelId:          channelId,
		publishChannel:     make(chan string),
	}
	go pubnubCommunicator.listenForReads()
	go publishOnEvent(
		pubnubCommunicator.channelId,
		pubnubCommunicator.publishChannel,
	)
	return pubnubCommunicator
}

func (p *PubnubCommunicator) listenForReads() {
	incomingMessages := make(chan string)
	go extractMessagesFromChannel(p.channelId, incomingMessages)
	for {
		payload := <-incomingMessages
		p.NotifyObservers(payload)
	}
}

func (p *PubnubCommunicator) SendMessage(message string) {
	p.publishChannel <- message
}
