package repository

import (
	"fmt"

	"github.com/feliperromao/imersao-7-codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(t *model.Transaction) error {
	err := r.Db.Create(t).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) Save(t *model.Transaction) error {
	err := r.Db.Save(t).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var t model.Transaction
	r.Db.Preload("AccountFrom.Bank").First(&t, "id = ?", id)

	if t.ID == "" {
		return nil, fmt.Errorf("transaction not found")
	}

	return &t, nil
}