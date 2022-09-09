package authorization

type Authorization[REQ, RES any] interface {
	GetToken(REQ) RES
}
