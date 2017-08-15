package main

type Event struct {
	Actions []Action
	Team Team
	User User
}

type Action struct {
	Name, Type, Value string
}

type Team struct {
	Id string
	Domain string
}

type User struct {
	Id string
	Name string
}