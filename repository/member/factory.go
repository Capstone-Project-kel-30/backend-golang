package member

import (
	"gym-app/business/member"
	"gym-app/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) member.MemberRepository {
	var membershipRepository member.MemberRepository

	if dbCon.Driver == util.PostgreSQL {
		membershipRepository = NewMemberPostgresRepo(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return membershipRepository
}
