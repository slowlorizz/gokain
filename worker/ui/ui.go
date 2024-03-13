package ui

// tui "github.com/gizak/termui/v3"
import (
	"github.com/gizak/termui/v3/widgets"
)

const (
	X_SPACER int = 1
	Y_SPACER int = 1
)

type (
	Component interface {
		Render()
	}

	TextComponent struct {
		widget *widgets.Paragraph
	}

	ThreadComponent struct {
		HashPH *widgets.Paragraph
		TxtPH  *widgets.Paragraph
	}

	Handler struct {
		Components []Component
	}
)

func New_ThreadComponent() *ThreadComponent {

}
