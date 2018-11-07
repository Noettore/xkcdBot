package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	err := getFlags()
	if err != nil {
		log.Fatalf("Error in parsing command line flags: %v", err)
	}

	err = redisInit(cmdFlags.redisAddr, cmdFlags.redisPwd, cmdFlags.redisDB)
	if err != nil {
		log.Fatalf("Error in initializing redis instance: %v", err)
	}
	defer redisClient.Close()

	err = botInit(cmdFlags.botToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}
	defer bot.Stop()

	updateDBCron := cron.New()
	updateDBCron.AddFunc("0 0 * * * MON,WED,FRI", updateDB)
}
