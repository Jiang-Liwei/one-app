package link

import (
    "forum/pkg/app"
    "forum/pkg/database"
    "forum/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (link Link) {
    database.DB.Where("id", id).First(&link)
    return
}

func GetBy(field, value string) (link Link) {
    database.DB.Where("? = ?", field, value).First(&link)
    return
}

func All() (links []Link) {
    database.DB.Find(&links)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Link{}),
        &links,
        app.URL(database.TableName(&Link{})),
        perPage,
    )
    return
}