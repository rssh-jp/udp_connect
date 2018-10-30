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
func (c *Connect)Read(cb func(data []byte))error{
    buf := make([]byte, 255)
    for{
        n, err := c.conn.Read(buf)
        if err != nil{
            return
        }
        cb(buf)
    }
    return nil
}

func Create(addr string)(*net.UDPConn, error){
    raddr, err := net.ResolveUDPAddr("udp", addr)
    if err != nil{
        fmt.Println(err)
        return, nil, err
    }

    conn, err := net.DialUDP("udp", nil, raddr)
    if err != nil{
        fmt.Println(err)
        return, nil, err
    }

}

