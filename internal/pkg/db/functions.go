package db

// GetRestaurant returns a Restaurant object from the database based on the
// restaurant's name.
func GetRestaurant(name string) (*Restaurant, error) {
    db, err := get()
    if err != nil {
        return nil, err
    }

    var r Restaurant
    err = db.Where("name = ?", name).First(&r).Error

    return &r, err
}

// PostRestaurant creates an entry and returns the corresponding Restaurant
// object.
func PostRestaurant(r Restaurant) (*Restaurant, error) {
    db, err := get()
    if err != nil {
        return nil, err
    }

    err = db.Create(&r).Error

    return &r, err
}
