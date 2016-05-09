package testdata

var ConvertDestination1 = "/tmp/test-env-root-json/"

var ConvertSource1 = "testdata/test-env-root-json"

var ConvertDestination2 = "/tmp/gostrufig-json/"

var ConvertSource2 = "testdata/gostrufig-json/"

var ValidEtcdHostUrl = "http://localhost:4002/v2/keys"

var ValidJsonConfig1 = `
    {
        "*":{
            "services":["well", "hi", "there"],
            "enableLogging": false
        },
        "production":{
            "enableLogging":true
        }
    }
`

var ValidYamlMigration1 = `
---
entries:
- path: "/default/services"
  value: "[well hi there]"
`

var ValidJsonConfig2 = `
{
    "*": {
        "logging": {
            "c2fo": {
                "level": "INFO",
                "appenders": [
                    {
                        "wrapStyle": false,
                        "type": "ConsoleAppender",
                        "pattern": "\u001b[36m{name}\u001b[0m {gid} {pid} {levelNameColored} - {message}"
                    }
                ]
            }
        }
    }
}

`

type GostrufigTestConfig1 struct {
	Environment string `cfg-ns:"true" cfg-def:"default"`
	Alice       []string
	Lewis       string
	Lyrics      string
}
