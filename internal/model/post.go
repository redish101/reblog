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

	IPFSURL     string `json:"ipfs_url"`

	OwnerAddress string `gorm:"index" json:"owner_address"`
	BlockNumber  uint   `json:"block_number"`
	CreateTXHash string `json:"create_tx_hash"`
	UpdateTXHash string `json:"update_tx_hash"`
}

type PostOnIPFS struct {
	Title    string   `json:"title"`
	Slug     string   `json:"slug"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Summary  string   `json:"summary"`
	Content  string   `json:"content"`
}

func (PostOnIPFS) FromPostModel(post *PostModel) *PostOnIPFS {
	var tagNames []string
	for _, tag := range post.Tags {
		tagNames = append(tagNames, tag.Name)
	}
	return &PostOnIPFS{
		Title:    post.Title,
		Slug:     post.Slug,
		Category: post.Category.Name,
		Tags:     tagNames,
		Summary:  post.Summary,
		Content:  post.Content,
	}
}

func (PostModel) TableName() string {
	return "posts"
}
