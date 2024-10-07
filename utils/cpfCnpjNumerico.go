package utils

import "strings"

func CpfCnpjNumerico(cpfCnpj string) string {
	cpfCnpj = strings.ReplaceAll(cpfCnpj, ".", "")
	cpfCnpj = strings.ReplaceAll(cpfCnpj, "/", "")
	cpfCnpj = strings.ReplaceAll(cpfCnpj, "-", "")
	return cpfCnpj
}
