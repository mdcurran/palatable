package db

import (
    "time"

    "github.com/jinzhu/gorm"
    "github.com/segmentio/ksuid"
)

// Restaurant defines a restaurant, cafe, bar, or anywhere that sells items
// that could be recorded.
type Restaurant struct {
    ID        string    `db:"id" gorm:"primary_key" json:"id,omitempty"`
    Name      string    `db:"name" json:"name,omitempty"`
    CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (r *Restaurant) BeforeCreate(scope *gorm.Scope) error {
    id := "rest_" + ksuid.New().String()
    scope.SetColumn("id", id)
    return nil
}

// RestaurantAddress contains the address of a restaurant. Storing addresses
// separately from restaurants accommodates restaurants with multiple
// locations.
type RestaurantAddress struct {
    ID           string    `db:"id" gorm:"primary_key" json:"id,omitempty"`
    RestaurantID string    `db:"restaurant_id" json:"restaurant_id,omitempty"`
    LineOne      string    `db:"line_one" json:"line_one,omitempty"`
    LineTwo      string    `db:"line_two" json:"line_two,omitempty"`
    City         string    `db:"city" json:"city,omitempty"`
    Postcode     string    `db:"postcode" json:"postcode,omitempty"`
    Country      string    `db:"country" json:"country,omitempty"`
    CreatedAt    time.Time `db:"created_at" json:"created_at,omitempty"`
    UpdatedAt    time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (a *RestaurantAddress) BeforeCreate(scope *gorm.Scope) error {
    id := "rest_addr_" + ksuid.New().String()
    scope.SetColumn("id", id)
    return nil
}
