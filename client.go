package main
import(
    "fmt"
    "net"
)
func main(){
    udpaddr, err := net.ResolveUDPAddr("udp", "localhost:5454")
    if err != nil{
        fmt.Println(err)
        return
    }

    //laddr, err := net.ResolveUDPAddr("udp", "localhost:5455")
    //if err != nil{
    //    fmt.Println(err)
    //    return
    //}

    conn, err := net.DialUDP("udp", nil, udpaddr)
    if err != nil{
        fmt.Println(err)
        return
    }
    defer conn.Close()

    fmt.Println("Create connection")

    conn.Write([]byte("data"))
    buf := make([]byte, 255, 255)

    // 最初の読み込みではアクセス先を取得
    //buf := make([]byte, 255)
    //n, err := conn.Read(buf)
    //if err != nil{
    //    fmt.Println(err)
    //    return
    //}

    for{
        n, err := conn.Read(buf)
        fmt.Println(n, err)
        if err != nil{
            fmt.Println(err)
            return
        }

        fmt.Println(string(buf[:n]))
    }
}
