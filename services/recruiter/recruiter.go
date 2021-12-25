package recruiter

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/services/email"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type IRecruiterService interface {
	AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error
	UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error
	GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error)
	GetProfileRecruiter(ctx context.Context, id int64) (*models.Recruiter, error)
	GetAllRecruiterForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListRecruiter, error)
	UpdateReciuterByAdmin(ctx context.Context, updateRequest *models.RequestUpdateRecruiterAdmin) error
	UpdateStatusReciuter(ctx context.Context, updateRequest *models.RequestUpdateStatusRecruiter, recruiter_id int64) error
	GetAllRecruiterForCandidate(ctx context.Context, recruiterName string, skillName string, address string, page int64, size int64) (*models.ResponsetListRecruiterForCandidate, error)
	DeleteRecruiterSkill(ctx context.Context, recruiter_skill_id int64) error
	SearchRecruiter(ctx context.Context, req models.RequestSearchRecruiter) (*models.ResponseSearchRecruiter, error)
	GetAllRecruiter(ctx context.Context, req models.RequestSearchRecruiter) (*models.ResponseSearchRecruiter, error)
	CountRecruiter(ctx context.Context) (int64, error)
	CheckPremium(ctx context.Context, recruiterID int64) (bool, error)
}

type Recruiter struct {
	RecruiterGorm IRecruiterDatabase
	Email         email.IMailService

	Conf   *config.Config
	Logger *zap.Logger
}

func NewRecruiterCategory(recruiterGorm *RecruiterGorm, emailSV *email.SGMailService, conf *config.Config, logger *zap.Logger) *Recruiter {
	return &Recruiter{
		RecruiterGorm: recruiterGorm,
		Email:         emailSV,
		Conf:          conf,
		Logger:        logger,
	}
}

func (r *Recruiter) GetProfileRecruiter(ctx context.Context, id int64) (*models.Recruiter, error) {
	rec, err := r.RecruiterGorm.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (r *Recruiter) UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error {
	recruiterModels := &models.RequestUpdateRecruiter{
		Name:             recruiter.Name,
		Address:          recruiter.Address,
		Avartar:          recruiter.Avartar,
		Banner:           recruiter.Banner,
		Phone:            recruiter.Phone,
		Website:          recruiter.Website,
		Description:      recruiter.Description,
		EmployeeQuantity: recruiter.EmployeeQuantity,
		ContacterName:    recruiter.ContacterName,
		ContacterPhone:   recruiter.ContacterPhone,
		Media:            recruiter.Media}
	err := r.RecruiterGorm.UpdateProfile(ctx, recruiterModels, recruiter_id)
	if err != nil {
		return err
	}
	return nil
}

// recruiterSkill
func (r *Recruiter) AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error {
	RecruiterSkillModels := &models.RecruiterSkill{
		Id:          recruiterSkill.Id,
		RecruiterId: recruiterSkill.RecruiterId,
		SkillId:     recruiterSkill.SkillId}
	err := r.RecruiterGorm.AddRecruiterSkill(ctx, RecruiterSkillModels)
	if err != nil {
		return err
	}
	return nil
}

func (r *Recruiter) GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error) {
	acc, err := r.RecruiterGorm.GetRecruiterSkill(ctx, recruiter_id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (r *Recruiter) DeleteRecruiterSkill(ctx context.Context, recruiter_skill_id int64) error {

	err := r.RecruiterGorm.DeleteRecruiterSkill(ctx, recruiter_skill_id)
	if err != nil {
		r.Logger.Error("Can not delete to MySQL", zap.Error(err))
		return err
	}
	return nil
}

//recruiter admin
func (r *Recruiter) GetAllRecruiterForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListRecruiter, error) {
	acc, err := r.RecruiterGorm.GetAllRecruiterForAdmin(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (r *Recruiter) UpdateReciuterByAdmin(ctx context.Context, updateRequest *models.RequestUpdateRecruiterAdmin) error {
	updateData := map[string]interface{}{}
	err1 := mapStructureDecodeWithTextUnmarshaler(updateRequest, &updateData)
	if err1 != nil {
		r.Logger.Error("Can not convert to map", zap.Error(err1))
		return err1
	}

	err := r.RecruiterGorm.UpdateRecruiterByAdmin(ctx, updateRequest.RecruiterID, updateData)
	if err != nil {
		r.Logger.Error("Can not Update to MySQL", zap.Error(err))
		return err
	}
	return nil
}

func (r *Recruiter) UpdateStatusReciuter(ctx context.Context, updateRequest *models.RequestUpdateStatusRecruiter, recruiter_id int64) error {
	recruiterModels := &models.RequestUpdateStatusRecruiter{
		Status: updateRequest.Status}
	err := r.RecruiterGorm.UpdateStatusRecruiter(ctx, recruiterModels, recruiter_id)
	if err != nil {
		return err
	}
	if recruiterModels.Status == true {
		recruiter, err := r.RecruiterGorm.GetProfile(ctx, recruiter_id)
		linkReset := r.Conf.Domain + "recruiter/login"

		from := "lamvhhe130764@fpt.edu.vn"
		to := []string{recruiter.Email}
		subject := "Approved your Company on NodeHub"
		mailType := email.Approve
		mailData := models.MailData{
			Link: linkReset,
		}

		mailReq := r.Email.NewMail(from, to, subject, mailType, &mailData)
		err = r.Email.SendMail(mailReq)
		if err != nil {
			r.Logger.Error("Cannot send email", zap.Error(err))
			return err
		}
	}

	return nil
}

func mapStructureDecodeWithTextUnmarshaler(input, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     output,
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

//recruiter for candidate
func (r *Recruiter) GetAllRecruiterForCandidate(ctx context.Context, recruiterName string, skillName string, address string, page int64, size int64) (*models.ResponsetListRecruiterForCandidate, error) {
	acc, err := r.RecruiterGorm.GetAllRecruiterForCandidate(ctx, recruiterName, skillName, address, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (r *Recruiter) SearchRecruiter(ctx context.Context, req models.RequestSearchRecruiter) (*models.ResponseSearchRecruiter, error) {
	offset := (req.Page - 1) * req.Size
	var recruiters []*models.Recruiter
	var total int64
	var err error
	if req.Text == "" {
		recruiters, total, err = r.RecruiterGorm.GetAllRecruiter(ctx, offset, req.Size)
	} else {
		recruiters, total, err = r.RecruiterGorm.SearchRecruiter(ctx, req.Text, offset, req.Size)
	}
	if err != nil {
		r.Logger.Error("Search recruiter error", zap.Error(err))
		return nil, err
	}
	recruitersWithSkills := make([]models.RecruiterWithSkill, 0)
	for _, recruiter := range recruiters {
		skills, err := r.RecruiterGorm.GetAllSkillByRecruiterID(ctx, recruiter.RecruiterID)
		if err != nil {
			r.Logger.Error(err.Error(), zap.Int64("Recruiter ID", recruiter.RecruiterID))
			continue
		}
		rwk := models.RecruiterWithSkill{
			Recruiter: recruiter,
			Skills:    skills,
		}
		recruitersWithSkills = append(recruitersWithSkills, rwk)
	}
	resp := models.ResponseSearchRecruiter{
		Total:      total,
		Recruiters: recruitersWithSkills,
	}

	return &resp, nil
}

func (r *Recruiter) GetAllRecruiter(ctx context.Context, req models.RequestSearchRecruiter) (*models.ResponseSearchRecruiter, error) {
	offset := (req.Page - 1) * req.Size
	recruiters, total, err := r.RecruiterGorm.GetAllRecruiter(ctx, offset, req.Size)
	if err != nil {
		r.Logger.Error("Search recruiter error", zap.Error(err))
		return nil, err
	}
	recruitersWithSkills := make([]models.RecruiterWithSkill, 0)
	for _, recruiter := range recruiters {
		skills, err := r.RecruiterGorm.GetAllSkillByRecruiterID(ctx, recruiter.RecruiterID)
		if err != nil {
			r.Logger.Error(err.Error(), zap.Int64("Recruiter ID", recruiter.RecruiterID))
			continue
		}
		rwk := models.RecruiterWithSkill{
			Recruiter: recruiter,
			Skills:    skills,
		}
		recruitersWithSkills = append(recruitersWithSkills, rwk)
	}
	resp := models.ResponseSearchRecruiter{
		Total:      total,
		Recruiters: recruitersWithSkills,
	}

	return &resp, nil
}

func (r *Recruiter) CountRecruiter(ctx context.Context) (int64, error) {
	return r.RecruiterGorm.Count(ctx)
}

func (r *Recruiter) CheckPremium(ctx context.Context, recruiterID int64) (bool, error) {
	return r.RecruiterGorm.GetPremiumField(ctx, recruiterID)
}
