package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

func QueryID(id int) (int, error) {
	if id < 0 {
		return 0, errors.Wrapf(sql.ErrNoRows, "no this id record: %+v", id)
	}
	return id, nil
}
func QueryName(name string) (string, error) {
	if name == "" {
		return "", errors.Wrapf(sql.ErrNoRows, "no this name record: %+v", name)
	}
	return "", nil
}
