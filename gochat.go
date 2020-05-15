package main

import (
    "fmt"
    "os"
    "net"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "4444"
    CONN_TYPE = "tcp"
)

func server() {
    fmt.Println("calling server")


    listen, error := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if error != nil {
        fmt.Println("Error listening:", error.Error())
        os.Exit(1)    //TODO: modfiy this?
    }

    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    connection, error := listen.Accept()
    if error != nil {
        fmt.Println("Error accepting: ", error.Error())
        os.Exit(1)
    }
   // handle_request(connection)
    connection.Close() //should be done only AFTER req handled


    listen.Close()
    fmt.Println("at end of call to server")
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
        server()
    } else if actor == "client" {
        fmt.Println("is client")

    } else {
        fmt.Println("ERROR: arg 0 malformed; must be 'client' or 'server")
        fmt.Println("returning from main prematurely")
        return
    }

    fmt.Println("returning from main successfully")

}
