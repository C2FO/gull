package gull

import (
	"fmt"

	"github.com/franela/goreq"
)

type EtcdMigrationTarget struct {
	EtcdHostUrl string
}

func NewEtcdMigrationTarget(hostUrl string) *EtcdMigrationTarget {
	return &EtcdMigrationTarget{
		EtcdHostUrl: hostUrl,
	}
}

func (emt *EtcdMigrationTarget) Set(path string, value string) error {
	if emt.EtcdHostUrl == "" {
		return fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := emt.EtcdHostUrl + path
	response, err := goreq.Request{
		Method: "PUT",
		Uri:    storageUrl,
		Body:   value,
	}.Do()
	if response != nil {
		defer func() { _ = response.Body.Close() }()
		if response.Response != nil {
			statusCode := response.Response.StatusCode
			if statusCode != 200 && statusCode != 201 {
				return fmt.Errorf("etcd returned HTTP %v on a PUT for [%v]->[%v]", statusCode, path, value)
			}
		}
	} else {
		return fmt.Errorf("etcd did not sent a response on a PUT for [%v]->[%v]", path, value)
	}
	return err
}