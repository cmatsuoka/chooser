package mptmenu

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type MptMenuOption struct {
	text string
	key  byte
}

type MptMenu struct {
	title       string
	description string
	prompt      string
	options     []MptMenuOption
	cursor      rune
	useHotkeys  bool
	hasTopOpt   bool
	optionsLine int
	current     int
}

func New(title, description, prompt string, opts []string, hasTopOpt bool) MptMenu {
	m := MptMenu{
		title:       title,
		description: description,
		prompt:      prompt,
		options:     make([]MptMenuOption, len(opts)),
		hasTopOpt:   hasTopOpt,
		cursor:      '>',
	}

	num := byte(0)
	if !m.hasTopOpt {
		num++
	}

	for i, opt := range opts {
		m.options[i] = MptMenuOption{text: opt, key: '0' + num}
		num++
	}

	return m
}

func printAt(y, x int, fg, bg termbox.Attribute, f string, parm ...interface{}) {
	maxX, maxY := termbox.Size()
	if x > maxX || y > maxY {
		return
	}
	s := fmt.Sprintf(f, parm...)
	l := len(s)
	if x+l > maxX {
		l = maxX - x
	}
	for i := 0; i < l; i++ {
		termbox.SetCell(x+i, y, rune(s[i]), fg, bg)
	}
}

func (m *MptMenu) Show() {
	line := 0
	index := 0

	fg := termbox.ColorWhite
	bg := termbox.ColorBlack

	if len(m.title) > 0 {
		printAt(line, 0, fg, bg, m.title)
		line += 2
	}

	if m.hasTopOpt {
		m.PrintOption(line, fg, bg, index)
		line += 2
		index++
	}

	if len(m.description) > 0 {
		printAt(line, 0, fg, bg, m.description)
		line += 2
	}

	if len(m.prompt) > 0 {
		printAt(line, 0, fg, bg, m.prompt)
		line++
	}

	m.optionsLine = line
	for index < len(m.options) {
		m.PrintOption(line, fg, bg, index)
		index++
		line++
	}

	termbox.Flush()
}

func (m *MptMenu) PrintOption(line int, fg, bg termbox.Attribute, index int) {
	opt := m.options[index]
	isCurrent := index == m.current

	cursor := ' '
	if isCurrent {
		cursor = m.cursor
		fg, bg = bg, fg
	}
	printAt(line, 0, fg, bg, "%c %c. %s", cursor, opt.key, opt.text)
}

func (m *MptMenu) Current() int {
	return m.current
}

func (m *MptMenu) Prev() bool {
	if m.current > 0 {
		m.current--
		return true
	} else {
		return false
	}
}

func (m *MptMenu) Next() bool {
	if m.current+1 < len(m.options) {
		m.current++
		return true
	} else {
		return false
	}
}

func (m *MptMenu) CheckKey(key string) int {
	for i, opt := range m.options {
		if string([]byte{opt.key}) == key {
			m.current = i
			return i
		}
	}

	switch key {
	case "enter":
		return m.current
	case "down":
		m.Next()
	case "up":
		m.Prev()
	case "esc":
		return 0
	}

	return -1
}
