package serial

import (
	"github.com/huin/goserial"
	"io/ioutil"
	"strings"
)

//TODO need to work on this library (or find something already done to talk to the base)

// findArduino looks for the file that represents the Arduino
// serial connection. Returns the fully qualified path to the
// device if we are able to find a likely candidate for an
// Arduino, otherwise an empty string if unable to find
// something that 'looks' like an Arduino device.
func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "ttyACM") ||
			strings.Contains(f.Name(), "ttyUSB") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}

func main() {
	// Find the device that represents the arduino serial
	// connection.
	c := &goserial.Config{Name: findArduino(), Baud: 9600}
	s, _ := goserial.OpenPort(c)
}
