package config

type Settings struct {
	Port     int    `json:"port"`
	Postgres string `json:"postgres"`
}
