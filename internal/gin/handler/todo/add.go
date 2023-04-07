package todo

import (
	"bytes"
	"deploy-test/domain"
	"deploy-test/internal/gin/handler"
	"deploy-test/internal/provider"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type requestAdd struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	IsDone  *bool  `json:"isDone"`
}

func Add(c *gin.Context) {
	var req requestAdd
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.BindWith(&req, binding.FormMultipart); err != nil {
		zap.S().Infof("AddTodo - err: %v, req: %v", err, string(reqBody))
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	svc, err := provider.NewService()
	if err != nil {
		zap.S().Infof("AddTodo - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	err = svc.AddTodo(req.Title, req.Content, req.IsDone)
	if err != nil {
		zap.S().Infof("AddTodo - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	handler.Success(c, nil)
}
