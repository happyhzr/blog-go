package model

type Tag struct {
	Model
	Name string `json:"name"`
}

func (t *Tag) AfterFind() (err error) {
	t.CreatedAtTS = t.CreatedAt.Unix()
	t.UpdatedAtTS = t.UpdatedAt.Unix()
	return nil
}
