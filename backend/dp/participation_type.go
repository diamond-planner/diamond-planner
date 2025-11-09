package dp

import "fmt"

type ParticipationType string

const (
	In    ParticipationType = "in"
	Out   ParticipationType = "out"
	Maybe ParticipationType = "maybe"
)

func (pt ParticipationType) String() string {
	return string(pt)
}

func (pt ParticipationType) IsValid() bool {
	switch pt {
	case In, Out, Maybe:
		return true
	default:
		return false
	}
}

func ParseParticipationType(s string) (ParticipationType, error) {
	switch ParticipationType(s) {
	case In, Out, Maybe:
		return ParticipationType(s), nil
	default:
		return "", fmt.Errorf("invalid ParticipationType: %s", s)
	}
}
