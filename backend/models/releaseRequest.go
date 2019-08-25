package models

import (
	//"encoding/json"
)

//type ReleaseRequest struct {
//	RequestID string `json:"RequestID"`
//	Workload []struct {
//		Container string `json:"Container"`
//		Current   string `json:"Current"`
//		Target    string `json:"Target"`
//	} `json:"default:deployment/flux-web"`
//	Cause       struct {
//		Message string `json:"Message"`
//		User    string `json:"User"`
//	} `json:"Cause"`
//}

type ReleaseRequest struct {
	RequestID string `json:"RequestID"`
}

func NewReleseRequest(data string) (ReleaseRequest) {
	var r ReleaseRequest
	r.RequestID = data
	return r
}

//func NewReleseRequest(data []byte) (ReleaseRequest, error) {
//	var r ReleaseRequest
//	err := json.Unmarshal(data, &r)
//	return r, err
//}