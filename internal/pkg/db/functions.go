package db

func GetRestaurant(name string) (*Restaurant, error) {
    db, err := get()
    if err != nil {
        return nil, err
    }

    var r Restaurant
    err = db.Where("name = ?", name).First(&r).Error

    return &r, err
}

func PostRestaurant(r Restaurant) (*Restaurant, error) {
    db, err := get()
    if err != nil {
        return nil, err
    }

    err = db.Create(&r).Error

    return &r, err
}
