package main

import(
    "encoding/json"
    "strconv"
    "unicode/utf16"
    "sort"  
)

func textUserMessage(id int) string{
    text := "UserId: " + strconv.Itoa(id) + " " 
    text = text + stringToHtml(MapUserData[id].FirstName)
    if MapUserData[id].LastName != ""{
        text = text + " " + stringToHtml(MapUserData[id].LastName)
    }
    if MapUserData[id].Username != ""{
        text = text + " @" +  MapUserData[id].Username
    }
    text = text + "\n🔝: " + strconv.Itoa(MapUserData[id].Rating) + " 💬: " + strconv.Itoa(MapUserData[id].SumMessageas)
    return text
}

func sendMessageLevel(id int, chatId int){
    var botMessage BotSendMessage
    botMessage.ChatId = chatId
    botMessage.Text = textUserMessage(id)
    botMessage.ParseMode = "HTML"
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/sendMessage", buf)
}

func (m *Message) forwMesOne(){
    id := m.From.Id
    var botMessage BotMessage
    botMessage.ChatId = admIds[0]
    botMessage.Text = m.Text
    botMessage.Text = botMessage.Text + "\n\n___________________________________\n" + "t.me/AlmetChat/"+ strconv.Itoa(m.MessageId)  + "\n" + textUserMessage(id)
    botMessage.ParseMode = "HTML"
    var inlineKey [][]InlineKeyboard
    var inTop []InlineKeyboard
    inTop = append(inTop, inlKeySet("Ban", "b" + strconv.Itoa(m.From.Id)), inlKeySet("Unban", "u" + strconv.Itoa(m.From.Id)))
    inlineKey = append(inlineKey, inTop)
    var inRow InlineKeyboardRow
    inRow.InlineKeyboardRow = inlineKey
    botMessage.Reply_Markup = inRow

    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/sendMessage", buf)    
}

func (c *CallbackQuery) admCallbackQuery(){
    userId, err := strconv.Atoi(c.Data[1:])
    _, ok := MapUserData[userId]
    if err != nil {
        LogToFile(err.Error())
    }
    if c.Data[0] == 'b' {
        banUser(groupId, userId)
        if ok {
            delete(MapUserData, userId)
            fileWritingMaps()
        }
    }
    if c.Data[0] == 'u' {
        if MapUserData[userId].Rating == 0 && ok{
            var u User
            u.Id = userId
            u.Username = MapUserData[userId].Username
            u.IsBot = MapUserData[userId].IsBot
            u.FirstName = MapUserData[userId].FirstName
            u.LastName = MapUserData[userId].LastName
            u.UserDataToFile(1,0,0)
            go u.sendMessageRating("востановлен", groupId)
        }
    }
    delMessage(c.Message.Chat.Id, c.Message.MessageId)
}

func (u *Update)admComand(){
    if u.Message.Text == "/r"{
        spisok(u.Message.Chat.Id)
        delMessage(u.Message.Chat.Id, u.Message.MessageId)
        return
    }
    if u.Message.Text == "/reset"{
        MapUserData = nil
        fileReadMaps()
        sendMessage("Данные перезагружены", u.Message.Chat.Id)
        delMessage(u.Message.Chat.Id, u.Message.MessageId)
        return
    }
    if u.Message.Text == "/clear"{
        for key, userData := range MapUserData{
            if userData.Rating == 0{
                delete(MapUserData, key)
            }else if userData.SumMessageas == 0 && userData.Date  + 432000 > u.Message.Date{
                delete(MapUserData, key)
            }
        }
        go fileWritingMaps()
        sendMessage("Таблица очищена", u.Message.Chat.Id)
        delMessage(u.Message.Chat.Id, u.Message.MessageId)
        return
    }
    if u.Message.ReplyToMessage.From.Id != 0{
        if u.Message.Text == "+" {
            u.Message.ReplyToMessage.From.UserDataToFile(MapUserData[u.Message.ReplyToMessage.From.Id].Rating + 1, MapUserData[u.Message.From.Id].SumMessageas, u.Message.Date)
            delMessage(u.Message.Chat.Id, u.Message.MessageId)
            go u.Message.ReplyToMessage.From.sendMessageRating("повышен", u.Message.Chat.Id)
            return
        }
        if u.Message.Text == "-" {
            u.Message.ReplyToMessage.From.UserDataToFile(MapUserData[u.Message.ReplyToMessage.From.Id].Rating - 1, 0 ,u.Message.Date)
            delMessage(u.Message.Chat.Id, u.Message.MessageId)
            go u.Message.ReplyToMessage.From.sendMessageRating("понижен", u.Message.Chat.Id)
            return
        }
    }
}

func spisok(chatId int){
    var text string
    var botMessage BotSendMessage
    botMessage.ChatId =  chatId
    botMessage.ParseMode = "HTML"

    var nameId= make(map[string]int)
    keysNameId := make([]string, 0, len(MapUserData))
    for key, userData := range MapUserData{
        nameId[userData.FirstName + userData.LastName + strconv.Itoa(key)] = key
        keysNameId = append(keysNameId, userData.FirstName + userData.LastName + strconv.Itoa(key))
    }
    sort.Strings(keysNameId)

    for _, k := range keysNameId{
        if MapUserData[nameId[k]].Rating > 6{
           continue
        }
        var name string
        name = MapUserData[nameId[k]].FirstName + " "
        if MapUserData[nameId[k]].LastName != ""{
            name = name + MapUserData[nameId[k]].LastName + " "
        }
        for i := len(utf16.Encode([]rune(name))); i < 23; i++{
            name = name + "."
        }
        text = name + "\t🔝: " + strconv.Itoa(MapUserData[nameId[k]].Rating) + "\t💬: " + strconv.Itoa(MapUserData[nameId[k]].SumMessageas)
        
        if len(utf16.Encode([]rune(botMessage.Text))) > 4000 {
            buf, err := json.Marshal(botMessage)
            if err != nil {
               LogToFile(err.Error())
            }
            postRequestGetResponse("/sendMessage", buf)
            botMessage.Text = ""
            continue
        }
        botMessage.Text = botMessage.Text +  "<pre><code class=\"language-golang\">" + stringToHtml(text) + "</code></pre>\n"

    }
    
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/sendMessage", buf)
}