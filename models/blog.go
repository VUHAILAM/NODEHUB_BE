package models

import "time"

type Blog struct {
	Blog_id      int64     `json:"blog_id" gorm:"primaryKey"`
	Category_id  int64     `json:"category_id"`
	CategoryName string    `json:"category_name" gorm:"-"`
	Title        string    `json:"title"`
	Icon         string    `json:"icon"`
	Excerpts     string    `json:"excerpts"`
	Description  string    `json:"description"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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

type ResponseBlog struct {
	Blog_id       int64     `json:"blog_id" gorm:"primaryKey"`
	Category_name string    `json:"category_name"`
	Category_id   int64     `json:"category_id"`
	Title         string    `json:"title"`
	Icon          string    `json:"icon"`
	Excerpts      string    `json:"excerpts"`
	Description   string    `json:"description"`
	Status        bool      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ResponsetListBlog struct {
	TotalBlog   int64          `json:"totalBlog"`
	TotalPage   float64        `json:"totalPage"`
	CurrentPage int64          `json:"currentPage"`
	Data        []ResponseBlog `json:"data"`
}

type RequestGetListBlog struct {
	Title       string `json:"title"`
	Category_id int64  `json:"category_id"`
	Page        int64  `json:"page"`
	Size        int64  `json:"size"`
}
type RequestGetListBlogByCategoryId struct {
	Category_id int64 `json:"category_id"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type RequestGetDetailBlog struct {
	BlogID int64 `json:"blog_id"`
}
