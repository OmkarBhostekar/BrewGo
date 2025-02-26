package rabbitmq

const (
	EventsExchange = "events"

	QueueNewOrders       = "new.orders"
	QueueOrderItemStatus = "order.item.status"
	QueueNotifications   = "notifications"

	NewOrderItemReceived   = "new.order.item.received"
	OrderItemStatusUpdated = "order.item.status.updated"
	SendNotification       = "send.notification"
)

type NewOrderItemEventData struct {
	OrderId     int32  `json:"order_id"`
	OrderItemId int32  `json:"order_item_id"`
	ProductId   int32  `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	Notes       string `json:"notes"`
}

type OrderItemStatusEventData struct {
	OrderId     int32  `json:"order_id"`
	OrderItemId int32  `json:"order_item_id"`
	Status      string `json:"status"`
}

type NotificationEventData struct {
	Type      string            `json:"type"`
	EmailData EmailNotification `json:"email_data"`
	SmsData   SmsNotification   `json:"sms_data"`
}

type EmailNotification struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type SmsNotification struct {
	To      string `json:"to"`
	Message string `json:"message"`
}
