package main

import (

    "net/http"
//	"errors"
	"github.com/gin-gonic/gin"

    "context"
    "fmt"
    "log"

    // Importing the general purpose Cosmos blockchain client
    "github.com/ignite/cli/ignite/pkg/cosmosclient"

    // Importing the types package of blockchain
    "mitoblockchaindev/x/mitoblockchaindev/types"
)


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}



type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

    var addressPrefix = "mito"

    var cosmos, err = cosmosclient.New(
        context.Background(),
        cosmosclient.WithAddressPrefix(addressPrefix),
    )

    // Account `alice` was initialized during `ignite chain serve`
    var accountName = "bob"

    // Get account from the keyring
    var account, accounterr = cosmos.Account(accountName)

    var addr, addresserr = account.Address(addressPrefix)

    // Instantiate a query client for the blockchain
    var queryClient = types.NewQueryClient(cosmos.Context())



func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, addr)
}


func getTokens(c *gin.Context) {

    // Query the blockchain using the client's `DiscountTokens` method
    // to get all tokens store all tokens in queryResp
    queryResp, qerr := queryClient.DiscountTokens(context.Background(), &types.QueryDiscountTokensRequest{})
    if qerr != nil {
        log.Fatal(err)
    }

    // Print response from querying all the tokens
    fmt.Print("\n\nAll tokens:\n\n")
    // fmt.Println(queryResp)

    fmt.Print((queryResp.GetDiscountToken))
	c.IndentedJSON(http.StatusOK, queryResp)
}


func createDiscountMembershipToken(c *gin.Context) {
    // msg := &types.MsgCreateDiscountToken{
    msg := &types.MsgCreateMembershipToken{
        Creator: addr,
	 	Timestamp: "timestamp",
        ActivityName:      "Weekly leaderboard",
        Score:             "10",
        Message:           "Impresionante",
        MembershipDuration:     "3",
        ExpiryDate:        "5th Dec 2022",
    }
    // Broadcast a transaction from account `alice` with the message
    // to create a post store response in txResp
    txResp, transerr := cosmos.BroadcastTx(account, msg)
    if transerr != nil {
        log.Fatal(err)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgCreateMemberToken:\n\n")
    fmt.Println(txResp)

	c.IndentedJSON(http.StatusOK, txResp)

}



// func createDiscountToken(c *gin.Context) {

func createDiscountBurritoToken(c *gin.Context) {

        // Define a message to create a discount token
    // msg := &types.MsgCreateDiscountToken{
    //     Creator: addr,
	// 	Timestamp: "timestamp",
	// 	ActivityName:      "ActivityName",
	// 	Score:             "Score",
	// 	Message:           "Message",
	// 	DiscountValue:     "DiscountValue",
	// 	EligibleCompanies: "EligibleCompanies",
	// 	ItemType:          "ItemType",
	// 	ExpiryDate:        "ExpiryDate",
    // }
        // Define a message to create a discount token
    msg := &types.MsgCreateDiscountToken{
        Creator: addr,
	 	Timestamp: "timestamp",
        ActivityName:      "Learn to make tacos",
        Score:             "10",
        Message:           "Excelente",
        DiscountValue:     "5",
        ItemType:          "protein burrito cooking class",
        ExpiryDate:        "5th Dec 2022",
    }

    // Broadcast a transaction from account `alice` with the message
    // to create a post store response in txResp
    txResp, transerr := cosmos.BroadcastTx(account, msg)
    if transerr != nil {
        log.Fatal(err)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgCreateDiscountToken:\n\n")
    fmt.Println(txResp)

	c.IndentedJSON(http.StatusOK, txResp)
}

func main() {
    if err != nil {
        log.Fatal(err)
    }

    if accounterr != nil {
        log.Fatal(accounterr)
    }

    if addresserr != nil {
        log.Fatal(addresserr)
    }

    router := gin.Default()
    router.Use(CORSMiddleware())

	router.GET("/books", getBooks)
    // router.GET("/token", createDiscountToken)
    router.GET("/tokens", getTokens)
    router.GET("/user", getUser)
    router.GET("/discountBurritoToken", createDiscountBurritoToken)
    router.GET("/discountMembershipToken", createDiscountMembershipToken)
	router.Run("localhost:8080")

}
