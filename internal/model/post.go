package model

type PostModel struct {
	BaseModel

	Slug    string `gorm:"uniqueIndex" json:"slug"`
	Title   string `json:"title"`
	Summary string `json:"summary"`

	CategoryID uint          `gorm:"index" json:"category_id"`
	Category   CategoryModel `gorm:"foreignKey:CategoryID" json:"category"`

	Tags []TagModel `gorm:"many2many:post_tags" json:"tags"`

	IsDraft bool `json:"is_draft"`

	Content string `json:"content"`

	IPFSURL string `json:"ipfs_url"`

	CreateTXHash string `json:"create_tx_hash"`
	UpdateTXHash string `json:"update_tx_hash"`
}

func (PostModel) TableName() string {
	return "posts"
}
