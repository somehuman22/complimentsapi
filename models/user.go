package models

//User ... user model
type User struct {
	UserID   int    `json:"userId"`
	Fullname string `json:"fullName"`
	Username string `json:"username"`
	Password string
}

//GetUserByID ... fetches a user from the database
func GetUserByID(id int) (User, error) {
	var user User
	err = DB.QueryRow("select userid, fullname, username, password from users where userid = $1", id).Scan(&user.UserID, &user.Fullname, &user.Username, &user.Password)
	return user, err
}

//GetUserByUsername ... fetches a user from the database by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err = DB.QueryRow("select userid, fullname, username, password from users where username = $1", username).Scan(&user.UserID, &user.Fullname, &user.Username, &user.Password)
	return user, err
}
