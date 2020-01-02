package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// All lib you need to build this will define here. To keep main function always clean and clear.
type ServiceContext struct {
	Config      *viper.Viper
	RedisClient *redis.Client
	Mongo       *mongo.Client
	Context     context.Context
	Logger      *logrus.Logger
}

// Create service context where centralize  all third party libs
func InitServiceContext() *ServiceContext {
	var serviceContext ServiceContext
	var err error
	//context
	serviceContext.Context = context.Background()
	// Load file configure
	serviceContext.Config = viper.New()
	serviceContext.Config.SetConfigName("local")
	serviceContext.Config.AddConfigPath("./config")
	if err := serviceContext.Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//Logger
	serviceContext.Logger = logrus.New()
	serviceContext.Logger.SetFormatter(&logrus.JSONFormatter{})

	//Mongo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	serviceContext.Mongo, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("Failed to init mongo: %s\n", err.Error()))
	}

	// Redis
	serviceContext.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := serviceContext.RedisClient.Ping().Result()
	fmt.Println(pong, err)
	if serviceContext.RedisClient == nil {
		panic("Failed to init Redis service")
	}

	return &serviceContext
}
