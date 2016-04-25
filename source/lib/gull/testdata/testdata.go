package testdata

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
content:
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
