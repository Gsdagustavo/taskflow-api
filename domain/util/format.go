package util

import "fmt"

// FormatDocument takes a raw document string and returns it properly formatted
//
// Will format it to CPF or CNPJ depending on size (11 or 14 respectively)
//
// Returns an empty string if length is invalid
func FormatDocument(document string) string {
	if len(document) == 11 {
		return FormatCPF(document)
	} else if len(document) == 14 {
		return FormatCNPJ(document)
	}

	return ""
}

// FormatPhoneNumber takes a raw phone number string and returns it properly formatted,
// independent of having 9-prefix or not
//
// Returns an empty string if length is invalid
func FormatPhoneNumber(phoneNumber string) string {
	if len(phoneNumber) == 10 {
		return formatPhoneNumber(phoneNumber, false)
	} else if len(phoneNumber) == 11 {
		return formatPhoneNumber(phoneNumber, true)
	}

	return ""
}

func formatPhoneNumber(s string, prefix bool) string {
	if prefix {
		return fmt.Sprintf("(%s) %s-%s", s[0:2], s[2:7], s[7:])
	} else {
		return fmt.Sprintf("(%s) %s-%s", s[0:2], s[2:6], s[6:])
	}
}

// FormatCNPJ takes a raw CNPJ string and returns it properly formatted
//
// String has to be exactly 14 characters of length, otherwise will return an empty string
func FormatCNPJ(cnpj string) string {
	if len(cnpj) != 14 {
		return ""
	}

	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[0:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
}

// FormatCPF takes a raw CPF string and returns it properly formatted
//
// String has to be exactly 11 characters of length, otherwise will return an empty string
func FormatCPF(cpf string) string {
	if len(cpf) != 11 {
		return ""
	}

	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])
}
