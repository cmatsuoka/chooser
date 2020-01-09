package mptmenu

import (
	nc "github.com/rthornton128/goncurses"
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

func (m *MptMenu) Show(scr *nc.Window, y, x int) {
	line := y
	index := 0

	if len(m.title) > 0 {
		scr.MovePrintf(line, 0, m.title)
		line += 2
	}

	if m.hasTopOpt {
		m.PrintOption(scr, line, index)
		line += 2
		index++
	}

	if len(m.description) > 0 {
		scr.MovePrintf(line, 0, m.description)
		line += 2
	}

	if len(m.prompt) > 0 {
		scr.MovePrintf(line, 0, m.prompt)
		line++
	}

	m.optionsLine = line
	for index < len(m.options) {
		m.PrintOption(scr, line, index)
		index++
		line++
	}

	scr.Refresh()
}

func (m *MptMenu) PrintOption(scr *nc.Window, line, index int) {
	opt := m.options[index]
	isCurrent := index == m.current

	cursor := ' '
	if isCurrent {
		cursor = m.cursor
		scr.AttrOn(nc.A_REVERSE)
	}
	scr.MovePrintf(line, 0, "%c %c. %s", cursor, opt.key, opt.text)
	scr.AttrOff(nc.A_REVERSE)
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

func (m *MptMenu) CheckKey(key string) (int, bool) {
	for i, opt := range m.options {
		if string([]byte{opt.key}) == key {
			m.current = i
			return i, true
		}
	}

	switch key {
	case "enter":
		return m.current, true
	case "down":
		m.Next()
		return 0, false
	case "up":
		m.Prev()
		return 0, false
	}

	return 0, false
}
