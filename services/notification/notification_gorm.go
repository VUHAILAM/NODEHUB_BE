package notification

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableAccount = "notification"

type NotificationGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewNotificationGorm(db *gorm.DB, logger *zap.Logger) *NotificationGorm {
	return &NotificationGorm{
		db:     db,
		logger: logger,
	}
}

/*Create Notification*/
func (n *NotificationGorm) Create(ctx context.Context, notification *models.Notification) error {
	db := n.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(notification).Error
	if err != nil {
		n.logger.Error("NotificationGorm: Create notification error", zap.Error(err))
		return err
	}
	return nil
}

/*Get Notification*/
func (n *NotificationGorm) GetListNotificationByAccount(ctx context.Context, account_id int64, page int64, size int64) (*models.ResponsetListNotification, error) {
	db := n.db.WithContext(ctx)
	arr := []models.Notification{}
	resutl := models.ResponsetListNotification{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select account_id, title, content, 'key', check_read, created_at, updated_at 
	FROM nodehub.notification 
	where account_id = ?  ORDER BY created_at desc LIMIT ?, ?`, account_id, offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.notification where account_id = ?`, account_id).Scan(&total)
	if err != nil {
		n.logger.Error("NotificationGorm: Get notification error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.Total = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}
