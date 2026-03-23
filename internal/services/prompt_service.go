package services

import "log"

func Generation(query string) string {
	result, err := SendGenerateNewPrompt(query)

	if err != nil {
		log.Println("Prompt generation failed", err)
	}

	return result
}