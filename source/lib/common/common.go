package common

var DefaultGullDirectory = "_gull/"

var DefaultEtcdServerUrl = "http://localhost:2379/v2/keys"

var ApplicationVersion string

func GetApplicationVersion() string {
	if ApplicationVersion == "" {
		return "'development'"
	}
	return ApplicationVersion
}
