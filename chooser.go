package chooser

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"

	"github.com/cmatsuoka/chooser/mptmenu"
)

func Init() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("cannot initialize the terminal: %v", err)
	}
	termbox.SetOutputMode(termbox.OutputNormal)
	return nil
}

func Deinit() {
	termbox.Close()
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func getKey() int {
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == 0 {
				return int(ev.Ch)
			} else {
				return int(ev.Key)
			}
		}
	}
}

type Handler func() error

type MenuOption struct {
	Text    string
	Handler Handler
}

type Menu struct {
	mptmenu.MptMenu
	handlers []Handler
}

func NewMenu(title, desc, prompt string, options []MenuOption, topOption bool) *Menu {
	items := make([]string, len(options))
	handlers := make([]Handler, len(options))
	for i, option := range options {
		items[i] = option.Text
		handlers[i] = option.Handler
	}

	return &Menu{
		MptMenu:  mptmenu.New(title, desc, prompt, items, topOption),
		handlers: handlers,
	}
}

func (m *Menu) Choose() {
	for {
		num := 0
		for {
			clear()
			m.Show()
			key := getKey()
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
		if err := handler(); err != nil {
			termbox.Close()
			log.Fatalf("internal error: %v", err)
		}
	}
}
