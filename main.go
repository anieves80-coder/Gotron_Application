package main

import (
	// "fmt"
	"github.com/Equanox/gotron"
	// "encoding/json"
)

var x = DataInfo{}

//CustomEvent is the struct sent to the front end
type CustomEvent struct {
	*gotron.Event
	CustomAttribute []string `json:"eventData"`
}

func main() {

	// Create a window instance
	window, err := gotron.New("./ui/")
	if err != nil {
		panic(err)
	}

	// Change the default window size.
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980

	// Start the browser window.
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Use dev tools for development.
	// Comment out for production
	// Needs to set after starting the window/browser
	//window.OpenDevTools()

	window.On(&gotron.Event{Event: "get-all"}, func(bin []byte) {
		d := x.getAll()
		window.Send(&CustomEvent{
			Event:           &gotron.Event{Event: "get-all"},
			CustomAttribute: d,
		})

	})

	// Wait for the application to close
	<-done
}
