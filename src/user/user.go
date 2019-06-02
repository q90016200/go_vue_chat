package user

import "fmt"

// 先定義一個 User 建構體，然後有個叫做 uid 的字串成員
type User struct {
	uid string
}

// 用來建構 User 的假建構子
func NewUser(id string) (user *User) {  
    user = &User{uid: id}

    // 這裡會回傳一個型態是 *User 建構體的 user 變數
    return
}

func (u *User) Login() {  
    fmt.Println("uid:",u.uid)
}



func Test(){
    fmt.Println("Hi")
}