package jwt_prac2

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	users []*User
}

//consider prepopulating this
func NewUserStore() *UserStore {
	var us []*User

	return &UserStore{
		users: us,
	}
}

//userstore methods
func (us *UserStore) CreateUser(u *User) error {

	err := HashPw(u)

	if err != nil {
		return err
	}

	us.users = append(us.users, u)

	return nil
}

func (us *UserStore) GetUserByName(username string) (*User, error) {
	for _, u := range us.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("user doesn't exist")
}

//helper to hash pw
func HashPw(u *User) error {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hp)

	return nil
}
