package gorm

import (
	"deploy-test/domain"
	model "deploy-test/repo/model"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (repo *repository) CreateTodo(todo *domain.Todo) error {
	db := repo.db

	modelTodo := model.TodoDomainToModel(todo)

	if err := db.Create(modelTodo).Error; err != nil {
		zap.S().Warnf("CreateTodo - err %v\n", errors.WithStack(err))
		return err
	}

	return nil
}

func (repo *repository) GetTodos(title string, isDone *bool) ([]*domain.Todo, error) {
	db := repo.db

	var (
		modelTodos []*model.Todo
		outputs    []*domain.Todo
	)

	if title != "" {
		db = db.Where("`title`=?", title)
	}

	if isDone != nil {
		db = db.Where("`is_done`=?", isDone)
	}

	if err := db.Find(modelTodos).Error; err != nil {
		zap.S().Warnf("GetTodos - err %v\n", errors.WithStack(err))
		return nil, err
	}

	for i := range modelTodos {
		outputs = append(outputs, model.TodoModelToDomain(modelTodos[i]))
	}

	return outputs, nil
}

func (repo *repository) GetTodo(id uint) (*domain.Todo, error) {
	db := repo.db

	var modelTodo *model.Todo

	if err := db.Where("`id`=?").First(modelTodo).Error; err != nil {
		zap.S().Warnf("GetTodo - err %v\n", errors.WithStack(err))
		return nil, err
	}

	output := model.TodoModelToDomain(modelTodo)

	return output, nil
}

func (repo *repository) UpdateTodo(condition *domain.Todo, instance *domain.Todo) error {
	db := repo.db

	if err := db.Where(condition).Updates(instance).Error; err != nil {
		zap.S().Warnf("UpdateTodo - err %v\n", errors.WithStack(err))
		return err
	}

	return nil
}

func (repo *repository) DeleteTodo(id uint) error {
	db := repo.db

	if err := db.Where("`id`=?", id).Delete(&model.Todo{}).Error; err != nil {
		zap.S().Warnf("DeleteTodo - err %v\n", errors.WithStack(err))
		return err
	}

	return nil
}
