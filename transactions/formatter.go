package transactions

import "survorest/user"

type TransactionFormatter struct {
	ID        int             `json:"id"`
	UserID    int             `json:"user_id"`
	Amount    int             `json:"amount"`
	Type      string          `json:"type"`
	Status    string          `json:"status"`
	User      UserTransaction `json:"user"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}
type UserTransaction struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:        transaction.ID,
		UserID:    transaction.UserId,
		Amount:    transaction.Amount,
		Type:      transaction.Type,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: transaction.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}
func FormatTransactionList(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:        transaction.ID,
		UserID:    transaction.UserId,
		Amount:    transaction.Amount,
		Type:      transaction.Type,
		Status:    transaction.Status,
		User:      FormatUserTransaction(transaction.User),
		CreatedAt: transaction.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: transaction.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}
func FormatListTransaction(transaction []Transaction) []TransactionFormatter {
	formatter := make([]TransactionFormatter, len(transaction))
	for i, v := range transaction {
		formatter[i] = FormatTransaction(v)
	}
	return formatter

}
func FormatDashboardTransaction(transaction []Transaction) []TransactionFormatter {
	formatter := []TransactionFormatter{}
	for _, v := range transaction {
		formatter = append(formatter, FormatTransactionList(v))

	}
	return formatter
}
func FormatUserTransaction(user user.User) UserTransaction {
	userFormatter := UserTransaction{}
	userFormatter.ID = int(user.Id)
	userFormatter.FullName = user.FullName
	return userFormatter
}
