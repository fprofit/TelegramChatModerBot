package main

import(
	"net/http"
	"encoding/json"
	"bytes"
	"strconv"
    "strings"
	"io/ioutil"
    "time" 
)

func forwMessage(chatId int, messageId int){
    var forMessage ForMessage
    forMessage.ChatId = 201014823
    forMessage.FromChatId = chatId
    forMessage.MessageId = messageId
    buf, err := json.Marshal(forMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/forwardMessage", buf)
    
}

func delMessage(delChatId int, delMesId int) {
    var botMessage BotDelMessage
    botMessage.ChatId = delChatId
    botMessage.MessageId = delMesId
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    
    postRequestGetResponse("/deleteMessage", buf)
}

func inlKeySet(text string, callData string)(InlineKeyboard){
    var inlKey InlineKeyboard
    inlKey.Text = text
    inlKey.CallData = callData
    return inlKey

}

func stringToHtml(str string) string{
    var str2 string
    for _, r := range []rune(str){
        if r == '<'{
            str2 = str2 + "&lt;"
            continue
        }else if r == '>' {
            str2 = str2 + "&gt;"
            continue
        }else if r == '&'{
            str2 = str2 + "&amp;"
            continue
        }else{
        str2 = str2 + string(r)
        }
    }
    return str2
}

func banUser(chatId int, userId int){
    var banChatMember BanChatMember
    banChatMember.ChatId = chatId
    banChatMember.UserId = userId
    banChatMember.RevokeMessages = true
    buf, err := json.Marshal(banChatMember)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/banChatMember", buf)
}

func unbanUser(chatId int, userId int){
    var unbanChatMember UnbanChatMember
    unbanChatMember.ChatId = chatId
    unbanChatMember.UserId = userId
    unbanChatMember.OnlyIfBanned = false
    buf, err := json.Marshal(unbanChatMember)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/unbanChatMember", buf)
}

func sendMessage(text string, chatId int){
    var botMessage BotSendMessage
    botMessage.ChatId = chatId
    botMessage.Text = text
    botMessage.ParseMode = "MarkdownV2"
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/sendMessage", buf)
}

func postRequestGetResponse(method string, buf []byte) PostResponse{
    var postResponse PostResponse
    for{
        resp, errCon := http.Post(botTelegramUrlToken + method,  "application/json", bytes.NewBuffer(buf))
        if errCon != nil {
            LogToFile(errCon.Error())
            time.Sleep(10 * time.Second)
            continue
        }
        defer resp.Body.Close()
        body, errBody := ioutil.ReadAll(resp.Body)
        if errBody != nil {
            LogToFile(errBody.Error())
        }
        if strings.EqualFold(string(body), "{\"ok\":true,\"result\":true}"){
            break
        }else {
            err := json.Unmarshal(body, &postResponse)
            if err != nil {
                LogToFile(err.Error())
            }
            break
        }
    }
    return postResponse    
}

func getUpdates(offset int) ([]Update) {
    for {    
        resp, err := http.Get(botTelegramUrlToken + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
        if err != nil {
            LogToFile(err.Error())
            time.Sleep(10 * time.Second)
            continue
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            LogToFile(err.Error())
        }
        var restResponse RestResponse
        err = json.Unmarshal(body, &restResponse)
        if err != nil {
            LogToFile(err.Error())
        }
        return restResponse.Result
    }
}