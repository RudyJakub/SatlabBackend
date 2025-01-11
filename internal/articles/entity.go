package articles

type Image struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	UploadedAt    string `json:"uploaded_at"`
	ImageLocation string `json:"image_location"`
}

type Article struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Public      bool    `json:"public"`
	Images      []Image `json:"images"`
}
