package main

import "fmt"

/*

Get
corr-id



post
codd-id
content-type


*/

func HeaderCreater(mp map[string]string) func(key, val []string) map[string]string {

	return func(key, val []string) map[string]string {
		if len(key) != len(val) {
			return nil
		}
		for i := range key {
			mp[key[i]] = val[i]
		}

		return mp
	}
}

func main() {
	getmap := map[string]string{
		"codrr-id": "sdfghjklkjhgfd",
	}

	postmap := map[string]string{
		"codrr-id": "sdfkhkghjklkjhgfd",
		"content":  "ssssfce",
	}

	getHeader, postHeader := HeaderCreater(getmap), HeaderCreater(postmap)
	fmt.Println(getHeader([]string{"name", "age"}, []string{"alpha", "12"}))
	fmt.Println(postHeader([]string{"name", "age", "payload"}, []string{"alpha", "12", "jshgdjfgajhdgsjfa"}))
	fmt.Println("hello world again")

}
