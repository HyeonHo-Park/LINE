package handler

import (
	"log"
	"os"

	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
)

func doPing(info PingInfo) {
	// Set Log Path
	// logPath := "/tmp/pingLog/" + info.Hostname + ".txt"
	logPath := "/Users/hyeonho/Desktop/LINE/pingLog/" + info.Hostname + ".txt"

	// Log File Delete
	os.Remove(logPath)

	// Ping
	p := func(addr string) {
		dst, dur, err := Ping(addr)
		if err != nil {
			f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer f.Close()
			log.SetOutput(f)
			log.SetFlags(0)
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
	}

	for i := 0; i < info.Count; i++ {
		p(info.Hostname)
	}
}
