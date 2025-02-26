package workers

import (
	"encoding/json"

	"github.com/omkarbhostekar/brewgo/notification/handlers"
	"github.com/omkarbhostekar/brewgo/notification/util"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func NotificationWorker(msg amqp.Delivery, config util.Config) error {
	data := rabbitmq.NotificationEventData{}
	err := json.Unmarshal(msg.Body, &data)

	if err != nil {
		log.Error().Err(err).Msg("cannot unmarshal data")
		return err
	}

	if data.Type == "email" {
		log.Info().Msgf("sending email to %s with %s", data.EmailData.To, data.EmailData.Body)
		handlers.SendEmail(data.EmailData, config)
	} else if data.Type == "sms" {
		log.Info().Msgf("sending sms to %s", data.SmsData.To)
	} else {
		log.Error().Msg("invalid notification type")
	}

	msg.Ack(false)
	return nil
}