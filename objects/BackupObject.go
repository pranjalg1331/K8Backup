package objects

import (
	// "K8Backup/internal"
	"context"
	// "encoding/json"
	"fmt"
	// "os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Backup represents a backup of Kubernetes resources.
type Backup struct {
    // id  primitive.ObjectID
    Name      string
    Namespace string
    Resource  string
    CreatedAt time.Time
    FilePath  string

}

var collection *mongo.Collection
var ctx = context.TODO() 

func connectMongo()(){
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// log.Fatal(err)
        panic("mongo connection failed")
	}
    collection = client.Database("K8Backups").Collection("backups")
    // return collection
}


// ListBackups lists all backup objects stored in the system.
func ListBackups() (error){
    connectMongo()
    var backups []Backup
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var backup Backup
        err := cursor.Decode(&backup)
        if err != nil {
            return err
        }
        backups = append(backups, backup)
    }

    if err := cursor.Err(); err != nil {
        return err
    }

    // return backups, nil
    for _, backup := range backups {
        fmt.Printf("Name: %s, Namespace: %s, Resource: %s, CreatedAt: %s, FilePath: %s\n",
            backup.Name, backup.Namespace, backup.Resource, backup.CreatedAt.Format(time.RFC3339), backup.FilePath)
    }
    return nil
    
}

func GetBackup(filepath string) (*Backup, error) {
    connectMongo() // Ensure connection to MongoDB
    var backup Backup

    // Create the filter
    filter := bson.M{
        "filepath":      filepath,

    }

    // Find one document in the collection that matches the filter
    err := collection.FindOne(ctx, filter).Decode(&backup)
    if err != nil {
        return nil, err
    }
    fmt.Println(backup)
    // fmt.Println(backup.FilePath)

    return &backup, nil
}


func CreateBackup(name, namespace, resource,filePath string) (*Backup,error) {
    connectMongo()
    var flag bool
    flag=true
    fmt.Println(filePath)
    backup,_:=GetBackup(filePath)
    // fmt.Println(backup)
    if(backup!=nil){
        fmt.Println("backup with this name already exists")
        flag=AskForConfirmation()
    }

    if(flag){
        if(backup!=nil){
            DeleteBackup(filePath)
        }
        backup= &Backup{
            Name:      name,
            Namespace: namespace,
            Resource:  resource,
            CreatedAt: time.Now(),
            FilePath:  filePath,
        }
    
        _, err := collection.InsertOne(ctx, backup)
        return backup,err
    }
    return nil,nil
    
}

func DeleteBackup(filepath string) error {
    connectMongo()
    filter := bson.M{"filepath": filepath}
    
    result, err := collection.DeleteMany(context.TODO(), filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return fmt.Errorf("no backup found filepath : %s", filepath)
    }
    return nil
}



