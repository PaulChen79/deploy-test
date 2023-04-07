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

type requestList struct {
	Title  string `json:"title"`
	IsDone *bool  `json:"isDone"`
}

func List(c *gin.Context) {
	var req requestList
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.BindWith(&req, binding.FormMultipart); err != nil {
		zap.S().Infof("ListTodo - err: %v, req: %v", err, string(reqBody))
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	svc, err := provider.NewService()
	if err != nil {
		zap.S().Infof("ListTodo - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	todos, err := svc.ListTodo(req.Title, req.IsDone)
	if err != nil {
		zap.S().Infof("ListTodo - err: %v", err)
		handler.Failed(c, domain.ErrorServer, err.Error())
		return
	}

	handler.Success(c, todos)
}
