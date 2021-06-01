package mail

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailHandler struct {
	service *gmail.Service
}

func NewOAuthGmailHandler(clientId string, clientSecret string, accessToken string, refreshToken string) *GmailHandler {
	config := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	service, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	if service != nil {
		fmt.Println("Email service is initialized")
	}

	return &GmailHandler{
		service: service,
	}
}

func (g *GmailHandler) SendEmailOAUTH2(to string, data interface{}, template string) (bool, error) {
	emailBody, err := parseTemplate(template, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "Notification Email from Bills Gate" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = g.service.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}
