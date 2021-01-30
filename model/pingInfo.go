package model

type PingInfo struct {
	Hostname string `json:"hostname"`
	Count    int    `json:"count"`
}
