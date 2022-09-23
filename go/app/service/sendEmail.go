package service

import "net/smtp"

type MailService struct{}

func (MailService) SendMail(subject, message string, recipient string) error {
	auth := smtp.PlainAuth(
		"",
		"tatu2425@gmail.com", // 送信に使うアカウント
		"obgultdwobmcdeou",   // アカウントのパスワード or アプリケーションパスワード
		"smtp.gmail.com",
	)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"tatu2425@gmail.com", // 送信元
		[]string{recipient},  // 送信先
		[]byte(
			"To: recipient\r\n"+
				"Subject:"+subject+"\r\n"+
				"\r\n"+
				message),
	)
}
