package main

import (
	"errors"
	"os"
	"fmt"
    "strconv"

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
        Pointer: promptui.PipeCursor,
    }
    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    // fmt.Printf("Input: %s\n", result)
    return result
}

func getIntInput(err_msg string, label string) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid number")
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
        Pointer: promptui.PipeCursor,
    }
    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    // fmt.Printf("Input: %s\n", result)
    res, _ := strconv.Atoi(result)
    return res
}