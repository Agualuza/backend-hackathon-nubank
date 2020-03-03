package conf

type Configuration struct {
	DB 		DB
	Server 	Server
}

type DB struct {
	Driver	 	string
	ConnQuery 	string
}

type Server struct {
	Port 	string
}