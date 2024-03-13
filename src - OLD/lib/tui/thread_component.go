package tui

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ThreadComponent struct {
	Id        int
	HashPH    *widgets.Paragraph
	TxtPH     *widgets.Paragraph
	SeedPH    *widgets.Paragraph
	YPos      int
	XPos      int
	Height    int
	HashWidth int
	TxtWidth  int
	SeedWidth int
}

func New_ThreadComponent(id int, seed string, x int, y int) *ThreadComponent {
	component := ThreadComponent{Id: id, HashPH: widgets.NewParagraph(), TxtPH: widgets.NewParagraph(), SeedPH: widgets.NewParagraph(), XPos: x, YPos: y, HashWidth: 70, Height: 3, TxtWidth: 15, SeedWidth: 30}

	xSpace := 2

	component.SeedPH.Title = fmt.Sprintf("[%d] Seed", component.Id)
	component.SeedPH.Text = seed
	component.SeedPH.SetRect(component.XPos, component.YPos, component.XPos+component.SeedWidth, component.YPos+component.Height)
	component.SeedPH.TextStyle.Fg = ui.ColorWhite
	component.SeedPH.BorderStyle.Fg = ui.ColorWhite

	component.HashPH.Text = " "
	component.HashPH.SetRect(component.XPos+component.SeedWidth+xSpace, component.YPos, component.XPos+component.SeedWidth+xSpace+component.HashWidth, component.YPos+component.Height)
	component.HashPH.TextStyle.Fg = ui.ColorWhite
	component.HashPH.BorderStyle.Fg = ui.ColorWhite

	component.TxtPH.Text = " "
	component.TxtPH.SetRect(component.XPos+component.SeedWidth+xSpace+component.HashWidth+xSpace, component.YPos, component.XPos+component.SeedWidth+xSpace+component.HashWidth+xSpace+component.TxtWidth, component.YPos+component.Height)
	component.TxtPH.TextStyle.Fg = ui.ColorWhite
	component.TxtPH.BorderStyle.Fg = ui.ColorWhite

	return &component
}

func (component *ThreadComponent) Render() {
	ui.Render(component.SeedPH, component.HashPH, component.TxtPH)
}

func (component *ThreadComponent) SetHashText(txt string) {
	component.HashPH.Text = txt
}

func (component *ThreadComponent) SetPlainText(txt string) {
	component.TxtPH.Text = txt
}

func (component *ThreadComponent) SetBordersGreen() {
	component.SeedPH.BorderStyle.Fg = ui.ColorGreen
	component.HashPH.BorderStyle.Fg = ui.ColorGreen
	component.TxtPH.BorderStyle.Fg = ui.ColorGreen
}
