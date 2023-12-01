package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rootwarp/key-tools/eth"
)

var ethCmd = &cobra.Command{
	Use: "eth",
}

var newWalletCmd = &cobra.Command{
	Use: "new",
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic, err := eth.NewMnemonic()
		if err != nil {
			return err
		}

		fmt.Println("***** WARNING KEEP THIS SECRET *****")
		fmt.Println(mnemonic)

		_ = mnemonic

		return nil
	},
}

var getAccountCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic := viper.GetString("mnemonic")
		index := viper.GetInt("index")
		file := viper.GetString("file")

		manager, err := eth.NewManager(mnemonic)
		if err != nil {
			return err
		}

		account, hdpath, err := manager.GetAccount(index)
		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("- Address: %s", account.Address.Hex()))
		fmt.Println(fmt.Sprintf("- HDPath: %s", hdpath.String()))

		return manager.ExportAccount(account, hdpath, file)
	},
}

func getEthCommands() *cobra.Command {
	ethCmd.AddCommand(newWalletCmd)

	_ = getAccountCmd.Flags().String("mnemonic", "", "Your Mnemonic")
	viper.BindPFlag("mnemonic", getAccountCmd.Flags().Lookup("mnemonic"))

	_ = getAccountCmd.Flags().String("index", "", "HDPath index")
	viper.BindPFlag("index", getAccountCmd.Flags().Lookup("index"))

	_ = getAccountCmd.Flags().String("file", "", "filename to export")
	viper.BindPFlag("file", getAccountCmd.Flags().Lookup("file"))

	ethCmd.AddCommand(getAccountCmd)
	return ethCmd
}
