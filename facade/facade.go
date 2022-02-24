package facade

import "fmt"

// walletFacade
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

// newWalletFacade
//  @param accountId
//  @param code
//  @return *walletFacade
func newWalletFacade(accountId string, code int) *walletFacade {
	fmt.Println("Starting crate account")
	walletFacade := &walletFacade{
		account:      newAccount(accountId),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

// addMoneyToWallet
//  @receiver w
//  @param accountID
//  @param securityCode
//  @param amount
//  @return error
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err

	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

// account
type account struct {
	name string
}

// newAccount
//  @param accountName
//  @return *account
func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

// checkAccount
//  @receiver a
//  @param accountName
//  @return error
func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

// securityCode
type securityCode struct {
	code int
}

// newSecurityCode
//  @param code
//  @return *securityCode
func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

// checkCode
//  @receiver s
//  @param incomingCode
//  @return error
func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}

	fmt.Println("SecurityCode Verified")
	return nil
}

// wallet
type wallet struct {
	balance int
}

// newWallet
//  @return *wallet
func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

// creditBalance
//  @receiver w
//  @param amount
func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfuly")
	return
}

// debitBalance
//  @receiver w
//  @param amount
//  @return error
func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

// notification
type notification struct {
}

// sendWalletDebitNotification
//  @receiver n
func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

// sendWalletCreditNotification
//  @receiver n
func (n *notification) sendWalletCreditNotification() {

}

// ledger
type ledger struct {
}

// makeEntry
//  @receiver s
//  @param accountID
//  @param txnType
//  @param amount
func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
