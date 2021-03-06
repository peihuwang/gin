package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gin/config"
	"os"
	"fmt"
)

var MainLogger *zap.Logger
//var GatewayLogger *zap.Logger
//var RedisLogger *zap.Logger
//var LmdbLogger *zap.Logger
//var HttpLogger *zap.Logger

func init() {

	filename :=config.GetConfigPath()+`/../../log/logs/main.log`
	fmt.Println(filename)
    os.Create(filename)

	MainLogger = NewLogger(filename, zapcore.InfoLevel, 128, 30, 7, true, "Main")
	//GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
	//RedisLogger = NewLogger("./logs/redis.log", zapcore.InfoLevel, 128, 30, 7, true, "redis")
	//LmdbLogger = NewLogger("./logs/lmdb.log", zapcore.InfoLevel, 128, 30, 7, true, "lmdb")
	//HttpLogger = NewLogger("./logs/http.log", zapcore.InfoLevel, 128, 30, 7, true, "http")
}