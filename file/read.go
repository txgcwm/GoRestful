package main


import (
    "os"
    "fmt"
)

func main() {
    userFile := "test.txt"
    fin, err := os.Open(userFile)
    defer fin.Close()
    if err != nil {
        fmt.Println(userFile, err)
        return
    }

    buf := make([]byte, 1024)
    for {
        n, _ := fin.Read(buf)
        if 0 == n {
            break
        }

        os.Stdout.Write(buf[:n])
    }

    err = os.Remove(userFile)
}


// http://blog.csdn.net/yatere/article/details/8025371