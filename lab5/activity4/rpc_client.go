package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// Args holds the arguments for arithmetic operations
type Args struct {
	A, B int
}

func main() {
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	defer client.Close()

	fmt.Println("Connected to RPC server successfully!")
	fmt.Println("=====================================\n")

	// Test Add operation
	fmt.Println("Testing Add operation...")
	args := Args{A: 10, B: 5}
	var reply int

	call := client.Go("Calculator.Add", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			log.Printf("RPC error: %v\n", call.Error)
		} else {
			fmt.Printf("Result of %d + %d = %d\n", args.A, args.B, reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	// Test Subtract operation
	fmt.Println("\nTesting Subtract operation...")
	args = Args{A: 20, B: 8}
	call = client.Go("Calculator.Subtract", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			log.Printf("RPC error: %v\n", call.Error)
		} else {
			fmt.Printf("Result of %d - %d = %d\n", args.A, args.B, reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	// Test Multiply operation
	fmt.Println("\nTesting Multiply operation...")
	args = Args{A: 6, B: 7}
	call = client.Go("Calculator.Multiply", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			log.Printf("RPC error: %v\n", call.Error)
		} else {
			fmt.Printf("Result of %d * %d = %d\n", args.A, args.B, reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	// Test Divide operation
	fmt.Println("\nTesting Divide operation...")
	args = Args{A: 15, B: 3}
	call = client.Go("Calculator.Divide", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			log.Printf("RPC error: %v\n", call.Error)
		} else {
			fmt.Printf("Result of %d / %d = %d\n", args.A, args.B, reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	// Test error handling - Division by zero
	fmt.Println("\nTesting error handling (division by zero)...")
	args = Args{A: 10, B: 0}
	call = client.Go("Calculator.Divide", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			fmt.Printf("Expected error caught: %v\n", call.Error)
		} else {
			fmt.Printf("Result: %d\n", reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	// Retrieve the last result
	fmt.Println("\nRetrieving last result from server...")
	args = Args{A: 0, B: 0} // Not used for GetLastResult
	call = client.Go("Calculator.GetLastResult", &args, &reply, nil)
	select {
	case <-call.Done:
		if call.Error != nil {
			log.Printf("RPC error: %v\n", call.Error)
		} else {
			fmt.Printf("Last result stored on server: %d\n", reply)
		}
	case <-time.After(2 * time.Second):
		log.Println("RPC call timed out")
	}

	fmt.Println("\n=====================================")
	fmt.Println("All operations completed!")
}
