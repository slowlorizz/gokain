package ui

// tui "github.com/gizak/termui/v3"
import (
	"fmt"
	"log"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	X_SPACER    int = 1
	Y_SPACER    int = 1
	LINE_HEIGHT int = 3
)

type (
	Component interface {
		Init(int, int)
		Render()
	}

	TextComponent struct {
		Widget *widgets.Paragraph
		Width  int
		X      int
		Y      int
	}

	ThreadComponent struct {
		ID    int
		HashC TextComponent
		PtC   TextComponent
		X     int
		Y     int
	}

	Handler struct {
		Components []Component
		Events     <-chan tui.Event
		Y          int
		X          int
	}
)

var HANDLER Handler = Handler{X: X_SPACER, Y: 0}

func Clear() {
	HANDLER.Components = make([]Component, 0)
}

func Append(c Component) {
	c.Init(HANDLER.X, HANDLER.Y+Y_SPACER)
	HANDLER.Y += Y_SPACER + LINE_HEIGHT
	HANDLER.Components = append(HANDLER.Components, c)
}

func Render() {
	for _, v := range HANDLER.Components {
		v.Render()
	}
}

func Init() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer tui.Close()

	HANDLER.Components = make([]Component, 0)
	HANDLER.Events = tui.PollEvents()
}

func (H *Handler) HandleEvent(e *tui.Event) bool { // true if event specifies an exit
	switch e.ID {
	case "q", "<C-c>":
		return true
	}

	return false
}

// ---------------------------------------------------------------------------------------------------- //

func New_TextComponent(title string) *TextComponent {
	tc := &TextComponent{Widget: widgets.NewParagraph(), Width: 2}

	tc.Widget.Title = title

	return tc
}

func (tc *TextComponent) Init(txt string, x int, y int) {
	tc.X = x
	tc.Y = y
	tc.SetText(txt)
	tc.SetTextColor(tui.ColorWhite)
	tc.SetBorderColor(tui.ColorWhite)
}

func (tc *TextComponent) SetText(txt string) {
	if len(txt)+2 != tc.Width {
		tc.Width = len(txt) + 2                                        // add 2 bc. of border
		tc.Widget.SetRect(tc.X, tc.Y, tc.X+tc.Width, tc.Y+LINE_HEIGHT) // line height is always 3 --> 2x Border 1x Text
	}

	tc.Widget.Text = txt
}

func (tc *TextComponent) SetBorderColor(color tui.Color) {
	tc.Widget.BorderStyle.Fg = color
}

func (tc *TextComponent) SetTextColor(color tui.Color) {
	tc.Widget.TextStyle.Fg = color
}

func (tc *TextComponent) Render() {
	tui.Render(tc.Widget)
}

func New_ThreadComponent(id int) *ThreadComponent {
	thc := &ThreadComponent{
		ID:    id,
		HashC: *New_TextComponent(fmt.Sprintf("[Hash - %d]", id)),
		PtC:   *New_TextComponent(fmt.Sprintf("[Text - %d]", id)),
	}

	Append(thc)

	return thc
}

func (thc *ThreadComponent) Init(x int, y int) {
	thc.X = x
	thc.Y = y

	thc.HashC.Init("", thc.X, thc.Y)
	thc.PtC.Init("", thc.X+thc.HashC.Width+X_SPACER, thc.Y)
}

func (thc *ThreadComponent) Render() {
	thc.HashC.Render()
	thc.PtC.Render()
}

func (thc *ThreadComponent) SetStyleFound() {
	thc.HashC.SetBorderColor(tui.ColorGreen)
	thc.PtC.SetBorderColor(tui.ColorGreen)
}
