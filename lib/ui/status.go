package ui

import "log"

func (o *Options) Status() {
	log.Printf("The environment name is [%s] and the etcdHost is [%s]", o.Environment, o.EtcdHost)
	log.Printf("TODO: Implement the Status subcommand")
}
