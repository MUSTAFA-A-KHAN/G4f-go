package g4f

import (
	"fmt"
	"os/exec"
	"strings"
)

func G4f(textInput string, MType string) (bool, string) {
	// Define the text input to pass to the Python script
	// textInput := "Explain quantum computing briefly"
	// Mtype := "text"
	var output []byte
	var err error
	MType = strings.ToLower(MType)

	switch MType {
	case "text":
		// Command to run the Python script with the text input as an argument
		cmd := exec.Command("python3", "./G4f/gpt-python/gpt-python.py", textInput)

		// Run the command and capture output
		output, err = cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Error executing Python script: %v", err)
		}
		return true, string(output)
	case "voice":
		// Command to run the Python script with the text input as an argument
		cmd := exec.Command("python3", "./G4f/gpt-python/generate-audio.py", textInput)

		// Run the command and capture output
		output, err = cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Error executing Python script: %v", err)
		}
		return true, string(output)
	case "custom-voice":
		// Command to run the Python script with the text input as an argument
		cmd := exec.Command("python3", "./G4f/gpt-python/CustomStyle.py", textInput)

		// Run the command and capture output
		output, err = cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Error executing Python script: %v", err)
		}
		return true, string(output)
	case "image":
		// Command to run the Python script with the text input as an argument
		cmd := exec.Command("python", "./G4f/gpt-python/generate-image.py", textInput)

		// Run the command and capture output
		output, err = cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Error executing Python script: %v", err)
		}
		return true, string(output)
	}

	// Print the output from the Python script
	fmt.Printf("Output from Python script:\n%s", output)

	return false, "No MType found"
}
