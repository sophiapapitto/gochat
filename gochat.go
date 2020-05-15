package main

import (
    "fmt"
    "os"
    "net"
    "bytes"
    "io"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "4444"
    CONN_TYPE = "tcp"
)

func handle_connection(conn net.Conn) {
    fmt.Println("handling connection")
   // status, err := bufio.NewReader(conn).ReadString('\n')

    var buf bytes.Buffer
    io.Copy(&buf, conn)
   // fmt.Println("total size:", buf.Len())
    fmt.Printf("buffer: %s\n", buf.String())


    fmt.Println("after status") 


    conn.Close()

}

func server() {
    fmt.Println("calling server")


    listen, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)    //TODO: modfiy this?
    }

    defer listen.Close()  //Does this ever actually happen? TODO: exit gracefully.
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handle_connection(conn)
    }
    
    listen.Close()
    fmt.Println("at end of call to server")
}

func client() {
    fmt.Println("calling client")

    conn, err := net.Dial(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("Error dialing:", err.Error())
        os.Exit(1)    //TODO: modfiy this?
    }
    fmt.Fprintf(conn, "first input\n")
  //  status, err := bufio.NewReader(conn).ReadString('\n')

    fmt.Println("at end of call to client")
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
        client()

    } else {
        fmt.Println("ERROR: arg 0 malformed; must be 'client' or 'server")
        fmt.Println("returning from main prematurely")
        return
    }

    fmt.Println("returning from main")

}
