package main

import (
    "encoding/json"
    "os"
    "strconv"
    "time"
    "sync"
)

type UserData struct {
	Rating int 			`json:"rating"`
	SumMessageas int 	`json:"sum_messageas"`
	Date int 			`json:"date"`
    Username string     `json:"username"`
    FirstName string    `json:"first_name"`
    LastName string     `json:"last_name"`
    IsBot bool          `json:"is_bot"`
}

var (
    fileMutex sync.Mutex
    mapMutex sync.Mutex
    MapUserData = make(map[int]UserData)
)

func (u *User) leftMember(){
    _, ok := MapUserData[u.Id]
    if ok {
        delete(MapUserData, u.Id)
        fileWritingMaps()
    }
    return
}

func fileReadMaps() {
    fileMutex.Lock()
    defer fileMutex.Unlock()
    _, err := os.Stat(strconv.Itoa(groupId) + ".txt")
    if err != nil{
        _, err := os.Create(strconv.Itoa(groupId) + ".txt")
        if err != nil {
            LogToFile(err.Error())
            return
        }
    }else{
        file, err := os.ReadFile(strconv.Itoa(groupId) + ".txt")
        if err != nil {
            LogToFile(err.Error())
            return
        }
        if len(file) != 0{
            err = json.Unmarshal(file, &MapUserData)
            if err != nil {
                LogToFile(err.Error())
                return
            }
        }
    }
}

func fileWritingMaps (){
    if MapUserData == nil {
        return
    }
    buf, err := json.Marshal(MapUserData)
     if err != nil {
         LogToFile(err.Error())
    }

    fileMutex.Lock()
    defer fileMutex.Unlock()
    file, err := os.OpenFile(strconv.Itoa(groupId) + ".txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
    if err != nil {
        LogToFile(err.Error())
        return
    }
    defer file.Close()
    if _, err = file.Write(buf); err != nil {
        LogToFile(err.Error())
    }
}


func (user *User) UserDataToFile(rating, sumMessageas, date int){
    if user.Id == 0 {
        return
    }
    mapMutex.Lock()
    defer mapMutex.Unlock()
    var UserData UserData
	UserData.Rating = rating
    UserData.SumMessageas =  sumMessageas
    UserData.Date = date
    UserData.Username = user.Username
    UserData.FirstName = user.FirstName
    UserData.LastName = user.LastName
    UserData.IsBot = user.IsBot
    MapUserData[user.Id] = UserData
    fileWritingMaps()
    setRestrictChatMember(groupId, user.Id)
}

func backUp(){
    for{
        os.MkdirAll("backUp/", 0777)
        file, err := os.OpenFile("backUp/" + strconv.Itoa(groupId) + "_" + time.Now().Format("15.04_02.01.2006") + ".txt", os.O_CREATE|os.O_WRONLY, 0777)
        if err != nil {
            LogToFile(err.Error())
            time.Sleep(30 * time.Minute)
            continue
        }
        defer file.Close()
        if MapUserData == nil {
            time.Sleep(30 * time.Minute)
            continue
        }
        buf, err := json.Marshal(MapUserData)
         if err != nil {
            LogToFile(err.Error())
            time.Sleep(30 * time.Minute)
            continue
        }
        _, err = file.Write(buf)
        if err != nil {
            LogToFile(err.Error())
            time.Sleep(30 * time.Minute)
            continue
        }
        time.Sleep(6 * time.Hour)
    }
}