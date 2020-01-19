package log

import (
	"fmt"
)

func DemoLog() {
	fmt.Println("------log---------")
	MainLogger.Debug("hello main Debug")
	MainLogger.Info("hello main Info")
	GatewayLogger.Debug("Hi Gateway Im Debug")
	GatewayLogger.Info("Hi Gateway  Im Info")
}