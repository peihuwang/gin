/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package bootstrap

import (
	"gin/config"
	"gin/mysqld"
	"gin/redis"
)

func Bootstrap()  {
	//读取配置文件
	config.RedisConfig()

	//链接mysql
	mysqld.DbConnect()

	//链接redis
	redis.Setup()
}