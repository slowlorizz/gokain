package logs

import (
	"errors"
	"fmt"
)

func SendLog(ch MessageChannel, log_msg string, data []any) {
	dm := make(map[string]string)

	if len(data)%2 != 0 {
		fmt.Println(errors.New("Message Data has a key without value --> invalid pair").Error())
	}

	if len(data) > 0 {
		k := ""
		for i, d := range data {
			if i%2 == 0 {
				k = fmt.Sprint(d)
			} else {
				dm[k] = fmt.Sprint(d)
			}
		}
	}

	ch <- &Message{Msg: log_msg, Data: dm}
}

func Debug(log_msg string, data ...any) {
	SendLog(Handler.Channel.Debug, log_msg, data)
}

func Info(log_msg string, data ...any) {
	SendLog(Handler.Channel.Info, log_msg, data)
}

func Warn(log_msg string, data ...any) {
	SendLog(Handler.Channel.Warn, log_msg, data)
}

func Error(log_msg string, data ...any) {
	SendLog(Handler.Channel.Error, log_msg, data)
}

func Fatal(log_msg string, data ...any) {
	SendLog(Handler.Channel.Fatal, log_msg, data)
}
