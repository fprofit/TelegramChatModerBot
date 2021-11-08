package main

import (
    "strings"
)

func (m *Message) messageOne(){
    if MapUserData[m.From.Id].Rating <= 1 {
        if m.findUrl(){
            go m.linkTrueBlock()
            return
        }
        if MapUserData[m.From.Id].SumMessageas <= 3{
            m.forwMesOne()
            return
        }
    }
    if MapUserData[m.From.Id].Rating == 2  {
        if m.findUrl(){
            m.forwMesOne()
            //delMessage(m.Chat.Id, m.MessageId)
            return
        }
    }
}

// find url & @
func (m *Message) findUrl() bool{
    if validText.MatchString(m.Text) || validText.MatchString(m.Caption) {
        return true
    }
    if len(m.Entities) > 0{
        return urlEntities(m.Text,  m.Entities)
    }
    if len(m.CaptionEntities) > 0{
        return urlEntities(m.Caption,  m.CaptionEntities)
    }
    if validUserName.MatchString(m.Text) || validUserName.MatchString(m.Caption) {
       return true
    }
    return false
}

// find url & username v MessageEntition
func urlEntities(text string, e []Entities) bool{
    for _, entitie := range e {
        if entitie.Type == "text_link" {
            return true
        }
        if entitie.Type == "text_mention"{
             _, ok := MapUserData[entitie.User.Id]
            if ok {
                return false
            }
        }
        if entitie.Type == "mention" {
            r := []rune(text)
            mention := string(r[entitie.Offset + 1:entitie.Offset + entitie.Length])
            for _, userData := range MapUserData {
                if strings.EqualFold(userData.Username, mention){
                    return false
                }
            }
        }
    }
    return true
}