package creds

// If is an interface to handle the different sources of credentials.
type CredsIf interface {
	DesiredState(account string, usernames []string) (map[string]string, error)
	DeleteAccount(account string) error
}

// Creds contains everything a Credentials uses
type Creds struct {
	Credentials map[string]map[string]string
}
