package account

type AccountClient[REQ, RES any] interface {
	FetchAccountList(req REQ) (*AccountList[RES], error)
}
