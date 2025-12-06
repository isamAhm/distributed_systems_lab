package main

import (
	"fmt"
	"log"
	"net/rpc"
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
	err = client.Call("Calculator.Add", &args, &reply)
	if err != nil {
		log.Fatal("Error calling RPC:", err)
	}
	fmt.Printf("Result of %d + %d = %d\n", args.A, args.B, reply)

	// Test Subtract operation
	fmt.Println("\nTesting Subtract operation...")
	args = Args{A: 20, B: 8}
	err = client.Call("Calculator.Subtract", &args, &reply)
	if err != nil {
		log.Fatal("Error calling RPC:", err)
	}
	fmt.Printf("Result of %d - %d = %d\n", args.A, args.B, reply)

	// Test Divide operation
	fmt.Println("\nTesting Divide operation...")
	args = Args{A: 15, B: 3}
	err = client.Call("Calculator.Divide", &args, &reply)
	if err != nil {
		log.Fatal("Error calling RPC:", err)
	}
	fmt.Printf("Result of %d / %d = %d\n", args.A, args.B, reply)

	fmt.Println("\n=====================================")
	fmt.Println("All operations completed!")
}
