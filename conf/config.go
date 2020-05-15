package conf

const (
	USER     string = "root"
	PASSWORD string = "root"
	PROTOCOL string = "tcp"
	HOST     string = "localhost"
	PORT     string = "3306"
	DATABASE string = "gome"
)

const (
	Init       string = "./db/000-init.sql"
	CoursesSQL string = "./db/001-courses.sql"
	UsersSQL   string = "./db/002-users.sql"
)
