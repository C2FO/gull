package gull

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/franela/goreq"
)

type EtcdMigrationTarget struct {
	EtcdHostUrl   string
	Application   string
	Environment   string
	FullMigration bool
	logger        ILogger
}

type etcdGetResponse struct {
	Node etcdPair
}

type etcdPair struct {
	Key   string
	Value string
}

func NewEtcdMigrationTarget(hostUrl string, application string, environment string, performFullMigration bool, logger ILogger) *EtcdMigrationTarget {
	hostUrl = formatEtcdHostUrl(hostUrl)
	return &EtcdMigrationTarget{
		EtcdHostUrl:   hostUrl,
		Application:   application,
		Environment:   environment,
		FullMigration: performFullMigration,
		logger:        logger,
	}
}

func formatEtcdHostUrl(url string) string {
	etcdTwoDefaultPort := 2379
	// Add a port if one wasn't specified
	if !strings.Contains(url, ":") {
		head := strings.Split(url, "/")[0]
		headWithPort := fmt.Sprintf("%v:%v", head, etcdTwoDefaultPort)
		url = strings.Replace(url, head, headWithPort, 1)
	}
	// Default to HTTP if none specified
	if !strings.Contains(url, "://") {
		url = fmt.Sprintf("http://%v", url)
	}
	// Add the resource locator etcd2 expects
	etcdTwoResource := "v2/keys"
	if !strings.Contains(url, etcdTwoResource) {
		lastSlash := strings.LastIndex(url, "/")
		if lastSlash != len(url)-1 {
			etcdTwoResource = fmt.Sprintf("/%v", etcdTwoResource)
		}
		url = url + etcdTwoResource
	}
	return url
}

func (emt *EtcdMigrationTarget) Set(path string, value string) error {
	if emt.EtcdHostUrl == "" {
		return fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := emt.EtcdHostUrl + url.QueryEscape(emt.getAppPath(path))
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
				return fmt.Errorf("etcd @ [%v] returned HTTP %v on a PUT for [%v]->[%v]", emt.EtcdHostUrl, statusCode, emt.getAppPath(path), value)
			}
		}
	} else {
		return fmt.Errorf("etcd [%v] did not sent a response on a PUT for [%v]->[%v]", emt.EtcdHostUrl, path, value)
	}
	return err
}

func (emt *EtcdMigrationTarget) Get(path string) (string, error) {
	if emt.EtcdHostUrl == "" {
		return "", fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := emt.EtcdHostUrl + url.QueryEscape(emt.getAppPath(path))
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

func (emt *EtcdMigrationTarget) getAppPath(path string) string {
	return "/" + emt.Application + path
}

func (emt *EtcdMigrationTarget) DeleteEnvironment() error {
	return emt.remove(emt.GetApplication() + "/" + emt.GetEnvironment())
}

func (emt *EtcdMigrationTarget) DeleteApplication() error {
	return emt.remove(emt.GetApplication())
}

func (emt *EtcdMigrationTarget) remove(root string) error {
	if emt.EtcdHostUrl == "" {
		return fmt.Errorf("EtcdMigrationTarget's EtcdHostUrl cannot be empty")
	}
	storageUrl := fmt.Sprintf("%v/%v%v", emt.EtcdHostUrl, url.QueryEscape(root), "?dir=true&recursive=true")
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
				return fmt.Errorf("etcd @ [%v] returned HTTP %v on a DELETE for [%v]", emt.EtcdHostUrl, statusCode, root)
			}
		}
	}
	return nil
}

func (emt *EtcdMigrationTarget) GetEnvironment() string {
	return emt.Environment
}

func (emt *EtcdMigrationTarget) GetApplication() string {
	return emt.Application
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
	return fmt.Sprintf("/%v/_gull/state", emt.GetEnvironment())
}

func (emt *EtcdMigrationTarget) IsPerformingFullMigration() bool {
	return emt.FullMigration
}

func (emt *EtcdMigrationTarget) GetMigrationTip() (*Migration, error) {
	migrationState, err := emt.GetStatus()
	if err != nil {
		return nil, err
	}
	return migrationState.Migrations.Last()
}

func (emt *EtcdMigrationTarget) GetLogger() ILogger {
	return emt.logger
}
