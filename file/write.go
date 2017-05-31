package main

import (
    "os"
    "fmt"
)

func main() {
    userFile := "test.txt"
    fout, err := os.Create(userFile)
    defer fout.Close()
    if err != nil {
        fmt.Println(userFile,err)
        return
    }

    for i:= 0; i < 10; i++ {
        fout.WriteString("Just a test!\r\n")
        fout.Write([]byte("Just a test!\r\n"))
    }
}