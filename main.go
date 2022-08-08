package main

import (
	"context"

	"hexagonal/common/cache"
	"hexagonal/common/logs"
	"hexagonal/core/handlers"
	"hexagonal/core/repositories"
	"hexagonal/core/services"

	"hexagonal/middlewares"

	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "hexagonal/docs"

	swagger "github.com/swaggo/fiber-swagger"
)

func init() {
	initTime()
	initConfig()
}

func initDB_mongodb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongodb.uri")))
	if err != nil {
		panic(err)
	}
	// # Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client.Database(viper.GetString("mongodb.dbname"))
}

func initCache_redis() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Refs https://pkg.go.dev/github.com/go-redis/redis#ParseURL
	opt, err := redis.ParseURL(viper.GetString("redis.uri"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	// # Check the connection
	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return client
}

func initTime() {
	// # INITIAL TIME ZONE IN APPLICATION --------------------------
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// # Default value ////////////////////////////
	viper.SetDefault("app.port", 3000)
	viper.SetDefault("app.env", "production")
	// # //////////////////////////////////////////

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Hexagonal API
// @version 1.0.0
// @description เป็นตัวอย่างการใช้งาน Hexagonal Architecture ด้วย Go Lang
// @termsOfService http://somewhere.com/
// @host localhost:3000
// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https http
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	// # inti Global Techonology
	db := initDB_mongodb()
	ch := initCache_redis()

	// # init Global Adapter
	log := logs.NewAppLogs()
	cache := cache.NewAppCache(ch)

	// userRepository := repositories.NewUserRepositoryMock() // # Data Layer Mock
	userRepository := repositories.NewUserRepositoryDB(db)             // # Data Layer
	userService := services.NewUserService(log, cache, userRepository) // # Business Layer
	userHandler := handlers.NewUserHandler(userService)                // # Presentation Layer

	menuRepository := repositories.NewMenuRepositoryMock()
	// menuRepository := repositories.NewMenuRepositoryDB(db)      // # Data Layer
	menuService := services.NewMenuService(log, menuRepository) // # Business Layer
	menuHandler := handlers.NewMenuHandler(menuService)         // # Presentation Layer

	// # Handler APIs
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.FiberWrapHandler())

	app.Get("/api/v1/users", middlewares.ValToken, userHandler.GetUsers)
	app.Get("/api/v1/user/:userid/account", middlewares.ValToken, middlewares.ValPer([]int{1, 2, 3}), userHandler.GetUser)
	app.Post("/api/v1/signin", userHandler.SignIn)
	app.Post("/api/v1/signup", userHandler.SignUp)

	app.Get("/api/v1/menus", menuHandler.GetMenu)
	app.Post("/api/v1/menu", menuHandler.CreateMenu)

	app.Listen("localhost:3000")
}
