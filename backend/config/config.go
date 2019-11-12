package config

import "os"
var (
	User     string
	Pwd      string
	Database string
	Host     string
	//@Env value　dev or prod
	Env = "dev"
)

func DbConfig() {
	if Env == "dev" {
		Host = "localhost"
		User = "root"
		Pwd = "123456"
		Database = "config_manager"
	}
	if Env == "prod" {
		Host = os.Getenv("BLOG_MYSQL_HOST")
		User = os.Getenv("BLOG_MYSQL_USER")
		Pwd = os.Getenv("BLOG_MYSQL_PWD")
		Database = os.Getenv("BLOG_MYSQL_DATABASE")
	}
}
