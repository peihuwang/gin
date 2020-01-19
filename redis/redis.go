/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2019/11/11
 * Time: 10:56
 */
package redis

import (
	"fmt"
	"gin/config"
	"github.com/garyburd/redigo/redis"
	_"log"
)

var RedisConn redis.Conn

func Setup(){
	var errs error
	address :=config.ConfSet.Redis.Host +":"+fmt.Sprintf("%d",config.ConfSet.Redis.Port)

	RedisConn, errs = redis.Dial("tcp", address)

	if errs !=nil{
		fmt.Println(errs)
	}

	password :=config.ConfSet.Redis.Password
	_,err := RedisConn.Do("auth", password)

	if err !=nil{
		fmt.Println(err)
	}

	db :=config.ConfSet.Redis.Db

	if db !=0{
		SelectDb(db)
	}


}

func SelectDb(port int)  {
	//切换DB
	RedisConn.Send("select",port)
	RedisConn.Flush()
}

func Get(key string) ([]byte,error) {
	//defer RedisConn.Close()
	bytes, e := redis.Bytes(RedisConn.Do("get", key))
	return bytes, e
}

func Set(key string,value string,time int) bool{
	//defer RedisConn.Close()
	RedisConn.Do("set", key, value, "ex", time)

	return true
}

//需要其他redis命令自行在这里扩充