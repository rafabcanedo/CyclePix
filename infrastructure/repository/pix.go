package repository

import (
	"fmt"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/rafabcanedo/CyclePix/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) findKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("No key was found")
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDb) findAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("No account found")
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) findBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.Db.First(&bank, "id = ?", id)

	if bank.id == "" {
		return nil, fmt.Errorf("No bank found")
	}
	return &bank, nil
}
