package eth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	mnemonic, err := NewMnemonic()
	assert.Nil(t, err)

	creator, err := NewManager(mnemonic)
	assert.Nil(t, err)

	account, err := creator.GetAccount(0)
	assert.Nil(t, err)

	fmt.Println(account.Address.Hex())
}

func TestLoadWallet(t *testing.T) {
	manager, err := NewManager(testMnemonic)
	assert.Nil(t, err)

	account, err := manager.GetAccount(0)
	assert.Nil(t, err)

	fmt.Println(account.Address.Hex())

	// Reproducible
	newManager, err := NewManager(testMnemonic)
	assert.Nil(t, err)

	account2, err := newManager.GetAccount(0)
	assert.Nil(t, err)

	assert.Equal(t, account.Address.Hex(), account2.Address.Hex())
}
