package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v2"
)

func addUser(user *tb.User) error {
	err := redisClient.SAdd(usersID, user.ID).Err()
	if err != nil {
		log.Printf("Error in adding user ID: %v", err)
		return errors.Wrap(err, "adding userID in redisDB failed")
	}
	jsonUser, err := json.Marshal(&user)
	if err != nil {
		log.Printf("Error in marshalling user to json: %v", err)
		return errors.Wrap(err, "marshalling userInfo failed")
	}
	err = redisClient.HSet(usersInfo, strconv.Itoa(user.ID), jsonUser).Err()
	if err != nil {
		log.Printf("Error adding user info in hash: %v", err)
		return errors.Wrap(err, "adding userInfo in redisDB failed")
	}

	return nil
}

func isUser(userID int) (bool, error) {
	user, err := redisClient.SIsMember(usersID, strconv.Itoa(userID)).Result()
	if err != nil {
		log.Printf("Error checking if ID is bot user: %v", err)
		return false, errors.Wrap(err, "checking if userID is known failed")
	}
	return user, nil
}

func getUserInfo(userID int) (*tb.User, error) {
	user, err := redisClient.HGet(usersInfo, strconv.Itoa(userID)).Result()
	if err != nil {
		log.Printf("Error retriving user info from hash: %v", err)
		return nil, errors.Wrap(err, "retriving userInfo from redisDB failed")
	}
	jsonUser := &tb.User{}
	err = json.Unmarshal([]byte(user), jsonUser)
	if err != nil {
		log.Printf("Error unmarshalling user info: %v", err)
		return nil, errors.Wrap(err, "unmarshalling userInfo failed")
	}
	return jsonUser, nil
}

func isStartedUser(userID int) (bool, error) {
	started, err := redisClient.SIsMember(startedUsers, strconv.Itoa(userID)).Result()
	if err != nil {
		log.Printf("Error checking if user is started: %v", err)
		return false, errors.Wrap(err, "checking if userID is started failed")
	}
	return started, nil
}

func startUser(userID int, start bool) error {
	if start {
		err := redisClient.SAdd(startedUsers, strconv.Itoa(userID)).Err()
		if err != nil {
			log.Printf("Error adding token to set: %v", err)
			return errors.Wrap(err, "adding started userID to redisDB failed")
		}
	} else {
		err := redisClient.SRem(startedUsers, strconv.Itoa(userID)).Err()
		if err != nil {
			log.Printf("Error removing token from set: %v", err)
			return errors.Wrap(err, "removing stopped userID from redisDB failed")
		}
	}
	return nil
}

func getUserDescription(u *tb.User) (string, error) {
	msg := "\xF0\x9F\x91\xA4 *INFORMAZIONI UTENTE*" +
		"\n- *Nome*: " + u.FirstName +
		"\n- *Username*: " + u.Username +
		"\n- *ID*: " + strconv.Itoa(u.ID)

	return msg, nil
}
