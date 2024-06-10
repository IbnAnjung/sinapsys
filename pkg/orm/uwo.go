package orm

import (
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Uow interface {
	Recovery(ctx context.Context)
	Begin(ctx context.Context) error
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	GetDB() *gorm.DB
}
