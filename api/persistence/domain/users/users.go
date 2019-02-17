package users

import (
	"errors"
	"github.com/decebal/payments-api-fleet/api/auth"
	"strings"
	"time"
)

var Users = []User{
	User{ID: 1, Username: "cbrown", password: auth.HashAndSalt("cbrown123"), Created: time.Now(), Updated: time.Now()},
	User{ID: 2, Username: "bender", password: auth.HashAndSalt("bender123"), Created: time.Now(), Updated: time.Now()},
	User{ID: 3, Username: "goofy", password: auth.HashAndSalt("goofy123"), Created: time.Now(), Updated: time.Now()},
	User{ID: 4, Username: "ecartman", password: auth.HashAndSalt("ecartman123"), Created: time.Now(), Updated: time.Now()},
	User{ID: 5, Username: "sylvester", password: auth.HashAndSalt("sylvester123"), Created: time.Now(), Updated: time.Now()},
}

type User struct {
	ID       int
	Username string
	password string
	Created  time.Time
	Updated  time.Time
}

func getNewUserID() int {
	if len(Users) == 0 {
		return 1
	}

	return Users[len(Users)-1].ID + 1
}

func filter(test func(User) bool) (ret []User) {
	for _, u := range Users {
		if test(u) {
			ret = append(ret, u)
		}
	}

	return
}

// GetUserByID returns the user requested by the id, or returns an error
func GetUserByID(id int) (User, error) {
	var user User

	test := func(u User) bool { return u.ID == id }

	u := filter(test)

	if len(u) == 0 {
		return user, errors.New("user not found: " + string(id))
	}

	user = u[0]

	return user, nil
}

// GetAllUsers from the database and returns them in a list
func GetAllUsers() []User {
	return Users
}

// GetUserByUsername returns a single user if the username exists, otherwise an error
func GetUserByUsername(username string) (User, error) {
	var user User

	test := func(u User) bool {
		return u.Username == username
	}

	u := filter(test)

	if len(u) == 0 {
		return user, errors.New("user not found: " + username)
	}

	user = u[0]

	return user, nil
}

// RemoveUser takes an id and removes that user if found
func RemoveUser(id int) error {
	_, e := GetUserByID(id)

	if e != nil {
		return e
	}

	test := func(u User) bool { return u.ID != id }
	Users = filter(test)

	return nil
}

// AddUser to the main users table
func AddUser(username, password string) (User, error) {
	u := User{
		ID:       getNewUserID(),
		Username: username,
		password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	_, e := GetUserByUsername(strings.ToLower(u.Username))

	if e == nil {
		return User{}, errors.New("username already exists: " + username)
	}

	Users = append(Users, u)

	return u, nil
}
