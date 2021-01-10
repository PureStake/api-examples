// Example of accesing a remote API with a custom token key.
// In this case, the API is expecting the key "X-API-Key" instead of the
// default "X-Algo-API-Token".
// This is done by using the SDK built in Header type, that allows us to
// define custom additional header, and create a new algod client using
// a special function for this case: MakeClientWithHeaders.

// DEPRECATED 12/30/20 with Algod v1 API Shutdown

package main

import (
	"encoding/json"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/algod"
)

const algodAddress = "https://testnet-algorand.api.purestake.io/ps1"
const psToken = "......"

func main() {
	var headers []*algod.Header
	headers = append(headers, &algod.Header{"X-API-Key", psToken})
	// Create an algod client
	algodClient, err := algod.MakeClientWithHeaders(algodAddress, "", headers)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}

	// Print algod status
	fmt.Println("Status")
	nodeStatus, err := algodClient.Status()
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}

	fmt.Printf("algod last round: %d\n", nodeStatus.LastRound)
	fmt.Printf("algod time since last round: %d\n", nodeStatus.TimeSinceLastRound)
	fmt.Printf("algod catchup: %d\n", nodeStatus.CatchupTime)
	fmt.Printf("algod latest version: %s\n", nodeStatus.LastVersion)

	// Fetch block information
	lastBlock, err := algodClient.Block(nodeStatus.LastRound)
	if err != nil {
		fmt.Printf("error getting last block: %s\n", err)
		return
	}

	// Print the block information
	fmt.Printf("\n-----------------Block Information-------------------\n")
	blockJSON, err := json.MarshalIndent(lastBlock, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall block data: %s\n", err)
	}
	fmt.Printf("%s\n", blockJSON)
}
