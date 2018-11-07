package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const (
	usersID        = "usersID"
	usersInfo      = "usersInfo"
	startedUsers   = "startedUsers"
	lastMsgPerUser = "lastMsgPerUser"
	stripIDs       = "stripIDs"
	stripTitles    = "stripTitles"
	stripPaths     = "stripPaths"
	mediaPath      = "mediaPath"
)

var redisClient *redis.Client

func redisInit(addr string, pwd string, db int) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	err := redisClient.Ping().Err()
	if err != nil {
		log.Printf("Error in connecting to redis instance: %v", err)
		return errors.Wrap(err, "redis connection failed")
	}
	return nil
}
