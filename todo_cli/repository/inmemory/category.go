package inmemory

import "todo_cli/entity"

type Category struct {
	categories []entity.Category
}

func NewCategoryStore() *Category {
	return &Category{
		[]entity.Category{},
	}
}

func (c Category) UserIDAndCategoryIDValidation(userID, categoryID int) bool {

	isFound := false
	for _, c := range c.categories {
		if c.ID == categoryID && c.UserID == userID {
			isFound = true

			break
		}
	}

	return isFound
}

func (c Category) CreateNewCategory(category entity.Category) (entity.Category, error) {

	category.ID = len(c.categories) + 1
	c.categories = append(c.categories, category)

	return category, nil

}