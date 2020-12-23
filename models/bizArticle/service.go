package bizArticle

import (
	"log"
	"service_post/db"
)

func CountByTagId(tagId int) (count int, err error) {
	dataSql := `
select count(1)
from biz_article_tags
where tag_id = ?
`
	if err := db.Db.QueryRow(dataSql, tagId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func CountByTypeId(typeId int) (count int, err error) {
	dataSql := `
select count(1)
from biz_article
where type_id = ?
`
	if err := db.Db.QueryRow(dataSql, typeId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticleList(req GetArticleListRequest) (list []ArticleModel, count int) {
	dataSql := `
select a.id,
       title,
       user_id,
       cover_image,
       qrcode_path,
       is_markdown,
       top,
       type_id,
       bt.name type_name,
       status,
       recommended,
       original,
       comment,
       a.create_time,
       a.update_time
from biz_article a
inner join biz_type bt on a.type_id = bt.id
where 1=1
`

	var params = make([]interface{}, 0)
	var filter string

	if len(req.Keyword) != 0 {
		filter += " and a.title like ? "
		params = append(params, "%"+req.Keyword+"%")
	}

	if req.TypeId > 0 {
		filter += " and a.type_id = ? "
		params = append(params, req.TypeId)
	}

	dataSql += filter + " order by a.id desc limit ?, ?;"
	countSql := `
select count(1) from biz_article a
inner join biz_type bt on a.type_id = bt.id
where 1=1
` + filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count biz_article err: ", err.Error())
	}

	offset, limit := req.GetLimit()
	params = append(params, offset)
	params = append(params, limit)

	if err := db.Db.Select(&list, dataSql, params...); err != nil {
		log.Panicln("select biz_article err: " + err.Error())
	}

	return
}
