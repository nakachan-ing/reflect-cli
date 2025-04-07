package model

import "fmt"

type subTypeInvalidError struct {
	Message string
}

func (e *subTypeInvalidError) Error() string {
	return e.Message
}

type sourceUnspecifiedError struct {
	Message string
}

func (e *sourceUnspecifiedError) Error() string {
	return e.Message
}

func IsSubType(subTypeInput string) (SubType, error) {
	subType := SubType(subTypeInput)

	if !AllowedSubType[subType] {
		// type is invalid の処理
		errorMsg := fmt.Sprintln("Type is invalid")
		errorMsg += fmt.Sprintln("Allowed types:")
		for subType := range AllowedSubType {
			errorMsg += fmt.Sprintf("  ・%v\n", subType)
		}
		return "", &subTypeInvalidError{
			Message: errorMsg,
		}
		// return "", fmt.Errorf("%s is invalid type\n\n%s", subTypeInput, AllowedSubTypesStringList()) // AllowedSubTypesStringListを使った場合、こちらをreturn

	} else {
		return subType, nil
	}
}

func IsSourceSpecified(subType SubType, source string) error {
	if subType == SubType("literature") {
		if source == "" {
			errorMsg := fmt.Sprintln("Source is required when type is literature")
			return &sourceUnspecifiedError{
				Message: errorMsg,
			}
		}
	}
	return nil
}

// 簡略化したい場合、こちらに切り替え
// func AllowedSubTypesStringList() string {
// 	var keys []string
// 	for k := range AllowedSubType {
// 		keys = append(keys, string(k))
// 	}
// 	sort.Strings(keys)

// 	msg := "Allowed types:\n"
// 	for _, k := range keys {
// 		msg += fmt.Sprintf("  ・%s\n", k)
// 	}
// 	return msg
// }
