package v2lstructs

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"someData"`
}

// DomainLevel 用于存储域名及其级别
type DomainLevel struct {
	Domain string
	Level  int
}

type Decrypt struct {
	Password   string
	Encryption string
}
