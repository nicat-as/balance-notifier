package provider

import (
	"os"
	"strings"
)

type Provider string

const (
	Kapital Provider = "KAPITAL"
)

const (
	kapitalBase       = "PROVIDER.KAPITAL."
	kapitalBaseUrl    = kapitalBase + "BASE_URL"
	kapitalDomainPath = kapitalBase + "DOMAIN_PATH"

	KapitalLogin       = kapitalBase + "LOGIN"
	KapitalRefresh     = kapitalBase + "REFRESH"
	KapitalAccountList = kapitalBase + "ACCOUNT_LIST"
)

func get(envs ...string) string {
	sb := strings.Builder{}
	for _, env := range envs {
		sb.WriteString(os.Getenv(env))
	}
	return sb.String()
}

func GetEndpoint(endpoint string) string {
	return get(kapitalBaseUrl, kapitalDomainPath, endpoint)
}
