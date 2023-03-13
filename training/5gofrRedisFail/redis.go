package main

import (
	"context"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	k := gofr.New()
	ctx := context.Background()

	key := "key"
	fmt.Println(k, "ljnkj")
	// Set accepts a key-value pair, and sets those in Redis, if expiration is non-zero value, it sets a expiration(TTL)
	// on those keys, if expiration is 0, then the keys have no expiration time

	err := k.Redis.Set(ctx, key, "value", 0)

	strCmd := k.Redis.Get(ctx, key) // return the *StringCmd
	// or
	val, er := k.Redis.Get(ctx, key).Result() // returns string,error
	e := k.Redis.Del(ctx, key)
	if err != nil {
		if er != nil {
			if e != nil {
			}
		}
	}
	fmt.Println(strCmd, "is strcmd")
	fmt.Println(val, "is val")
}
