package model

import "time"

func BuildInsertWallet(customerID string) InsertWallet {
	return InsertWallet{
		OwnedBy: customerID,
		Status:  StatusWalletDisabled,
	}
}

func BuildAPIResponseError(err error) GeneralAPIResponse {
	return GeneralAPIResponse{
		Status: StatusFailedAPI,
		Error:  err.Error(),
	}
}

func BuildAPIResponseSuccess(data interface{}) GeneralAPIResponse {
	return GeneralAPIResponse{
		Status: StatusSuccessAPI,
		Data:   data,
	}
}

func BuildUpdateWalletStatusEnabled(customerID string, now time.Time) UpdateWalletStatus {
	return UpdateWalletStatus{
		Status:     StatusWalletEnabled,
		CustomerID: customerID,
		EnabledAt:  &now,
	}
}

func BuildUpdateWalletStatusDisabled(customerID string, now time.Time) UpdateWalletStatus {
	return UpdateWalletStatus{
		Status:     StatusWalletDisabled,
		CustomerID: customerID,
		DisabledAt: &now,
	}
}

func BuildUpdateWalletBalance(customerID string, increment float64) UpdateWalletBalance {
	return UpdateWalletBalance{
		CustomerID: customerID,
		Increment:  increment,
	}
}

func BuildInsertDeposit(req DepositRequest, customerID string, now time.Time) InsertDeposit {
	return InsertDeposit{
		DepositedBy: customerID,
		Status:      StatusDepositSuccess,
		ReferenceID: req.ReferenceID,
		Amount:      req.Amount,
		DepositedAt: now,
	}
}

func BuildInsertWithdrawal(req WithdrawalRequest, customerID string, now time.Time) InsertWithdrawal {
	return InsertWithdrawal{
		WithdrawnBy: customerID,
		Status:      StatusWithdrawalSuccess,
		ReferenceID: req.ReferenceID,
		Amount:      req.Amount,
		WithdrawnAt: now,
	}
}
