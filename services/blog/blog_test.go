package blog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

type MockBlogGorm struct {
	mock.Mock
}

func (g *MockBlogGorm) GetDetailBlog(ctx context.Context, blogID int64) (*models.Blog, error) {
	args := g.Called(ctx, blogID)
	return args.Get(0).(*models.Blog), args.Error(1)
}

func (g *MockBlogGorm) GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error) {
	args := g.Called(ctx, title, page, size)
	return args.Get(0).(*models.ResponsetListBlog), args.Error(1)
}

func (g *MockBlogGorm) GetListBlogUser(ctx context.Context, title string, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error) {
	args := g.Called(ctx, title, category_id, page, size)
	return args.Get(0).(*models.ResponsetListBlog), args.Error(1)
}

func (g *MockBlogGorm) Create(ctx context.Context, blog *models.Blog) error {
	args := g.Called(ctx, blog)
	return args.Error(0)
}

func (g *MockBlogGorm) Update(ctx context.Context, blog *models.RequestUpdateBlog, Blog_id int64) error {
	args := g.Called(ctx, blog, Blog_id)
	return args.Error(0)
}

func (g *MockBlogGorm) GetListBlogByCategory(ctx context.Context, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error) {
	args := g.Called(ctx, category_id, page, size)
	return args.Get(0).(*models.ResponsetListBlog), args.Error(1)
}

func TestBlog_GetDetailBlog(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Blog
		BlogID      int64
		ExpectedRes *models.Blog
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Blog{
				BlogGorm: &MockBlogGorm{},
				Logger:   zap.L(),
			},
			BlogID: 1,
			ExpectedRes: &models.Blog{
				Blog_id: 1,
			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockBlogGorm)
			mockObj.On("GetDetailBlog", context.Background(), test.BlogID).Return(&models.Blog{
				Blog_id: 1,
			}, nil)
			test.TestObj.BlogGorm = mockObj
			blog, err := test.TestObj.GetDetailBlog(context.Background(), test.BlogID)
			assert.Equal(t, test.ExpectedRes, blog)
			assert.Nil(t, err)
		})
	}
}

func TestBlog_GetListBlog(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Blog
		Title       string
		Page        int64
		Size        int64
		ExpectedRes *models.ResponsetListBlog
		ExpectedErr error
	}{
		{
			Name: "Happy Case",
			TestObj: Blog{
				BlogGorm: &MockBlogGorm{},
				Logger:   zap.L(),
			},
			Title: "Blog",
			Page:  1,
			Size:  5,
			ExpectedRes: &models.ResponsetListBlog{
				TotalBlog:   1,
				TotalPage:   1,
				CurrentPage: 1,
				Data: []models.ResponseBlog{
					{
						Blog_id: 1,
						Title:   "Blog",
					},
				},
			},
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockBlogGorm)
			mockObj.On("GetListBlog", context.Background(), test.Title, test.Page, test.Size).
				Return(&models.ResponsetListBlog{
					TotalBlog:   1,
					TotalPage:   1,
					CurrentPage: 1,
					Data: []models.ResponseBlog{
						{
							Blog_id: 1,
							Title:   "Blog",
						},
					},
				}, nil)
			test.TestObj.BlogGorm = mockObj
			resp, err := test.TestObj.GetListBlog(context.Background(), test.Title, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}
