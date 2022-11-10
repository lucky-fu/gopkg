package middleware

import (
	"fmt"
	"os"
)

var (
	ENV      string
	Port     string
	RootPath string
)

// InitVariables ...
func InitVariables() {
	var (
		err error
		ok  bool
	)

	if RootPath, err = os.Getwd(); err != nil {
		panic(any(fmt.Sprintf("get root path failed : %v", err.Error())))
	}

	if ENV, ok = os.LookupEnv("ENV"); !ok {
		panic(any("get env failed"))
	}

	if Port, ok = os.LookupEnv("PORT"); !ok {
		panic(any("get port failed"))
	}

	fmt.Printf("初始化变量成功\n")
}
