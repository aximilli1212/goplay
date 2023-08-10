package main

import "log"

type Coordinate struct {
	X, Y int
}

func (coord Coordinate) shiftBy(x, y int) Coordinate {

	c1 := coord.X + x
	c2 := coord.Y + y

	return Coordinate{c1, c2}
	//coord.X += x
	//coord.Y += y
}

func main() {

	mycoord := Coordinate{5, 5}
	log.Println(mycoord.shiftBy(1, 2))

	log.Println(mycoord)

}

//Yes, there are intrinsic benefits to defining methods on a type (also known as receiver methods) compared to defining them freely (as standalone functions). Here are some advantages of using receiver methods:
//
//Encapsulation: When you define methods on a type, you can control which fields of the type are accessible from those methods. This allows you to encapsulate the internal workings of your type, keeping the implementation details hidden and providing a clear interface for interacting with it.
//
//Code Organization: Grouping related methods together with the type they operate on makes the code more organized and easier to understand. It's like keeping all the tools needed for a specific job in one toolbox.
//
//Contextual Information: Methods can access and work with the fields of the type directly. This provides contextual information to the method, making it easier to work with and manipulate the data within the type.
//
//Method Chaining: When methods are defined on a type, you can chain multiple method calls together. This is not possible with standalone functions, as they don't have access to the specific instance they're operating on.
//
//Interface Implementation: Receiver methods allow a type to satisfy interfaces. This means that if your type has the necessary methods, it can automatically be treated as implementing the interface without any extra steps.
//
//Inheritance and Polymorphism: In object-oriented programming, defining methods on types facilitates concepts like inheritance and polymorphism. Subtypes can override methods defined on their parent types to provide specialized behavior.
//
//Readability: When you see a method call on an instance, it's clear that the operation is closely related to the type of the instance. This improves the readability of the code, as it indicates the intention of the operation.
//
//Consistency: Defining methods on types provides a consistent way to interact with instances of that type. Developers using your type will expect certain methods to be available, leading to a more intuitive and consistent experience.
//
//While there are benefits to using receiver methods, there are cases where standalone functions make more sense. For example, if a function doesn't need access to the internal state of a type and is more general-purpose, it might be better suited as a standalone function.
//
//In essence, receiver methods offer a way to create more structured, organized, and contextual code that's closely tied to the data it operates on. They promote better design practices and contribute to code maintainability and readability.
