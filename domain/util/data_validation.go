package util

import (
	"strconv"
	"strings"
)

var (
	personWeights  = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	companyWeights = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// GetRawDocument returns a CPF or CNPJ without any dots, dashes and slashes
func GetRawDocument(document string) string {
	document = strings.TrimSpace(document)
	document = strings.ReplaceAll(document, ".", "")
	document = strings.ReplaceAll(document, "-", "")
	document = strings.ReplaceAll(document, "/", "")

	return document
}

// GetRawPhoneNumber returns a phone number without any parenthesis, dashes and spaces
func GetRawPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.TrimSpace(phoneNumber)
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")

	return phoneNumber
}

// IsValidDocument checks if a raw document string is a valid CPF or CNPJ.
//
// The input parameter `document` should be a raw document string, meaning
// it must not contain any dots, dashes or slashes. Use the GetRawDocument
// function to preprocess the input if needed.
func IsValidDocument(document string) bool {
	return IsValidCPF(document) || IsValidCNPJ(document)
}

// IsValidCPF checks if a raw document string is a valid CPF.
//
// The input parameter `document` should be a raw document string, meaning
// it must not contain any dots, dashes or slashes. Use the GetRawDocument
// function to preprocess the input if needed.
func IsValidCPF(document string) bool {
	// Check for valid length
	strLen := len(document)
	if strLen != 11 {
		return false
	}

	// TODO: Maybe this is wrong... need better testing
	// Check if all chars are digits
	//diff := false
	//last := int32(document[0])
	//
	//for _, char := range document {
	//	if char < '0' || char > '9' {
	//		return false
	//	}
	//
	//	if last != char {
	//		diff = true
	//	}
	//	last = char
	//}

	// CPF with all 11 equal digits passes in the algorithm, but is considered invalid
	//if strLen == 11 && !diff {
	//	return false
	//}

	// Treat as CPF
	baseDocument := document[:strLen-2]

	// Calculating expected verification numbers
	ver1 := calcVerificationNumber(baseDocument, personWeights)
	ver2 := calcVerificationNumber(baseDocument+strconv.Itoa(ver1), personWeights)

	// Assembling the generated valid document
	generated := baseDocument + strconv.Itoa(ver1) + strconv.Itoa(ver2)

	// The original document must be the same as the generated
	return document == generated
}

// IsValidCNPJ checks if a raw document string is a valid CNPJ.
//
// The input parameter `document` should be a raw document string, meaning
// it must not contain any dots, dashes or slashes. Use the GetRawDocument
// function to preprocess the input if needed.
func IsValidCNPJ(document string) bool {
	// Check for valid length
	strLen := len(document)
	if strLen != 14 {
		return false
	}

	// Calculating expected verification numbers
	baseDocument := document[:strLen-2]
	ver1 := calcVerificationNumber(baseDocument, companyWeights)
	ver2 := calcVerificationNumber(baseDocument+strconv.Itoa(ver1), companyWeights)

	// Assembling the generated valid document
	generated := baseDocument + strconv.Itoa(ver1) + strconv.Itoa(ver2)

	// The original document must be the same as the generated
	return document == generated
}

// IsValidPhoneNumber checks if a string is a valid Brazil phone number.
//
// The input parameter `phoneNumber` should be a raw phone number string,
// meaning it must not contain any parenthesis, dashes or spaces. Use the
// GetRawPhoneNumber function to preprocess the input if needed.
func IsValidPhoneNumber(phoneNumber string) bool {
	// Len should be 10 or 11, depending on the prefix 9
	strLen := len(phoneNumber)
	if strLen != 10 && strLen != 11 {
		return false
	}

	// All digits should be numbers
	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			return false
		}
	}

	// First two digits are DDD, must be between 11 and 99
	ddd, err := strconv.Atoi(phoneNumber[0:2])
	if err != nil {
		return false
	}

	if ddd < 11 || ddd > 99 {
		return false
	}

	return true
}

func calcVerificationNumber(document string, weights []int) int {
	var sum = 0
	for i := len(document) - 1; i >= 0; i-- {
		number, _ := strconv.Atoi(document[i : i+1])
		sum += number * weights[len(weights)-len(document)+i]
	}

	check := 11 - sum%11
	if check > 9 {
		return 0
	} else {
		return check
	}
}
