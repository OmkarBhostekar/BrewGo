package handlers

import (
	"github.com/omkarbhostekar/brewgo/notification/util"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/rs/zerolog/log"
)

func SendEmail(emailData rabbitmq.EmailNotification, config util.Config) {
	sender := util.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	err := sender.SendEmail(
		emailData.Subject,
		emailData.Body,
		[]string{emailData.To},
		nil,
		nil,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg("cannot send email")
	}
}