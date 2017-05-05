/*
Smart Contract for PoC - Track & Trace Use Case
*/
package main

import (
	"errors"
	"fmt"
	//"strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/crypto/primitives"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//iQOS Changes starts ----------------------------------------------------------------
// Assembly comprises of one Serial ID and multiple Batch IDs
type AssemblyLine struct{	
	AssemblyLineId string `json:"assemblyLineId"`
	SerialId string `json:"serialId"`
	OriginalFilamentBatchId string `json:"originalFilamentBatchId"`
	OriginalLedBatchId string `json:"originalLedBatchId"`
	OriginalCircuitBoardBatchId string `json:"originalCircuitBoardBatchId"`
	OriginalWireBatchId string `json:"originalWireBatchId"`
	OriginalCasingBatchId string `json:"originalCasingBatchId"`
	OriginalAdaptorBatchId string `json:"originalAdaptorBatchId"`
	OriginalStickPodBatchId string `json:"originalStickPodBatchId"`
	AssemblyLineStatus string `json:"assemblyLineStatus"`
	}

// GetAssemblyLineStatus is for storing retreived Assembly Line Status
type GetAssemblyLineStatus struct{	
	AssemblyLineStatus string `json:"assemblyLineStatus"`
}

//iQOS Changes ends ----------------------------------------------------------------


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	/*
	// Check if table already exists
	_, err = stub.GetTable("AssemblyLine")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
   */
	// Create application Table
	err := stub.CreateTable("AssemblyLine", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "assemblyLineId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "serialId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalFilamentBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalLedBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalCircuitBoardBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalWireBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalCasingBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalAdaptorBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "originalStickPodBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyLineStatus", Type: shim.ColumnDefinition_STRING, Key: false},

	})
	if err != nil {
		return nil, errors.New("Failed creating AssemblyLine.")
	}

	return nil, nil
}



//startAssemblyLine to start an Assemblyline
func (t *SimpleChaincode) startAssemblyLine(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
if len(args) != 9 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 9. Got: %d.", len(args))
		}
		
		assemblyLineId:=args[0]
		serialId:=args[1]
		originalFilamentBatchId:=args[2]
		originalLedBatchId:=args[3]
		originalCircuitBoardBatchId:=args[4]
		originalWireBatchId:=args[5]
		originalCasingBatchId:=args[6]
		originalAdaptorBatchId:=args[7]
		originalStickPodBatchId:=args[8]
		assemblyLineStatus:= "InProgress"

		// Insert a row
		ok, err := stub.InsertRow("AssemblyLine", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: assemblyLineId}},
				&shim.Column{Value: &shim.Column_String_{String_: serialId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalFilamentBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalLedBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalCircuitBoardBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalWireBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalCasingBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalAdaptorBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalStickPodBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: assemblyLineStatus}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//Update AssemblyLine status
func (t *SimpleChaincode) updateAssemblyLineStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	} 
	//else {
	//	return nil, errors.New("Correct number of arguments. Got 2.")
	//}


	assemblyLineId := args[0]
	assemblyLineStatus := args[1]
	
	

	// Get the row pertaining to this AssemblyLineId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assemblyLineId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssemblyLine", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving AssemblyLine with assemblyLineId %s. Error %s", assemblyLineId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	// Delete the row pertaining to this assemblyLineId
	err = stub.DeleteRow(
		"Assemblyline",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

		//assemblyLineId:=row.Columns[0].GetString_()
		serialId:=row.Columns[1].GetString_()
		originalFilamentBatchId:=row.Columns[2].GetString_()
		originalLedBatchId:=row.Columns[3].GetString_()
		originalCircuitBoardBatchId:=row.Columns[4].GetString_()
		originalWireBatchId:=row.Columns[5].GetString_()
		originalCasingBatchId:=row.Columns[6].GetString_()
		originalAdaptorBatchId:=row.Columns[7].GetString_()
		originalStickPodBatchId:=row.Columns[8].GetString_()
		//assemblyLineStatus:= assemblyLineStatus


		// Insert a row
		ok, err := stub.InsertRow("AssemblyLine", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: assemblyLineId}},
				&shim.Column{Value: &shim.Column_String_{String_: serialId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalFilamentBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalLedBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalCircuitBoardBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalWireBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalCasingBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalAdaptorBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: originalStickPodBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: assemblyLineStatus}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists in Assemblyline.")
		}
		
	return nil, nil

}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "startAssemblyLine" {
		fmt.Printf("Function is startAssemblyLine")
		return t.startAssemblyLine(stub, args)
	} else if function == "updateAssemblyLineStatus" {
		fmt.Printf("Function is updateAssemblyLineStatus")
		return t.updateAssemblyLineStatus(stub, args)
	} 

	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "startAssemblyLine" {
		fmt.Printf("Function is startAssemblyLine")
		return t.startAssemblyLine(stub, args)
	}  else if function == "updateAssemblyLineStatus" {
		fmt.Printf("Function is updateAssemblyLineStatus")
		return t.updateAssemblyLineStatus(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

//get the AssemblyLine against ID
func (t *SimpleChaincode) getAssemblyLineByID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting AssemblyLineID to query")
	}

	assemblyLineID := args[0]
	

	// Get the row pertaining to this assemblyLineID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assemblyLineID}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssemblyLine", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the assemblyLineID " + assemblyLineID + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the assemblyLineID " + assemblyLineID + "\"}"
		return nil, errors.New(jsonResp)
	}

	//return []byte (row), nil
	 mapB, _ := json.Marshal(row)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


//get the status against the AssemblyLineID
func (t *SimpleChaincode) getAssemblyLineStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting AssemblyLineID to query")
	}

	assemblyLineID := args[0]
	

	// Get the row pertaining to this assemblyLineID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assemblyLineID}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssemblyLine", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the assemblyLineID " + assemblyLineID + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the assemblyLineID " + assemblyLineID + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := GetAssemblyLineStatus{}
	
	res2E.AssemblyLineStatus = row.Columns[9].GetString_()
	
  /*  
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	return mapB, nil
  */	
 
    return []byte (res2E.AssemblyLineStatus), nil

}

//get all AssemblyLines
func (t *SimpleChaincode) getAllAssemblyLines(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
var columns []shim.Column

	rows, err := stub.GetRows("AssemblyLine", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
 
   
		
	res2E:= []*AssemblyLine{}	
	
	for row := range rows {		
		newApp:= new(AssemblyLine)
		newApp.AssemblyLineId = row.Columns[0].GetString_()
		newApp.SerialId = row.Columns[1].GetString_()
		newApp.OriginalFilamentBatchId = row.Columns[2].GetString_()
		newApp.OriginalLedBatchId = row.Columns[3].GetString_()
		newApp.OriginalCircuitBoardBatchId = row.Columns[4].GetString_()
		newApp.OriginalWireBatchId = row.Columns[5].GetString_()
		newApp.OriginalCasingBatchId = row.Columns[6].GetString_()
		newApp.OriginalAdaptorBatchId = row.Columns[7].GetString_()
		newApp.OriginalStickPodBatchId = row.Columns[8].GetString_()
		newApp.AssemblyLineStatus = row.Columns[9].GetString_()
		
		if len(newApp.AssemblyLineId) > 0{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// query queries the chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")

	if function == "getAssemblyLineStatus" { 
		t := SimpleChaincode{}
		return t.getAssemblyLineStatus(stub, args)
	} else if function == "getAssemblyLineByID" { 
		t := SimpleChaincode{}
		return t.getAssemblyLineByID(stub, args)
	} else if function == "getAllAssemblyLines" { 
		t := SimpleChaincode{}
		return t.getAllAssemblyLines(stub, args)
	}
	
	return nil, errors.New("Received unknown function query")
}
//iQOS Changes ends------------------------------------------------------------------------------------------

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
