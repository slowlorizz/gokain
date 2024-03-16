package ui

// tui "github.com/gizak/termui/v3"
import (
	"fmt"
	"log"
	"time"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/slowlorizz/gokain/worker/src/thread/combination"
)

type (
	TextComponent struct {
		Widget *widgets.Paragraph
		Width  int
		X      int
		Y      int
	}

	ThreadComponent struct {
		ID          int
		HashC       *widgets.Paragraph
		PtC         *widgets.Paragraph
		X           int
		Y           int
		HashC_width int
		PtC_width   int
	}

	RuntimeClock struct {
		StartTime time.Time
		X         int
		Y         int
		W         int
		Widget    *widgets.Paragraph
		Stop      bool
	}
)

const (
	X_SPACER    int = 1
	Y_SPACER    int = 1
	LINE_HEIGHT int = 3
)

var (
	Events     <-chan tui.Event
	Components []*ThreadComponent = make([]*ThreadComponent, 0)
	Clock      RuntimeClock       = RuntimeClock{}
)

func Render() {
	for _, v := range Components {
		v.Render()
	}

	Clock.Render()
}

func Init() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer tui.Close()

	Components = make([]*ThreadComponent, 0)
}

// --------------------------------------------------------------------------

func New_ThreadComponent(id int, hashType combination.HashType) *ThreadComponent {
	thc := &ThreadComponent{
		ID:          id,
		X:           X_SPACER,
		Y:           (LINE_HEIGHT + Y_SPACER) * (id - 1),
		HashC:       widgets.NewParagraph(),
		PtC:         widgets.NewParagraph(),
		HashC_width: 40,
		PtC_width:   15,
	}

	switch hashType {
	case combination.SHA1:
		thc.HashC_width = 40 + 2
	case combination.SHA256:
		thc.HashC_width = 64 + 2
	case combination.SHA512:
		thc.HashC_width = 128 + 2
	case combination.MD5:
		thc.HashC_width = 32 + 2
	}

	thc.HashC.Title = fmt.Sprintf("[Thread - %d]", id)
	thc.HashC.Text = ""
	thc.HashC.SetRect(thc.X, thc.Y, thc.X+thc.HashC_width, thc.Y+LINE_HEIGHT)
	thc.HashC.Border = true
	thc.HashC.BorderStyle.Fg = tui.ColorWhite
	thc.HashC.TextStyle.Fg = tui.ColorWhite

	thc.PtC.Text = ""
	thc.PtC.SetRect(thc.X+thc.HashC_width+X_SPACER, thc.Y, thc.X+thc.HashC_width+X_SPACER+thc.PtC_width, thc.Y+LINE_HEIGHT)
	thc.PtC.Border = true
	thc.PtC.BorderStyle.Fg = tui.ColorWhite
	thc.PtC.TextStyle.Fg = tui.ColorWhite

	return thc
}

func (thc *ThreadComponent) Render() {
	tui.Render(thc.HashC, thc.PtC)
}

func (thc *ThreadComponent) SetStyleFound() {
	thc.HashC.BorderStyle.Fg = tui.ColorGreen
	thc.PtC.BorderStyle.Fg = tui.ColorGreen
}

// ---------------------------------------------------------------------------------------------

func (rtc *RuntimeClock) Init() {
	Clock = RuntimeClock{
		X:         X_SPACER,
		Y:         (LINE_HEIGHT + Y_SPACER) * len(Components),
		W:         20,
		StartTime: time.Now(),
		Widget:    widgets.NewParagraph(),
		Stop:      false,
	}

	rtc.Widget.Title = "Duration"
	rtc.Widget.Text = ""
	rtc.Widget.SetRect(rtc.X, rtc.Y, rtc.X+rtc.W, rtc.Y+LINE_HEIGHT)
	rtc.Widget.Border = true
	rtc.Widget.BorderStyle.Fg = tui.ColorWhite
	rtc.Widget.TextStyle.Fg = tui.ColorWhite
}

func (rtc *RuntimeClock) Tick() {
	if !rtc.Stop {
		rtc.Widget.Text = (time.Since(rtc.StartTime) % time.Second).String()
	}
}

func (rtc *RuntimeClock) Render() {
	rtc.Tick()
	tui.Render(rtc.Widget)
}
