package ui

// tui "github.com/gizak/termui/v3"
import (
	"fmt"
	"log"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type (
	TextComponent struct {
		Widget *widgets.Paragraph
		Width  int
		X      int
		Y      int
	}

	ThreadComponent struct {
		ID      int
		HashC   *TextComponent
		PtC     *TextComponent
		widgets []*widgets.Paragraph
		X       int
		Y       int
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
)

func Render() {
	for _, v := range Components {
		fmt.Printf("Render [%d]\n", v.ID)
		v.Render()
	}
}

func Init() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer tui.Close()
}

// ---------------------------------------------------------------------------------------------------- //

func New_TextComponent(title string, txt string, x int, y int) (*TextComponent, *widgets.Paragraph) {
	tc := TextComponent{Width: 2, X: x, Y: y}

	tc.Widget = widgets.NewParagraph()
	tc.Widget.Title = title
	tc.Widget.SetRect(tc.X, tc.Y, tc.X+tc.Width, tc.Y+LINE_HEIGHT)

	tc.SetText(txt)
	tc.SetTextColor(tui.ColorWhite)
	tc.SetBorderColor(tui.ColorWhite)

	return &tc, tc.Widget
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
	fmt.Println("Render TextComponent")
	fmt.Printf("\n COMPONENT:\nTitle: \"%s\"\nText: \"%s\"\nWidth: %d\nX: %d\nY: %d\n", tc.Widget.Title, tc.Widget.Text, tc.Width, tc.X, tc.Y)
	tui.Render(tc.Widget)
	fmt.Println("Renderd")
}

func New_ThreadComponent(id int, x int, y int) *ThreadComponent {
	thc := &ThreadComponent{
		ID:      id,
		X:       x,
		Y:       y,
		widgets: make([]*widgets.Paragraph, 2),
	}

	thc.HashC, thc.widgets[0] = New_TextComponent(fmt.Sprintf("[Hash - %d]", id), " ", thc.X, thc.Y)
	thc.PtC, thc.widgets[1] = New_TextComponent(fmt.Sprintf("[Text - %d]", id), " ", thc.X+thc.HashC.Width+X_SPACER, thc.Y)

	return thc
}

func (thc *ThreadComponent) Render() {
	fmt.Println("Render Thread Component")
	tui.Render(thc.widgets[0], thc.widgets[1])
}

func (thc *ThreadComponent) SetStyleFound() {
	thc.HashC.SetBorderColor(tui.ColorGreen)
	thc.PtC.SetBorderColor(tui.ColorGreen)
}
