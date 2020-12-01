package conf

import (
	"flag"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	RedisCheckListKey = "im:active"
	RedisCheckIncKey  = "im:inc:"
	RedisCheckHashKey = "im:hash:"
)

var C *Conf

type Conf struct {
	IoPort      string `yaml:"io_port"`
	RpcPort     string `yaml:"rpc_port"`
	RedisHost   string `yaml:"redis_host"`
	RedisPwd    string `yaml:"redis_pwd"`
	RedisPrefix int64  `yaml:"redis_prefix"`
}

func init() {
	confPath := flag.String("c", "./conf/config.yml", "config path")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*confPath)
	if err != nil {
		panic("conf init fail")
	}
	C = &Conf{}
	if err = yaml.Unmarshal(yamlFile, C); err != nil {
		panic("conf Unmarshal fail")
	}
}

func GetHostStr() string {
	return "localhost:" + C.RpcPort
}
