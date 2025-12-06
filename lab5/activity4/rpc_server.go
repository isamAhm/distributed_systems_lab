package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// Args holds the arguments for arithmetic operations
type Args struct {
	A, B int
}

// Calculator provides methods for arithmetic operations with state management
type Calculator struct {
	lastResult int
	mu         sync.Mutex
}

// Multiply multiplies two integers and returns the result
func (c *Calculator) Multiply(args *Args, reply *int) error {
	if args.A == 0 || args.B == 0 {
		return errors.New("multiplication by zero is not allowed")
	}

	c.mu.Lock()
	*reply = args.A * args.B
	c.lastResult = *reply
	c.mu.Unlock()

	return nil
}

// Add adds two integers and returns the result
func (c *Calculator) Add(args *Args, reply *int) error {
	c.mu.Lock()
	*reply = args.A + args.B
	c.lastResult = *reply
	c.mu.Unlock()

	return nil
}

// Subtract subtracts two integers and returns the result
func (c *Calculator) Subtract(args *Args, reply *int) error {
	c.mu.Lock()
	*reply = args.A - args.B
	c.lastResult = *reply
	c.mu.Unlock()

	return nil
}

// Divide divides two integers and returns the result
func (c *Calculator) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero is not allowed")
	}

	c.mu.Lock()
	*reply = args.A / args.B
	c.lastResult = *reply
	c.mu.Unlock()

	return nil
}

// GetLastResult retrieves the last result from the server
func (c *Calculator) GetLastResult(args *Args, reply *int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	*reply = c.lastResult
	return nil
}

func main() {
	// Register the Calculator service
	calc := new(Calculator)
	rpc.Register(calc)

	// Start listening for incoming RPC connections
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting RPC server:", err)
		return
	}

	fmt.Println("RPC server is listening on port 1234...")
	fmt.Println("Server supports: Add, Subtract, Multiply, Divide, GetLastResult")

	// Accept connections and handle them concurrently
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle each client in a separate goroutine for concurrency
		go rpc.ServeConn(conn)
	}
}
