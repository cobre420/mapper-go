package domain

type Config struct {
	Services *[]Service `json:"services"`
}

type Service struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Prefix   string `json:"prefix"`
	UserName string `json:"user"`
	Password string `json:"password"`
}
