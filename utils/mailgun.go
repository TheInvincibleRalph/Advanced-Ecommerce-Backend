package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go"
)

func InitializeMailgun() *mailgun.MailgunImpl {
	return mailgun.NewMailgun(
		os.Getenv("MAILGUN_DOMAIN"),
		os.Getenv("MAILGUN_API_KEY"),
	)
}

func SendOrderConfirmationEmail(toEmail string, orderDetails string) error {
	mg := InitializeMailgun()

	sender := "no-reply@ecommerce" // Replace with your Mailgun sender email
	subject := "Order Confirmation"
	body := fmt.Sprintf("Thank you for your order!\n\nOrder Details:\n%s", orderDetails)

	message := mg.NewMessage(sender, subject, body, toEmail) /

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	return err
}
