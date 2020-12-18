package bizType

import (
	"errors"
	"log"
	"service_post/db"
)

func (t *BizType) ExistByName() bool {
	dataSql := `
select ifnull((select 1
               from biz_type
               where name = ?
               limit 1), 0) a;
`
	res := 0
	if err := db.Db.QueryRow(dataSql, t.Name).Scan(&res); err != nil {
		log.Panicln("BizType ExistByName err: ", err.Error())
	}
	return res > 0
}

func GetTypeList() (list []BizType) {
	sql := `
select id,
       pid,
       name,
       description,
       sort,
       icon,
       available,
       create_time,
       update_time
from biz_type
`
	if err := db.Db.Select(&list, sql); err != nil {
		log.Panicln("select biz_type err: " + err.Error())
	}
	return
}

func AddBizType(request AddBizTypeRequest) (err error) {
	bizType := BizType{Name: request.Name}
	if bizType.ExistByName() {
		err = errors.New("分类已存在")
		return
	}

	insertSql := `
insert into biz_type
    (pid, name, description, sort, icon, available, create_time)
    value
    (:pid, :name, :description, :sort, :icon, :available, now())
`
	if _, err := db.Db.NamedExec(insertSql, request); err != nil {
		return err
	}
	return nil
}
