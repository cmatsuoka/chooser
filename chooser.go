package chooser

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"

	"github.com/cmatsuoka/chooser/mptmenu"
)

type Chooser struct {
}

func (c *Chooser) Init() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.OutputNormal)
	return nil
}

func (c *Chooser) Deinit() {
	termbox.Close()
}

func (c *Chooser) Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (c *Chooser) GetKey() string {
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				return "esc"
			case termbox.KeyArrowUp:
				return "up"
			case termbox.KeyArrowDown:
				return "down"
			case termbox.KeyEnter:
				return "enter"
			default:
				if ev.Key > '0' && ev.Key < 'z' {
					return fmt.Sprintf("%c", ev.Key)
				}
			}
		}
	}
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
		num := 0
		for {
			m.c.Clear()
			m.Show()
			key := m.c.GetKey()
			if num = m.CheckKey(key); num >= 0 {
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
