/**
 * Created by sunyh-vm on 2019/12/19
 * Description:
 * Reference: https://github.com/mongodb/mongo-go-driver/blob/master/examples/documentation_examples/examples.go
 */

package mongo_demo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoCollection(collStr string) (coll *mongo.Collection) {
	//clientOptions := options.Client().ApplyURI("mongodb://user:pwd@host:port")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	coll = client.Database("ai-company").Collection(collStr)
	return
}

func QueryNullMissingFieldsExamples(coll *mongo.Collection) {
	//{
	//	// Start Example 38
	//
	//	docs := []interface{}{
	//		bson.D{
	//			{"_id", 1},
	//			{"item", nil},
	//		},
	//		bson.D{
	//			{"_id", 2},
	//		},
	//	}
	//
	//	result, err := coll.InsertMany(context.Background(), docs)
	//
	//	// End Example 38
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Print(result.InsertedIDs)
	//}

	{
		// Start Example 39

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"item", nil},
			})

		// End Example 39

		log.Printf("------------------  %d  -----------------", 39)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 40
		// reference: https://www.runoob.com/mongodb/mongodb-operators-type.html
		// $type 10 代表 Null
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"item", bson.D{
					{"$type", 10},
				}},
			})

		// End Example 40

		log.Printf("------------------  %d  -----------------", 40)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 41

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"item", bson.D{
					{"$exists", false},
				}},
			})

		// End Example 41

		log.Printf("------------------  %d  -----------------", 41)
		queryManyResultPrint(cursor, err)
	}
}

func MongoQueryEmbeddedArray(coll *mongo.Collection) {
	//{
	//	// Start Example 29
	//
	//	docs := []interface{}{
	//		bson.D{
	//			{"item", "journal"},
	//			{"instock", bson.A{
	//				bson.D{
	//					{"warehouse", "A"},
	//					{"qty", 5},
	//				},
	//				bson.D{
	//					{"warehouse", "C"},
	//					{"qty", 15},
	//				},
	//			}},
	//		},
	//		bson.D{
	//			{"item", "notebook"},
	//			{"instock", bson.A{
	//				bson.D{
	//					{"warehouse", "C"},
	//					{"qty", 5},
	//				},
	//			}},
	//		},
	//		bson.D{
	//			{"item", "paper"},
	//			{"instock", bson.A{
	//				bson.D{
	//					{"warehouse", "A"},
	//					{"qty", 60},
	//				},
	//				bson.D{
	//					{"warehouse", "B"},
	//					{"qty", 15},
	//				},
	//			}},
	//		},
	//		bson.D{
	//			{"item", "planner"},
	//			{"instock", bson.A{
	//				bson.D{
	//					{"warehouse", "A"},
	//					{"qty", 40},
	//				},
	//				bson.D{
	//					{"warehouse", "B"},
	//					{"qty", 5},
	//				},
	//			}},
	//		},
	//		bson.D{
	//			{"item", "postcard"},
	//			{"instock", bson.A{
	//				bson.D{
	//					{"warehouse", "B"},
	//					{"qty", 15},
	//				},
	//				bson.D{
	//					{"warehouse", "C"},
	//					{"qty", 35},
	//				},
	//			}},
	//		},
	//	}
	//
	//	result, err := coll.InsertMany(context.Background(), docs)
	//
	//	// End Example 29
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Print(result.InsertedIDs)
	//}

	{
		// Start Example 30
		// 严格顺序匹配
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock", bson.D{
					{"warehouse", "A"},
					{"qty", 5},
				}},
			})

		// End Example 30

		log.Printf("------------------  %d  -----------------", 30)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 31

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock", bson.D{
					{"qty", 5},
					{"warehouse", "A"},
				}},
			})

		// End Example 31

		log.Printf("------------------  %d  -----------------", 31)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 32

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock.0.qty", bson.D{
					{"$lte", 20},
				}},
			})

		// End Example 32

		log.Printf("------------------  %d  -----------------", 32)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 33

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock.qty", bson.D{
					{"$lte", 20},
				}},
			})

		// End Example 33

		log.Printf("------------------  %d  -----------------", 33)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 34
		// 不严格要求字段顺序
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock", bson.D{
					{"$elemMatch", bson.D{
						{"qty", 5},
						{"warehouse", "A"},
					}},
				}},
			})

		// End Example 34

		log.Printf("------------------  %d  -----------------", 34)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 35

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock", bson.D{
					{"$elemMatch", bson.D{
						{"qty", bson.D{
							{"$gt", 10},
							{"$lte", 20},
						}},
					}},
				}},
			})

		// End Example 35

		log.Printf("------------------  %d  -----------------", 35)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 36

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock.qty", bson.D{
					{"$gt", 10},
					{"$lte", 20},
				}},
			})

		// End Example 36

		log.Printf("------------------  %d  -----------------", 36)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 37

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"instock.qty", 5},
				{"instock.warehouse", "A"},
			})

		// End Example 37

		log.Printf("------------------  %d  -----------------", 37)
		queryManyResultPrint(cursor, err)
	}
}

func MongoQueryArray(coll *mongo.Collection) {
	//{
	//	// Start Example 20
	//
	//	docs := []interface{}{
	//		bson.D{
	//			{"item", "journal"},
	//			{"qty", 25},
	//			{"tags", bson.A{"blank", "red"}},
	//			{"dim_cm", bson.A{14, 21}},
	//		},
	//		bson.D{
	//			{"item", "notebook"},
	//			{"qty", 50},
	//			{"tags", bson.A{"red", "blank"}},
	//			{"dim_cm", bson.A{14, 21}},
	//		},
	//		bson.D{
	//			{"item", "paper"},
	//			{"qty", 100},
	//			{"tags", bson.A{"red", "blank", "plain"}},
	//			{"dim_cm", bson.A{14, 21}},
	//		},
	//		bson.D{
	//			{"item", "planner"},
	//			{"qty", 75},
	//			{"tags", bson.A{"blank", "red"}},
	//			{"dim_cm", bson.A{22.85, 30}},
	//		},
	//		bson.D{
	//			{"item", "postcard"},
	//			{"qty", 45},
	//			{"tags", bson.A{"blue"}},
	//			{"dim_cm", bson.A{10, 15.25}},
	//		},
	//	}
	//
	//	result, err := coll.InsertMany(context.Background(), docs)
	//
	//	// End Example 20
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Print(result.InsertedIDs)
	//}

	{
		// Start Example 21
		// 顺序,个数都匹配
		cursor, err := coll.Find(
			context.Background(),
			bson.D{{"tags", bson.A{"red", "blank"}}},
		)

		// End Example 21
		log.Printf("------------------  %d  -----------------", 21)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 22
		// 顺序.个数不限制
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"tags", bson.D{{"$all", bson.A{"red", "blank"}}}},
			})

		// End Example 22

		log.Printf("------------------  %d  -----------------", 22)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 23
		// tags数组中包含red
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"tags", "red"},
			})

		// End Example 23

		log.Printf("------------------  %d  -----------------", 23)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 24
		// dim_cm数组中有元素大于25
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"dim_cm", bson.D{
					{"$gt", 25},
				}},
			})

		// End Example 24

		log.Printf("------------------  %d  -----------------", 24)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 25
		// dim_cm数组中有元素大于15小于20
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"dim_cm", bson.D{
					{"$gt", 15},
					{"$lt", 20},
				}},
			})

		// End Example 25

		log.Printf("------------------  %d  -----------------", 25)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 26
		// todo ???
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"dim_cm", bson.D{
					{"$elemMatch", bson.D{
						{"$gt", 22},
						{"$lt", 30},
					}},
				}},
			})

		// End Example 26

		log.Printf("------------------  %d  -----------------", 26)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 27
		// dim_cm数组第二个元素大于25
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"dim_cm.1", bson.D{
					{"$gt", 25},
				}},
			})

		// End Example 27

		log.Printf("------------------  %d  -----------------", 27)
		queryManyResultPrint(cursor, err)
	}

	{
		// Start Example 28
		// tags数组长度为3
		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"tags", bson.D{
					{"$size", 3},
				}},
			})

		// End Example 28

		log.Printf("------------------  %d  -----------------", 28)
		queryManyResultPrint(cursor, err)
	}
}

func queryManyResultPrint(cursor *mongo.Cursor, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(result)
		log.Print(result["_id"])
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

func MongoQueryEmbedded(coll *mongo.Collection) {
	//{
	//	docs := []interface{}{
	//		bson.D{
	//			{"item", "journal"},
	//			{"qty", 25},
	//			{"size", bson.D{
	//				{"h", 14},
	//				{"w", 21},
	//				{"uom", "cm"},
	//			}},
	//			{"status", "A"},
	//		},
	//		bson.D{
	//			{"item", "notebook"},
	//			{"qty", 50},
	//			{"size", bson.D{
	//				{"h", 8.5},
	//				{"w", 11},
	//				{"uom", "in"},
	//			}},
	//			{"status", "A"},
	//		},
	//		bson.D{
	//			{"item", "paper"},
	//			{"qty", 100},
	//			{"size", bson.D{
	//				{"h", 8.5},
	//				{"w", 11},
	//				{"uom", "in"},
	//			}},
	//			{"status", "D"},
	//		},
	//		bson.D{
	//			{"item", "planner"},
	//			{"qty", 75},
	//			{"size", bson.D{
	//				{"h", 22.85},
	//				{"w", 30},
	//				{"uom", "cm"},
	//			}},
	//			{"status", "D"},
	//		},
	//		bson.D{
	//			{"item", "postcard"},
	//			{"qty", 45},
	//			{"size", bson.D{
	//				{"h", 10},
	//				{"w", 15.25},
	//				{"uom", "cm"},
	//			}},
	//			{"status", "A"},
	//		},
	//	}
	//
	//	result, err := coll.InsertMany(context.Background(), docs)
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//
	//	log.Print(result)
	//}

	//{
	//	cursor, err := coll.Find(
	//		context.Background(),
	//		bson.D{
	//			{"size", bson.D{
	//				{"h", 14},
	//				{"w", 21},
	//				{"uom", "cm"},
	//			}},
	//		})
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//	defer cursor.Close(context.Background())
	//	for cursor.Next(context.Background()){
	//		var result bson.D
	//		err := cursor.Decode(&result)
	//		if err != nil{
	//			log.Fatal(err)
	//		}
	//		log.Print(result.Map())
	//		log.Print(result.Map()["_id"])
	//	}
	//}

	//{
	//	cursor, err := coll.Find(
	//		context.Background(),
	//		bson.D{
	//			{"size", bson.D{
	//				{"w", 21},
	//				{"h", 14},
	//				{"uom", "cm"},
	//			}},
	//		})
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//	defer cursor.Close(context.Background())
	//	for cursor.Next(context.Background()){
	//		var result bson.D
	//		err := cursor.Decode(&result)
	//		if err != nil{
	//			log.Fatal(err)
	//		}
	//		log.Print(result.Map())
	//		log.Print(result.Map()["_id"])
	//	}
	//}

	{
		//cursor, err := coll.Find(
		//	context.Background(),
		//	bson.D{{"size.uom", "in"}},
		//)

		//cursor, err := coll.Find(
		//	context.Background(),
		//	bson.D{
		//		{"size.h", bson.D{
		//			{"$lt", 15},
		//		}},
		//	})

		cursor, err := coll.Find(
			context.Background(),
			bson.D{
				{"size.h", bson.D{
					{"$lt", 15},
				}},
				{"size.uom", "in"},
				{"status", "D"},
			})

		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var result bson.D
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(result.Map())
			log.Print(result.Map()["_id"])
		}
	}
}

func MongoQueryTop(coll *mongo.Collection) {
	//cursor, err := coll.Find(
	//	context.Background(),
	//	bson.D{},
	//)

	//cursor, err := coll.Find(
	//	context.Background(),
	//	bson.D{{"item", "canvas"}},
	//)

	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"qty", bson.D{{"$in", bson.A{25, 26}}}}})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(result)
		log.Print(result["_id"])
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

func MongoInsertMany(coll *mongo.Collection) {
	result, err := coll.InsertMany(
		context.Background(),
		[]interface{}{
			bson.D{
				{"item", "journal"},
				{"qty", int32(25)},
				{"tags", bson.A{"blank", "red"}},
				{"size", bson.D{
					{"h", 14},
					{"w", 21},
					{"uom", "cm"},
				}},
			},
			bson.D{
				{"_id", "im001"},
				{"int32", int32(25)},
				{"array", bson.A{'1', '2'}},
				{"doc", bson.D{}},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result.InsertedIDs)
}

func MongoConnect() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func MongoInsert(coll *mongo.Collection) {
	//clientOptions := options.Client().ApplyURI("mongodb://user:pwd@host:port")
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//client, err := mongo.Connect(context.Background(),clientOptions)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//coll := client.Database("ai-company").Collection("ygrsy")

	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"_id", "323"},
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton", "sasas"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"sizem", bson.E{"qwer", "asdf"}},
			bson.E{"qwer1", "asdf1"},
		})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
	log.Print(result.InsertedID)
}
