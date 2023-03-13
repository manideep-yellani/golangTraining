package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6371", // can be any port for local machine but if that port is empty, error-->
		Password: "",               // connection refused, if that con is configed with docker
		DB:       0,                //but redis port is something else the EOF error, also default in redis is 6379
	})

	pong, err := client.Ping().Result()
	fmt.Println("pong:", pong, "error:", err)

	val, _ := client.Set("mani", "money", 0).Result()
	val, _ = client.Set("mani1", "money1", 0).Result()
	fmt.Println(val)
	val, _ = client.Get("mani").Result()
	fmt.Println(val)
	val, _ = client.Get("mani1").Result()
	fmt.Println(val)

	//jason, _ := json.Marshal(Movie{"jersey"})
	//client.Set("movie", jason, 0)
	strCmd := client.Get("movie")
	strCmdResult, _ := strCmd.Result()
	fmt.Printf("%T \n", strCmd)
	fmt.Println("strcmd:", strCmd)
	fmt.Println("strcmdResult:", strCmdResult)
	result, _ := client.Get("movie").Result()
	fmt.Println(result)

	fmt.Printf("\n \n")

	k := client.Set("movie1", 5, 0)
	k = client.Set("movie1", 6, 0)
	fmt.Println(k)
	fmt.Println(k.Args())
	s := k.Args()
	fmt.Printf("%T \n", s)
	fmt.Printf("%v %T \n", s[2], s[2])
	strCmd = client.Get("movie1")
	s = strCmd.Args()
	fmt.Printf("%v %T \n", s, s)
	strCmdResult, _ = strCmd.Result()
	fmt.Println("strcmd:", strCmd)
	fmt.Println("strcmdResult:", strCmdResult)
	result, _ = client.Get("movie1").Result()
	fmt.Printf("%T", result)

}

type Movie struct {
	Name string `json:"name"`
}
