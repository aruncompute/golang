package main
import (
	"fmt"
	"encoding/json"
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

	b, err := json.Marshal(users)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Println(string(b))	
	/////Json decode
	fmt.Println("================Json decode =============")
	decodeUsers	:= []user{};
	fmt.Println(b)	
	json.Unmarshal([]byte(b), &decodeUsers)
	fmt.Println(decodeUsers);

	
}
