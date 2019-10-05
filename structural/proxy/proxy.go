package proxy

import (
	"errors"
	"fmt"
)

type finder interface {
	find(id int) user
}

type user struct {
	id int
}

type userRecords []user

func (u userRecords) find(id int) (user, error) {
	for _, v := range u {
		if v.id == id {
			return v, nil
		}
	}
	return user{}, errors.New("user doesn't exist")
}

type userRecordsProxy struct {
	userDB    *userRecords
	cache     userRecords
	cacheSize int
}

func newUserRecordsProxy(db *userRecords, cacheSize int) userRecordsProxy {
	return userRecordsProxy{
		userDB:    db,
		cache:     make(userRecords, 0),
		cacheSize: cacheSize,
	}
}

func (u userRecordsProxy) String() string {
	return fmt.Sprintln("cache:", u.cache, "\ncache size:", len(u.cache))
}

func (u *userRecordsProxy) find(id int) (user, error) {
	if v, err := u.cache.find(id); err == nil {
		return v, err
	}

	v, err := u.userDB.find(id)
	if err == nil {
		if (len(u.cache) + 1) > u.cacheSize {
			u.cache = u.cache[1:]
		}
		u.cache = append(u.cache, v)
	}

	return v, err
}
