package chooser

import (
	"log"

	nc "github.com/rthornton128/goncurses"

	"github.com/cmatsuoka/chooser/mptmenu"
)

type Chooser struct {
	scr *nc.Window
}

func (c *Chooser) Init() error {
	var err error
	c.scr, err = nc.Init()
	if err != nil {
		return err
	}

	nc.Raw(true)
	nc.Echo(false)
	nc.Cursor(0)
	c.scr.Timeout(0)
	c.scr.Clear()
	c.scr.Keypad(true)

	return nil
}

func (c *Chooser) Deinit() {
	nc.End()
}

type Handler func(*Chooser) error

type MenuOption struct {
	Text    string
	Handler Handler
}

type Menu struct {
	mptmenu.MptMenu
	handlers []Handler
	c        *Chooser
}

func NewMenu(c *Chooser, title, desc, prompt string, options []MenuOption, topOption bool) *Menu {
	items := make([]string, len(options))
	handlers := make([]Handler, len(options))
	for i, option := range options {
		items[i] = option.Text
		handlers[i] = option.Handler
	}

	return &Menu{
		MptMenu:  mptmenu.New(title, desc, prompt, items, topOption),
		handlers: handlers,
		c:        c,
	}
}

func (m *Menu) Choose() {
	for {
		m.c.scr.Clear()
		num := 0
		for {
			m.Show(m.c.scr, 0, 0)

			ch := m.c.scr.GetChar()
			key := nc.KeyString(ch)

			var ok bool
			num, ok = m.CheckKey(key)
			if ok {
				break
			}
		}

		handler := m.handlers[num]
		if handler == nil {
			// return to previous menu
			break
		}

		// handle this option
		if err := handler(m.c); err != nil {
			m.c.Deinit()
			log.Fatalf("internal error: %v", err)
		}
	}
}
