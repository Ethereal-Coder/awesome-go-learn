/**
 * Created by sunyh-vm on 2019/12/20
 * Description:
 * Reference: https://github.com/go-redis/redis
 */

package redis_demo

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func InitRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.104:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Print("----------------")
	fmt.Println(pong, err)
	return client
}

func RedisKey(client *redis.Client) {
	// 存在返回1,否则返回0
	valInt, err := client.Exists("key").Result()
	log.Println(valInt, err)

	// 删除成功返回1,否则返回0
	valInt1, err1 := client.Del("key").Result()
	log.Println(valInt1, err1)

	valInt2, err2 := client.Del("key").Result()
	log.Println(valInt2, err2)

	valStr3, err3 := client.Type("key1").Result()
	log.Println(valStr3, err3)
	log.Println(valStr3 == "string")

	strList4, err4 := client.Keys("").Result()
	log.Println(strList4, err4)

	// 返回list长度
	valInt5, err5 := client.LLen("list_test").Result()
	log.Println(valInt5, err5)

	// 返回set长度
	valInt6, err6 := client.SCard("set_test").Result()
	log.Println(valInt6, err6)
}

func RedisString(client *redis.Client) {
	err := client.Set("key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func RedisList(client *redis.Client) {
	key := "list_test"
	result := client.RPush(key, "hello")
	//log.Println(result)
	//log.Println(result.Val())
	//log.Println(result.Uint64())
	//log.Println(result.Args())
	//log.Println(result.Name())
	//log.Println(result.Result())
	//log.Println(result.String())
	//log.Println(result.Err())
	if err := result.Err(); err != nil {
		panic(err)
	}

	val, err := client.LPop(key).Result()
	if err == redis.Nil {
		fmt.Printf("%s does not exist", key)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(key, val)
	}
}

func RedisSet(client *redis.Client) {
	key := "set_test"
	setMembers := []string{"1", "2"}
	valInt1, err1 := client.SAdd(key, setMembers).Result()
	log.Println(valInt1, err1)

	// 删除名称为key的set中的元素member
	//client.SRem(key, setMembers)

	// 随机返回并删除名称为key的set中一个元素
	//client.SPop(key)

	// 返回名称为key的set的基数
	//client.SCard(key)

	// 返回名称为key的set的所有元素
	//client.SMembers(key)

	// 随机返回名称为key的set的一个元素
	//client.SRandMember(key)
}
