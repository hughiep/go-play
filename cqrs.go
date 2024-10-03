package main

// CQRS in Go
// Command Query Responsibility Segregation (CQRS) is a software architectural pattern that separates the read and write operations of a system.
// The write operations are handled by commands, while the read operations are handled by queries.
// CQRS is based on the idea that the read and write operations have different requirements and should be handled separately.
// In this example, we will implement CQRS in Go using channels.

// Command
// A command is a request to change the state of the system.
// In this example, we will define a command type that represents a command to increment a counter.
type Command struct {
	Increment int
}

// Query
// A query is a request to read the state of the system.
// In this example, we will define a query type that represents a query to get the current value of the counter.
type Query struct {
	Result chan int
}

// Counter
// The counter is a simple struct that represents a counter.
// The counter has a value field that stores the current value of the counter.
// The counter also has a command channel that receives commands to increment the counter.
// The counter has a query channel that receives queries to get the current value of the counter.
type Counter struct {
	value   int
	command chan Command
	query   chan Query
}

// NewCounter
// The NewCounter function creates a new counter and starts a goroutine to handle commands and queries.
func NewCounter() *Counter {
	counter := &Counter{
		value:   0,
		command: make(chan Command),
		query:   make(chan Query),
	}

	go func() {
		for {
			select {
			case cmd := <-counter.command:
				counter.value += cmd.Increment
			case qry := <-counter.query:
				qry.Result <- counter.value
			}
		}
	}()

	return counter
}

// Increment
// The Increment method sends a command to increment the counter by the specified value.
func (c *Counter) Increment(inc int) {
	c.command <- Command{Increment: inc}
}

// Value
// The Value method sends a query to get the current value of the counter.
// The Value method returns the current value of the counter.
func (c *Counter) Value() int {
	qry := Query{Result: make(chan int)}
	c.query <- qry
	return <-qry.Result
}

// Main function
func CQRS() {
	// Create a new counter
	counter := NewCounter()

	// Increment the counter by 1
	counter.Increment(1)

	// Get the current value of the counter
	value := counter.Value()
	println(value)
}

// Output
// 1
// We defined a command type to represent a command to increment a counter and a query type to represent a query to get the current value of the counter.
// We created a Counter struct that represents a counter with a value field, a command channel, and a query channel.
// We implemented the NewCounter function to create a new counter and start a goroutine to handle commands and queries.
// We implemented the Increment method to send a command to increment the counter and the Value method to send a query to get the current value of the counter.
// Finally, we created a new counter, incremented it by 1, and got the current value of the counter.
// The output of the program is 1, which is the current value of the counter after incrementing it by 1.
// This example demonstrates how to implement CQRS in Go using channels to separate the read and write operations of a system.
// The CQRS pattern can be useful for building scalable and maintainable systems by separating the concerns of read and write operations.
// The CQRS pattern can also help improve performance by optimizing read and write operations separately.

// Why need go routine in CQRS?
// Go routines are lightweight threads managed by the Go runtime.
// Go routines allow us to run concurrent code in a simple and efficient way.
// In this example, we used a Go routine to handle commands and queries for the counter.
// The Go routine listens for commands and queries on the command and query channels and updates the counter state accordingly.
// Using a Go routine allows us to handle commands and queries concurrently without blocking the main thread.
// This concurrency model is well-suited for implementing CQRS in Go, where we need to handle read and write operations separately.
