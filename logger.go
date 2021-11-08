package main
import(
    "time"
    "os"
    "runtime"
    "strconv"
    "log"
)

func LogToFile (error string){
    pc, fi, line, _ := runtime.Caller(1) 
    os.MkdirAll("logs", 0777)
    file, err := os.OpenFile("logs/" + time.Now().Format("02.01.2006") + ".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
    log.Println(runtime.FuncForPC(pc).Name() +"\t" + error + "\t" + strconv.Itoa(line) + " " + fi +"\n")
}

