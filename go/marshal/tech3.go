package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Dog struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	BornAt Time   `json:"born_at"`
}

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	t.Time = time.Unix(i, 0)
	return nil
}

func main() {
	dog := Dog{1, "bowser", "husky", Time{time.Now()}}
	b, err := json.Marshal(dog)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	b = []byte(`{
    "id":1,
    "name":"bowser",
    "breed":"husky",
    "born_at":1480979203}`)
	dog = Dog{}
	json.Unmarshal(b, &dog)
	fmt.Println(dog)
}
