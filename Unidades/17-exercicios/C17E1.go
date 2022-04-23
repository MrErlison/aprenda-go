//C17E1
// - Partindo do [código](https://play.golang.org/p/U0jea43X55), utilize marshal para transformar []user em JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, error := json.Marshal(users)
	if error != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

}
