package helper

import (
	json "github.com/json-iterator/go"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func SetCache(ctx *fiber.Ctx, key string, data interface{}) (err error) {
	byte, err := json.Marshal(data)
	if err != nil {
		return
	}

	response := database.Redis.Set(ctx.Context(), key, byte, 0)
	if response.Err() != nil {
		err = response.Err()
		return
	}

	return
}

func GetCache(ctx *fiber.Ctx, key string) (response *redis.StringCmd) {
	response = database.Redis.Get(ctx.Context(), key)
	return
}
