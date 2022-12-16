package main

import (
	"database/sql"

	"Demo-Port-n-Adapter/config"
	"Demo-Port-n-Adapter/database"
	"Demo-Port-n-Adapter/server"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var mysql_db *sql.DB
var redis_cache *redis.Client

func init() {
	config.ConfigInit()
	mysql_db = database.NewDatabase().MySQLInit()
	redis_cache = database.NewDatabase().RedisInit()
}

func main() {

	defer mysql_db.Close()
	defer redis_cache.Close()

	app := fiber.New()

	server := server.NewServer(mysql_db, redis_cache)
	server.Router(app)

	app.Listen(viper.GetString("demo.port"))

}
