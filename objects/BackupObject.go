package objects

import (
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

func CreateBackup(name, namespace, resource,filePath string) (*Backup,error) {
    connectMongo()
    backup:= &Backup{
        Name:      name,
        Namespace: namespace,
        Resource:  resource,
        CreatedAt: time.Now(),
        FilePath:  filePath,
    }

    _, err := collection.InsertOne(ctx, backup)
    return backup,err
}

func DeleteBackup(name string) error {
    connectMongo()
    filter := bson.M{"name": name}
    
    result, err := collection.DeleteMany(context.TODO(), filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return fmt.Errorf("no backup found with name: %s", name)
    }
    return nil
}


// func LoadBackups() ([]Backup, error) {
//     data, err := os.ReadFile("backups\\database.json")
//     if err != nil {
//         return nil, err
//     }

//     var backups []Backup
//     err = json.Unmarshal(data, &backups)
//     if err != nil {
//         return nil, err
//     }

//     return backups, nil
// }

// func AddBackup(newBackup Backup) error {
//     // Load existing backups
//     backups, err := LoadBackups()
//     if err != nil && !os.IsNotExist(err) {
//         return err
//     }

//     // Add the new backup to the list
//     backups = append(backups, newBackup)

//     // Save the updated list back to the file
//     return SaveBackups( backups)
// }


// func SaveBackups( backups []Backup) error {
//     fmt.Println("calling file")
//     data, err := json.MarshalIndent(backups, "", "  ")
//     if err != nil {
//         return err
//     }
//     return os.WriteFile("backups\\database.json", data, 0644)
// }

