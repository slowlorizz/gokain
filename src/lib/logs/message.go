package logs

type Message struct {
	Msg  string
	Data map[string]string
}

func Msg(msg string, data *map[string]string) (*Message, error) {
	var err error
	err = nil

	message := Message{Msg: msg, Data: *data}

	return &message, err
}
