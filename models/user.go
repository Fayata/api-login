
package models

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var Users = map[int]*User{}
var Seq = 1