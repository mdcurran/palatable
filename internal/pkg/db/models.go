package db

import (
    "time"

    "github.com/jinzhu/gorm"
    "github.com/segmentio/ksuid"
)

// Restaurant defines a restaurant, cafe, bar, or anywhere that sells items
// that could be recorded, along with its address.
type Restaurant struct {
    ID        string    `db:"id" gorm:"primary_key" json:"id,omitempty"`
    Name      string    `db:"name" json:"name,omitempty"`
    LineOne   string    `db:"line_one" json:"line_one,omitempty"`
    LineTwo   string    `db:"line_two" json:"line_two,omitempty"`
    City      string    `db:"city" json:"city,omitempty"`
    Postcode  string    `db:"postcode" json:"postcode,omitempty"`
    Country   string    `db:"country" json:"country,omitempty"`
    CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (r *Restaurant) BeforeCreate(scope *gorm.Scope) error {
    id := "rest_" + ksuid.New().String()
    scope.SetColumn("id", id)
    return nil
}

// Item is any kind of dish obtained in a restaurant.
type Item struct {
    ID           string    `db:"id" gorm:"primary_key" json:"id,omitempty"`
    RestaurantID string    `db:"restaurant_id" json:"restaurant_id,omitempty"`
    Title        string    `db:"title" json:"title,omitempty"`
    Description  string    `db:"description" json:"description,omitempty"`
    Review       string    `db:"review" json:"review,omitempty"`
    Date         time.Time `db:"date" json:"date,omitempty"`
    CreatedAt    time.Time `db:"created_at" json:"created_at,omitempty"`
    UpdatedAt    time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (i *Item) BeforeCreate(scope *gorm.Scope) error {
    id := "item_" + ksuid.New().String()
    scope.SetColumn("id", id)
    return nil
}
