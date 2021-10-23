package models

type MailType int

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
	Mtype   MailType
	Data    *MailData
}

type MailData struct {
	Username string
	Code     string
	Link     string
}
