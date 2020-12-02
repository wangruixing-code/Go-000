package biz

import (
	"Week02/dao"
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

var errMsg = map[error]string{
	sql.ErrNoRows: "have no this record",
}

func AddData(id int, name string) string {
	var getID int
	var getName string
	var err error
	if getID, err = dao.QueryID(id); err != nil {
		log.Println(fmt.Sprintf("err with %+v", err))
		return errMsg[errors.Cause(err)]
	}
	if getName, err = dao.QueryName(name); err != nil {
		log.Println(fmt.Sprintf("err with %+v", err))
		return errMsg[errors.Cause(err)]
	}
	log.Println(fmt.Sprintf("%d, %s", getID, getName))
	return "ok"

}
