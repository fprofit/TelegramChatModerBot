package main

import (
    "io/ioutil"
    "encoding/json"
)
type Settings struct {
	BotToken string `json:"botToken"`
	AdmIds []int `json:"admId"`
    GroupId int `json:"groupId"`
}
var (
	botTelegramUrlToken string
	admIds []int
    groupId int
)

func fileReadSettings () bool{
    buf, _ := ioutil.ReadFile("settings.txt")
    var settings Settings
    err := json.Unmarshal(buf, &settings)
    if err != nil {
        LogToFile(err.Error())
        return false

    }
    botTelegramUrlToken = "https://api.telegram.org/bot" + settings.BotToken
    admIds = settings.AdmIds
    groupId = settings.GroupId
    return true
}

func admIdBool(id int) bool{
    for _, admId := range admIds{
        if id == admId{
            return true
        }
    }
    return false
}