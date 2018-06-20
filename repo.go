package main

import (
	"fmt"
	"strconv"
)

func NewPost(post Post) {
	pool := NewPool()
	c := pool.Get()
	defer c.Close()

	i, err := c.Do("INCR", "post:id")
	if err != nil {
		panic(err)
	}

	prin
	k := "post:" + strconv.Itoa(i)
	fmt.Println(k)

	c.Send("SET", k, post)

	r, err := c.Do("EXEC")
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
