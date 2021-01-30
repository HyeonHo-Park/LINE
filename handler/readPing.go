package handler

import (
	"io/ioutil"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadPingLog(hostname string) string {
	// Set Log Path
	// logPath := "/tmp/pingLog/" + hostname + ".txt"
	logPath := "/Users/hyeonho/Desktop/LINE/pingLog/" + hostname + ".txt"

	// Read Log File
	data, err := ioutil.ReadFile(logPath)
	errCheck(err)

	return string(data)
}
