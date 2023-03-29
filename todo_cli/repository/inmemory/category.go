package inmemory

import (
	"fmt"
	"todo_cli/entity"
)

type Category struct {
	categories []entity.Category
}

func NewCategoryStore() *Category {
	return &Category{
		[]entity.Category{},
	}
}

func (c *Category) UserIDAndCategoryIDValidation(userID, categoryID int) bool {

	isFound := false
	fmt.Println(c.categories)
	for _, c := range c.categories {
		fmt.Println(c.ID, categoryID, c.UserID, userID)
		if c.ID == categoryID && c.UserID == userID {
			isFound = true

			break
		}
	}

	return isFound
}

func (c *Category) CreateNewCategory(category entity.Category) (entity.Category, error) {

	category.ID = len(c.categories) + 1
	c.categories = append(c.categories, category)

	return category, nil

}
