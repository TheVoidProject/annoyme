package main

import (
	"errors"
	"os"
	"fmt"

	"github.com/manifoldco/promptui"
)

func getInput(err_msg string, label string) string {
    validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(err_msg)
        }
        return nil
    }

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }
    prompt := promptui.Prompt{
        Label:     label,
        Templates: templates,
        Validate:  validate,
    }
    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    // fmt.Printf("Input: %s\n", result)
    return result
}