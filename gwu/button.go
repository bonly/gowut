// Copyright (C) 2013 Andras Belicza. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Button component interface and implementation.

package gwu

// Button interface defines a clickable button.
//
// Suggested event type to handle actions: ETypeClick
//
// Default style class: "gwu-Button"
type Button interface {
	// Button is a component.
	Comp

	// Button has text.
	HasText

	// Button can be enabled/disabled.
	HasEnabled
}

// Button implementation.
type buttonImpl struct {
	compImpl       // Component implementation
	hasTextImpl    // Has text implementation
	hasEnabledImpl // Has enabled implementation
}

// NewButton creates a new Button.
func NewButton(text string) Button {
	c := newButtonImpl(nil, text)
	c.Style().AddClass("gwu-Button")
	return &c
}

// newButtonImpl creates a new buttonImpl.
func newButtonImpl(valueProviderJs []byte, text string) buttonImpl {
	return buttonImpl{newCompImpl(valueProviderJs), newHasTextImpl(text), newHasEnabledImpl()}
}

var (
	strButtonOp = []byte(`<button type="button"`) // `<button type="button"`
	strButtonCl = []byte("</button>")             // "</button>"
)

func (c *buttonImpl) Render(w Writer) {
	w.Write(strButtonOp)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	c.renderEnabled(w)
	w.Write(strGT)

	c.renderText(w)

	w.Write(strButtonCl)
}
