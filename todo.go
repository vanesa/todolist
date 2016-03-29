package main

import "time"

//  basic Todo model to send and retrieve data
type Todo struct {
	Name string
	Completed bool
	Due time.Time
}

type Todos []Todo