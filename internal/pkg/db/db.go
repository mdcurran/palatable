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
func get() (*gorm.DB, error) {
    return fetchConnection(defaultConnectionPool)
}

func fetchConnection(pool string) (*gorm.DB, error) {
    lock.Lock()
    defer lock.Unlock()

    if connections[pool] != nil {
        return connections[pool], nil
    }

    // If the POSTGRES_URL environment variable is set then override the default
    // connection string.
    connectionString := defaultConnectionString
    env := os.Getenv("POSTGRES_URL")
    if env != "" {
        connectionString = env
    }

    // If no connection exists then one is created.
    c, err := connect(connectionString)
    if err != nil {
        return nil, err
    }

    // Add the newly opened connection to the pool.
    connections[pool] = c
    return connections[pool], nil
}

func connect(connectionString string) (*gorm.DB, error) {
    c, err := gorm.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }
    // Set the maximum number of concurrently open connections. Setting this to
    // less than or equal to 0 will mean there is no maximum limit (which is
    // also the default setting).
    c.DB().SetMaxOpenConns(maximumConnections)
    // Set the maximum number of concurrently idle connections. Setting this to
    // less than or equal to 0 will mean that no idle connections are retained.
    c.DB().SetMaxIdleConns(maximumConnections)
    // Set the maximum lifetime of a connection. Setting it to 0 means that
    // there is no maximum lifetime and the connection is reused forever (which
    // is the default behavior).
    c.DB().SetConnMaxLifetime(time.Hour)
    return c, nil
}
