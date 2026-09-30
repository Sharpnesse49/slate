package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func main() {
	var file string

	keyMap := huh.NewDefaultKeyMap()

	keyMap.Text.NewLine = key.NewBinding(
		key.WithKeys("enter"),
	)
	keyMap.Text.Submit = key.NewBinding(
		key.WithKeys("alt+enter"),
		key.WithHelp("alt+enter", "to close"),
	)

	keyMap.Text.Next = key.NewBinding(key.WithDisabled())
	keyMap.Text.Prev = key.NewBinding(key.WithDisabled())
	keyMap.Text.Editor = key.NewBinding(key.WithDisabled())

	filepicker := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("\033[1mTargeted file :\033[22m").
				Value(&file),
		),
	).WithTheme(huh.ThemeBase())

	filepicker.Run()

	var content string

	_, err := os.Stat(file)

	if err == nil {
		data, err := os.ReadFile(file)
		if err == nil {
			content = string(data)
		}
	} else {
		content = ""
		var new bool
		new_choice := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("\033[1m" + file + " does not exist, create it?\n\033[22m").
					Value(&new),
			),
		).WithTheme(huh.ThemeBase())

		new_choice.Run()

		if !new {
			fmt.Print("\033[1m✔ slate has been succesfully aborted\n\033[22m")
			return
		}
	}

	slate := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("\033[1mEditing " + file + "...\033[22m").
				Lines(15).
				Value(&content),
		),
	).WithTheme(huh.ThemeBase()).
		WithKeyMap(keyMap)

	slate.Run()

	var save bool

	save_choice := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("\033[1mSave the file?\n\033[22m").
				Value(&save),
		),
	).WithTheme(huh.ThemeBase())

	save_choice.Run()

	if save {
		err = os.WriteFile(file, []byte(content+"\n"), 0644)
		if err == nil {
			fmt.Print("\033[1m✔ " + file + " has been saved\n\033[22m")

		} else {
			fmt.Print("\033[1m✗ error " + file + " cannot be saved\n\033[22m")
		}
	} else {
		fmt.Print("\033[1m✔ " + file + " has been succesfully discarded\n\033[22m")
	}
}
