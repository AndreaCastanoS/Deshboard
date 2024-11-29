package models

type Workclass int

const (
	StateGov Workclass = iota
	SelfEmpNotInc
	Private
	FederalGov
	LocalGov
)

func (w Workclass) String() string {
	return [...]string{"State-gov", "Self-emp-not-inc", "Private", "Federal-gov", "Local-gov"}[w]
}

func ParseWorkclass(s string) Workclass {
	switch s {
	case "State-gov":
		return StateGov
	case "Self-emp-not-inc":
		return SelfEmpNotInc
	case "Private":
		return Private
	case "Federal-gov":
		return FederalGov
	case "Local-gov":
		return LocalGov
	default:
		return -1 
	}
}
