package email

import (
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/models"
	"go.uber.org/zap"
)

// List of Mail Types we are going to send.
const (
	MailConfirmation models.MailType = iota + 1
	PassReset
	Approve
	Company
)

type IMailService interface {
	CreateMail(mailReq *models.Mail) []byte
	SendMail(mailReq *models.Mail) error
	NewMail(from string, to []string, subject string, mailType models.MailType, data *models.MailData) *models.Mail
}

type SGMailService struct {
	Logger  *zap.Logger
	Configs *config.Config
}

func NewSGMailService(logger *zap.Logger, conf *config.Config) *SGMailService {
	return &SGMailService{
		Logger:  logger,
		Configs: conf,
	}
}

func (ms *SGMailService) CreateMail(mailReq *models.Mail) []byte {
	m := mail.NewV3Mail()

	from := mail.NewEmail("NodeHub", mailReq.From)
	m.SetFrom(from)
	m.Subject = mailReq.Subject
	if mailReq.Mtype == MailConfirmation {
		m.SetTemplateID(ms.Configs.MailVerifTemplateID)
	} else if mailReq.Mtype == PassReset {
		m.SetTemplateID(ms.Configs.PassResetTemplateID)
	} else if mailReq.Mtype == Approve {
		m.SetTemplateID(ms.Configs.ApproveTemplateID)
	} else if mailReq.Mtype == Company {
		m.SetTemplateID(ms.Configs.CompanyTemplateID)
	}

	p := mail.NewPersonalization()

	tos := make([]*mail.Email, 0)
	for _, to := range mailReq.To {
		tos = append(tos, mail.NewEmail("user", to))
	}

	p.AddTos(tos...)

	p.SetDynamicTemplateData("Link", mailReq.Data.Link)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

func (ms *SGMailService) SendMail(mailReq *models.Mail) error {
	request := sendgrid.GetRequest(ms.Configs.SendGridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = ms.CreateMail(mailReq)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		ms.Logger.Error("unable to send mail", zap.Error(err))
		return err
	}
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		ms.Logger.Error("Mail sent failed")
		return errors.New("Request send mail failed: " + response.Body)
	}
	ms.Logger.Info("mail sent successfully", zap.Reflect("sent status code", response.StatusCode))
	return nil
}

func (ms *SGMailService) NewMail(from string, to []string, subject string, mailType models.MailType, data *models.MailData) *models.Mail {
	return &models.Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Mtype:   mailType,
		Data:    data,
	}
}
