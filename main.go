package main

import "fmt"

func main() {
	db := readFromJson(getPath())
	t := db.Add("Get a new job")

	res, err := db.Find(t.Id)
	if err != nil {
		print("Unaable to find %d", t.Id)
		return
	}

	fmt.Println("Found: ", res)
}
