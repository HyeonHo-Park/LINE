package test

import (
	"testing"

	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
)

func TestPing(t *testing.T) {
	// Set UP
	count1 := 10
	data1 := PingInfo{hostname1, count1}

	// Act
	p := func(addr string, seq int) {
		dst, dur, err := Ping(seq, addr)
		if err != nil {
			t.Logf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		t.Logf("Response from %s: icmp_seq=%d time=%s\n", dst, seq, dur)
	}

	for i := 0; i < data1.Count; i++ {
		p(data1.Hostname, (i + 1))
	}
}
