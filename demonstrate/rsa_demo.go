package demonstrate

import (
	"fmt"
	"os"

	"github.com/ingoaf/rsa-example/encode"
	"github.com/tcnksm/go-input"
)

// StartRsa starts the demo of RSA Algorithm, starting with asking the user for an input message, ending with returning the decrypted message by the receiver
// The steps are briefly explained in the README.md, for further information see the presentation folder
func StartRsa() {

	message, err := askForMessage()
	if err != nil {
		panic("Message not provided: " + err.Error())
	}

	e := encode.NewService()
	encodedMessage := e.EncodeMessage(message)

	fmt.Println(encodedMessage)
}

// askForMessage asks the user for a message, with which the rsa-algorithm will be demonstrated
func askForMessage() (string, error) {

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "What is your message?"
	message, err := ui.Ask(query, &input.Options{
		Default:  "Hi!",
		Required: true,
		Loop:     true,
	})

	if err != nil {
		return "", err
	}

	return message, nil
}
