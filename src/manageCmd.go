package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

var commands = map[string]bool{
	"/start":    true,
	"/stop":     true,
	"/menu":     true,
	"/userInfo": true,
	"/config":   true,
	"/botInfo":  true,
	"/help":     true,
}

func startCmd(u *tb.User, newMsg bool) {
	var msg string

	isUser, err := isUser(u.ID)
	if err != nil {
		log.Printf("Error checking if ID is bot user: %v", err)
	}

	started, err := isStartedUser(u.ID)
	if err != nil {
		log.Printf("Error checking if user is started: %v", err)
	}
	if !started {
		err = startUser(u.ID, true)
		if err != nil {
			log.Printf("Error starting user: %v", err)
		}
		if isUser {
			msg = restartMsg
		} else {
			err := addUser(u)
			if err != nil {
				log.Printf("Error adding user: %v", err)
			}
			msg = startMsg
		}
	} else {
		msg = alreadyStartedMsg
	}

	err = sendMsgWithMenu(u, msg, newMsg)
	if err != nil {
		log.Printf("Error sending message to started user: %v", err)
	}
}

func stopCmd(u *tb.User) {
	err := startUser(u.ID, false)
	if err != nil {
		log.Printf("Error starting user: %v", err)
	}
	err = sendMsgWithSpecificMenu(u, stopMsg, startMenu, false)
	if err != nil {
		log.Printf("Error sending message to stopped user: %v", err)
	}
}
