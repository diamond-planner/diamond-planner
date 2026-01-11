package dp

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/SherClockHolmes/webpush-go"
)

const (
	VAPIDSubscriberEnvName = "VAPID_SUBSCRIBER"
	VAPIDPublicKeyEnvName  = "VAPID_PUBLIC_KEY"
	VAPIDPrivateKeyEnvName = "VAPID_PRIVATE_KEY"
)

type PushMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tag   string `json:"tag"`
}

type PushService interface {
	handleTestPush(sub *webpush.Subscription) error
	SendPushMessage(msg *PushMessage, sub *webpush.Subscription) error
}

type PushServiceImpl struct {
	Subscriber      string
	VAPIDPublicKey  string
	VAPIDPrivateKey string
}

func NewPushService() PushService {
	s := PushServiceImpl{
		Subscriber:      os.Getenv(VAPIDSubscriberEnvName),
		VAPIDPublicKey:  os.Getenv(VAPIDPublicKeyEnvName),
		VAPIDPrivateKey: os.Getenv(VAPIDPrivateKeyEnvName),
	}

	if s.Subscriber == "" || s.VAPIDPublicKey == "" || s.VAPIDPrivateKey == "" {
		log.Fatal("Missing VAPID environment variables, exiting")
	}
	return &s
}

// SendPushMessage sends a push notification to the given subscription
func (p PushServiceImpl) SendPushMessage(msg *PushMessage, sub *webpush.Subscription) error {
	blob, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := webpush.SendNotification(blob, sub, &webpush.Options{
		Subscriber:      p.Subscriber,
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
		TTL:             30,
	})
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}

// handleTestPush sends a test push notification to the given subscription
func (p PushServiceImpl) handleTestPush(sub *webpush.Subscription) error {
	msg := PushMessage{
		Title: "Hello there",
		Body:  "Real men test in production",
		Tag:   "Test Tag",
	}

	return p.SendPushMessage(&msg, sub)
}
