package config

type mailerConf struct {
	Smtp_host     string
	Smtp_port     int
	Smtp_username string
	Smtp_password string
	Sender        string
	TemplateDir   string
}
