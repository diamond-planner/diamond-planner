package dp

import (
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

type PushService interface {
	handleTestPush(sub *webpush.Subscription) error
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

func (p PushServiceImpl) handleTestPush(sub *webpush.Subscription) error {
	resp, err := webpush.SendNotification([]byte("Web Push Test"), sub, &webpush.Options{
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
