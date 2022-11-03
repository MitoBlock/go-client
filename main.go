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

var addressPrefix = "mito"

   
var cosmos, err = cosmosclient.New(
    context.Background(),
    cosmosclient.WithAddressPrefix(addressPrefix),
)

var accountName = "bob"
  
var account, accounterr = cosmos.Account(accountName)

var addr, addrerr = account.Address(addressPrefix)
	
var queryClient = types.NewQueryClient(cosmos.Context())

// discount token from angular
type discountToken struct {
	Timestamp         string `json:"timestamp"`
	ActivityName      string `json:"activity_name"`
	Score             string `json:"score"`
	Message           string `json:"message"`
	DiscountValue     string `json:"discount_value"`
	EligibleCompanies string `json:"eligible_companies"`
	ItemType          string `json:"item_type"`
	ExpiryDate        string `json:"expiry_date"`
}

// discount token status from angular
type DiscountTokenStatus struct {
	TokenID   uint64 `json:"token_id"`
	ID        uint64 `json:"id"`
	Timestamp string `json:"timestamp"`
	status    string `json:"status"`
}

// membership token from angular
type membershipToken struct {
	Timestamp          string `json:"timestamp"`
	ActivityName       string `json:"activity_name"`
	Score              string `json:"score"`
	Message            string `json:"message"`
	MembershipDuration string `json:"membership_duration"`
	EligibleCompanies  string `json:"eligible_companies"`
	ExpiryDate         string `json:"expiry_date"`
}

// membership token status from angular
type MembershipTokenStatus struct {
	TokenID   uint64 `json:"token_id"`
	ID        uint64 `json:"id"`
	Timestamp string `json:"timestamp"`
	status    string `json:"status"`
}

var tokens = []discountToken{
	{
		Timestamp:         "timestamp",
		ActivityName:      "activity_name",
		Score:             "10",
		Message:           "message",
		DiscountValue:     "discount_value",
		EligibleCompanies: "eligible_companies",
		ItemType:          "item_type",
		ExpiryDate:        "expiry_date",
	},
}

func createMembershipToken(c *gin.Context) {
	var newMembershipToken membershipToken

	if err := c.BindJSON(&newMembershipToken); err != nil {
		return
	}

	msg := &types.MsgCreateMembershipToken{
		Creator:            addr,
		Timestamp:          newMembershipToken.Timestamp,
		ActivityName:       newMembershipToken.ActivityName,
		Score:              newMembershipToken.Score,
		Message:            newMembershipToken.Message,
		MembershipDuration: newMembershipToken.MembershipDuration,
		EligibleCompanies:  newMembershipToken.EligibleCompanies,
		ExpiryDate:         newMembershipToken.ExpiryDate,
	}

	txResp, transerr := cosmos.BroadcastTx(account, msg)
	if transerr != nil {
		log.Fatal(transerr)
	}

	fmt.Print("MsgCreateMembershipToken:\n\n")
	fmt.Println(txResp)

	msgStatus := &types.MsgCreateMembershipTokenStatus{
		Creator:   addr,
		TokenID:   0,
		Timestamp: newMembershipToken.Timestamp,
		Status:    "Valid",
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
	var newDiscountToken discountToken

	if err := c.BindJSON(&newDiscountToken); err != nil {
		return
	}

	// Define a message to create a discount token
	msg := &types.MsgCreateDiscountToken{
		Creator:           addr,
		Timestamp:         newDiscountToken.Timestamp,
		ActivityName:      newDiscountToken.ActivityName,
		Score:             newDiscountToken.Score,
		Message:           newDiscountToken.Message,
		DiscountValue:     newDiscountToken.DiscountValue,
		EligibleCompanies: newDiscountToken.EligibleCompanies,
		ItemType:          newDiscountToken.ItemType,
		ExpiryDate:        newDiscountToken.ExpiryDate,
	}

	txResp, transerr := cosmos.BroadcastTx(account, msg)
	if transerr != nil {
		log.Fatal(transerr)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreateDiscountToken:\n\n")
	fmt.Println(txResp)

	msgStatus := &types.MsgCreateDiscountTokenStatus{
		Creator:   addr,
		TokenID:   0,
		Timestamp: newDiscountToken.Timestamp,
		Status:    "Valid",
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

	var newDiscountTokenStatus DiscountTokenStatus

	if err := c.BindJSON(&newDiscountTokenStatus); err != nil {
		return
	}

	msgStatus := &types.MsgDeleteDiscountTokenStatus{
		Creator:               addr,
		DiscountTokenStatusID: newDiscountTokenStatus.ID,
		TokenID:               newDiscountTokenStatus.TokenID,
	}

	txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
	if transerrStatus != nil {
		log.Fatal(transerrStatus)
	}

	fmt.Print("MsgDeleteDiscountTokenStatus:\n\n")
	fmt.Println(txRespStatus)

	msgNewStatus := &types.MsgCreateDiscountTokenStatus{
		Creator:   addr,
		TokenID:   newDiscountTokenStatus.TokenID,
		Timestamp: newDiscountTokenStatus.Timestamp,
		Status:    newDiscountTokenStatus.Status,
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

	var newMembershipTokenStatus MembershipTokenStatus

	if err := c.BindJSON(&newMembershipTokenStatus); err != nil {
		return
	}

	msgStatus := &types.MsgDeleteMembershipTokenStatus{
		Creator:                 addr,
		MembershipTokenStatusID: newMembershipTokenStatus.ID,
		TokenID:                 newMembershipTokenStatus.TokenID,
	}

	txRespStatus, transerrStatus := cosmos.BroadcastTx(account, msgStatus)
	if transerrStatus != nil {
		log.Fatal(transerrStatus)
	}

	fmt.Print("MsgDeleteMembershipTokenStatus:\n\n")
	fmt.Println(txRespStatus)

	msgNewStatus := &types.MsgCreateMembershipTokenStatus{
		Creator:   addr,
		TokenID:   newMembershipTokenStatus.TokenID,
		Timestamp: newMembershipTokenStatus.Timestamp,
		Status:    newMembershipTokenStatus.Status,
	}

	txRespNewStatus, transerrNewStatus := cosmos.BroadcastTx(account, msgNewStatus)
	if transerrNewStatus != nil {
		log.Fatal(transerrStatus)
	}

	fmt.Print("MsgCreateMembershipTokenStatus:\n\n")
	fmt.Println(txRespNewStatus)

	c.IndentedJSON(http.StatusOK, txRespStatus)
}

func getAddr(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, addr)
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/addr", getAddr)
	router.POST("/discountToken", createDiscountToken)
	router.POST("/membershipToken", createMembershipToken)
	router.POST("/discountTokenStatus", deleteDiscountTokenStatus)
	router.POST("/membershipTokenStatus", deleteMembershipTokenStatus)
	router.Run("localhost:8080")
}