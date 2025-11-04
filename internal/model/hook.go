package model

import (
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/rid"
	"github.com/onexstack/onexstack/pkg/authn"
	"gorm.io/gorm"
)

// AfterCreate 在数据库创建记录完成后新增资源 ID
func (p *Post) AfterCreate(tx *gorm.DB) error {
	p.PostID = rid.PostID.New(uint64(p.ID))
	return tx.Save(p).Error
}

// AfterCreate 在数据库创建记录完成后新增资源 ID
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.UserID = rid.UserID.New(uint64(u.ID))
	return tx.Save(u).Error
}

// BeforeCreate 在新增数据库记录前对密钥进行加密
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = authn.Encrypt(u.Password)
	if err != nil {
		return err
	}
	return nil
}
