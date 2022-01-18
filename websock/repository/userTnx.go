package repository

import "github.com/SarthakJha/distr-websock/models"

func (usr *UserRepository) QueryUserForEmail(email string) (models.User, error) {
	var user models.User
	err := usr.Table.Get("email", email).One(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (usr *UserRepository) QueryUserForName(name string) (models.User, error) {
	var user models.User
	err := usr.Table.Get("name", name).One(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (usr *UserRepository) QueryUserForUsername(username string) (models.User, error) {
	var user models.User
	err := usr.Table.Get("username", username).One(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (usr *UserRepository) QueryUserForID(id string) (models.User, error) {
	var user models.User
	err := usr.Table.Get("user_id", id).One(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
