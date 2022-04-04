package prompt

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	// "strings"

	// "github.com/TheVoidProject/annoyme/pkg/reminder"
	"github.com/manifoldco/promptui"
)

// type Reminder struct {
// 	Title 			string
// 	Message 		string
// 	Datetime 		time.Time
// 	Recurring 	bool
// 	Repeat 			int
// 	Delay 			time.Duration
// 	Sound 			bool
// }


func GetString(err_msg string, label string) string {
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

func GetInt(err_msg string, label string) int {
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

func GetTime(err_msg string) string {
	validate := func(input string) error {
        timeRegex := regexp.MustCompile(`(?m)^\d{1,2}:\d\d$`)
        if timeRegex.Match([]byte(input)) {
            return nil
        } else {
            return errors.New("Invalid Time Format")
        }
	}

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }
    prompt := promptui.Prompt{
        Label:     "What time? format: HH:MM",
        Templates: templates,
        Validate:  validate,
        Pointer: promptui.PipeCursor,
    }
    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    return result
}
func GetBool(err_msg string, label string) bool {
	validate := func(input string) error {
        if strings.ToLower(input) == "y" || strings.ToLower(input) == "n" {
            return nil
        } else {
            return errors.New("Not y/n")
        }
	}

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }
    prompt := promptui.Prompt{
        Label: label +" [y/n]",
        Templates: templates,
        Validate:  validate,
        Pointer: promptui.PipeCursor,
    }
    in, err := prompt.Run()
    result := strings.ToLower(in)
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }
    if result == "y" {
        return true
    } else {
        return false
    }
}

func GetDay() string {
    days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"}
    templates := &promptui.SelectTemplates{
        Inactive:  "{{ . }} ",
        Active:   "{{ . | green }} ",
        // Invalid: "{{ . | red }} ",
        Selected: "{{ . | bold }} ",
    }
    // templates := &promptui.SelectTemplates{
	// 	Label:    "{{ . }}?",
	// 	// Active:   "\U0001F336 {{ .Name | cyan }})",
	// 	// Inactive: "  {{ .Name | white }}",
	// 	// Selected: "\U0001F336 {{ .Name | red | cyan }}",
	// }
	// searcher := func(input string, index int) bool {
	// 	// day := days[index]
	// 	name := strings.Replace(strings.ToLower(days[index]), " ", "", -1)
	// 	input = strings.Replace(strings.ToLower(input), " ", "", -1)

	// 	return strings.Contains(name, input)
	// }
    prompt := promptui.Select{
		Label: "Select Day",
		Items: days,
        Templates: templates,
        IsVimMode: true,
        // Size: 7,
        // Searcher:  searcher,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	fmt.Printf("You choose %q\n", result)
    return result
}