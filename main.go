package main
import(
    "fmt"
    "net"
    "time"

    "udp_conn/connection/protocol"
)
func main(){

    data, err := protocol.SerializeUser("nice")
    data2, err := protocol.Deserialize(data)
    //protocol.Deserialize([]byte(`{"protocol":1,"data":{"name":"aaaa","test":{"test2":10,"test3":"test3","test4":[11,12],"test5":[{"test6":13},{"test6":14}]}}}`))

    fmt.Println(data, err)
    fmt.Printf("%+v\n", data2)


    var udpaddr *net.UDPAddr
    udpaddr, err = net.ResolveUDPAddr("udp", ":5454")
    if err != nil{
        fmt.Println(err)
        return
    }

    conn, err := net.ListenUDP("udp", udpaddr)
    if err != nil{
        fmt.Println(err)
        return
    }
    defer conn.Close()

    fmt.Println(conn.LocalAddr())

    buf := make([]byte, 255, 255)

    go func(){
        for{
            conn.Write([]byte("from server"))
            time.Sleep(time.Second)
        }
    }()

    for{
        n, remoteaddr, err := conn.ReadFrom(buf)
        if err != nil{
            fmt.Println(err)
            return
        }
        fmt.Println(n, remoteaddr, err)

        fmt.Println(string(buf[:n]))

        conn.WriteTo(buf[:n], remoteaddr)
    }
}

