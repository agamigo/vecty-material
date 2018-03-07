// dialog implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// D is a material dialog component.
type D struct {
	mdc    *base.Component
	IsOpen bool `js:"open"`
}

// New returns a new component.
func New() *D {
	c := &D{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *D) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{
		"checked": c.IsOpen,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *D) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *D) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCDialog",
				MDCCamelCaseName: "dialog",
			},
		}
	}
	return c.mdc.Component()
}

// Open shows the dialog. If the dialog is already open then Open is a no-op.
func (c *D) Open() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("show")
	return err
}

// Close removes the dialog from view. If the dialog is already closed then
// Close is a no-op.
func (c *D) Close() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("close")
	return err
}
