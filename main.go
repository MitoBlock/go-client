package main

import (
    "context"
    "fmt"
    "log"

    // Importing the general purpose Cosmos blockchain client
    "github.com/ignite/cli/ignite/pkg/cosmosclient"

    // Importing the types package of blockchain
    "mitoblockchaindev/x/mitoblockchaindev/types"
)

func main() {
    // Prefix to use for account addresses.
    // The address prefix was assigned to the blog blockchain
    // using the `--address-prefix` flag during scaffolding.
    addressPrefix := "mito"

    // Create a Cosmos client instance
    cosmos, err := cosmosclient.New(
        context.Background(),
        cosmosclient.WithAddressPrefix(addressPrefix),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Account `alice` was initialized during `ignite chain serve`
    accountName := "alice"

    // Get account from the keyring
    account, err := cosmos.Account(accountName)
    if err != nil {
        log.Fatal(err)
    }

    addr, err := account.Address(addressPrefix)
    if err != nil {
        log.Fatal(err)
    }

    // Define a message to create a discount token
    msg := &types.MsgCreateDiscountToken{
        Creator: addr,
		Timestamp: "timestamp",
		ActivityName:      "ActivityName",
		Score:             "Score",
		Message:           "Message",
		DiscountValue:     "DiscountValue",
		EligibleCompanies: "EligibleCompanies",
		ItemType:          "ItemType",
		ExpiryDate:        "ExpiryDate",
    }

    // Broadcast a transaction from account `alice` with the message
    // to create a post store response in txResp
    txResp, err := cosmos.BroadcastTx(account, msg)
    if err != nil {
        log.Fatal(err)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgCreateDiscountToken:\n\n")
    fmt.Println(txResp)

    // Instantiate a query client for the blockchain
    queryClient := types.NewQueryClient(cosmos.Context())

    // Query the blockchain using the client's `DiscountTokens` method
    // to get all tokens store all tokens in queryResp
    queryResp, err := queryClient.DiscountTokens(context.Background(), &types.QueryDiscountTokensRequest{})
    if err != nil {
        log.Fatal(err)
    }

    // Print response from querying all the tokens
    fmt.Print("\n\nAll tokens:\n\n")
    fmt.Println(queryResp)
}