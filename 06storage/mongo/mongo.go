package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	containerPort = "27027/tcp"
)

var mongoURI string

const defaultMongoURI = "mongodb://localhost:27017"

type Trainer struct {
	Name string
	Age  int
	City string
}

// 新建客户端连接
func NewClient(c context.Context) (*mongo.Client, error) {
	cliOptions := options.Client().ApplyURI(defaultMongoURI)
	cli, err := mongo.Connect(context.TODO(), cliOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = cli.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to MongoDB!")
	return cli, err
}

// 关闭客户端连接
func DisClient(c context.Context, cli *mongo.Client) error {
	err := cli.Disconnect(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to mongodb closed")
	return err
}

func Insert(c context.Context, col *mongo.Collection) error {
	ash := Trainer{"Ash", 10, "wuhan"}
	misty := Trainer{"Misty", 20, "shanghai"}
	brock := Trainer{"Brock", 30, "shenzhen"}

	rst, err := col.InsertOne(c, ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("insert one doc to db success:", rst.InsertedID)

	trainers := []interface{}{misty, brock}
	manyRst, err := col.InsertMany(c, trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("insert many rst:", manyRst.InsertedIDs)
	return err
}

func Update(c context.Context, col *mongo.Collection) {
	filter := bson.D{
		{"name", "Ash"},
	}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	rst, err := col.UpdateOne(c, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update db rst:", rst)
}

func FindOne(c context.Context, col *mongo.Collection) {

	var rst Trainer
	filter := bson.D{{"name", "Ash"}}
	frst := col.FindOne(c, filter)

	err := frst.Decode(&rst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("find rst:%v\n", rst)
}

func FindMany(c context.Context, col *mongo.Collection) {
	var rst []*Trainer
	curr, err := col.Find(c, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for curr.Next(c) {
		var elem Trainer
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		rst = append(rst, &elem)
	}

	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	curr.Close(c)

	fmt.Printf("find many rst:%+v\n", rst)
	for _, v := range rst {
		fmt.Printf("%v\n", *v)
	}
}

func Delete(c context.Context, col *mongo.Collection) {
	//如果有多个name=ash的文档，则会删除多个
	rst, err := col.DeleteMany(c, bson.D{{"name", "Ash"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete ash:", rst.DeletedCount)

	//delete all
	rst, err = col.DeleteMany(c, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete many:", rst.DeletedCount)
}

func main() {

	ctx := context.Background()
	cli, _ := NewClient(ctx)
	defer DisClient(ctx, cli)

	col := cli.Database("mydb").Collection("trainers")

	//insert data to db
	Insert(ctx, col)

	Update(ctx, col)

	//find many
	FindOne(ctx, col)

	FindMany(ctx, col)

	Delete(ctx, col)
}
