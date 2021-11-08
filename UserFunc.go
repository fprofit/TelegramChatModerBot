package main

import (
	"encoding/json"
    "strconv"
    "time"
)

// message text up/down level
func ratingText (rating int) string{
    switch rating{
    case 1:
        return "•Только текст"
    case 2:
        return "•Текст и ссылки\n•Фото\n•Видео\n•Голосовые и видеосообщения"
    case 3:
        return "•Текст и ссылки\n•Фото\n•Видео\n•Голосовые и видеосообщения\n•Стикеры\n•GIF'ки"
    case 4:
        return "•Текст и ссылки\n•Фото\n•Видео\n•Голосовые и видеосообщения\n•Стикеры\n•GIF'ки\n•Опросы" 
    default:
        return ""
    }
}

func (user *User)sendMessageRating(text string, chatId int){
    var botMessage BotSendMessage
    botMessage.ChatId = chatId
    text =  "<a href=\"tg://user?id=" + strconv.Itoa(user.Id) + "\">" + stringToHtml(user.FirstName +  " " + user.LastName) + "</a> ваш ретинг " + text + "\nСейча он равен " + strconv.Itoa(MapUserData[user.Id].Rating)
    if ratingText(MapUserData[user.Id].Rating) != ""{
   		text = text + "\nСейчас ты можешь отправлять\n" + ratingText(MapUserData[user.Id].Rating)
   	}
    botMessage.Text = text
    botMessage.ParseMode = "HTML"
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postResponse := postRequestGetResponse("/sendMessage", buf)
    time.Sleep(5 * time.Minute)
    delMessage(groupId, postResponse.Result.PostMessageId)
}

func callbackQuery(update Update){
	var answerCallbackQuery AnswerCallbackQuery 
    answerCallbackQuery.CallbackQueryId = update.CallbackQuery.Id
    if strconv.Itoa(update.CallbackQuery.From.Id) == update.CallbackQuery.Data {
        update.CallbackQuery.From.UserDataToFile(1,0,0)
        go setRestrictChatMember(update.CallbackQuery.Message.Chat.Id, update.CallbackQuery.From.Id)
        answerCallbackQuery.Text = "Твой рейтинг повышен!\nПока ты можешь отправлять только текстовые сообщения.\nБудь активнее и ты сможешь отправлять\n•Ссылки\n•Фото\n•Видео\n•Голосовые и видеосообщения\n•Стикеры\n•GIF'ки"
    	answerCallbackQuery.ShowAlert = true
        go sendMessageLevel(update.CallbackQuery.From.Id, admIds[0])
        go delMessage(update.CallbackQuery.Message.Chat.Id, update.CallbackQuery.Message.MessageId)
    }
    if strconv.Itoa(update.CallbackQuery.From.Id) != update.CallbackQuery.Data{
        answerCallbackQuery.Text = "Эта кнопка не для тебя🔒"
    	answerCallbackQuery.ShowAlert = false
    }
    buf, err := json.Marshal(answerCallbackQuery)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/answerCallbackQuery", buf) 
}

func (m *Message) userBoolToMap() bool{
    _, ok := MapUserData[m.From.Id]
    if !ok && !admIdBool(m.From.Id){
        m.From.UserDataToFile(0,0,0)
        go sendKeyImNotBot(m.Chat.Id, m.From.Id, m.From.FirstName + " " + m.From.LastName)
        go delMessage(m.Chat.Id, m.MessageId)
        return false
    }
    return true
}


func sendKeyImNotBot(idGroup int, idUser int, name string){
    var botMessage BotMessage
    botMessage.ChatId = idGroup
    botMessage.Text = "<a href=\"tg://user?id=" + strconv.Itoa(idUser) + "\">" + stringToHtml(name) + "</a> ваш ретинг равен 0\nВы пока не можете отправлять сообщения\nЧтобы отправлять сообщения и повысить свой рейтинг нажмите кнопку \"Я не бот\""
    botMessage.ParseMode = "HTML"

    var inlineKey [][]InlineKeyboard
    
    var inTop []InlineKeyboard
    inTop = append(inTop, inlKeySet("Я не бот", strconv.Itoa(idUser)))
    inlineKey = append(inlineKey, inTop)

    var inRow InlineKeyboardRow
    inRow.InlineKeyboardRow = inlineKey

    botMessage.Reply_Markup = inRow
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postResponse := postRequestGetResponse("/sendMessage", buf)

    //dell
    time.Sleep(15 * time.Minute)
    if MapUserData[idUser].Rating < 1 {
    delMessage(idGroup, postResponse.Result.PostMessageId)
    unbanUser(idGroup, idUser)
    }

}

func setChatPermissions(userId int) ChatPermissions{
    var chatPermissions ChatPermissions
    if MapUserData[userId].Rating >= 0 {
        chatPermissions.Can_send_messages = false
        chatPermissions.Can_send_media_messages = false
        chatPermissions.Can_send_other_messages = false
        chatPermissions.Can_send_polls = false
        chatPermissions.Can_add_web_page_previews = false
        chatPermissions.Can_invite_users = false
        chatPermissions.Can_pin_messages = false
    }
    if MapUserData[userId].Rating >= 1 {
        chatPermissions.Can_send_messages = true
    }
    if MapUserData[userId].Rating >= 2 {
        chatPermissions.Can_send_media_messages = true
    }
    if MapUserData[userId].Rating >= 3 {
        chatPermissions.Can_send_other_messages = true
    }
    if MapUserData[userId].Rating >= 4 {
        chatPermissions.Can_send_polls = true
    }
    if MapUserData[userId].Rating >= 5 {
        chatPermissions.Can_add_web_page_previews = true
    }
    if MapUserData[userId].Rating >= 6 {
        chatPermissions.Can_invite_users = true
    }
    return chatPermissions
}

func reSetRestrictChatMember(){
    for key, _ := range MapUserData{
            setRestrictChatMember(groupId, key)
    }
}

func setRestrictChatMember (restChatId int, restUserId int){
    if MapUserData[restUserId].Rating > 6 {
        return
    //}
    // if UserIdRating[restUserId] >= 4 {
    // setComnd(restChatId, restUserId)
    }else{
        delComnd(restChatId, restUserId)
    }
    var restrictChatMember RestrictChatMember 
    restrictChatMember.Chat_id = restChatId
    restrictChatMember.User_id = restUserId
    restrictChatMember.Permissions = setChatPermissions(restUserId)
    buf, err := json.Marshal(restrictChatMember)
    if err != nil {
        LogToFile(err.Error())
    }
    postRequestGetResponse("/restrictChatMember", buf)
}

func (m *Message) linkTrueBlock(){
    go m.forwMesOne()
    delMessage(m.Chat.Id, m.MessageId)
    m.From.UserDataToFile(0,0,0)
    var botMessage BotSendMessage
    botMessage.ChatId = m.Chat.Id
    botMessage.Text = "<a href=\"tg://user?id=" + strconv.Itoa(m.From.Id) + "\">" + stringToHtml(m.From.FirstName +  " " + m.From.LastName) + "</a> ваше сообщение подозрительно, поэтому оно удалено и будет проверено модератором\n\nА пока ты не cможешь отправлять сообщения"
    botMessage.ParseMode = "HTML"
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err.Error())
    }
    postResponse := postRequestGetResponse("/sendMessage", buf)
    time.Sleep(90 * time.Second)
    delMessage(m.Chat.Id, postResponse.Result.PostMessageId)
}