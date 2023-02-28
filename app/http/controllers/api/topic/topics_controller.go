package controller

import (
	"forum/app/http/controllers/api"
	"forum/app/models/topic"
	"forum/app/requests"
	topicRequests "forum/app/requests/topic"
	"forum/pkg/auth"
	"forum/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	api.Controller
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := topicRequests.TopicRequest{}
	if ok := requests.Validate(c, &request, topicRequests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.UserID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
