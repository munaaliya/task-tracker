package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	c.filebasedDb.StoreCategory(*Category)
	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	if err := c.filebasedDb.UpdateCategory(id, category); err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *categoryRepository) Delete(id int) error {
if err := c.filebasedDb.DeleteCategory(id); err != nil{
	return fmt.Errorf("not implement") // TODO: replace this
}
return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	category, err := c.filebasedDb.GetCategoryByID(id)

	return category, err
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	categories, err := c.filebasedDb.GetCategories()
	if err != nil{
		return nil, err
	}
	return categories, nil // TODO: replace this
}