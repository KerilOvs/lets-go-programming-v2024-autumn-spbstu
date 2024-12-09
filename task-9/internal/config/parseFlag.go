package config

import "flag"

var fileName string

func init() {

	flag.StringVar(&fileName, "config", "../../configs/config.yaml", "Read file with configuration data")
	flag.Parse()
}