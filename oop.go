package main

// OOP in Go
// Go is not an object-oriented language, but it has some features that allow us to implement OOP.

// Structs
// Go has a special type called struct, which is a collection of fields.
// Struct fields are accessed using a dot.
// A struct can also have methods.
// Methods are functions that belong to a struct.
// Here is an example of a struct with a method:
type Person struct {
	Name string
	Age  int
}

// Method
// The method is a function that belongs to the Person struct.
// The method is called with the dot operator.
// The method can modify the struct.
func (p *Person) SayHello() {
	println("Hello", p.Name)
}

// Composition
// Go does not have inheritance, but we can use composition to achieve similar results.
// Composition is a way to combine structs.
// Here is an example of composition:
type Employee struct {
	Person
	Company string
}

// Method Overriding
// We can override methods by defining a new method with the same name.
// Here is an example of method overriding:
func (e *Employee) SayHello() {
	println("Hello", e.Name, "from", e.Company)
}

// Polymorphism
// Polymorphism is the ability to use a common interface for different types.
// We can achieve polymorphism in Go using interfaces.
// Here is an example of polymorphism:
type Greeter interface {
	SayHello()
}

// The SayHello method is defined in the Person and Employee structs.
// Both Person and Employee implement the Greeter interface.
// We can use the Greeter interface to call the SayHello method on both Person and Employee.
func Greet(g Greeter) {
	g.SayHello()
}

func Run() {
	// Create a Person
	p := Person{Name: "Alice", Age: 30}
	p.SayHello()

	// Create an Employee
	e := Employee{Person: Person{Name: "Bob", Age: 25}, Company: "Google"}
	e.SayHello()

	// Polymorphism
	Greet(&p)
	Greet(&e)

	// Output:
	// Hello Alice
	// Hello Bob from Google
	// Hello Alice
	// Hello Bob from Google

	// Note: Go does not have inheritance, but we can use composition and interfaces to achieve similar results.

	// Summary:
	// - Go has structs, which are collections of fields.
	// - Structs can have methods.
	// - Go does not have inheritance, but we can use composition to combine structs.
	// - We can use interfaces to achieve polymorphism in Go.
	// - Go does not have classes, but we can use structs and interfaces to implement OOP concepts.

}

// Inheritance, Encapsulation, and Abstraction
// Go does not have inheritance, encapsulation, and abstraction like traditional OOP languages.
// Go encourages composition over inheritance.
// Go does not have access modifiers like public, private, and protected.
// Go uses naming conventions to indicate visibility.
// - Capitalized names are exported (public).
// - Lowercase names are not exported (private).
// Go does not have classes, but we can use structs and interfaces to achieve similar results.
// Go is a simple and pragmatic language that focuses on simplicity and readability.

// Composition over Inheritance
// Go encourages composition over inheritance.
// Composition is a way to combine structs.
// Composition is more flexible than inheritance.
// Composition allows us to combine multiple structs to create new structs.
// Composition allows us to reuse code without creating complex inheritance hierarchies.
// Composition is a powerful feature of Go that enables code reuse and maintainability.
