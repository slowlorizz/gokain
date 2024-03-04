package logs

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type RGB struct {
	R uint8
	G uint8
	B uint8
}

//----------------------------------------------------------------------------------------------//

type handler struct {
	Lvl       slog.Level
	Channel   Channels
	CtrlCh    (chan string)
	TickSpeed time.Duration // Milliseconds to sleep if no message is present in any channel
	colors    []RGB
	tags      []string
}

func newHandler(tickspeed time.Duration) *handler {
	hnd := handler{
		Lvl:       slog.LevelInfo,
		Channel:   Channels{},
		CtrlCh:    make(chan string),
		TickSpeed: tickspeed,
		colors:    []RGB{RGB{255, 115, 0}, RGB{132, 222, 2}, RGB{255, 200, 0}, RGB{255, 40, 30}, RGB{255, 0, 100}},
		tags:      []string{"DBG", "INF", "WRN", "ERR", "FAT"},
	}

	hnd.Channel.init()

	return &hnd
}

func (h *handler) Start(verbose bool) {
	if verbose {
		h.Lvl = slog.LevelDebug
	}

	go h.handle()
}

func (h *handler) writeMsg(i int, msg *Message) {
	out_str := time.Now().Format("15:04:05")
	out_str += " "

	out_str += fmt.Sprintf("%s[%s]\033[0m", "\033[38;2;"+fmt.Sprint(h.colors[i].R)+";"+fmt.Sprint(h.colors[i].G)+";"+fmt.Sprint(h.colors[i].B)+"m", h.tags[i])

	out_str += fmt.Sprintf("  %s", msg.Msg)

	if len(msg.Data) > 0 {
		val_pairs := make([]string, 0, len(msg.Data)/2)

		for k := range msg.Data {
			val_pairs = append(val_pairs, strings.Join([]string{k, msg.Data[k]}, "="))
		}

		out_str += "\t"
		out_str += strings.Join(val_pairs, ", ")
	}

	fmt.Println(out_str)
}

func (h *handler) handle() {
parent_loop:
	for {
		select {
		case v := <-h.CtrlCh:
			if v == "#Join" {
				break parent_loop
			}

		case v := <-h.Channel.Debug:
			if h.Lvl == slog.LevelDebug {
				h.writeMsg(0, v)
			}

		case v := <-h.Channel.Info:
			h.writeMsg(1, v)

		case v := <-h.Channel.Warn:
			h.writeMsg(2, v)

		case v := <-h.Channel.Error:
			h.writeMsg(3, v)

		case v := <-h.Channel.Fatal:
			h.writeMsg(4, v)

		default:
			time.Sleep(h.TickSpeed * time.Millisecond)
		}
	}

	h.CtrlCh <- "#Joined"
}

func (h *handler) Join() {
	h.CtrlCh <- "#Join"

	for {
		v := <-h.CtrlCh

		if v == "#Joined" {
			break
		}
	}
}

var Handler = newHandler(10)
