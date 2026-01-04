package dp

import (
	"github.com/pocketbase/pocketbase/core"
)

const (
	PushSubscriptionsCollection = "pushsubscriptions"
)

var _ core.RecordProxy = (*PushSubscription)(nil)

// PushSubscription RecordProxy for collection `pushsubscriptions`.
// Provides type-safe struct access to push subscription records.
type PushSubscription struct {
	core.BaseRecordProxy
}

func (p *PushSubscription) ID() string {
	return p.Id
}

func (p *PushSubscription) User() string {
	return p.GetString("user")
}

func (p *PushSubscription) SetUser(userId string) {
	p.Set("user", userId)
}

func (p *PushSubscription) Endpoint() string {
	return p.GetString("endpoint")
}

func (p *PushSubscription) SetEndpoint(endpoint string) {
	p.Set("endpoint", endpoint)
}

func (p *PushSubscription) KeyP256dh() string {
	return p.GetString("key_p256dh")
}

func (p *PushSubscription) SetKeyP256dh(key string) {
	p.Set("key_p256dh", key)
}

func (p *PushSubscription) KeyAuth() string {
	return p.GetString("key_auth")
}

func (p *PushSubscription) SetKeyAuth(key string) {
	p.Set("key_auth", key)
}

func (p *PushSubscription) Encoding() string {
	return p.GetString("encoding")
}

func (p *PushSubscription) SetEncoding(encoding string) {
	p.Set("encoding", encoding)
}
