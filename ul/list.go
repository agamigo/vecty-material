package ul

import (
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// L is a vecty-material list component.
type L struct {
	*base.Base
	*State
}

type State struct {
	Items          []vecty.ComponentOrHTML
	Dense          bool
	Avatar         bool
	NonInteractive bool
	ClickHandler   func(thisL *L, thisI *Item, e *vecty.Event)
	GroupSubheader string
}

// Item is a vecty-material list-item component.
type Item struct {
	*base.Base
	*ItemState
}

type ItemState struct {
	Primary      vecty.ComponentOrHTML
	Secondary    vecty.ComponentOrHTML
	Graphic      vecty.ComponentOrHTML
	Meta         vecty.ComponentOrHTML
	Selected     bool
	Activated    bool
	ClickHandler func(i *Item, e *vecty.Event)
	Href         string
}

// Group is a vecty-material list-group component.
type Group struct {
	*base.Base
	*GroupState
}

type GroupState struct {
	Lists []vecty.ComponentOrHTML
}

func New(p *base.Props, s *State) *L {
	c := &L{}
	if s == nil {
		s = &State{}
	}
	c.State = s
	c.Base = base.New(p, nil)
	return c
}

func NewItem(p *base.Props, s *ItemState) *Item {
	c := &Item{}
	if s == nil {
		s = &ItemState{}
	}
	c.ItemState = s
	c.Base = base.New(p, nil)
	return c
}

func NewGroup(p *base.Props, s *GroupState) *Group {
	c := &Group{}
	if s == nil {
		s = &GroupState{}
	}
	c.GroupState = s
	c.Base = base.New(p, nil)
	return c
}

// Render implements the vecty.Component interface.
func (c *L) Render() vecty.ComponentOrHTML {
	twoLine := false
	for _, li := range c.Items {
		lItem, ok := li.(*Item)
		if ok && lItem.Secondary != nil {
			twoLine = true
		}
	}
	return c.Base.Render(elem.UnorderedList(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-list"),
			vecty.MarkupIf(twoLine,
				vecty.Class("mdc-list--two-line")),
			vecty.MarkupIf(c.Dense,
				vecty.Class("mdc-list--dense")),
			vecty.MarkupIf(c.Avatar,
				vecty.Class("mdc-list--avatar-list")),
			vecty.MarkupIf(c.NonInteractive,
				vecty.Class("mdc-list--non-interactive")),
		),
		c.itemList(),
	))
}

// Copy implements the vecty.Copier interface
func (c *L) Copy() vecty.Component {
	newL := New(
		&base.Props{
			ID:     c.ID,
			Ripple: c.Ripple,
			Markup: make([]vecty.Applyer, len(c.Markup)),
		},
		&State{
			Avatar:         c.Avatar,
			Dense:          c.Dense,
			GroupSubheader: c.GroupSubheader,
			NonInteractive: c.NonInteractive,
			ClickHandler:   c.ClickHandler,
			Items:          make([]vecty.ComponentOrHTML, len(c.Items)),
		},
	)
	copy(newL.Markup, c.Markup)
	return newL
}

// Render implements the vecty.Component interface.
func (c *Item) Render() vecty.ComponentOrHTML {
	tag := "li"
	if c.Href != "" {
		tag = "a"
	}
	graphic := setupGraphicOrMeta(c.Graphic)
	if graphic != nil {
		if g, ok := graphic.(*vecty.HTML); ok {
			vecty.Class("mdc-list-item__graphic").Apply(g)
			vecty.Attribute("role", "presentation").Apply(g)
		}
	}
	meta := setupGraphicOrMeta(c.Meta)
	if meta != nil {
		if g, ok := meta.(*vecty.HTML); ok {
			vecty.Class("mdc-list-item__meta").Apply(g)
			vecty.Attribute("role", "presentation").Apply(g)
		}
	}
	return c.Base.Render(vecty.Tag(tag,
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-list-item"),
			vecty.MarkupIf(c.Selected,
				vecty.Class("mdc-list-item--selected")),
			vecty.MarkupIf(c.Activated,
				vecty.Class("mdc-list-item--activated")),
			vecty.MarkupIf(c.ClickHandler != nil,
				event.Click(c.wrapClickHandler()),
			),
			vecty.MarkupIf(c.Href != "", prop.Href(c.Href)),
		),
		graphic,
		vecty.If(c.Primary != nil && c.Secondary == nil, c.Primary),
		vecty.If(c.Secondary != nil,
			elem.Span(vecty.Markup(vecty.Class("mdc-list-item__text")),
				vecty.If(c.Primary != nil, c.Primary),
				elem.Span(vecty.Markup(
					vecty.Class("mdc-list-item__secondary-text")),
					c.Secondary,
				)),
		),
		meta,
	))
}

// Copy implements the vecty.Copier interface
func (c *Item) Copy() vecty.Component {
	newI := NewItem(
		&base.Props{
			ID:     c.ID,
			Ripple: c.Ripple,
			Markup: make([]vecty.Applyer, len(c.Markup)),
		},
		&ItemState{
			Primary:      c.Primary,
			Secondary:    c.Secondary,
			Graphic:      c.Graphic,
			Meta:         c.Meta,
			Selected:     c.Selected,
			Activated:    c.Activated,
			Href:         c.Href,
			ClickHandler: c.ClickHandler,
		},
	)
	copy(newI.Markup, c.Markup)
	return newI
}

// Render implements the vecty.Component interface.
func (c *Group) Render() vecty.ComponentOrHTML {
	return c.Base.Render(elem.Div(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-list-group"),
		),
		c.listList(),
	))
}

func ListDivider() *vecty.HTML {
	return elem.HorizontalRule(
		vecty.Markup(vecty.Class("mdc-list-divider")),
	)
}

func ListDividerInset() *vecty.HTML {
	d := ListDivider()
	vecty.Class("mdc-list-divider--inset").Apply(d)
	return d
}

func ItemDivider() *vecty.HTML {
	return elem.ListItem(vecty.Markup(
		vecty.Class("mdc-list-divider"),
		vecty.Attribute("role", "separator"),
	))
}

func ItemDividerInset() *vecty.HTML {
	d := ItemDivider()
	vecty.Class("mdc-list-divider--inset").Apply(d)
	return d
}

func (c *L) itemList() vecty.List {
	items := make(vecty.List, len(c.Items))
	for i, item := range c.Items {
		if lItem, ok := item.(*Item); ok {
			switch {
			case lItem.ClickHandler != nil:
			case c.ClickHandler != nil:
				lItem.ClickHandler = c.wrapClickHandler()
			}
		}
		items[i] = item
	}
	return items
}

func (c *Group) listList() vecty.List {
	lists := make(vecty.List, len(c.Lists)*2)
	for _, cList := range c.Lists {
		if list, ok := cList.(*L); ok {
			if list.GroupSubheader != "" {
				lists = append(lists,
					elem.Heading3(vecty.Markup(
						vecty.Class("mdc-list-group__subheader")),
						vecty.Text(list.GroupSubheader)),
				)
			}
			lists = append(lists, list)
		}
	}
	return lists
}

func (c *L) wrapClickHandler() func(i *Item, e *vecty.Event) {
	return func(i *Item, e *vecty.Event) {
		c.ClickHandler(c, i, e)
	}
}

func (c *Item) wrapClickHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.ClickHandler(c, e)
	}
}

func setupGraphicOrMeta(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	var graphic vecty.ComponentOrHTML
	if c != nil {
		graphic = c
		switch t := c.(type) {
		case vecty.Component:
			if h, ok := t.Render().(*vecty.HTML); ok {
				graphic = h
			}
		}
		if js.InternalObject(graphic).Get("tag").String() != "img" {
			graphic = elem.Span(graphic)
		}
	}
	return graphic
}
