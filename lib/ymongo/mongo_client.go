package ymongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	//uri := "mongodb://localhost/argos?replicaSet=replset&authSource=admin"
	//mongodb://mongodb0.example.com:27017/admin
	uri := "mongodb://root:password@192.168.0.99:27017/test?&authSource=admin"
	//uri := "mongodb://root:password@192.168.0.99:27017/test"
	//if os.Getenv("DATABASE_URL") != "" {
	//	uri = os.Getenv("DATABASE_URL")
	//}

	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return client, nil
}
