package main

func RemoveByHostname(hostname string) {
	for i := range pingList {
		if pingList[i].Hostname == hostname {
			pingList = append(pingList[:i], pingList[i+1:]...)
		}
	}
}
