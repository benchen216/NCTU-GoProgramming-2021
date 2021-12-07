package cw

type Account struct {
	Ip string
}

type PTTAccount struct {
	Account
	Accounts []string
}

type FBAccount struct {
	Account
	Accounts []string
}
