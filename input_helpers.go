package terminal

import (
	"bufio"
	"fmt"
	"strings"
)

func AskString(message string, validator func(string) (string, bool)) string {
	return AskStringDefault(message, "", validator)
}

func AskStringDefault(message, def string, validator func(string) (string, bool)) string {
	if !Stdin.IsInteractive() {
		return def
	}
	if def != "" {
		message = fmt.Sprintf("%s <question>[%s]</> ", message, def)
	}

	reader := bufio.NewReader(Stdin)
	for {
		Print(message)
		answer, readError := reader.ReadString('\n')
		if readError != nil {
			continue
		}
		answer = strings.TrimRight(answer, "\r\n")
		answer = strings.Trim(answer, " ")
		if answer == "" {
			answer = def
		}
		if answer, isValid := validator(answer); !isValid {
			continue
		} else {
			return answer
		}
	}
}

func AskConfirmation(message string, def bool) bool {
	if !Stdin.IsInteractive() {
		return def
	}
	defaultHelp, defaultAnswer := "Y/n", "yes"
	if !def {
		defaultHelp, defaultAnswer = "y/N", "no"
	}
	message = fmt.Sprintf("%s <question>[%s]</> ", message, defaultHelp)

	answer := AskString(message, func(answer string) (string, bool) {
		answer = strings.ToLower(answer)
		if answer == "" {
			return defaultAnswer, true
		}
		if answer == "y" || answer == "yes" {
			return "yes", true
		}
		if answer == "n" || answer == "no" {
			return "no", true
		}

		return answer, false
	})

	return answer == "yes"
}
