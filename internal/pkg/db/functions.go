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
func PostRestaurant(r *Restaurant) error {
    db, err := get()
    if err != nil {
        return err
    }

    err = db.Create(&r).Error

    return err
}

// GetItems returns all items for a given restaurant.
func GetItems(restaurantID string) ([]*Item, error) {
    db, err := get()
    if err != nil {
        return nil, err
    }

    var items []*Item
    err = db.Where("restaurant_id = ?", restaurantID).Find(&items).Error

    return items, err
}

// PostItem creates an entry and returns the corresponding Item object.
func PostItem(i *Item) error {
    db, err := get()
    if err != nil {
        return err
    }

    err = db.Create(&i).Error

    return err
}
