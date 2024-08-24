// Code generated by "enumer -type Reason -trimprefix Reason -transform=lower"; DO NOT EDIT.

package api

import (
	"fmt"
	"strings"
)

const _ReasonName = "unknownwaitingforauthorization"

var _ReasonIndex = [...]uint8{0, 7, 30}

const _ReasonLowerName = "unknownwaitingforauthorization"

func (i Reason) String() string {
	if i < 0 || i >= Reason(len(_ReasonIndex)-1) {
		return fmt.Sprintf("Reason(%d)", i)
	}
	return _ReasonName[_ReasonIndex[i]:_ReasonIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ReasonNoOp() {
	var x [1]struct{}
	_ = x[ReasonUnknown-(0)]
	_ = x[ReasonWaitingForAuthorization-(1)]
}

var _ReasonValues = []Reason{ReasonUnknown, ReasonWaitingForAuthorization}

var _ReasonNameToValueMap = map[string]Reason{
	_ReasonName[0:7]:       ReasonUnknown,
	_ReasonLowerName[0:7]:  ReasonUnknown,
	_ReasonName[7:30]:      ReasonWaitingForAuthorization,
	_ReasonLowerName[7:30]: ReasonWaitingForAuthorization,
}

var _ReasonNames = []string{
	_ReasonName[0:7],
	_ReasonName[7:30],
}

// ReasonString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ReasonString(s string) (Reason, error) {
	if val, ok := _ReasonNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ReasonNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Reason values", s)
}

// ReasonValues returns all values of the enum
func ReasonValues() []Reason {
	return _ReasonValues
}

// ReasonStrings returns a slice of all String values of the enum
func ReasonStrings() []string {
	strs := make([]string, len(_ReasonNames))
	copy(strs, _ReasonNames)
	return strs
}

// IsAReason returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Reason) IsAReason() bool {
	for _, v := range _ReasonValues {
		if i == v {
			return true
		}
	}
	return false
}
