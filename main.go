package main

import (
	"encoding/json"
	"fmt"
	"github.com/Equanox/gotron"
	"strconv"
)

var x = DataInfo{}

//CustomEvent is the struct sent to the front end
type CustomEvent struct {
	*gotron.Event
	CustomAttribute []string `json:"eventData"`
}

//GetEvent is the struct get data from the front end
type GetEvent struct {
	// *gotron.Event
	Event string            `json:"event"`
	Data  map[string]string `json:"data"`
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
		d := x.search("Select * from rmaData")
		window.Send(&CustomEvent{
			Event:           &gotron.Event{Event: "get-all"},
			CustomAttribute: d,
		})
	})
	window.On(&gotron.Event{Event: "add-one"}, func(bin []byte) {
		var ge GetEvent
		var d DataInfo
		json.Unmarshal(bin, &ge)
		d.Comment = ge.Data["comment"]
		d.Date = ge.Data["date"]
		d.Sn1 = ge.Data["sn1"]
		d.Sn2 = ge.Data["sn2"]
		i, err := strconv.Atoi(ge.Data["rma"])
		if err != nil {
			fmt.Println("send a msg to frontend about rma being only numbers")
		}
		d.Rma = i
		d.addUser()
	})
	window.On(&gotron.Event{Event: "get-searchByRMA"}, func(bin []byte) {
		var ge GetEvent
		var d DataInfo
		json.Unmarshal(bin, &ge)	
		i, err := strconv.Atoi(ge.Data["rma"])
		if err != nil {
			fmt.Println("send a msg to frontend about rma being only numbers")
		}	
		query := fmt.Sprintf(`SELECT * FROM rmaData WHERE RMA = %d`, i)
		// query := fmt.Sprintf(`SELECT * FROM rmaData WHERE RMA = %d AND SN1 = "%s" AND DATE = "%s"`, d.Rma, d.Sn1, d.Date)
		res := d.search(query)
		fmt.Println(res)
	})
	window.On(&gotron.Event{Event: "get-searchBySN"}, func(bin []byte) {})
	window.On(&gotron.Event{Event: "get-searchByDATE"}, func(bin []byte) {})

	// Wait for the application to close
	<-done
}
