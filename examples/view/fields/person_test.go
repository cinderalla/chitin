package person_test

import (
	"fmt"
	"log"

	"chitin.io/chitin/examples/view/fields"
)

func Example() {
	input := []byte{
		// age
		0, 21,
		// siblings
		0, 3,
		// len of name +1
		5,
		// name
		'J', 'a', 'n', 'e',
	}
	view, err := person.NewPersonV2View(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q is %d years old.\n", view.Name(), view.Age())

	// Output:
	// "Jane" is 21 years old.
}
