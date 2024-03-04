package logs

import (
	"errors"
	"fmt"
)

func Debug(log_msg string, data ...any) {

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

	msg, err := Msg(log_msg, &dm)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler.Channel.Debug <- msg
}

func Info(log_msg string, data ...any) {
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

	msg, err := Msg(log_msg, &dm)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler.Channel.Info <- msg
}

func Warn(log_msg string, data ...any) {
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

	msg, err := Msg(log_msg, &dm)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler.Channel.Warn <- msg
}

func Error(log_msg string, data ...any) {
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

	msg, err := Msg(log_msg, &dm)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler.Channel.Error <- msg
}

func Fatal(log_msg string, data ...any) {
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

	msg, err := Msg(log_msg, &dm)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler.Channel.Fatal <- msg
}
