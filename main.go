package main

import (
	"os"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/Theraine/theraine-sdk-go/ascii"
	parent "github.com/Theraine/theraine-sdk-go/parent-sdk"
	platform "github.com/Theraine/theraine-sdk-go/platform-sdk"
)

func main() {
	ascii.Theraine()

	choice := ""
    prompt := &survey.Select{
        Message: "Choose an operation:",
        Options: []string{
			"createPlatform", 
			"runPlatform", 
			"exit",
		},
    }
    survey.AskOne(prompt, &choice, nil)

	switch choice {
	case "createPlatform":
		parent.Theraine()
	case "runPlatform":
		platform.Theraine()
	case "exit":
		os.Exit(0)
	}
}