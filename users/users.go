package users

import (
	"strconv"
	"task4/model"
)

const getUsers = "ref/users/"

func (p *Provider) GetAllUsers() ([]model.Users, error) {
	usersCache := p.cache.GetCache(getUsers)
	if usersCache == nil {
		users, err := p.storage.GetAllUsersFromDB()
		if err != nil {
			return nil, err
		}
		p.cache.AddToCache(getUsers, users)
		return users, nil
	}

	return usersCache.([]model.Users), nil
}

func (p *Provider) GetUsersWithMinAge(age int) ([]model.Users, error) {

	return p.storage.GetUsersWithMinAgeFromDB(age)
}

func (p *Provider) GetUserById(id int) (model.Users, error) {
	usersCache := p.cache.GetCache(getUsers + strconv.Itoa(id))

	if usersCache == nil {
		user, err := p.storage.GetUserByIdDB(id)
		if err != nil {
			p.cache.DeleteCache(getUsers + strconv.Itoa(id))
			return model.Users{}, err
		}
		p.cache.AddToCache(getUsers +  strconv.Itoa(id), user)

		return user, nil
	}

	return usersCache.(model.Users), nil
}

func (p *Provider) CreateUser(user model.Users) error {
	p.cache.ChangeActualStatus(getUsers)
	return p.storage.CreateUserDB(user)
}

func (p *Provider) UpdateUser(id int, user model.Users) error {
	p.cache.ChangeActualStatus(getUsers)
	p.cache.ChangeActualStatus(getUsers + strconv.Itoa(id))
	return p.storage.ChangeUserDB(id, user)
}

func (p *Provider) DeleteUser(id int) error {
	p.cache.ChangeActualStatus(getUsers)
	p.cache.ChangeActualStatus(getUsers + strconv.Itoa(id))
	return p.storage.DeleteUserDB(id)
}

func (p *Provider) GetAllCache() interface{}{
	return p.cache.GetAllCache()
}
