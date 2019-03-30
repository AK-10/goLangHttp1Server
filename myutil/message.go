package myutil

type Message struct {
	Handle string
	Value string
}

func NewMessage(handle string, value string) *Message {
	return &Message{
		Handle: handle,
		Value: value,
	}
}