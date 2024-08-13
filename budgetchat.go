package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"protohackers/server"
)

type User struct {
	Name string
	Conn net.Conn
}

var users map[string]User

func budgetchat(conn net.Conn) {

	// Setting the user's name
	user := User{
		Name: "",
		Conn: conn,
	}
	fmt.Fprintln(conn, "Welcome to budgetchat! What shall I call you?")
	_, err := fmt.Fscanln(conn, &user.Name)
	if err != nil {
		fmt.Fprintln(conn, err.Error())
		return
	}
	if _, ok := users[user.Name]; ok {
		fmt.Fprintln(conn, "Name already taken")
		return
	}
	for _, r := range user.Name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			fmt.Fprintln(conn, "Name is not alphanumeric")
			return
		}
	}
	users[user.Name] = user

	// A user joins - broadcast message
	var names []string
	for _, u := range users {
		if u.Name == user.Name {
			continue
		}
		fmt.Fprintf(u.Conn, "* %s has entered the room\n", user.Name)
		names = append(names, u.Name)
	}

	// Presence notification
	fmt.Fprintf(conn, "* The room contains: %s\n", strings.Join(names, ", "))

	// Chat messages
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		for _, u := range users {
			if u.Name == user.Name {
				continue
			}
			fmt.Fprintf(u.Conn, "[%s] %s\n", user.Name, scanner.Text())
		}
	}

	// A user leaves - broadcast message
	delete(users, user.Name)
	for _, u := range users {
		fmt.Fprintf(u.Conn, "* %s has left the room\n", user.Name)
	}

}

func main() {
	users = make(map[string]User)
	server.Run(budgetchat)
}
