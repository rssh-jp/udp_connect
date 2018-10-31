package protocol

import(
    "encoding/json"
)

type Connect struct{
}

func DeserializeConnect(data []byte)(Connect, error){
    var ret Connect 
    err := json.Unmarshal(data, &ret)
    if err != nil{
        return ret, err
    }
    return ret, nil
}

func SerializeConnect()([]byte, error){
    data := Connect{
    }
    return serialize(data)
}

type AccessPoint struct{
    Address string  `json:"address"`
}

func DeserializeAccessPoint(data []byte)(AccessPoint, error){
    var ret AccessPoint 
    err := json.Unmarshal(data, &ret)
    if err != nil{
        return ret, err
    }
    return ret, nil
}

func SerializeAccessPoint(address string)([]byte, error){
    data := AccessPoint{
        Address : address,
    }
    return serialize(data)
}

type User struct{
    Name string `json:"name"`
}

func DeserializeUser(data []byte)(User, error){
    var ret User
    err := json.Unmarshal(data, &ret)
    if err != nil{
        return ret, err
    }
    return ret, nil
}

func SerializeUser(name string)([]byte, error){
    data := User{
        Name : name,
    }
    return serialize(data)
}

type Message struct{
    Message string `json:"message"`
}

func DeserializeMessage(src []byte)(Message, error){
    var ret Message
    err := json.Unmarshal(src, &ret)
    if err != nil{
        return ret, err
    }
    return ret, nil
}

func SerializeMessage(message string)([]byte, error){
    data := Message{
        Message : message,
    }
    return serialize(data)
}


func serialize(src interface{})([]byte, error){
    var proto protocol
    switch src.(type){
    case User:
        proto.Protocol = AppUser
    case Message:
        proto.Protocol = AppMessage
    }
    proto.Data = src

    dest, err := json.Marshal(proto)
    if err != nil{
        return nil, err
    }

    return dest, nil
}
