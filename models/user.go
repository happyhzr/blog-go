package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Insert() error {
	db := getDb()
	query := "INSERT INTO users(name, password) VALUES($1, $2) RETURNING id"
	err := db.QueryRow(query, u.Name, u.Password).Scan(&u.ID)
	return err
}

func getUser(query string, args ...interface{}) (*User, error) {
	user := &User{}
	db := getDb()
	err := db.QueryRow(query, args...).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}

func GetUserByName(name string) (*User, error) {
	query := "SELECT id, name, password FROM users WHERE name=$1"
	user, err := getUser(query, name)
	return user, err
}

func GetUserByID(id int) (*User, error) {
	query := "SELECT id, name, password FROM users WHERE id=$1"
	user, err := getUser(query, id)
	return user, err
}
