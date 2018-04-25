package model

type Post struct {
	Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     int64  `json:"user_id"`
	User       *User
	CategoryID int64 `json:"category_id"`
	Category   *Category
	Tags       []*Tag `gorm:"many2many:post_tags;"`
}

func (p *Post) AfterFind() (err error) {
	p.CreatedAtTS = p.CreatedAt.Unix()
	p.UpdatedAtTS = p.UpdatedAt.Unix()
	return nil
}

func (p *Post) Create() error {
	return GetDB().Set("gorm:association_autocreate", false).Create(p).Error
}

func GetPostByID(id int64) (*Post, error) {
	post := &Post{}
	err := GetDB().Where("id = ?", id).Preload("User").Preload("Category").First(post).Error
	if err != nil {
		return nil, err
	}
	return post, err
}

func GetPosts(page int, pageSize int) ([]*Post, int, error) {
	posts := []*Post{}
	err := GetDB().Preload("User").Preload("Category").Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	var count int
	err = GetDB().Table("posts").Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return posts, count, nil
}
