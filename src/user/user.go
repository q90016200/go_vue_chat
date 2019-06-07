package user

import "fmt"
import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


// 先定義一個 User 建構體，然後有個叫做 uid 的字串成員
type User struct {
	uid string
}

type TestStruct struct {
    id string
}


// 用來建構 User 的假建構子
func NewUser(id string) (ts *TestStruct) {  
    ts = &TestStruct{id: id}
    // 這裡會回傳一個型態是 *User 建構體的 user 變數
    return
}

// 連接到 mongo
func connectionMongo () (*mongo.Collection){
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    collection := client.Database("go_chat").Collection("user")

    return collection

}

// 登入
func Login() {  
    // collection := connectionMongo()

    // fmt.Println(mongo)
}

// 註冊
func Register(username string) {
    collection := connectionMongo()

    userExists := checkUserExists(username)

    if userExists == false {
        data := make(map[string]string)

        data["username"] = username

        // fmt.Println(data)
        // fmt.Println(len(data))

        insertResult, err := collection.InsertOne(context.TODO(), data)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    } else {
        fmt.Println("User Exists")
    }

}

// 檢查 user 是否存在
func checkUserExists(username string) (bool) {
    collection := connectionMongo()

    result := make(map[string]interface{})  

    filter := bson.D{{"username", username}}

    err := collection.FindOne(context.TODO(), filter).Decode(&result)

    fmt.Println(result)

    if err != nil {
        // log.Fatal(err)
        return false
    }

    
    return true
    
}


func (ts *TestStruct) Test(){
    fmt.Println("Hi",ts.id)
}