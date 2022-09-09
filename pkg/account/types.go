package account

type AccountList[T any] struct {
	Data []Account[T]
}

type Account[T any] struct {
	AccountNo        string
	CurrentBalance   float64
	Currency         string
	AdditionalFields T
}
