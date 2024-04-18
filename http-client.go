package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlipResult struct {
    Slip struct {
        ID     int    `json:"id"`
        Advice string `json:"advice"`
    } `json:"slip"`
}

func main() {

    resp, err := http.Get("https://api.adviceslip.com/advice")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    result := ""
    for i := 0; scanner.Scan() && i < 5; i++ {
        result += scanner.Text()
        fmt.Println(result)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    slip := SlipResult{}
    json.Unmarshal([]byte(result), &slip)

    fmt.Println("The advice is \"" + slip.Slip.Advice + "\"")
}
