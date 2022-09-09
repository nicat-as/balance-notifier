package account

type KapitalAccountClient struct {
}

type AccountList struct {
	AccountsList []Account `json:"accountsList"`
}

type Account struct {
	AccountNo     string `json:"custAcNo"`
	Currency      string `json:"ccy"`
	CurrentAmount string `json:"currAmt"`
}
