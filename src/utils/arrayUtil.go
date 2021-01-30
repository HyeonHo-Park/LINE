package utils

type pingInfo struct {
	Hostname string `json:"hostname"`
	Count    int    `json:"count"`
}

func RemoveByHostname(pingList []pingInfo, hostname string) []pingInfo {
	for i := range pingList {
		if pingList[i].Hostname == hostname {
			return append(pingList[:i], pingList[i+1:]...)
		}
	}
	return pingList
}
