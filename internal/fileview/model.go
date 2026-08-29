package fileview

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
)

type Model struct {
	index  int
	fields []inputField
}

type inputField struct {
	kind              int             // kind == 0 (Single line), kind == 1 (Multi Line)
	singleLnContainer textinput.Model // Active if kind == 0
	multiLnContainer  textarea.Model  // Active if kind == 1
}

func (m Model) New() Model {

}
