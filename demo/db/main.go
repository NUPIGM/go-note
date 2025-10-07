package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func LeveldbBasic() {
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// db.Put([]byte("user"), []byte("john"), nil)
	date, _ := db.Get([]byte("user"), nil)

	fmt.Println(string(date))
}
