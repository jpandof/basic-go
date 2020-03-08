package main

// Use pointers to share data
// & address of
// * value that the pointer to

func main() {
	example1()
	println("====================")
	//example2()
}

func example1() {

	count := 10

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// pass the "value of" the count
	incrementValueOf(count)

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	incrementPointer(&count)

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

func incrementValueOf(inc int) {
	inc++
	println("inc:\tValue of[", inc, "]\tAddr of[", &inc, "]")
}

func incrementPointer(inc *int) {
	*inc++
	println("inc:\tValue of[", inc, "]\tAddr of[", &inc, "]")
}

// example 2
type user struct {
	name  string
	email string
}

func example2() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}

func createUserV1() user {
	u := user{
		name:  "Iker",
		email: "iker@gmail.com",
	}
	println("V1", &u)

	return u
}

func createUserV2() *user {
	u := user{
		name:  "Iker",
		email: "iker@gmail.com",
	}
	println("V1", &u)

	return &u
}
