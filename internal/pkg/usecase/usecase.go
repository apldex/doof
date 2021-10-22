/*
 * @Author: Adrian Faisal
 * @Date: 08/10/21 9.15 PM
 */

package usecase

import (
	"fmt"
	"github.com/apldex/doof/internal/pkg/model"
	"github.com/apldex/doof/internal/pkg/resource/db"
)

type Usecase interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(name, email string) error
}

type usecase struct {
	persistentDB db.Persistent
}

func New(persistentDB db.Persistent) Usecase {
	return &usecase{persistentDB: persistentDB}
}

func (u *usecase) GetUserByID(id int) (*model.User, error) {
	return u.persistentDB.GetUserByID(id)
}

func (u *usecase) CreateUser(name, email string) error {
	if len(name) < 3 {
		return fmt.Errorf("name should be more than 3 chars")
	}

	return u.persistentDB.CreateUser(name, email)
}
