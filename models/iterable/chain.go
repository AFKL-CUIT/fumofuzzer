package iterable

import (
	"afkl/fumofuzzer/models/payload"
)

type ChainIterator struct {
	isEnd bool
	token []string

	channel chan []string
}

func NewChainIterator() *ChainIterator {
	return &ChainIterator{
		isEnd:   false,
		channel: make(chan []string),
	}
}

func (iter *ChainIterator) Exec(Payloads []payload.Payload) {
	go func() {
		defer close(iter.channel)
		for _, payload := range Payloads {
			for _, data := range payload.Value {
				iter.channel <- []string{data}
			}
		}
	}()
}

func (iter *ChainIterator) IsEnd() bool {
	return iter.isEnd
}

func (iter *ChainIterator) Scan() bool {
	if iter.IsEnd() {
		return false
	}

	if data, ok := <-iter.channel; ok {
		iter.token = data
		return true
	} else {
		iter.isEnd = true
		return false
	}
}

func (iter *ChainIterator) Value() []string {
	return iter.token
}