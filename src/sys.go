package main

import (
	"flag"
	"log"
	"os"
)

type flags struct {
	redisAddr string
	redisPwd  string
	redisDB   int
	botToken  string
	mediaPath string
}

var cmdFlags flags

func getFlags() error {
	const (
		defaultAddr      = "127.0.0.1:6379"
		addrUsage        = "The address of the redis instance"
		defaultPwd       = ""
		pwdUsage         = "The password of the redis instance"
		defaultDB        = 0
		dbUsage          = "The database to be selected after connecting to redis instance"
		defaultBotToken  = ""
		botTokenUsage    = "A bot token to be added to the set of tokens"
		defaultMediaPath = ""
		mediaPathUsage   = "A path to be used as media directory"
	)

	flag.StringVar(&(cmdFlags.redisAddr), "redisAddr", defaultAddr, addrUsage)
	flag.StringVar(&(cmdFlags.redisAddr), "a", defaultAddr, addrUsage+"(shorthand)")
	flag.StringVar(&(cmdFlags.redisPwd), "redisPwd", defaultPwd, pwdUsage)
	flag.StringVar(&(cmdFlags.redisPwd), "p", defaultPwd, pwdUsage+"(shorthand)")
	flag.IntVar(&(cmdFlags.redisDB), "redisDB", defaultDB, dbUsage)
	flag.IntVar(&(cmdFlags.redisDB), "d", defaultDB, dbUsage+"(shorthand)")
	flag.StringVar(&(cmdFlags.botToken), "botToken", defaultBotToken, botTokenUsage)
	flag.StringVar(&(cmdFlags.botToken), "t", defaultBotToken, botTokenUsage+"(shorthand")
	flag.StringVar(&(cmdFlags.mediaPath), "mediaPath", defaultMediaPath, mediaPathUsage)
	flag.StringVar(&(cmdFlags.mediaPath), "m", defaultMediaPath, mediaPathUsage+"(shorthand")

	flag.Parse()

	return nil
}

func exit() error {
	log.Printf("Stopping %s", bot.Me.Username)
	bot.Stop()
	log.Println("Bot stopped")

	log.Println("Closing redis instance")
	redisClient.Close()
	log.Println("Redis instance closed")

	log.Println("Exiting")
	os.Exit(0)

	return nil
}
