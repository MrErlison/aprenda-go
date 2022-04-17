//C17E5

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
	}

	sort.Sort(byAge(users))

	fmt.Println("\nSort by age")

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
	}

}
