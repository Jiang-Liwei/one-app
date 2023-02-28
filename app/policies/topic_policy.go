// Package policies 用户授权
package policies

import (
	"forum/app/models/topic"
	"forum/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.UserID(c) == _topic.UserID
}
