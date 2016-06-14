package testdata

var ConvertDestination1 = "/tmp/test-env-root-json/"

var ConvertSource1 = "testdata/test-env-root-json"

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
