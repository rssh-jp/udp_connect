package protocol

type protocol struct {
	Protocol int         `json:"protocol"`
	Data     interface{} `json:"data"`
}

type Connect struct {
}

type AccessPoint struct {
	Address string `json:"address"`
}

type User struct {
	Name string `json:"name"`
}

type Message struct {
	Message string `json:"message"`
}
