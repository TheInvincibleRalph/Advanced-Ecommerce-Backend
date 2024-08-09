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

	message := mg.NewMessage(sender, subject, body, toEmail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	message.SetDeliveryTime(time.Now().Add(time.Second * 10))
	message.SetReplyTo("no-reply@ecommerce")

	// A goroutine is used to send the email asynchronously since the method Send() does not support context.
	// Context is used to cancel the email sending operation if it takes too long to prevent the program from hanging.

	resultChan := make(chan error, 1) // Buffered channel of type error to receive the result of the email sending operation. The 1 indicates the buffer size of the channel which means it can only store one value.

	go func() {
		_, _, err := mg.Send(message)
		resultChan <- err
	}()

	select {
	case err := <-resultChan:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil

}
