package bizArticle

import (
	"log"
	"service_post/db"
	"service_post/models/bizTags"
	"service_post/models/bizType"
)

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

func GetById(id int) (article BizArticle) {
	dataSql := `
select id,
       title,
       user_id,
       cover_image,
       qrcode_path,
       is_markdown,
       content,
       content_md,
       top,
       type_id,
       status,
       recommended,
       original,
       description,
       keywords,
       comment,
       create_time,
       update_time
from biz_article
where id = ?
`
	if err := db.Db.Get(&article, dataSql, id); err != nil {
		log.Panicln("select biz_article err: " + err.Error())
	}
	return
}

func GetDetailById(id int) (detail ArticleDetail, err error) {
	article := GetById(id)
	detail = ArticleDetail{BizArticle: article}

	curType, typeErr := bizType.GetById(article.TypeId)
	if typeErr != nil {
		return ArticleDetail{}, typeErr
	}
	tags, tagErr := bizTags.GetByArticleId(id)
	if tagErr != nil {
		return ArticleDetail{}, tagErr
	}

	detail.BizType = curType
	detail.BizTags = tags
	return
}
