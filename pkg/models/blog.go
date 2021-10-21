package models

import "time"

type Blog struct {
	Blog_id     int64     `json:"blog_id" gorm:"primaryKey"`
	Category_id int64     `json:"category_id"`
	Title       string    `json:"title"`
	Icon        string    `json:"icon"`
	Excerpts    string    `json:"excerpts"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestCreateBlog struct {
	Blog_id     int64  `json:"blog_id" gorm:"primaryKey"`
	Category_id int64  `json:"category_id"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Excerpts    string `json:"excerpts"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type RequestUpdateBlog struct {
	Category_id int64  `json:"category_id"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Excerpts    string `json:"excerpts"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type ResponsetListBlog struct {
	TotalBlog   int64   `json:"totalBlog"`
	TotalPage   float64 `json:"totalPage"`
	CurrentPage int64   `json:"currentPage"`
	Data        []Blog  `json:"data"`
}

type RequestGetListBlog struct {
	Title string `json:"title"`
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
}
