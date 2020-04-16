package db

import (
    "os"
    "sync"
    "time"

    "github.com/jinzhu/gorm"
    // GORM requires the database driver to be imported.
    // https://gorm.io/docs/connecting_to_the_database.html#Connecting-to-database
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    connections             = map[string]*gorm.DB{}
    defaultConnectionString = "host=localhost user=postgres dbname=postgres password=postgres sslmode=disable"
    defaultConnectionPool   = "default"
    lock                    = &sync.Mutex{}
    maximumConnections      = 10
)

// Get fetches an open database connection from the default connection pool.
// If no connection exists then one is created and added to the pool.
// If the POSTGRES_URL environment variable is set then override the default
// connection string.
func get() (*gorm.DB, error) {
    return fetchConnection(defaultConnectionPool)
}

func fetchConnection(pool string) (*gorm.DB, error) {
    lock.Lock()
    defer lock.Unlock()

    if connections[pool] != nil {
        return connections[pool], nil
    }

    connectionString := defaultConnectionString
    env := os.Getenv("POSTGRES_URL")
    if env != "" {
        connectionString = env
    }

    c, err := connect(connectionString)
    if err != nil {
        return nil, err
    }

    connections[pool] = c
    return connections[pool], nil
}

func connect(connectionString string) (*gorm.DB, error) {
    c, err := gorm.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    c.DB().SetMaxOpenConns(maximumConnections)
    c.DB().SetMaxIdleConns(maximumConnections)
    c.DB().SetConnMaxLifetime(time.Hour)

    return c, nil
}
