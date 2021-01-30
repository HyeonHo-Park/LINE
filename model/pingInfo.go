package model

type pingInfo struct {
	Hostname string `json:"hostname"`
	Count    int    `json:"count"`
}
