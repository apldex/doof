/*
 * @Author: Adrian Faisal
 * @Date: 08/10/21 9.15 PM
 */

package usecase

import (
	"github.com/apldex/doof/internal/pkg/model"
	"github.com/apldex/doof/internal/pkg/resource/db"
)

type Usecase interface {
	GetUserByID(id int) (*model.User, error)
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
