package main

import (
	"fmt";
	"os"
	"github.com/charmbracelet/huh"
)

func main() {
	var file string 

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
	}

		slate := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
			Title("\033[1mEditing "+file+"...\033[22m").
			Value(&content),
		),
	).WithTheme(huh.ThemeBase())

	slate.Run()

	err = os.WriteFile(file, []byte(content+"\n"), 0644)
	if err == nil {
		fmt.Print("\033[1m✔ "+file+" has been saved\n\033[22m")

	} else {
		fmt.Print("\033[1m✗ error "+file+"wasn't saved\n\033[22m")
	}
}
