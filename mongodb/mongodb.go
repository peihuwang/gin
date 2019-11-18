/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

func SetUp(){
	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)

	client, e := mongo.Connect(timeout, options.Client().ApplyURI("mongodb://localhost:27017"))

	if e !=nil{
		fmt.Println(e)
	}

	fmt.Println(client)
}