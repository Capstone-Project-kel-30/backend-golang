package user

import (
	"gym-app/business/user"
	"gym-app/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.UserRepository {
	var userRepository user.UserRepository

	if dbCon.Driver == util.PostgreSQL {
		userRepository = NewPostgresRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
