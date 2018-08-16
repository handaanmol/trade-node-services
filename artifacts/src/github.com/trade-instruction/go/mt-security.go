package main 
import (
	"encoding/json"
	"fmt"
	"strings"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


// Implements a simple chaincode 
// TradeInstruction example simple Chaincode implementation
type TradeInstruction struct {
}

type brokerSSI struct {
	ObjectType string `json:"docType"`
	SsiId string `json:"ssiId"`
    InvestmentManager string `json:"investmentManager"`
	ExecutingBroker string `json:"executingBroker"`
	Product string `json:"product"`
	SubProduct string `json:"subProduct"`
	AgentBank string `json:"agentBank"`
	AgentBankName string `json:"agentBankName"`
	AgentBankAccount string `json:"agentBankAccount"`
	AgentBankBic string `json:"agentBankBic"`
	SettleLocation   string `json:"settleLocation"`
}

type allocation struct {
	ObjectType string `json:"docType"`
	AllocID string `json:"allocId"`
	TradeID string `json:"tradeId"`
	InvestmentManager string `json:"investmentManager"`
	ExecutingBroker string `json:"executingBroker"`
	AgentBank string `json:"agentBank"`
	AgentBankName string `json:"agentBankName"`
	AgentBankAccount string `json:"agentBankAccount"`
	AgentBankBic string `json:"agentBankBic"`
	CustodianBank string `json:"custodianBank"`
	CustodianBankName string `json:"custodianBankName"`
	CustodianBankAccount string `json:"custodianBankAccount"`
	CustodianBic string `json:"custodianBic"`
	EndAccount   string `json:"endAccount"`
	Product string `json:"product"`
	SubProduct string `json:"subProduct"`
	Security string `json:"security"`
	TradeDate string `json:"tradeDate"`
	SettleDate string`json:"settleDate"`
	SettleLocation string `json:"settleLocation"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
	StatusDetail string `json:"statusDetail"`
}




func main() {
	err :=shim.Start(new(TradeInstruction))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *TradeInstruction) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Smart contract is starting up...")
	return shim.Success(nil)
}


//Invoke method to match the calling function with the actual function
func (t *TradeInstruction) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is running " + function)

	// Handle different functions
	switch function {
	case "setAllocation":
		return t.setAllocation(stub, args)
	case "initBrokerSSI":
		return t.initBrokerSSI(stub, args)
	case "getAllocationByID":
		return t.getAllocationByID(stub, args)
	/*case "editBrokerSSI":
		return t.editBrokerSSI(stub, args)
	case "deleteBrokerSSI":
		return t.deleteBrokerSSI(stub, args)
	case "getBrokerSSI":
		return t.getBICgetBrokerSSIbyAccountNumber(stub, args)
	*/
	default:
		//error
		fmt.Println("invoke did not find function: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

// =============================================================================
// initSecurities - create a new security record on the ledger from the allocation
// ==============================================================================
func (t *TradeInstruction) setAllocation(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	//  0-allocID 1-tradeID 3-investmentManager 4-executingBroker 11-EndAccount 
	// 12-Product 13-SubProduct 14-Security 15-TradeDate 16-SettleDate 17-SettleLocation 
	// "asdf",  "blue",  "35",   "bob",   "99"
	if len(args) != 12{
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	// Shall we make this a for loop? Need to do more research on how to do that in Go. Or should we 
	// only provide the check for fields that really matter?
	fmt.Println("- start init securities")
	if len(args[0]) == 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) == 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) == 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) == 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) == 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) == 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) == 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	if len(args[7]) == 0 {
		return shim.Error("8th argument must be a non-empty string")
	}
	if len(args[8]) == 0 {
		return shim.Error("9th argument must be a non-empty string")
	}
	if len(args[9]) == 0 {
		return shim.Error("10th argument must be a non-empty string")
	}
	if len(args[10]) == 0 {
		return shim.Error("11th argument must be a non-empty string")
	}
	if len(args[11]) == 0 {
		return shim.Error("12th argument must be a non-empty string")
	}

	
	tradeID := args[0]
	investmentManager := strings.ToUpper(args[1])
	executingBroker := strings.ToUpper(args[2])
	endAccount := strings.ToUpper(args[3])
	product := strings.ToUpper(args[4])
	subProduct := strings.ToUpper(args[5])
	security := strings.ToUpper(args[6])
	tradeDate := args[7]
	settleDate := args[8]
	settleLocation := strings.ToUpper(args[9])
	quantity,err := strconv.Atoi(args[10])
	agentBank:= strings.ToUpper(args[11])
	allocID:= investmentManager + tradeID + endAccount

	if err != nil {
		return shim.Error("11th argument must be a numeric string")
	}
	status := "CREATED"
	statusDetail:= "Trade Instruction is created"
	privateCollection := "collectionAllocationFor" + executingBroker + investmentManager + agentBank
    // TODO: handle if investmentManager is nil (wildcard) for all existing Broker SSI collections for different investmentManager
    collectionAllocationAsBytes, err := stub.GetPrivateData(privateCollection, allocID)
    if err != nil {
        // e.g. no defined private collection for the combination of executingBroker, investmentManager, agentBank
        return shim.Error("Failed to get allocation: " + err.Error())
    } else if collectionAllocationAsBytes != nil {
        fmt.Println("This allocation already exists: " + allocID)
        return shim.Error("This allocation already exists: " + allocID)
    }

    // Create Broker SSI object and marshal to JSON
    objectType := "allocation"
	allocation := &allocation{objectType, allocID, tradeID, investmentManager, executingBroker, agentBank, "AB", "AB", "AB","AB", "AB", "AB", "", endAccount, product, subProduct, security, tradeDate, settleDate, settleLocation, quantity,status,statusDetail}
    allocationJSONasBytes, err := json.Marshal(allocation)
    if err != nil {
        return shim.Error(err.Error())
    }

    // Save Broker SSI to state
    err = stub.PutPrivateData(privateCollection, allocID, allocationJSONasBytes)
    if err != nil {
        // e.g. no defined private collection for the combination of executingBroker, investmentManager, agentBank
        return shim.Error(err.Error())
    }

    // Index the brokerSSI to enable key-based range queries, e.g. return all Broker SSI of an executingBroker, an investmentManager, a Product, a SubProduct, and a SettleLocation
    // (An 'index' is a normal key/value entry in state.
    //  The key is a composite key, with the elements that you want to range query on listed first.
    //  This will enable very efficient state range queries based on composite keys being matched.)
    indexName := "allocationKey"
    indexCompositeKey, err := stub.CreateCompositeKey(indexName, []string{allocation.InvestmentManager, allocation.TradeID, allocation.EndAccount})
    if err != nil {
        return shim.Error(err.Error())
    }
    // Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the Broker SSI.
    // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
    value := []byte{0x00}
    stub.PutPrivateData(privateCollection, indexCompositeKey, value)

    // Broker SSI saved and indexed. Return success.
    fmt.Println("- end set allocation")
    return shim.Success(nil)
	}


func (t *TradeInstruction) initBrokerSSI(stub shim.ChaincodeStubInterface, args []string) peer.Response {
// ==== Save security private details for agent ====
	var err error

	//  0-investmentManager 1-executingBroker 2-product 3-subProduct 4-agentBank
	// 5-agentBankName 6-agentBankAccount 7-agentBankBic 8-SsiId 9-SettleLocation
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// initBrokerSSI - create some new Broker SSI, store into chaincode state
	fmt.Println("- start init Broker SSI")
	if len(args[0]) == 0 {
        return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) == 0 {
        return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) == 0 {
        return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) == 0 {
        return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) == 0 {
        return shim.Error("5th argument must be a non-empty string")
    }
    if len(args[5]) == 0 {
        return shim.Error("6th argument must be a non-empty string")
    }
    if len(args[6]) == 0 {
        return shim.Error("7th argument must be a non-empty string")
    }
    if len(args[7]) == 0 {
        return shim.Error("8th argument must be a non-empty string")
    }
    if len(args[8]) == 0 {
        return shim.Error("9th argument must be a non-empty string")
    }
	
	investmentManager := strings.ToUpper(args[0])
	executingBroker := strings.ToUpper(args[1])
	product := strings.ToUpper(args[2])
	subProduct := strings.ToUpper(args[3])
	agentBank := strings.ToUpper(args[4])
	agentBankName := strings.ToUpper(args[5])
	agentBankAccount := strings.ToUpper(args[6])
	agentBankBic := strings.ToUpper(args[7])
	settleLocation := strings.ToUpper(args[8])
	ssiId := executingBroker + "\xbd" + investmentManager + "\xbd" + product + "\xbd" + subProduct + "\xbd" + settleLocation
	
	
	
	 // ==== Check if marble already exists and ====
	privateCollection := "collectionBrokerSSIFor" + executingBroker + investmentManager + agentBank
	// TODO: handle if investmentManager is nil (wildcard) for all existing Broker SSI collections for different investmentManager
	brokerSSIAsBytes, err := stub.GetPrivateData(privateCollection, ssiId)
	if err != nil {
		// e.g. no defined private collection for the combination of executingBroker, investmentManager, agentBank
		return shim.Error("Failed to get agentBankAccount: " + err.Error())
	} else if brokerSSIAsBytes != nil {
		fmt.Println("This Broker SSI already exists: " + ssiId)
        return shim.Error("This Broker SSI already exists: " + ssiId)
	}

	// Create Broker SSI object and marshal to JSON
	objectType := "brokerSSI"
	brokerSSI := &brokerSSI{objectType,ssiId, investmentManager, executingBroker, product, subProduct, agentBank, agentBankName, agentBankAccount, agentBankBic, settleLocation}
	brokerSSIJsonAsBytes, err := json.Marshal(brokerSSI)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Save Broker SSI to state
	err = stub.PutPrivateData("privateCollection", agentBankAccount, brokerSSIJsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Index the brokerSSI to enable key-based range queries, e.g. return all Broker SSI of an executingBroker, an investmentManager, a Product, a SubProduct, and a SettleLocation
    // (An 'index' is a normal key/value entry in state.
    //  The key is a composite key, with the elements that you want to range query on listed first.
    //  This will enable very efficient state range queries based on composite keys being matched.)
    indexName := "brokerSSIKey"
    indexCompositeKey, err := stub.CreateCompositeKey(indexName, []string{brokerSSI.ExecutingBroker, brokerSSI.InvestmentManager, brokerSSI.Product, brokerSSI.SubProduct, brokerSSI.SettleLocation})
    if err != nil {
        return shim.Error(err.Error())
    }
    // Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the Broker SSI.
    // Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
    value := []byte{0x00}
    stub.PutPrivateData(privateCollection, indexCompositeKey, value)

    // Broker SSI saved and indexed. Return success.
    fmt.Println("- end init Broker SSI")
    return shim.Success(nil)

}

// arg 0 - executingBroker name
func (t *TradeInstruction) getAllocationByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var investmentManager, executingBroker, agentBank, tradeId, endAccount, allocationId, jsonResp string
	var err error
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting allocationID") 
	}

	investmentManager = strings.ToUpper(args[0])
	executingBroker = strings.ToUpper(args[1])
	agentBank = strings.ToUpper(args[2])
	tradeId= strings.ToUpper(args[3])
	endAccount= strings.ToUpper(args[4])
	allocationId = investmentManager+ tradeId + endAccount
	valAsbytes, err := stub.GetPrivateData("collectionAllocationFor" + executingBroker + investmentManager + agentBank, allocationId)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + allocationId + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"allocation does not exist: " + allocationId + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

/*
func (t *TradeInstruction) getBrokerSSI(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var ssiid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting the ssiid to query")
	}
	ssiid := strings.ToUpper(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"brokerSSI\",\"ssiid\":\"%s\"}}", ssiid)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}


*/