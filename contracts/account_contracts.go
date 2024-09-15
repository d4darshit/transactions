package contracts

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}
type CreateAccountResponse struct {
	AccoundId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number""`
}

type GetAccountRequest struct {
	AccountId int `json:"account_d"`
}

type GetAccountResponse struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
