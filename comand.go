package main

import(
	"encoding/json"
)

type MyCommands struct{
	Commands []BotCommand `json:"commands"`
	Scope BotCommandScope `json:"scope"`
}

type BotCommand struct {
	Command string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope struct{
	Type string     `json:"type"`
	ChatId int `json:"chat_id"`
    UserId int `json:"user_id"`
}

type DelCommands struct{
	Scope BotCommandScope `json:"scope"`
}

func setComnd (chatId int, userId int){
	
	var botCommand BotCommand
	botCommand.Command = "ban"
	botCommand.Description = "Ответьте этой командой на сообщение"
	
	var botCommandScope BotCommandScope
	botCommandScope.Type = "chat_member"
	botCommandScope.ChatId = chatId
	botCommandScope.UserId = userId

	var myCommands MyCommands
	myCommands.Commands  = append(myCommands.Commands, botCommand)
	myCommands.Scope = botCommandScope

	buf, err := json.Marshal(myCommands)
    if err != nil {
        LogToFile(err.Error())
    }

    postRequestGetResponse("/setMyCommands", buf)
}

func delComnd (chatId int, userId int){
	var delCommands DelCommands
	delCommands.Scope.Type = "chat_member"
	delCommands.Scope.ChatId = chatId
	delCommands.Scope.UserId = userId

	buf, err := json.Marshal(delCommands)
    if err != nil {
        LogToFile(err.Error())
    }

    postRequestGetResponse("/deleteMyCommands", buf)

	}