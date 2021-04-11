package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitStoreCmd() *cobra.Command {
	storeCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print stores where you can buy beers",
		Run:   runStoreFn(),
	}
	storeCmd.Flags().StringP("id", "i", "", "id of the store")
	return storeCmd
}

func runStoreFn() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
