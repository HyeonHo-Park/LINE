package test

import (
	"testing"

	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
	"github.com/stretchr/testify/assert"
)

func TestCheckByHostname(t *testing.T) {
	// Set UP
	var pingList []PingInfo
	data1 := PingInfo{hostname1, count1}
	data2 := PingInfo{hostname2, count2}

	pingList = append(pingList, data1)
	pingList = append(pingList, data2)

	// Act
	case1 := CheckByHostname(pingList, hostname1)
	case2 := CheckByHostname(pingList, hostname2)
	case3 := CheckByHostname(pingList, "nothing")

	// Assertion
	assert.Equal(t, true, case1)
	assert.NotEqual(t, false, case2)
	assert.Equal(t, false, case3)
}

func TestGetIndexByHostname(t *testing.T) {
	// Set UP
	var pingList []PingInfo
	data1 := PingInfo{hostname1, count1}
	data2 := PingInfo{hostname2, count2}

	pingList = append(pingList, data1)
	pingList = append(pingList, data2)

	// Act
	case1 := GetIndexByHostname(pingList, hostname1)
	case2 := GetIndexByHostname(pingList, hostname2)
	case3 := GetIndexByHostname(pingList, "nothing")

	// Assertion
	assert.Equal(t, 0, case1)
	assert.Equal(t, 1, case2)
	assert.Equal(t, -1, case3)
}

func TestRemoveByHostname(t *testing.T) {
	// Set UP
	var pingList []PingInfo
	var case1List []PingInfo

	data1 := PingInfo{hostname1, count1}
	data2 := PingInfo{hostname2, count2}

	pingList = append(pingList, data1)
	pingList = append(pingList, data2)

	case1List = append(case1List, data2)

	// Act
	case1 := RemoveByHostname(pingList, hostname1)
	case2 := RemoveByHostname(case1, hostname2)

	// Assertion
	assert.Equal(t, case1List, case1)
	assert.Empty(t, case2)
}
