package main

import (

    "net/http"
	"github.com/gin-gonic/gin"

    "context"
    "fmt"
    "log"

    // Importing the general purpose Cosmos blockchain client
    "github.com/ignite/cli/ignite/pkg/cosmosclient"

    // Importing the types package of blockchain
    "mitoblockchaindev/x/mitoblockchaindev/types"
)

// middleware to allow incoming requests from all ports
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        // we do not allow options
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        // forward request to next handler
        c.Next()
    }
}

// Prefix to use for account addresses.
    // The address prefix was assigned to the blog blockchain
    // using the `--address-prefix` flag during scaffolding.
    var addressPrefix = "mito"

    // Create a Cosmos client instance
    var cosmos, err = cosmosclient.New(
        context.Background(),
        cosmosclient.WithAddressPrefix(addressPrefix),
    )

    // Account `bob` was initialized during `ignite chain serve`
    var accountName = "bob"

    // Get account from the keyring
    var account, accounterr = cosmos.Account(accountName)

    var addr, addrerr = account.Address(addressPrefix)
	
	var queryClient = types.NewQueryClient(cosmos.Context())

func createMembershipToken(c *gin.Context) {
    
    msg := &types.MsgCreateMembershipToken{
        Creator:           addr,
	 	Timestamp:         "2 NOv 2022",
        ActivityName:      "Weekly leaderboard",
        Score:             "10",
        Message:           "Impresionante",
        MembershipDuration:"3",
		EligibleCompanies: "Building Block Fitness",
        ExpiryDate:        "5th Dec 2022",
    }

    txResp, transerr := cosmos.BroadcastTx(account, msg)
    if transerr != nil {
        log.Fatal(transerr)
    }

    fmt.Print("MsgCreateMembershipToken:\n\n")
    fmt.Println(txResp)

	msgStatus := &types.MsgCreateMembershipTokenStatus{
        Creator:           addr,
		TokenID:           0,
		Timestamp:         "2 Nov 2022",
		Status:            "Valid",
    }

    txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
    if transerrStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgCreateMembershipTokenStatus:\n\n")
    fmt.Println(txRespStatus)

	c.IndentedJSON(http.StatusOK, txResp)

}

func createDiscountToken(c *gin.Context) {
    
    // Define a message to create a discount token
    msg := &types.MsgCreateDiscountToken{
        Creator:           addr,
	 	Timestamp:         "2 Nov 2022",
        ActivityName:      "Learn to make tacos",
        Score:             "10",
        Message:           "Excelente",
        DiscountValue:     "5",
		EligibleCompanies: "Cooking Academy",
        ItemType:          "protein burrito cooking class",
        ExpiryDate:        "5th Dec 2022",
    }

    txResp, transerr := cosmos.BroadcastTx(account, msg)
    if transerr != nil {
        log.Fatal(transerr)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgCreateDiscountToken:\n\n")
    fmt.Println(txResp)

	msgStatus := &types.MsgCreateDiscountTokenStatus{
        Creator:           addr,
		TokenID:           0,
		Timestamp:         "2 Nov 2022",
		Status:            "Valid",
    }

    txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
    if transerrStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgCreateDiscountTokenStatus:\n\n")
    fmt.Println(txRespStatus)

	c.IndentedJSON(http.StatusOK, txResp)
}

func deleteDiscountTokenStatus(c *gin.Context) {

	msgStatus := &types.MsgDeleteDiscountTokenStatus{
        Creator:               addr,
		DiscountTokenStatusID: 0,
		TokenID:               0,
    }

    txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
    if transerrStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgDeleteDiscountTokenStatus:\n\n")
    fmt.Println(txRespStatus)

	msgNewStatus := &types.MsgCreateDiscountTokenStatus{
        Creator:           addr,
		TokenID:           0,
		Timestamp:         "2 Nov 2022",
		Status:            "Invalid",
    }

    txRespNewStatus, transerrNewStatus := cosmos.BroadcastTx(account, msgNewStatus)
    if transerrNewStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgCreateDiscountTokenStatus:\n\n")
    fmt.Println(txRespNewStatus)

	c.IndentedJSON(http.StatusOK, txRespStatus)
}

func deleteMembershipTokenStatus(c *gin.Context) {

	msgStatus := &types.MsgDeleteMembershipTokenStatus{
        Creator:               addr,
		MembershipTokenStatusID: 0,
		TokenID:                 0,
    }

    txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
    if transerrStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgDeleteMembershipTokenStatus:\n\n")
    fmt.Println(txRespStatus)

	msgNewStatus := &types.MsgCreateMembershipTokenStatus{
        Creator:           addr,
		TokenID:           0,
		Timestamp:         "2 Nov 2022",
		Status:            "Invalid",
    }

    txRespNewStatus, transerrNewStatus := cosmos.BroadcastTx(account, msgNewStatus)
    if transerrNewStatus != nil {
        log.Fatal(transerrStatus)
    }

    fmt.Print("MsgCreateMembershipTokenStatus:\n\n")
    fmt.Println(txRespNewStatus)


	c.IndentedJSON(http.StatusOK, txRespStatus)
}

func main() {
    router := gin.Default()
    router.Use(CORSMiddleware())

    router.GET("/discountToken", createDiscountToken)
    router.GET("/membershipToken", createMembershipToken)
    router.GET("/deleteDiscountTokenStatus", deleteDiscountTokenStatus)
	router.GET("/deleteMembershipTokenStatus", deleteMembershipTokenStatus)
	router.Run("localhost:8080")
}
