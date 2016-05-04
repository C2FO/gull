package gull

import (
	"fmt"
	"net/url"

	"github.com/franela/goreq"
)

type EtcdMigrationTarget struct {
	EtcdHostUrl string
	Environment string
}

func NewEtcdMigrationTarget(hostUrl string, environment string) *EtcdMigrationTarget {
	return &EtcdMigrationTarget{
		EtcdHostUrl: hostUrl,
		Environment: environment,
	}
}

func (emt *EtcdMigrationTarget) Set(path string, value string) error {
	if emt.EtcdHostUrl == "" {
		return fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := emt.EtcdHostUrl + url.QueryEscape(path)
	value = fmt.Sprintf("value=%v", url.QueryEscape(value))
	fmt.Printf("Writing [%v]->[%v]\n", storageUrl, value)
	response, err := goreq.Request{
		Method:      "PUT",
		Uri:         storageUrl,
		Body:        value,
		ContentType: "application/x-www-form-urlencoded",
	}.Do()
	if response != nil {
		defer func() { _ = response.Body.Close() }()
		if response.Response != nil {
			statusCode := response.Response.StatusCode
			if statusCode != 200 && statusCode != 201 {
				return fmt.Errorf("etcd @ [%v]returned HTTP %v on a PUT for [%v]->[%v]", emt.EtcdHostUrl, statusCode, path, value)
			}
		}
	} else {
		return fmt.Errorf("etcd did not sent a response on a PUT for [%v]->[%v]", path, value)
	}
	return err
}

func (emt *EtcdMigrationTarget) GetEnvironment() string {
	return emt.Environment
}

func (emt *EtcdMigrationTarget) GetAll() map[string]string {
	return nil
}

func (emt *EtcdMigrationTarget) Debug() {

}
