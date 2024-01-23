package main

type System struct {
	Users map[string]*User
}

func (s *System) AddUser(id string, user *User) {
	if s.Users == nil {
		s.Users = make(map[string]*User, 0)
	}
	s.Users[id] = user
}

func (s *System) GetUser(id string) *User {
	return s.Users[id]
}
