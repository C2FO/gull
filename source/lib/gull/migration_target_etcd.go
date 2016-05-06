package gull

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/franela/goreq"
)

type EtcdMigrationTarget struct {
	EtcdHostUrl string
	Environment string
}

type etcdGetResponse struct {
	Node etcdPair
}

type etcdPair struct {
	Key   string
	Value string
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
	response, err := goreq.Request{
		Method:      "PUT",
		Uri:         storageUrl,
		Body:        value,
		ContentType: "application/x-www-form-urlencoded",
	}.Do()
	if response != nil {
		_ = response.Body.Close()
		if response.Response != nil {
			statusCode := response.Response.StatusCode
			if statusCode != 200 && statusCode != 201 {
				return fmt.Errorf("etcd @ [%v] returned HTTP %v on a PUT for [%v]->[%v]", emt.EtcdHostUrl, statusCode, path, value)
			}
		}
	} else {
		return fmt.Errorf("etcd did not sent a response on a PUT for [%v]->[%v]", path, value)
	}
	return err
}

func (emt *EtcdMigrationTarget) Get(path string) (string, error) {
	if emt.EtcdHostUrl == "" {
		return "", fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := emt.EtcdHostUrl + url.QueryEscape(path)
	response, err := goreq.Request{
		Method: "GET",
		Uri:    storageUrl,
	}.Do()
	if err != nil {
		return "", err
	}
	if response != nil {
		defer func() { _ = response.Body.Close() }()
		bodyBytes, err := ioutil.ReadAll(response.Response.Body)
		if err != nil {
			return "", err
		}
		var etcdResponse etcdGetResponse
		err = json.Unmarshal(bodyBytes, &etcdResponse)
		return etcdResponse.Node.Value, err
	}
	return "", fmt.Errorf("etcd did not send a response on a GET for [%v]", path)
}

func (emt *EtcdMigrationTarget) DeleteEnvironment() error {
	if emt.EtcdHostUrl == "" {
		return fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := fmt.Sprintf("%v/%v%v", emt.EtcdHostUrl, url.QueryEscape(emt.GetEnvironment()), "?dir=true&recursive=true")
	response, err := goreq.Request{
		Method: "DELETE",
		Uri:    storageUrl,
	}.Do()
	if err != nil {
		return err
	}
	if response != nil {
		_ = response.Body.Close()
		if response.Response != nil {
			statusCode := response.Response.StatusCode
			if statusCode != 200 && statusCode != 201 {
				return fmt.Errorf("etcd @ [%v] returned HTTP %v on a DELETE for [%v]", emt.EtcdHostUrl, statusCode, emt.GetEnvironment())
			}
		}
	}
	return nil
}

func (emt *EtcdMigrationTarget) GetEnvironment() string {
	return emt.Environment
}

func (emt *EtcdMigrationTarget) GetAll() map[string]string {
	return nil
}

func (emt *EtcdMigrationTarget) Debug() {

}

func (emt *EtcdMigrationTarget) SetStatus(state *MigrationState) error {
	yamlBytes, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return emt.Set(emt.getStatusPath(), string(yamlBytes))
}

func (emt *EtcdMigrationTarget) GetStatus() (*MigrationState, error) {
	migrationState := &MigrationState{}
	stateYaml, err := emt.Get(emt.getStatusPath())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(stateYaml), migrationState)
	return migrationState, err
}

func (emt *EtcdMigrationTarget) getStatusPath() string {
	return fmt.Sprintf("/%v/_gull/migration/state", emt.GetEnvironment())
}
