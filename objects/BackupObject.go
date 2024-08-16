
package objects

import (
    "fmt"
    "time"
    "encoding/json"
    "os"
)

// Backup represents a backup of Kubernetes resources.
type Backup struct {
    Name      string
    Namespace string
    Resource  string
    CreatedAt time.Time
    FilePath  string
}



// ListBackups lists all backup objects stored in the system.
func ListBackups(backups []Backup) {
    for _, backup := range backups {
        fmt.Printf("Name: %s, Namespace: %s, Resource: %s, CreatedAt: %s, FilePath: %s\n",
            backup.Name, backup.Namespace, backup.Resource, backup.CreatedAt.Format(time.RFC3339), backup.FilePath)
    }
    
}

func LoadBackups() ([]Backup, error) {
    data, err := os.ReadFile("backups\\database.json")
    if err != nil {
        return nil, err
    }

    var backups []Backup
    err = json.Unmarshal(data, &backups)
    if err != nil {
        return nil, err
    }

    return backups, nil
}

func AddBackup(newBackup Backup) error {
    // Load existing backups
    backups, err := LoadBackups()
    if err != nil && !os.IsNotExist(err) {
        return err
    }

    // Add the new backup to the list
    backups = append(backups, newBackup)

    // Save the updated list back to the file
    return SaveBackups( backups)
}


func SaveBackups( backups []Backup) error {
    fmt.Println("calling file")
    data, err := json.MarshalIndent(backups, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile("backups\\database.json", data, 0644)
}

func CreateBackup(name, namespace, resource,filePath string) (*Backup,error) {
   
    backup:= &Backup{
        Name:      name,
        Namespace: namespace,
        Resource:  resource,
        CreatedAt: time.Now(),
        FilePath:  filePath,
    }

    return  backup,nil
}
