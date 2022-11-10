package middleware

import (
	"fmt"

	"github.com/gopkg/define"
	"github.com/gopkg/util"
)

var (
	Config *define.Config
)

// InitConfig ...
func InitConfig() {
	Config = &define.Config{}

	if err := util.Parse.ParseYml(fmt.Sprintf("config/%s.yml", ENV), Config); err != nil {
		panic(any(fmt.Sprintf("parse yaml failed: %v", err.Error())))
	}

	fmt.Printf("初始化配置成功\n")
}
