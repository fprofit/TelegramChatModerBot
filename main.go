package main

import(
	"time"
    "regexp"
    //"fmt"
)

var validText = regexp.MustCompile(`(https?:\/\/)?([\w-]{1,32}\.[a-zA-Z]{2,32})[^\s]*`)
var validUserName = regexp.MustCompile(`@[\w-]{1,32}`)

var admId int

func main(){

    for{
        if fileReadSettings(){
            break
        }else{
            LogToFile("files Setting")
            time.Sleep(30 * time.Second)
        }
    }
    fileReadMaps()
    go backUp()
    hand()	
}

func hand(){
    offset := 0
    for {
        updates := getUpdates(offset)
        for _, update := range updates {
            offset = update.UpdateId + 1

            //left user
            if update.Message.LeftChatMember.Id != 0 && update.Message.Chat.Id == groupId{
                go update.Message.LeftChatMember.leftMember()
                go delMessage(update.Message.Chat.Id, update.Message.MessageId)
                continue
            }
            //new users
            if len(update.Message.NewChatMembers) > 0 && update.Message.Chat.Id == groupId{
                for _, newChatMember := range update.Message.NewChatMembers{
                    newChatMember.UserDataToFile(0, 0, update.Message.Date)
                    go sendMessageLevel(newChatMember.Id, admIds[0])
                    go sendKeyImNotBot(update.Message.Chat.Id, newChatMember.Id, newChatMember.FirstName + " " + newChatMember.LastName)
                }
                go delMessage(update.Message.Chat.Id, update.Message.MessageId)
                continue
            }
            if admIdBool(update.Message.From.Id){
                go update.admComand()
                continue
            }
            if admIdBool(update.CallbackQuery.From.Id) {
                go update.CallbackQuery.admCallbackQuery()
                continue
            }
            if update.Message.Chat.Id == groupId || update.EditMessage.Chat.Id == groupId{
                go update.handMessage()
                continue
            }
            if update.CallbackQuery.Message.Chat.Id  == groupId && !admIdBool(update.CallbackQuery.From.Id)  {
                callbackQuery(update)
                continue
            }
        }
        //fmt.Println(updates)
    }
}

func (u *Update)handMessage(){
    if u.Message.From.Id != 0 {
        if u.Message.userBoolToMap(){
            u.Message.messageOne()
        }
    } 
    if u.EditMessage.From.Id != 0 {
        u.EditMessage.messageOne()
    }

    if MapUserData[u.Message.From.Id].Rating > 0{
        u.Message.sumMessageas()
    }
}

func (m *Message)sumMessageas(){
    if len(m.Text) > 3 && !m.findUrl(){
        m.From.UserDataToFile(MapUserData[m.From.Id].Rating, MapUserData[m.From.Id].SumMessageas + 1,MapUserData[m.From.Id].Date)
        if MapUserData[m.From.Id].SumMessageas == 1 {
            m.From.UserDataToFile(MapUserData[m.From.Id].Rating, MapUserData[m.From.Id].SumMessageas, m.Date)
        }
        if MapUserData[m.From.Id].Rating == 1{
            m.checkSumMesTime(420, 259200) //3, 45)
        }else if MapUserData[m.From.Id].Rating == 2{
            m.checkSumMesTime( 840, 518400)//6, 75)
        }else if MapUserData[m.From.Id].Rating == 3{
            m.checkSumMesTime( 1260, 777600) //9, 105)
        }
    }
    return
}

func (m *Message)checkSumMesTime(sumMessag int, date int) {
        if MapUserData[m.From.Id].SumMessageas > sumMessag {
            if  m.Date > MapUserData[m.From.Id].Date + date{
                m.From.UserDataToFile(MapUserData[m.From.Id].Rating + 1, MapUserData[m.From.Id].SumMessageas, MapUserData[m.From.Id].Date)
                go sendMessageLevel(m.From.Id, admIds[0])
                go m.From.sendMessageRating("повышен", m.Chat.Id)
            }
        }
}