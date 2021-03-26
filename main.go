package main

import (
	"fmt"

	module "github.com/anggriawanrilda88/myboilerplate/app"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/migration"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/middleware"
	configuration "github.com/anggriawanrilda88/myboilerplate/config"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/session/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/expvar"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

//App struct for app
type App struct {
	*fiber.App

	Hasher  hashing.Driver
	Session *session.Session
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	config := configuration.New()
	app := App{
		App:     fiber.New(*config.GetFiberConfig()),
		Hasher:  hashing.New(config.GetHasherConfig()),
		Session: session.New(config.GetSessionConfig()),
	}

	// Add middleware
	app.registerMiddlewares(config)

	// Connected with database
	db, err := database.New(&database.DbConfig{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	// Auto-migrate database models
	if err != nil {
		fmt.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			fmt.Println("failed to connect to database: db variable is nil")
		} else {
			err = migration.AutoMigratePostgres(db)
			if err != nil {
				fmt.Println("failed to automigrate role model:", err.Error())
				return
			}
		}
	}

	// Create a /api endpoint from appFiber module
	api := app.Group("/api")
	module.RegisterRoute(api)

	// Listen on port from env
	port := config.GetString("APP_ADDR")
	log.Fatal(app.Listen(port))
}

func (app *App) registerMiddlewares(config *configuration.Config) {
	// Middleware - Custom Access Logger based on zap
	if config.GetBool("MW_ACCESS_LOGGER_ENABLED") {
		app.Use(middleware.AccessLogger(&middleware.AccessLoggerConfig{
			Type:        config.GetString("MW_ACCESS_LOGGER_TYPE"),
			Environment: config.GetString("APP_ENV"),
			Filename:    config.GetString("MW_ACCESS_LOGGER_FILENAME"),
			MaxSize:     config.GetInt("MW_ACCESS_LOGGER_MAXSIZE"),
			MaxAge:      config.GetInt("MW_ACCESS_LOGGER_MAXAGE"),
			MaxBackups:  config.GetInt("MW_ACCESS_LOGGER_MAXBACKUPS"),
			LocalTime:   config.GetBool("MW_ACCESS_LOGGER_LOCALTIME"),
			Compress:    config.GetBool("MW_ACCESS_LOGGER_COMPRESS"),
		}))
	}

	// Middleware - Force HTTPS
	if config.GetBool("MW_FORCE_HTTPS_ENABLED") {
		app.Use(middleware.ForceHTTPS())
	}

	// Middleware - Force trailing slash
	if config.GetBool("MW_FORCE_TRAILING_SLASH_ENABLED") {
		app.Use(middleware.ForceTrailingSlash())
	}

	// Middleware - HSTS
	if config.GetBool("MW_HSTS_ENABLED") {
		app.Use(middleware.HSTS(&middleware.HSTSConfig{
			MaxAge:            config.GetInt("MW_HSTS_MAXAGE"),
			IncludeSubdomains: config.GetBool("MW_HSTS_INCLUDESUBDOMAINS"),
			Preload:           config.GetBool("MW_HSTS_PRELOAD"),
		}))
	}

	// Middleware - Suppress WWW
	if config.GetBool("MW_SUPPRESS_WWW_ENABLED") {
		app.Use(middleware.SuppressWWW())
	}

	// Middleware - Recover
	if config.GetBool("MW_FIBER_RECOVER_ENABLED") {
		app.Use(recover.New())
	}

	// TODO: Middleware - Basic Authentication

	// Middleware - Cache
	if config.GetBool("MW_FIBER_CACHE_ENABLED") {
		app.Use(cache.New(cache.Config{
			Expiration:   config.GetDuration("MW_FIBER_CACHE_EXPIRATION"),
			CacheControl: config.GetBool("MW_FIBER_CACHE_CACHECONTROL"),
		}))
	}

	// Middleware - Compress
	if config.GetBool("MW_FIBER_COMPRESS_ENABLED") {
		lvl := compress.Level(config.GetInt("MW_FIBER_COMPRESS_LEVEL"))
		app.Use(compress.New(compress.Config{
			Level: lvl,
		}))
	}

	// Middleware - CORS
	if config.GetBool("MW_FIBER_CORS_ENABLED") {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     config.GetString("MW_FIBER_CORS_ALLOWORIGINS"),
			AllowMethods:     config.GetString("MW_FIBER_CORS_ALLOWMETHODS"),
			AllowHeaders:     config.GetString("MW_FIBER_CORS_ALLOWHEADERS"),
			AllowCredentials: config.GetBool("MW_FIBER_CORS_ALLOWCREDENTIALS"),
			ExposeHeaders:    config.GetString("MW_FIBER_CORS_EXPOSEHEADERS"),
			MaxAge:           config.GetInt("MW_FIBER_CORS_MAXAGE"),
		}))
	}

	// Middleware - CSRF
	if config.GetBool("MW_FIBER_CSRF_ENABLED") {
		app.Use(csrf.New(csrf.Config{
			TokenLookup: config.GetString("MW_FIBER_CSRF_TOKENLOOKUP"),
			Cookie: &fiber.Cookie{
				Name:     config.GetString("MW_FIBER_CSRF_COOKIE_NAME"),
				SameSite: config.GetString("MW_FIBER_CSRF_COOKIE_SAMESITE"),
			},
			CookieExpires: config.GetDuration("MW_FIBER_CSRF_COOKIE_EXPIRES"),
			ContextKey:    config.GetString("MW_FIBER_CSRF_CONTEXTKEY"),
		}))
	}

	// Middleware - ETag
	if config.GetBool("MW_FIBER_ETAG_ENABLED") {
		app.Use(etag.New(etag.Config{
			Weak: config.GetBool("MW_FIBER_ETAG_WEAK"),
		}))
	}

	// Middleware - Expvar
	if config.GetBool("MW_FIBER_EXPVAR_ENABLED") {
		app.Use(expvar.New())
	}

	// Middleware - Favicon
	if config.GetBool("MW_FIBER_FAVICON_ENABLED") {
		app.Use(favicon.New(favicon.Config{
			File:         config.GetString("MW_FIBER_FAVICON_FILE"),
			CacheControl: config.GetString("MW_FIBER_FAVICON_CACHECONTROL"),
		}))
	}

	// TODO: Middleware - Filesystem

	// Middleware - Limiter
	if config.GetBool("MW_FIBER_LIMITER_ENABLED") {
		app.Use(limiter.New(limiter.Config{
			Max:      config.GetInt("MW_FIBER_LIMITER_MAX"),
			Duration: config.GetDuration("MW_FIBER_LIMITER_DURATION"),
			// TODO: Key
			// TODO: LimitReached
		}))
	}

	// Middleware - Monitor
	if config.GetBool("MW_FIBER_MONITOR_ENABLED") {
		app.Use(monitor.New())
	}

	// Middleware - Pprof
	if config.GetBool("MW_FIBER_PPROF_ENABLED") {
		app.Use(pprof.New())
	}

	// TODO: Middleware - Proxy

	// Middleware - RequestID
	if config.GetBool("MW_FIBER_REQUESTID_ENABLED") {
		app.Use(requestid.New(requestid.Config{
			Header: config.GetString("MW_FIBER_REQUESTID_HEADER"),
			// TODO: Generator
			ContextKey: config.GetString("MW_FIBER_REQUESTID_CONTEXTKEY"),
		}))
	}

	// TODO: Middleware - Timeout
}
