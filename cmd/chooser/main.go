package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmatsuoka/chooser"
)

const (
	title        = "Ubuntu Core for YoyoDyne Retro Encabulator 550"
	goBackOption = "Go Back"
)

func previousHandler() error {
	options := []chooser.MenuOption{
		{goBackOption, nil},
		{"some version", nil},
	}

	menu := chooser.NewMenu("", "", "Start into a previous version:", options, true)
	menu.Choose()

	return nil
}

func recoverHandler() error {
	// set system to boot in recover mode
	return nil
}

func reinstallHandler() error {
	// reinstall system
	return nil
}

func main() {
	//output := flag.String("output", "/run/chooser.out", "Output file location")
	//seed := flag.String("seed", "/run/ubuntu-seed", "Ubuntu-seed location")
	flag.Parse()

	if err := chooser.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer chooser.Deinit()

	options := []chooser.MenuOption{
		{"Start normally", nil},
		{"Start into a previous version  >", previousHandler},
		{"Recover                        >", recoverHandler},
		{"Reinstall                      >", reinstallHandler},
	}

	menu := chooser.NewMenu(title, "", "Use arrow/number keys then Enter:", options, false)
	menu.Choose()
}
