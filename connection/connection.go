package connection
import(
    "fmt"
    "net"
)

type Connect struct{
    conn *net.UDPConn
}
func (c *Connect)Close(){
    if c.conn == nil{
        return
    }
    c.conn.Close()
}
func (c *Connect)Read(cb func([]byte, error)){
    buf := make([]byte, 255)
    go func(){
        for{
            n, addr, err := c.conn.ReadFrom(buf)
            if err != nil{
                cb(nil, err)
                return
            }
            go func(data []byte){
                cb(data, nil)
            }(buf[:n])
        }
    }()
    return
}
func (c *Connect)Disconnect(){
    if c.conn == nil{
        return
    }
    c.conn.Close()
}
func (c *Connect)Send(address string, data []byte)error{
    udpaddr, err := net.ResolveUDPAddr("udp", address)
    if err != nil{
        return err
    }

    conn, err := net.DialUDP("udp", nil, udpaddr)
    if err != nil{
        return err
    }

    conn.Write(data)

    return nil
}

func Create(addr string)(*Connect, error){
    raddr, err := net.ResolveUDPAddr("udp", addr)
    if err != nil{
        fmt.Println(err)
        return nil, err
    }

    conn, err := net.ListenUDP("udp", raddr)
    if err != nil{
        fmt.Println(err)
        return nil, err
    }

    ret := Connect{
        conn : conn,
    }
    fmt.Println(conn)

    return &ret, nil
}

