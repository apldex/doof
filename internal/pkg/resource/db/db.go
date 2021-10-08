/*
 * @Author: Adrian Faisal
 * @Date: 08/10/21 8.46 PM
 */

package db

import (
	"database/sql"
	"fmt"
	"github.com/apldex/doof/internal/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

type Persistent interface {
	GetUserByID(id int) (*model.User, error)
}

type persistent struct {
	conn *sql.DB
}

func NewPersistent(dataSourceName string) (Persistent, error) {
	c, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[resource.db] open sql connection failed: %v", err)
	}

	err = c.Ping()
	if err != nil {
		return nil, fmt.Errorf("[resource.db] ping db failed: %v", err)
	}

	return &persistent{conn: c}, nil
}

func (p *persistent) GetUserByID(id int) (*model.User, error) {
	row := p.conn.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err := row.Err()
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("[resource.db.persistent] query get user error: %v", err)
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("[resource.db.persistent] user not found")
	}

	var user model.User
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("[resource.db.persistent] scan row to struct error: %v", err)
	}

	return &user, nil
}
