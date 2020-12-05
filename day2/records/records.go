package records

import "strings"

const (
	CountPolicyType = iota
	PositionPolicyType
)

type Policy interface {
	ValidatePassword(string) bool
}

type CountPolicy struct {
	MustContain   rune
	MinOccurances int
	MaxOccurantes int
}

type PositionPolicy struct {
	MustContain rune
	PosA        int
	PosB        int
}

type Record struct {
	Password string
	Policy   Policy
}

func (p CountPolicy) ValidatePassword(password string) bool {
	charCount := strings.Count(password, string(p.MustContain))
	if p.MinOccurances <= charCount && charCount <= p.MaxOccurantes {
		return true
	}
	return false
}

func (p PositionPolicy) ValidatePassword(password string) bool {
	posAOK := password[p.PosA-1] == byte(p.MustContain)
	posBOK := password[p.PosB-1] == byte(p.MustContain)

	if (posAOK || posBOK) && !(posAOK && posBOK) {
		return true
	}
	return false
}
