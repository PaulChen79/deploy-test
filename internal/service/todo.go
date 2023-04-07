package service

import (
	"deploy-test/domain"

	"go.uber.org/zap"
)

func (svc *service) AddTodo(title, content string, isDone *bool) error {
	domainTodo := &domain.Todo{
		Title:   title,
		Content: content,
		IsDone:  isDone,
	}

	if err := svc.repo.CreateTodo(domainTodo); err != nil {
		zap.S().Warnf("AddTodo - err %v\n", err)
		return err
	}

	return nil
}

func (svc *service) ListTodo(title string, isDone *bool) ([]*domain.Todo, error) {

	outputs, err := svc.repo.GetTodos(title, isDone)
	if err != nil {
		zap.S().Warnf("ListTodo - err %v\n", err)
		return nil, err
	}

	return outputs, nil
}

func (svc *service) GetTodo(id uint) (*domain.Todo, error) {

	output, err := svc.repo.GetTodo(id)
	if err != nil {
		zap.S().Warnf("GetTodo - err %v\n", err)
		return nil, err
	}

	return output, nil
}

func (svc *service) UpdateTodo(id uint, title, content string, isDone *bool) error {

	condition := &domain.Todo{
		ID: id,
	}

	instance := &domain.Todo{}

	if title != "" {
		instance.Title = title
	}

	if content != "" {
		instance.Content = content
	}

	if isDone != nil {
		instance.IsDone = isDone
	}

	err := svc.repo.UpdateTodo(condition, instance)
	if err != nil {
		zap.S().Warnf("UpdateTodo - err %v\n", err)
		return err
	}

	return nil
}

func (svc *service) DeleteTodo(id uint) error {
	err := svc.repo.DeleteTodo(id)
	if err != nil {
		zap.S().Warnf("DeleteTodo - err %v\n", err)
		return err
	}

	return nil
}
