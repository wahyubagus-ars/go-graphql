package resolvers

import model "golang-graphql-medium/cmd/app/domain/dao"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users []*model.User
	Posts []*model.Post
}
