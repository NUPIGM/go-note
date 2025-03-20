package notes

import (
	"encoding/json"
	"fmt"
)

func Json() {
	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "admin",
		Password: "123456",
	}

	jsonData, _ := json.Marshal(credentials)
	fmt.Println(string(jsonData))

}
