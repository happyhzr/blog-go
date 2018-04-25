package model

type Category struct {
	Model
	Name string `json:"name"`
}

func (c *Category) AfterFind() error {
	c.CreatedAtTS = c.CreatedAt.Unix()
	c.UpdatedAtTS = c.UpdatedAt.Unix()
	return nil
}

func (c *Category) Create() error {
	return GetDB().Create(c).Error
}

func GetCategoryByID(id int64) (*Category, error) {
	c := &Category{}
	err := GetDB().Where("id = ?", id).First(c).Error
	return c, err
}
