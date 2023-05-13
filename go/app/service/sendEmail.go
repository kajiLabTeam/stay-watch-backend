package service

import "net/smtp"

type MailService struct{}

func (MailService) SendMail(subject, message string, recipient string) error {
	auth := smtp.PlainAuth(
		"",
		"staywatch108@gmail.com", // 送信に使うアカウント
		"Ait-StayWatch108",       // アカウントのパスワード or アプリケーションパスワード
		"smtp.gmail.com",
	)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"staywatch108@gmail.com", // 送信元
		[]string{recipient},      // 送信先
		[]byte(
			"To: recipient\r\n"+
				"Subject:"+subject+"\r\n"+
				"\r\n"+
				message),
	)
}
