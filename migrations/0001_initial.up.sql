BEGIN;
CREATE TABLE IF NOT EXISTS restaurants (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    line_one TEXT NOT NULL,
    line_two TEXT,
    city TEXT NOT NULL,
    postcode TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_restaurants_name ON restaurants USING btree (name);
CREATE UNIQUE INDEX IF NOT EXISTS uix_restaurants ON restaurants USING btree (name, line_one, postcode, country);

CREATE TABLE IF NOT EXISTS items (
    id TEXT PRIMARY KEY,
    restaurant_id TEXT REFERENCES restaurants(id) ON DELETE CASCADE NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    review TEXT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uix_items ON items USING btree (restaurant_id, title, date);
COMMIT;
