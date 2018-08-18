/**
 * $Date: 2018-08-13 12:16:58 -0700 (Mon, 13 Aug 2018) $
 * $Revision: 9461 $
 * $Author: dwong $
 * $HeadURL: http://svn/svn/incubator/blockchain/colorado/smart-contract/chaincode/fabric1.2/go/colorado.go $
 */

 package main

 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
	 "strings"
	 "time"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Colorado POC chaincode implementation
 type Colorado struct {
 }
 
 type trade struct {
	 ObjectType        string  `json:"docType"` // docType is used to distinguish the various types of objects in state database
	 TradeId           string  `json:"tradeId"` // the json fieldtags are needed to keep case from bouncing around
	 Source            string  `json:"source"` // the json fieldtags are needed to keep case from bouncing around
	 BuySide           string  `json:"buySide"` // the json fieldtags are needed to keep case from bouncing around
	 BuySideTradeId    string  `json:"buySideTradeId"` // the json fieldtags are needed to keep case from bouncing around
	 SellSide          string  `json:"sellSide"` // the json fieldtags are needed to keep case from bouncing around
	 SellSideTradeId   string  `json:"sellSideTradeId"` // the json fieldtags are needed to keep case from bouncing around
	 InvestmentManager string  `json:"investmentManager"` // the json fieldtags are needed to keep case from bouncing around
	 ExecutingBroker   string  `json:"executingBroker"` // the json fieldtags are needed to keep case from bouncing around
	 Product           string  `json:"product"` // the json fieldtags are needed to keep case from bouncing around
	 SubProduct        string  `json:"subProduct"` // the json fieldtags are needed to keep case from bouncing around
	 TradeDate         string  `json:"tradeDate"` // the json fieldtags are needed to keep case from bouncing around
	 Quantity          int     `json:"quantity"` // the json fieldtags are needed to keep case from bouncing around
	 Price             float64 `json:"price"` // the json fieldtags are needed to keep case from bouncing around
	 SecurityId        string  `json:"securityId"` // the json fieldtags are needed to keep case from bouncing around
	 Status            string  `json:"status"` // the json fieldtags are needed to keep case from bouncing around
	 StatusDetail      string  `json:"statusDetail"` // the json fieldtags are needed to keep case from bouncing around
 }
 
 type allocation struct {
	 ObjectType           string `json:"docType"` // docType is used to distinguish the various types of objects in state database
	 AllocationId         string `json:"allocationId"` // the json fieldtags are needed to keep case from bouncing around
	 InvestmentManager    string `json:"investmentManager"` // the json fieldtags are needed to keep case from bouncing around
	 ExecutingBroker      string `json:"executingBroker"` // the json fieldtags are needed to keep case from bouncing around
	 EndAccount           string `json:"endAccount"` // the json fieldtags are needed to keep case from bouncing around
	 Quantity             int    `json:"quantity"` // the json fieldtags are needed to keep case from bouncing around
	 Product              string `json:"product"` // the json fieldtags are needed to keep case from bouncing around
	 SubProduct           string `json:"subProduct"` // the json fieldtags are needed to keep case from bouncing around
	 SettleLocation       string `json:"settleLocation"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBank            string `json:"agentBank"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankName        string `json:"agentBankName"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankBIC         string `json:"agentBankBIC"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankAccount     string `json:"agentBankAccount"` // the json fieldtags are needed to keep case from bouncing around
	 CustodianBank        string `json:"custodianBank"` // the json fieldtags are needed to keep case from bouncing around
	 CustodianBankName    string `json:"custodianBankName"` // the json fieldtags are needed to keep case from bouncing around
	 CustodianBankBIC     string `json:"custodianBankBIC"` // the json fieldtags are needed to keep case from bouncing around
	 CustodianBankAccount string `json:"custodianBankAccount"` // the json fieldtags are needed to keep case from bouncing around
	 Status               string `json:"status"` // the json fieldtags are needed to keep case from bouncing around
	 StatusDetail         string `json:"statusDetail"` // the json fieldtags are needed to keep case from bouncing around
	 StatusTimestamp      time.Time `json:"statusTimestamp"` // the json fieldtags are needed to keep case from bouncing around
 }
 
 type brokerSSI struct {
	 ObjectType        string `json:"docType"` // docType is used to distinguish the various types of objects in state database
	 SSIId             string `json:"ssiId"` // the json fieldtags are needed to keep case from bouncing around
	 ExecutingBroker   string `json:"executingBroker"` // the json fieldtags are needed to keep case from bouncing around
	 InvestmentManager string `json:"investmentManager"` // the json fieldtags are needed to keep case from bouncing around
	 Product           string `json:"product"` // the json fieldtags are needed to keep case from bouncing around
	 SubProduct        string `json:"subProduct"` // the json fieldtags are needed to keep case from bouncing around
	 Currency          string `json:"currency"` // the json fieldtags are needed to keep case from bouncing around
	 SettleLocation    string `json:"settleLocation"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBank         string `json:"agentBank"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankName     string `json:"agentBankName"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankBIC      string `json:"agentBankBIC"` // the json fieldtags are needed to keep case from bouncing around
	 AgentBankAccount  string `json:"agentBankAccount"` // the json fieldtags are needed to keep case from bouncing around
 }
 
 func main() {
	 err := shim.Start(new(Colorado))
	 if err != nil {
		 fmt.Printf("Error starting Colorado chaincode: %s", err)
	 }
 }
 
 // Init - initialize chaincode
 func (t *Colorado) Init(stub shim.ChaincodeStubInterface) peer.Response {
	 return shim.Success(nil)
 }
 
 // Invoke - select entry point for invocations
 func (t *Colorado) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	 function, args := stub.GetFunctionAndParameters()
	 fmt.Println("invoke is running " + function)
 
	 // handle different invocations
	 switch function {
	 case "setTrade":
		 // create new trade
		 return t.setTrade(stub, args)
	 case "getMatchableTrades":
		 // get matchable trades
		 return t.getMatchableTrades(stub, args)
	 case "matchTradesByTradeIds":
		 // match trades
		 return t.matchTradesByTradeIds(stub, args)
	 case "getTradesByPartialCompositeKey":
		 // query trades
		 return t.getTradesByPartialCompositeKey(stub, args)
	 case "setBrokerSSI":
		 // create new Broker SSI
		 return t.setBrokerSSI(stub, args)
	 case "getBrokerSSIsByPartialCompositeKey":
		 // query Broker SSIs
		 return t.getBrokerSSIsByPartialCompositeKey(stub, args)
	 default:
		 // error
		 fmt.Println("invoke has no function named: " + function)
		 return shim.Error("Received unknown function invocation")
	 }
 }
 
 // setTrade - create new trade, store into chaincode state/indices, 
 // match trade by range query against a composite key index, and
 // return the Trade ID of this new trade in Response payload
 func (t *Colorado) setTrade(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 var err error
 
	 if len(args) != 13 {
		 return shim.Error("Incorrect number of arguments. Expecting 13.")
	 }
 
	 // input sanitation
	 fmt.Println("- start setting trade")
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Source) must be a non-empty string")
	 }
	 if len(args[1]) == 0 {
		 return shim.Error("2nd argument (Buy Side) must be a non-empty string")
	 }
	 if len(args[3]) == 0 {
		 return shim.Error("4th argument (Sell Side) must be a non-empty string")
	 }
	 if len(args[2]) != 0 && len(args[4]) != 0 ||
		len(args[2]) == 0 && len(args[4]) == 0 {
		 return shim.Error("3rd/5th argument (Buy Side Trade ID / Sell Side Trade ID) must be either, but not both, non-empty string")
	 }
	 if len(args[5]) == 0 {
		 return shim.Error("6th argument (Investment Manager) must be a non-empty string")
	 }
	 if len(args[6]) == 0 {
		 return shim.Error("7th argument (Executing Broker) must be a non-empty string")
	 }
	 if len(args[7]) == 0 {
		 return shim.Error("8th argument (Product) must be a non-empty string")
	 }
	 if len(args[8]) == 0 {
		 return shim.Error("9th argument (Sub-Product) must be a non-empty string")
	 }
	 if len(args[9]) == 0 {
		 return shim.Error("10th argument (Trade Date) must be a non-empty string")
	 }
	 if len(args[10]) == 0 {
		 return shim.Error("11th argument (Quantity) must be a non-empty string")
	 }
	 if len(args[11]) == 0 {
		 return shim.Error("12th argument (Price) must be a non-empty string")
	 }
	 if len(args[12]) == 0 {
		 return shim.Error("13th argument (Security ID) must be a non-empty string")
	 }
	 source := strings.ToUpper(args[0])
	 buySide := strings.ToUpper(args[1])
	 buySideTradeId := args[2]
	 sellSide := strings.ToUpper(args[3])
	 sellSideTradeId := args[4]
	 im := strings.ToUpper(args[5])
	 eb := strings.ToUpper(args[6])
	 if buySide != im && buySide != eb {
		 return shim.Error("2nd argument (Buy Side) must be the same as either 6th argument (Investment Manager) or 7th argument (Executing Broker)")
	 }
	 if sellSide != im && sellSide != eb {
		 return shim.Error("4th argument (Sell Side) must be the same as either 6th argument (Investment Manager) or 7th argument (Executing Broker)")
	 }
	 if buySide == sellSide {
		 return shim.Error("2nd argument (Buy Side) and 4th argument (Sell Side) must be different")
	 }
	 if source != buySide && source != sellSide {
		 return shim.Error("1st argument (Source) must be either the same as 2nd argument (Buy Side) or 4th argument (Sell Side)")
	 }
	 product := strings.ToUpper(args[7])
	 subProduct := strings.ToUpper(args[8])
	 tradeDate := strings.ToUpper(args[9])
	 quantity, err := strconv.Atoi(args[10])
	 if err != nil {
		 return shim.Error("11th argument (Quantity) must be an integer numeric string")
	 }
	 quantityAsString := args[10]
	 price, err := strconv.ParseFloat(args[11], 64)
	 if err != nil {
		 return shim.Error("12th argument (Price) must be a floating-point numeric string")
	 }
	 priceAsString := args[11]
	 securityId := args[12]
	 status := "CREATED"
	 statusDetail := ""
	 delimiter := "::"
	 thisTradeId := buySide + delimiter + buySideTradeId
	 if buySideTradeId == "" {
		 thisTradeId = sellSide + delimiter + sellSideTradeId
	 }
 
	 // check if trade already exists
	 privateCollection := "privateTradeFor" + im + eb
	 dupTradeAsBytes, err := stub.GetPrivateData(privateCollection, thisTradeId)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB
		 return shim.Error("Unable to get trade: " + err.Error())
	 } else if dupTradeAsBytes != nil {
		 fmt.Println("This trade already exists: " + thisTradeId)
		 return shim.Error("This trade already exists: " + thisTradeId)
	 }
 
	 // create trade object and marshal to JSON
	 objectType := "trade"
	 trade := &trade{objectType, thisTradeId, source, buySide, buySideTradeId, sellSide, sellSideTradeId, im, eb, product, subProduct, tradeDate, quantity, price, securityId, status, statusDetail}
	 tradeJSONAsBytes, err := json.Marshal(trade)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
 
	 // save trade to state
	 err = stub.PutPrivateData(privateCollection, thisTradeId, tradeJSONAsBytes)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB
		 return shim.Error(err.Error())
	 }
 
	 // Index the trade to enable key-based range queries,
	 // e.g. return all trades of a Source, a Buy Side, a Sell Side, a Product, a Sub-Product,
	 // a Trade Date, a Quantity, a Price, a Security ID, and a Status
	 // (An 'index' is a normal key/value entry in state.
	 //  The key is a composite key, with the elements that you want to range query on listed first.
	 //  This will enable very efficient state range queries based on composite keys being matched.)
	 indexName1 := "tradeMatchingCompositeKey"
	 indexCompositeKey1, err := stub.CreateCompositeKey(indexName1, []string{trade.Source, trade.BuySide, trade.SellSide, trade.Product, trade.SubProduct, trade.TradeDate, quantityAsString, priceAsString, trade.SecurityId, trade.Status})
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 // save index entry to state
	 stub.PutPrivateData(privateCollection, indexCompositeKey1, []byte(thisTradeId))
	 // If only the key name is needed, no need to store a duplicate copy of the trade.
	 // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	 // value := []byte{0x00}
	 // stub.PutPrivateData(privateCollection, indexCompositeKey1, value)
 
	 // Index the trade to enable key-based range queries with a partial composite key,
	 // e.g. return all trades of a Buy Side
	 // (An 'index' is a normal key/value entry in state.
	 //  The key is a composite key, with the elements that you want to range query on listed first.
	 //  This will enable very efficient state range queries based on composite keys being matched.)
	 indexName2 := "buySide~tradeId"
	 indexCompositeKey2, err := stub.CreateCompositeKey(indexName2, []string{trade.BuySide, trade.TradeId})
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 // save index entry to state
	 // If only the key name is needed, no need to store a duplicate copy of the trade.
	 // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	 value := []byte{0x00}
	 stub.PutPrivateData(privateCollection, indexCompositeKey2, value)
 
	 // Index the trade to enable key-based range queries with a partial composite key,
	 // e.g. return all trades of a Sell Side
	 // (An 'index' is a normal key/value entry in state.
	 //  The key is a composite key, with the elements that you want to range query on listed first.
	 //  This will enable very efficient state range queries based on composite keys being matched.)
	 indexName3 := "sellSide~tradeId"
	 indexCompositeKey3, err := stub.CreateCompositeKey(indexName3, []string{trade.SellSide, trade.TradeId})
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 // save index entry to state
	 // If only the key name is needed, no need to store a duplicate copy of the trade.
	 // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	 stub.PutPrivateData(privateCollection, indexCompositeKey3, value)
 
	 responseJSON := fmt.Sprintf("{\"investmentManager\":\"%s\",\"executingBroker\":\"%s\",\"tradeIdToBeMatched\":\"%s\",\"matchableTradeIds\":[%s]}", im, eb, thisTradeId, "")
	 fmt.Println("- end setting trade with response: " + responseJSON)
 
	 return shim.Success([]byte(responseJSON))
 }
 
 // setTrade - create new trade, store into chaincode state/indices, 
 // match trade by range query against a composite key index, and
 // return the Trade ID of this new trade in Response payload
 func (t *Colorado) getMatchableTrades(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1.")
	 }
 
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Request JSON) must be a non-empty string, i.e. {\"investmentManager\":\"IMi\",\"executingBroker\":\"EBj\",\"tradeIdToBeMatched\":\"IDn\",\"matchableTradeIds\":[\"IDx\",\"IDy\",\"IDz\"]}")
	 }
 
	 requestJSON := args[0]
	 fmt.Println("- start getting matchable trades with request: " + requestJSON)
 
	 type requstType struct {
		 InvestmentManager  string   `json:"investmentManager"` // the json fieldtags are needed to keep case from bouncing around
		 ExecutingBroker    string   `json:"executingBroker"` // the json fieldtags are needed to keep case from bouncing around
		 TradeIdToBeMatched string   `json:"tradeIdToBeMatched"` // the json fieldtags are needed to keep case from bouncing around
		 MatchableTradeIds  []string `json:"matchableTradeIds"` // the json fieldtags are needed to keep case from bouncing around
	 }
	 var request requstType
	 err := json.Unmarshal([]byte(requestJSON), &request)
	 if err != nil {
		 return shim.Error("Unable to parse request: " + err.Error())
	 }
 
	 if len(request.TradeIdToBeMatched) == 0 {
		 return shim.Success([]byte("There is no trade to be matched."))
	 }
 
	 // get trade to be matched
	 privateCollection := "privateTradeFor" + request.InvestmentManager + request.ExecutingBroker
	 tradeToBeMatchedAsBytes, err := stub.GetPrivateData(privateCollection, request.TradeIdToBeMatched)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB
		 return shim.Error("Unable to get trade: " + err.Error())
	 } else if tradeToBeMatchedAsBytes == nil {
		 fmt.Println("This trade does not exist: " + request.TradeIdToBeMatched)
		 return shim.Error("This trade does not exist: " + request.TradeIdToBeMatched)
	 }
	 tradeToBeMatched := trade{}
	 err = json.Unmarshal(tradeToBeMatchedAsBytes, &tradeToBeMatched) // unmarshal it aka JSON.parse()
	 if err != nil {
		 return shim.Error("Unable to parse Trade To Be Matched: " + err.Error())
	 }
 
	 // Query the tradeMatchingCompositeKey index by all fields in the composite key,
	 // e.g. return all trades of a Source, a Buy Side, a Sell Side, a Product, a Sub-Product,
	 // a Trade Date, a Quantity, a Price, a Security ID, and a Status
	 var buffer bytes.Buffer
	 delimiter := ","
	 sourceToBeMatched := tradeToBeMatched.BuySide
	 if tradeToBeMatched.Source == tradeToBeMatched.BuySide {
		 sourceToBeMatched = tradeToBeMatched.SellSide
	 }
	 quantityToBeMatchedAsString := strconv.Itoa(tradeToBeMatched.Quantity)
	 priceToBeMatchedAsString := strconv.FormatFloat(tradeToBeMatched.Price, 'f', -1, 64)
	 tradeResultsIterator, err := stub.GetPrivateDataByPartialCompositeKey(privateCollection, "tradeMatchingCompositeKey", []string{sourceToBeMatched, tradeToBeMatched.BuySide, tradeToBeMatched.SellSide, tradeToBeMatched.Product, tradeToBeMatched.SubProduct, tradeToBeMatched.TradeDate, quantityToBeMatchedAsString, priceToBeMatchedAsString, tradeToBeMatched.SecurityId, "CREATED"})
	 if err != nil {
		 return shim.Error("Unable to call GetPrivateDataByPartialCompositeKey: " + err.Error())
	 }
 
	 defer tradeResultsIterator.Close()
	 for tradeResultsIterator.HasNext() {
		 tradeResult, err := tradeResultsIterator.Next()
		 if err != nil {
			 return shim.Error("Unable to iterate through StateQueryIteratorInterface: " + err.Error())
		 }
		 if buffer.Len() > 0 {
			 buffer.WriteString(delimiter)
		 }
		 buffer.WriteString("\"" + string(tradeResult.Value) + "\"")
	 }
 
	 responseJSON := fmt.Sprintf("{\"investmentManager\":\"%s\",\"executingBroker\":\"%s\",\"tradeIdToBeMatched\":\"%s\",\"matchableTradeIds\":[%s]}", request.InvestmentManager, request.ExecutingBroker, request.TradeIdToBeMatched, buffer.String())
	 fmt.Println("- end getting matchable trades with response: " + responseJSON)
 
	 return shim.Success([]byte(responseJSON))
 }
 
 // matchTradesByTradeIds - 
 func (t *Colorado) matchTradesByTradeIds(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1.")
	 }
 
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Request JSON) must be a non-empty string, i.e. {\"investmentManager\":\"IMi\",\"executingBroker\":\"EBj\",\"tradeIdToBeMatched\":\"IDn\",\"matchableTradeIds\":[\"IDx\",\"IDy\",\"IDz\"]}")
	 }
 
	 requestJSON := args[0]
	 fmt.Println("- start matching trades by Trade IDs with request: " + requestJSON)
 
	 type requstType struct {
		 InvestmentManager  string   `json:"investmentManager"` // the json fieldtags are needed to keep case from bouncing around
		 ExecutingBroker    string   `json:"executingBroker"` // the json fieldtags are needed to keep case from bouncing around
		 TradeIdToBeMatched string   `json:"tradeIdToBeMatched"` // the json fieldtags are needed to keep case from bouncing around
		 MatchableTradeIds  []string `json:"matchableTradeIds"` // the json fieldtags are needed to keep case from bouncing around
	 }
	 var request requstType
	 err := json.Unmarshal([]byte(requestJSON), &request)
	 if err != nil {
		 return shim.Error("Unable to parse request: " + err.Error())
	 }
 
	 if len(request.MatchableTradeIds) == 0 {
		 return shim.Success([]byte("There is no matchable trade."))
	 }
 
	 // get trade to be matched
	 privateCollection := "privateTradeFor" + request.InvestmentManager + request.ExecutingBroker
	 tradeToBeMatchedAsBytes, err := stub.GetPrivateData(privateCollection, request.TradeIdToBeMatched)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB
		 return shim.Error("Unable to get trade: " + err.Error())
	 } else if tradeToBeMatchedAsBytes == nil {
		 fmt.Println("This trade does not exist: " + request.TradeIdToBeMatched)
		 return shim.Error("This trade does not exist: " + request.TradeIdToBeMatched)
	 }
	 tradeToBeMatched := trade{}
	 err = json.Unmarshal(tradeToBeMatchedAsBytes, &tradeToBeMatched) // unmarshal it aka JSON.parse()
	 if err != nil {
		 return shim.Error("Unable to parse Trade To Be Matched: " + err.Error())
	 }
 
	 // get and re-compare all matchable trades with the trade to be matched
	 var matchableTradeAsBytes []byte
	 for i := 0; i < len(request.MatchableTradeIds); i++ {
		 matchableTradeAsBytes, err = stub.GetPrivateData(privateCollection, request.MatchableTradeIds[i])
		 if err != nil {
			 // e.g. no defined private collection for the combination of IM, EB
			 return shim.Error("Unable to get trade: " + err.Error())
		 } else if matchableTradeAsBytes == nil {
			 fmt.Println("This trade does not exist: " + request.MatchableTradeIds[i])
			 return shim.Error("This trade does not exist: " + request.MatchableTradeIds[i])
		 }
		 matchableTrade := trade{}
		 err = json.Unmarshal(matchableTradeAsBytes, &matchableTrade) // unmarshal it aka JSON.parse()
		 if err != nil {
			 return shim.Error("Unable to parse Matchable Trade: " + err.Error())
		 }
 
		 // re-compare this matchable trade with the trade to be matched
		 if matchableTrade.Source != tradeToBeMatched.Source &&
			(matchableTrade.Source == matchableTrade.BuySide || matchableTrade.Source == matchableTrade.SellSide) &&
			(tradeToBeMatched.Source == tradeToBeMatched.BuySide || tradeToBeMatched.Source == tradeToBeMatched.SellSide) &&
			matchableTrade.BuySide == tradeToBeMatched.BuySide &&
			matchableTrade.SellSide == tradeToBeMatched.SellSide &&
			matchableTrade.Product == tradeToBeMatched.Product &&
			matchableTrade.SubProduct == tradeToBeMatched.SubProduct &&
			matchableTrade.TradeDate == tradeToBeMatched.TradeDate &&
			matchableTrade.Quantity == tradeToBeMatched.Quantity &&
			matchableTrade.Price == tradeToBeMatched.Price &&
			matchableTrade.SecurityId == tradeToBeMatched.SecurityId {
			 if matchableTrade.Status == "CREATED" && tradeToBeMatched.Status == "CREATED" {
				 // update both with matching details
				 matchableTrade.Status = "MATCHED"
				 tradeToBeMatched.Status = "MATCHED"
				 matchableTrade.StatusDetail = "Matched with Trade ID: " + tradeToBeMatched.TradeId
				 tradeToBeMatched.StatusDetail = "Matched with Trade ID: " + matchableTrade.TradeId
				 if matchableTrade.SellSideTradeId == "" {
					 matchableTrade.SellSideTradeId = tradeToBeMatched.TradeId
				 } else {
					 matchableTrade.BuySideTradeId = tradeToBeMatched.TradeId
				 }
				 if tradeToBeMatched.SellSideTradeId == "" {
					 tradeToBeMatched.SellSideTradeId = matchableTrade.TradeId
				 } else {
					 tradeToBeMatched.BuySideTradeId = matchableTrade.TradeId
				 }
 
				 // save the trade to be matched to state
				 tradeToBeMatchedJSONAsBytes, err := json.Marshal(tradeToBeMatched)
				 if err != nil {
					 return shim.Error("Unable to serialize Trade To Be Matched" + err.Error())
				 }
				 err = stub.PutPrivateData(privateCollection, tradeToBeMatched.TradeId, tradeToBeMatchedJSONAsBytes)
				 if err != nil {
					 // e.g. no defined private collection for the combination of IM, EB
					 return shim.Error(err.Error())
				 }
			 } else if matchableTrade.Status == "CREATED" && tradeToBeMatched.Status == "MATCHED" {
				 // put this matchable trade on-error because a previous matchable trade
				 // was already matched against the trade to be matched
				 matchableTrade.Status = "ON-ERROR"
				 matchableTrade.StatusDetail = "Too many matches with trade of Trade ID: " + tradeToBeMatched.TradeId
			 }
 
			 // save this matchable trade to state
			 matchableTradeJSONAsBytes, err := json.Marshal(matchableTrade)
			 if err != nil {
				 return shim.Error(err.Error())
			 }
			 err = stub.PutPrivateData(privateCollection, matchableTrade.TradeId, matchableTradeJSONAsBytes)
			 if err != nil {
				 // e.g. no defined private collection for the combination of IM, EB
				 return shim.Error(err.Error())
			 }
		 }
	 }
 
	 fmt.Println("- end matching trades by Trade IDs")
	 return shim.Success(nil)
 }
 
 // getTradesByPartialCompositeKey - get trades by range query against 2 composite key indices
 // Committing peers will re-execute range queries to guarantee that result sets are stable
 // between endorsement time and commit time. The transaction is invalidated by the
 // committing peers if the result set has changed between endorsement time and commit time.
 // Therefore, range queries are a safe option for performing update transactions based on query results.
 func (t *Colorado) getTradesByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1.")
	 }
 
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Entitled Organization) must be a non-empty string")
	 }
 
	 entitledOrg := strings.ToUpper(args[0])
	 fmt.Println("- start getting trades by partial composite key: " + entitledOrg)
 
	 // buffer is a JSON array containing query results
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 // 1. Query the buySide~tradeId index by partial fields (e.g. buy side IM1 only)
	 //    This will execute a key range query on all keys starting with 'IM1'
	 // 2. Query the sellSide~tradeId index by partial fields (e.g. sell side IM1 only)
	 //    This will execute a key range query on all keys starting with 'IM1'
	 allIMs := [1]string{"IM1"}
	 allEBs := [2]string{"EB1", "EB2"}
	 allIndices := [2]string{"buySide~tradeId", "sellSide~tradeId"}
	 for i := 0; i < len(allIMs); i++ {
		 for j := 0; j < len(allEBs); j++ {
			 for k := 0; k < len(allIndices); k++ {
				 tradeResultsIterator, err := stub.GetPrivateDataByPartialCompositeKey("privateTradeFor" + allIMs[i] + allEBs[j], allIndices[k], []string{entitledOrg})
				 if err != nil {
					 // not return error for non-entitled privateTradeForIMxEBy
					 // return shim.Error(err.Error())
				 } else {
					 defer tradeResultsIterator.Close()
					 for tradeResultsIterator.HasNext() {
						 tradeResult, err := tradeResultsIterator.Next()
 
						 // get the color and name from color~name composite key
						 _, compositeKeyParts, err := stub.SplitCompositeKey(tradeResult.Key)
						 if err != nil {
							 return shim.Error(err.Error())
						 }
						 // get split parts of "buySide~tradeId" and "sellSide~tradeId"
						 // tradeBuyOrSellSide := compositeKeyParts[0] // no need
						 tradeId := compositeKeyParts[1]
 
						 if err != nil {
							 return shim.Error(err.Error())
						 }
						 if buffer.Len() > 1 {
							 buffer.WriteString(",")
						 }
						 //buffer.WriteString("{\"compositeKey\":")
						 //buffer.WriteString("\"")
						 //buffer.WriteString(tradeResult.Key)
						 //buffer.WriteString("\"")
 
						 //buffer.WriteString(", \"trade\":")
						 tradeJSONAsBytes, err := stub.GetPrivateData("privateTradeFor" + allIMs[i] + allEBs[j], tradeId)
						 if err != nil {
							 // e.g. no defined private collection for the combination of IM, EB
							 return shim.Error("Unable to get trade: " + err.Error())
						 } else if tradeJSONAsBytes == nil {
							 fmt.Println("This trade does not exist: " + tradeId)
							 return shim.Error("This trade does not exist: " + tradeId)
						 }
						 buffer.WriteString(string(tradeJSONAsBytes))
 
						 //buffer.WriteString("}")
					 }
				 }
			 }
		 }
	 }
	 buffer.WriteString("]")
 
	 fmt.Println("- end getting trades with query result:\n" + buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 // setBrokerSSI - create new Broker SSI, store into chaincode state
 func (t *Colorado) setBrokerSSI(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 var err error
 
	 if len(args) != 10 {
		 return shim.Error("Incorrect number of arguments. Expecting 10.")
	 }
 
	 // input sanitation
	 fmt.Println("- start setting Broker SSI")
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Executing Broker) must be a non-empty string")
	 }
	 if len(args[1]) == 0 {
		 return shim.Error("2nd argument (Investment Manager) must be a non-empty string")
	 }
	 if len(args[6]) == 0 {
		 return shim.Error("7th argument (Agent Bank) must be a non-empty string")
	 }
	 if len(args[7]) == 0 {
		 return shim.Error("8th argument (Agent Bank Name) must be a non-empty string")
	 }
	 if len(args[8]) == 0 {
		 return shim.Error("9th argument (Agent Bank BIC) must be a non-empty string")
	 }
	 if len(args[9]) == 0 {
		 return shim.Error("10th argument (Agent Bank Account) must be a non-empty string")
	 }
	 eb := strings.ToUpper(args[0])
	 im := strings.ToUpper(args[1])
	 product := strings.ToUpper(args[2])
	 subProduct := strings.ToUpper(args[3])
	 currency := strings.ToUpper(args[4])
	 settleLocation := strings.ToUpper(args[5])
	 ab := strings.ToUpper(args[6])
	 agentBankName := strings.ToUpper(args[7])
	 agentBankBIC := strings.ToUpper(args[8])
	 agentBankAccount := strings.ToUpper(args[9])
	 delimiter := "::"
	 ssiId := eb + delimiter + im + delimiter + product + delimiter + subProduct + delimiter + currency + delimiter + settleLocation
 
	 // check if Broker SSI already exists
	 privateCollection := "privateBrokerSSIFor" + im + eb + ab
	 // TODO: handle if IM is nil (wildcard) for all "privateBrokerSSIFor" for different IMs
	 brokerSSIAsBytes, err := stub.GetPrivateData(privateCollection, ssiId)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB, AB
		 return shim.Error("Unable to get Broker SSI: " + err.Error())
	 } else if brokerSSIAsBytes != nil {
		 fmt.Println("This Broker SSI already exists: " + ssiId)
		 return shim.Error("This Broker SSI already exists: " + ssiId)
	 }
 
	 // create Broker SSI object and marshal to JSON
	 objectType := "brokerSSI"
	 brokerSSI := &brokerSSI{objectType, ssiId, eb, im, product, subProduct, currency, settleLocation, ab, agentBankName, agentBankBIC, agentBankAccount}
	 brokerSSIJSONAsBytes, err := json.Marshal(brokerSSI)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
 
	 // save Broker SSI to state
	 err = stub.PutPrivateData(privateCollection, ssiId, brokerSSIJSONAsBytes)
	 if err != nil {
		 // e.g. no defined private collection for the combination of IM, EB, AB
		 return shim.Error(err.Error())
	 }
 
	 // Index the brokerSSI to enable key-based range queries,
	 // e.g. return all Broker SSI of an EB, an IM, a Product, a Sub-Product, a Currency, and a Settle Location
	 // (An 'index' is a normal key/value entry in state.
	 //  The key is a composite key, with the elements that you want to range query on listed first.
	 //  This will enable very efficient state range queries based on composite keys being matched.)
	 indexName := "brokerSSICompositeKey"
	 indexCompositeKey, err := stub.CreateCompositeKey(indexName, []string{brokerSSI.ExecutingBroker, brokerSSI.InvestmentManager, brokerSSI.Product, brokerSSI.SubProduct, brokerSSI.Currency, brokerSSI.SettleLocation})
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 // save index entry to state
	 stub.PutPrivateData(privateCollection, indexCompositeKey, brokerSSIJSONAsBytes)
	 // If only the key name is needed, no need to store a duplicate copy of the Broker SSI.
	 // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	 // value := []byte{0x00}
	 // stub.PutPrivateData(privateCollection, indexCompositeKey, value)
 
	 // Broker SSI is saved and indexed
	 fmt.Println("- end setting Broker SSI")
	 return shim.Success(nil)
 }
 
 // getBrokerSSIsByPartialCompositeKey - get Broker SSIs by range query against a composite key index
 // Committing peers will re-execute range queries to guarantee that result sets are stable
 // between endorsement time and commit time. The transaction is invalidated by the
 // committing peers if the result set has changed between endorsement time and commit time.
 // Therefore, range queries are a safe option for performing update transactions based on query results.
 func (t *Colorado) getBrokerSSIsByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	 if len(args) != 6 {
		 return shim.Error("Incorrect number of arguments. Expecting 6.")
	 }
 
	 if len(args[0]) == 0 {
		 return shim.Error("1st argument (Executing Broker) must be a non-empty string")
	 }
 
	 eb := strings.ToUpper(args[0])
	 im := strings.ToUpper(args[1])
	 product := strings.ToUpper(args[2])
	 subProduct := strings.ToUpper(args[3])
	 currency := strings.ToUpper(args[4])
	 settleLocation := strings.ToUpper(args[5])
	 fmt.Println("- start getting Broker SSIs by partial composite key: " + eb + ", " + im + ", " + product + ", " + subProduct + ", " + currency + ", " + settleLocation)
 
	 // buffer is a JSON array containing query results
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 // Query the brokerSSICompositeKey index by partial fields (e.g. EB1 only)
	 // This will execute a key range query on all keys starting with 'EB1'
	 allIMs := [1]string{"IM1"}
	 allEBs := [2]string{"EB1", "EB2"}
	 allABs := [2]string{"AB1", "AB2"}
	 indexName := "brokerSSICompositeKey"
	 for i := 0; i < len(allIMs); i++ {
		 for j := 0; j < len(allEBs); j++ {
			 for k := 0; k < len(allABs); k++ {
				 brokerSSIResultsIterator, err := stub.GetPrivateDataByPartialCompositeKey("privateBrokerSSIFor" + allIMs[i] + allEBs[j] + allABs[k], indexName, []string{eb, im, product, subProduct, currency, settleLocation})
				 if err != nil {
					 // not return error for non-entitled privateBrokerSSIForIMxEByABz
					 // return shim.Error(err.Error())
				 } else {
					 defer brokerSSIResultsIterator.Close()
					 for brokerSSIResultsIterator.HasNext() {
						 brokerSSIResult, err := brokerSSIResultsIterator.Next()
						 if err != nil {
							 return shim.Error(err.Error())
						 }
						 if buffer.Len() > 1 {
							 buffer.WriteString(",")
						 }
						 buffer.WriteString("{\"compositeKey\":")
						 buffer.WriteString("\"")
						 buffer.WriteString(brokerSSIResult.Key)
						 buffer.WriteString("\"")
 
						 buffer.WriteString(", \"brokerSSI\":")
						 // record is a JSON object, so write as-is
						 buffer.WriteString(string(brokerSSIResult.Value))
						 buffer.WriteString("}")
					 }
				 }
			 }
		 }
	 }
	 buffer.WriteString("]")
 
	 fmt.Println("- end getting Broker SSIs with query result:\n" + buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }