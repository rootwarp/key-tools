package eth

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Manager interface {
	GetAccount(int) (accounts.Account, accounts.DerivationPath, error)
	ExportAccount(account accounts.Account, hdpath accounts.DerivationPath, filename string) error
}

type WalletManager struct {
	wallet *hdwallet.Wallet
}

func (a *WalletManager) loadMnemonic(mn string) error {
	w, err := hdwallet.NewFromMnemonic(mn)
	if err != nil {
		return err
	}

	a.wallet = w
	return nil
}

func (a *WalletManager) GetAccount(index int) (accounts.Account, accounts.DerivationPath, error) {
	hdpath, err := accounts.ParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := a.wallet.Derive(hdpath, false)
	if err != nil {
		return accounts.Account{}, nil, err
	}

	return account, hdpath, nil
}

func (a *WalletManager) ExportAccount(account accounts.Account, hdpath accounts.DerivationPath, filename string) error {
	privKey, err := a.wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	f.WriteString(fmt.Sprintf("- Address: %s\n", account.Address.Hex()))
	f.WriteString(fmt.Sprintf("- HDPath: %s\n", hdpath.String()))
	f.WriteString(fmt.Sprintf("- PrivateKey: 0x%s\n", privKey))

	return nil
}

func NewMnemonic() (string, error) {
	mn, err := hdwallet.NewMnemonic(256)
	if err != nil {
		return "", err
	}

	return mn, nil
}

func NewManager(mnemonic string) (Manager, error) {
	creator := WalletManager{}
	err := creator.loadMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	return &creator, nil
}
