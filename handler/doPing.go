package handler

import (
	"log"
	"os"

	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
)

func DoPing(pingList *[]PingInfo, info PingInfo) {
	// Set Log Path
	logPath := "/tmp/pingLog/" + info.Hostname + ".txt"

	// Log File Delete
	os.Remove(logPath)

	// set ping log
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(0)

	// Ping
	p := func(addr string, seq int) {
		dst, dur, err := Ping(seq, addr)
		if err != nil {
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		log.Printf("Response from %s: icmp_seq=%d time=%s\n", dst, seq, dur)
	}

	for i := 0; i < info.Count; i++ {
		p(info.Hostname, (i + 1))
	}

	*pingList = RemoveByHostname(*pingList, info.Hostname)
}
