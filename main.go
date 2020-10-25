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

	// Adds a new record to the db
	window.On(&gotron.Event{Event: "add-one"}, func(bin []byte) {
		var ge GetEvent
		var d DataInfo
		json.Unmarshal(bin, &ge)
		d.Comment = ge.Data["comment"]
		d.Date = ge.Data["frmDate"]
		d.Sn1 = ge.Data["sn1"]
		d.Sn2 = ge.Data["sn2"]
		i, err := strconv.Atoi(ge.Data["rma"])
		if err != nil {
			sendMsg(window, "Only numbers allowed in the RMA field.")
		} else {
			d.Rma = i
			d.addRec()
			sendMsg(window, "Data added.")
		}
	})

	// Composes a search query and fetches the data from the db
	window.On(&gotron.Event{Event: "get-searchBy"}, func(bin []byte) {
		var ge GetEvent
		var d DataInfo
		q := []string{"IS NOT NULL", "IS NOT NULL", "IS NOT NULL"}
		json.Unmarshal(bin, &ge)
		fmt.Println(ge.Data)
		if ge.Data["rma"] != "" {
			q[0] = fmt.Sprintf(`= %s`, ge.Data["rma"])
		}
		if ge.Data["sn1"] != "" {
			q[1] = fmt.Sprintf(`= "%s" OR SN2 = "%s"`, ge.Data["sn1"], ge.Data["sn1"])
		}
		if ge.Data["frmDate"] != "" {
			q[2] = fmt.Sprintf(`= "%s"`, ge.Data["frmDate"])
		}
		query := fmt.Sprintf(`SELECT * FROM rmaData WHERE RMA %s AND SN1 %s AND DATE %s`, q[0], q[1], q[2])
		res, r := d.search(query)
		if r != "" {
			sendMsg(window, r)
		} else {
			if ge.Data["update"] == "true" {
				sendBack(window, res, "show-results")
				sendBack(window, res, "show-inForm")
			} else {
				sendBack(window, res, "show-results")
			}
		}
	})

	// Updates a single record from the db
	window.On(&gotron.Event{Event: "update-one"}, func(bin []byte) {
		var ge GetEvent
		var d DataInfo
		json.Unmarshal(bin, &ge)
		d.Comment = ge.Data["comment"]
		d.Date = ge.Data["frmDate"]
		d.Sn1 = ge.Data["sn1"]
		d.Sn2 = ge.Data["sn2"]
		i, err := strconv.Atoi(ge.Data["rma"])
		if err != nil {
			sendMsg(window, "Only numbers allowed in the RMA field.")
		} else {
			d.Rma = i
			n, _ := strconv.Atoi(ge.Data["prev"])
			d.update(n)
			sendMsg(window, "Data updated.")
		}
	})

	// Wait for the application to close
	<-done
}

// Sends the results from the db query back to the front end.
func sendBack(window *gotron.BrowserWindow, d []string, e string) {
	window.Send(&CustomEvent{
		Event:           &gotron.Event{Event: e},
		CustomAttribute: d,
	})
}

// Sends a custom message back to the front end.
func sendMsg(window *gotron.BrowserWindow, m string) {
	type ErrorEvent struct {
		*gotron.Event
		CustomAttribute string `json:"msg"`
	}
	window.Send(&ErrorEvent{
		Event:           &gotron.Event{Event: "show-msg"},
		CustomAttribute: m,
	})
}
