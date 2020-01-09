package main

import (
	"flag"

	"github.com/cmatsuoka/chooser"
)

const (
	title        = "Ubuntu Core for YoyoDyne Retro Encabulator 550"
	goBackOption = "Go Back"
)

func previousHandler(c *chooser.Chooser) error {
	options := []chooser.MenuOption{
		{goBackOption, nil},
		{"some version", nil},
	}

	menu := chooser.NewMenu(c, "", "", "Start into a previous version:", options, true)
	menu.Choose()

	return nil
}

func recoverHandler(c *chooser.Chooser) error {
	// set system to boot in recover mode
	return nil
}

func reinstallHandler(c *chooser.Chooser) error {
	// reinstall system
	return nil
}

func main() {
	//output := flag.String("output", "/run/chooser.out", "Output file location")
	//seed := flag.String("seed", "/run/ubuntu-seed", "Ubuntu-seed location")
	flag.Parse()

	c := &chooser.Chooser{}
	c.Init()
	defer c.Deinit()

	options := []chooser.MenuOption{
		{"Start normally", nil},
		{"Start into a previous version  >", previousHandler},
		{"Recover                        >", recoverHandler},
		{"Reinstall                      >", reinstallHandler},
	}

	menu := chooser.NewMenu(c, title, "", "Use arrow/number keys then Enter:", options, false)
	menu.Choose()
}
