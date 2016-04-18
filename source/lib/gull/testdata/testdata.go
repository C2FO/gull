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

var ValidYamlMigration1 = `---
          up:
            entries:
              - path: "/default/services"
                value: "[well hi there]"
          down:
            entries:
              - path: "/default/services"
                value: "delete"
        `
