// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// CB is a material checkbox component.
type CB struct {
	mdc           *base.Component
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// New returns a new component.
func New() *CB {
	c := &CB{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *CB) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *CB) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *CB) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCCheckbox",
				MDCCamelCaseName: "checkbox",
			},
		}
		fallthrough
	case c.mdc.Object == nil:
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *CB) StateMap() base.StateMap {
	return base.StateMap{
		"checked":       c.Checked,
		"indeterminate": c.Indeterminate,
		"disabled":      c.Disabled,
		"value":         c.Value,
	}
}
