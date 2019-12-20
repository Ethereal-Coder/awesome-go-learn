/**
 * Created by sunyh-vm on 2019/12/19
 * Description:
 * Reference: https://github.com/mongodb/mongo-go-driver/blob/master/examples/documentation_examples/examples.go
 */

package main

import (
	"github.com/Ethereal-Coder/awesome-go-learn/storage-test/mongo_demo"
	"github.com/Ethereal-Coder/awesome-go-learn/storage-test/redis_demo"
)

func main() {
	//xMongo()
	xRedis()
}

func xRedis() {
	client := redis_demo.InitRedisClient()
	//redis_demo.RedisKey(client)
	//redis_demo.RedisString(client)
	//redis_demo.RedisList(client)
	redis_demo.RedisSet(client)
}

func xMongo() {
	coll := mongo_demo.GetMongoCollection("ygrsy")
	//mongo_demo.MongoConnect()
	//mongo_demo.MongoInsert(coll)
	//mongo_demo.MongoInsertMany(coll)
	//mongo_demo.MongoQueryTop(coll)
	//mongo_demo.MongoQueryEmbedded(coll)
	//mongo_demo.MongoQueryArray(coll)
	//mongo_demo.MongoQueryEmbeddedArray(coll)
	mongo_demo.QueryNullMissingFieldsExamples(coll)
}
