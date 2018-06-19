package main

import "time"

type Post struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	Tags         []string  `json:"tags"`
	CreationDate time.Time `json:"creation_date"`
}
type Posts []Post
