package proxyserver

type POWInterface interface {
	GetWork() string
	ValidateWork(string) bool
}
