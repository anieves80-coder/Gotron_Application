package main

import (
	"github.com/Equanox/gotron"
	"fmt"
)

var x = dataInfo{}

func main() {
	
	fmt.Println(x.getAll())

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
		fmt.Println(string(bin))
	})











	// Wait for the application to close
	<-done
}
