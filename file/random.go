package main


import (
    "fmt"
    "math/rand"
    "time"
    "bytes"
)

func main() {
    fmt.Println(randomString(10))
}

func randomString(l int) string {
    var result bytes.Buffer
    var temp string

    for i := 0; i < l;  {
        if string(randInt(65, 90)) != temp {
            temp = string(randInt(65, 90))
            result.WriteString(temp)
            i++
        }
    }

    return result.String()
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max - min)
}


