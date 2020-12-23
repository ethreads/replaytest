package config

import "flag"

type AppConfigs struct {
	InputDev string `json:"input-dev"`
}

var Configs AppConfigs

func init() {
	flag.StringVar(&Configs.InputDev, "input-dev", "", "监听的网卡设备 Example: eth0")
}
