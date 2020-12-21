package bizArticle

import "service_post/db"

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
