package configs

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pdnguyen1503/base-go/pkg/logging"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mongoOnce sync.Once
	// MongoDbName database name in mongodb
	MongoDbName    string
	connectTimeout time.Duration = 10 // sec
	mongoClient    *mongo.Client
)

type configDb struct {
	Host_Mysql     string `mapstructure:",HOST_MYSQL"`
	Port_Mysql     string `mapstructure:",PORT_MYSQL"`
	Username_Mysql string `mapstructure:",USERNAME_MYSQL"`
	Password_Mysql string `mapstructure:",PASSWORD_MYSQL"`
	Database_Mysql string `mapstructure:",DATABASE_MYSQL"`

	Port_Mongodb     string `mapstructure:",PORT_MONGODB"`
	Host_Mongodb     string `mapstructure:",HOST_MONGODB"`
	Username_Mongodb string `mapstructure:",USERNAME_MONGODB"`
	Password_Mongodb string `mapstructure:",PASSWORD_MONGODB"`
	Database_Mongodb string `mapstructure:",DATABASE_MONGODB"`
}

func LoadConfig(path string) (config configDb, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return configDb{}, err
	}

	err = viper.Unmarshal(&config)
	return
}

func ConnectMySql(config configDb) *gorm.DB {
	fmt.Println("config", config)
	connStrOptions := []string{
		"parseTime=True",
		"loc=Local",
	}
	// if tls {
	// 	connStrOptions = append(connStrOptions, "tls=true")
	// }

	options := strings.Join(connStrOptions, "&")
	protocol := "tcp"
	hostport := strings.TrimSpace(config.Host_Mysql)
	if config.Port_Mysql != "" {
		hostport = fmt.Sprintf("%v:%v", strings.TrimSpace(config.Host_Mysql), strings.TrimSpace(config.Port_Mysql))
	}
	mySQLServerStr := fmt.Sprintf(
		"%v:%v@%v(%v)/%v?%v",
		strings.TrimSpace(config.Username_Mysql),
		strings.TrimSpace(config.Password_Mysql),
		protocol,
		hostport,
		strings.TrimSpace(config.Database_Mysql),
		options,
	)

	fmt.Println(fmt.Printf("connecting to MySql... [%v]", mySQLServerStr))
	db, err := gorm.Open(mysql.Open(mySQLServerStr), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed connecting with mysql: %v", err))
	}
	fmt.Println(fmt.Sprintf("Connected with mysql: [%v]", mySQLServerStr))

	return db
}

func StartMongo(config configDb) {
	setMongoDbName(config)
	mongoOnce.Do(func() {
		ctx, cancelConnect := context.WithTimeout(context.Background(), connectTimeout*time.Second)
		defer cancelConnect()
		connString := fmt.Sprintf("mongodb://%v:%v", config.Host_Mongodb, config.Port_Mongodb)
		auth := options.Credential{
			Username: config.Username_Mongodb,
			Password: config.Password_Mongodb,
		}

		clientOptions := options.Client().SetAuth(auth)
		logging.Info("MongoDbName", MongoDbName)
		logging.Info("MongoDbName", clientOptions)
		// if strings.TrimSpace(cfg.ReplicaSetName) != "" {
		// 	clientOptions.SetReplicaSet(cfg.ReplicaSetName)
		// }
		clientOptions.ApplyURI(connString)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			panic(fmt.Sprintf("failed to connect mongo: %v", err))
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			panic(fmt.Sprintf("failed to ping mongo: %v", err))
		}

		logging.Info(fmt.Sprintf("connected to mongo: %v:%v (db: %v)", config.Host_Mongodb, config.Port_Mongodb, MongoDbName))
		mongoClient = client
	})
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func setMongoDbName(config configDb) {
	if dbFromEnv := os.Getenv("APP_MONGODB"); len(dbFromEnv) != 0 {
		MongoDbName = dbFromEnv
		return
	}

	if config.Database_Mongodb != "" {
		MongoDbName = config.Database_Mongodb
		return
	}

	panic("mongodb database name could notbe set")
}
