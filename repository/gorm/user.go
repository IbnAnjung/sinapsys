package repository

import (
	"context"
	"errors"

	"github.com/IbnAnjung/synapsis/entity"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/IbnAnjung/synapsis/pkg/orm"

	"github.com/IbnAnjung/synapsis/repository/gorm/model"

	"gorm.io/gorm"
)

type userRepository struct {
	uow orm.Uow
}

func NewGormUserRepository(
	uow orm.Uow,
) entity.UserRepository {
	return &userRepository{
		uow,
	}
}

func (r *userRepository) FindUserByPhoneNumber(ctx context.Context, phone_number string) (u entity.User, err error) {
	m := model.MUser{}
	db := r.uow.GetDB().WithContext(ctx)

	if err = db.WithContext(ctx).Where("phone_number = ?", phone_number).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "user not found")
		}

		return
	}

	u = m.ToEntity()

	return
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) (err error) {
	m := model.MUser{}
	m.FillFromEntity(*u)

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			e = coreerror.NewCoreError(coreerror.CoreErrorDuplicate, "user_id already registered")
		}

		err = e
		return
	}

	u.ID = m.ID

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id string) (user entity.User, err error) {
	m := model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "user tidak valid")
		}
		err = e
		return
	}

	user = m.ToEntity()
	return
}

func (r *userRepository) FindByIds(ctx context.Context, ids []string) (user []entity.User, err error) {
	m := []model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("id in (?)", ids).Find(&m).Error; err != nil {
		return
	}

	for _, v := range m {
		user = append(user, v.ToEntity())
	}

	return
}

func (r *userRepository) FindUsers(ctx context.Context, gender uint8, excludeUserIds []string) (user entity.User, err error) {
	m := model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("gender = ?", gender).
		Not(map[string]interface{}{"id": excludeUserIds}).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "no profile available")
		}
		return
	}

	user = m.ToEntity()
	return
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) (err error) {
	m := model.MUser{}
	m.FillFromEntity(*user)

	return r.uow.GetDB().WithContext(ctx).Updates(&m).Error
}
