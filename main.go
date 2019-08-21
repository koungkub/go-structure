package main

import (
	"os"
	"os/signal"

	"github.com/koungkub/go-structure/src/connection"
	"github.com/koungkub/go-structure/src/route"
	"github.com/koungkub/go-structure/src/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigName("env")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func main() {

	tracer, closer, _ := utils.GetGlobalTrace(viper.GetString("APP.NAME"))
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	log := connection.GetLogConnection()
	cache := connection.GetCacheConnection()
	db, _ := connection.GetDBConnection("mysql")

	router := route.GetRouter(log, cache, db)

	// Gracefull Shutdown
	go func() {
		if err := router.Start(viper.GetString("APP.PORT")); err != nil {
			log.Error(errors.WithMessage(err, "Graceful shutdown starting !!"))
		}
	}()

	graceful := make(chan os.Signal)
	signal.Notify(graceful, os.Interrupt)
	<-graceful

	ctx, cancel := utils.GetContext()
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		log.Fatal(errors.WithMessage(err, "Graceful shutdown timeout !!"))
	}
}
