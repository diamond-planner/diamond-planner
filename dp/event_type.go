package dp

import "fmt"

type EventType string

const (
	Game     EventType = "game"
	Practice EventType = "practice"
	Misc     EventType = "misc"
	All      EventType = "all"
)

func (et EventType) String() string {
	return string(et)
}

func (et EventType) IsValid() bool {
	switch et {
	case Game, Practice, Misc, All:
		return true
	default:
		return false
	}
}

func ParseEventType(s string) (EventType, error) {
	switch EventType(s) {
	case Game, Practice, Misc, All:
		return EventType(s), nil
	default:
		return "", fmt.Errorf("invalid EventType: %s", s)
	}
}
