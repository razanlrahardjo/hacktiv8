package migrations

import "github.com/razanlrahardjo/hacktiv8/app/model"

// ModelMigrations models to migrate
var ModelMigrations []interface{} = []interface{}{
	&model.Todo{},
	&model.Status{},
	&model.User{},
}
