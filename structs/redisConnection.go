package structs

type RedisConnection struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
