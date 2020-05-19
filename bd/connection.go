package

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoConnection is the object connection to database */
var MongoConnection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://lcmencia:mindcontrol@tuitta-bdhyy.mongodb.net/test?retryWrites=true&w=majority")


/* ConnectDB it is the function that allows to connect database */
func ConnectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err !=nil {
        log.Fatal(err)
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    log.Println("Successful connection to the database")
    return client
}

/* CheckConnection is the PING to the database */
func CheckConnection() int {
    err := MongoConnection.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}