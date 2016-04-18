package testdata

var ValidJsonConfig1 = "{\"*\":{\"services\":[\"well\",\"hi\",\"there\"],\"enableLogging\":false},\"production\": {\"enableLogging\":true}}"

var ValidYamlMigration1 = `---
          up:
            - path: "/default/services"
              value: "[well hi there]"
          down:
            - path: "/default/services"
              value: "delete"
        `
