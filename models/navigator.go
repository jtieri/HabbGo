package models

type Navigator struct {
	categories map[int]Category
}

type Category struct {
	ID             int
	ParentID       int
	Name           string
	IsNode         bool
	IsPublic       bool
	IsTrading      bool
	MinRankAccess  Rank
	MinRankSetFlat Rank
}

func NewNavigator(categories []Category) Navigator {
	nav := Navigator{
		categories: make(map[int]Category),
	}

	for _, cat := range categories {
		nav.categories[cat.ID] = cat
	}

	return nav
}

func (n *Navigator) Categories() []Category {
	var categories []Category

	for _, category := range n.categories {
		categories = append(categories, category)
	}

	return categories
}

func (n *Navigator) CategoryByID(id int) (Category, bool) {
	category, ok := n.categories[id]
	return category, ok
}

func (n *Navigator) CategoriesByParentID(id int) []Category {
	categories := make([]Category, 0)

	for _, category := range n.categories {
		if category.ParentID == id {
			categories = append(categories, category)
		}
	}

	return categories
}
