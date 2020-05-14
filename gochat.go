package main

import (
    "fmt"
    "os"
)

func server() {

}

func main() {
    fmt.Println("at beginning of main")

    args := os.Args[1:]
    //TODO throw errors differently? return
    if len(args) != 1 {
        fmt.Printf("ERROR: %d arguments given; should be %d\n", len(args), 1)
        fmt.Println("returning from main prematurely")
        return

    }


    actor := args[0]

    if actor == "server" {         
        fmt.Println("is server")
    } else if actor == "client" {
        fmt.Println("is client")

    } else {
        fmt.Println("ERROR: arg 0 malformed; must be 'client' or 'server")
        fmt.Println("returning from main prematurely")
        return
    }

    fmt.Println("returning from main successfully")

}
