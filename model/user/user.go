package user

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
}

var Table = "user"
