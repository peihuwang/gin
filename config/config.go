/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	_ "log"
	"path/filepath"
)

type Conf struct {
	Database struct{
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Redis struct{
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Db int `yaml:"db"`
	}
}

var ConfSet *Conf

func RedisConfig()  {
	filename :=`gin/config/config.yaml`

	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)



	configFile, err := ioutil.ReadFile(filename)

	if err != nil{
		fmt.Println(err)
		//log.Fatalln(err)
	}


	err = yaml.Unmarshal(configFile, &ConfSet)
	if err != nil{
		fmt.Println(err)
		//log.Fatalln(err)
	}
}