package utils

import . "github.com/HyeonHo-Park/LINE/model"

func RemoveByHostname(pingList []PingInfo, hostname string) []PingInfo {
	for i := range pingList {
		if pingList[i].Hostname == hostname {
			return append(pingList[:i], pingList[i+1:]...)
		}
	}
	return pingList
}
