package entity

type CategoryDescriptions struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Meta_description string `json:"meta_description"`
	Meta_keyword     string `json:"meta_keyword"`
	Meta_title       string `json:"meta_title"`
}

type Category struct {
	Category_id          string                 `json:"category_id"`
	Category_description []CategoryDescriptions `json:"category_descriptions"`
}

func NewCategoryDescription(name string, description string, meta_description string, meta_keyword string, meta_title string) (*CategoryDescriptions, error) {
	categoryDescription := &CategoryDescriptions{
		Name:             name,
		Description:      description,
		Meta_description: meta_description,
		Meta_keyword:     meta_keyword,
		Meta_title:       meta_title,
	}
	return categoryDescription, nil
}
func NewCategory(category_id string, categoryDescription []CategoryDescriptions) (*Category, error) {
	category := &Category{
		Category_id:          category_id,
		Category_description: categoryDescription,
	}
	return category, nil

}
