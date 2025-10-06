# Page 17

## Text Content

```
location_id UUID NOT NULL REFERENCES locations(id),
instructor_id UUID NOT NULL REFERENCES instructors(id),
total_seats INTEGER NOT NULL,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- Indexes for Events
CREATE INDEX index_events_on_event_date ON events(event_date);
CREATE INDEX index_events_on_location_id ON events(location_id);
CREATE INDEX index_events_on_instructor_id ON events(instructor_id);

-- =========================================
-- Migration 012: Create Users Table
-- =========================================
CREATE TABLE users (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
email VARCHAR NOT NULL UNIQUE,
password_digest VARCHAR NOT NULL,
first_name VARCHAR NOT NULL,
last_name VARCHAR NOT NULL,
phone VARCHAR,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================================
-- Migration 013: Create Attendees Table
-- =========================================
CREATE TABLE attendees (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
first_name VARCHAR NOT NULL,
last_name VARCHAR NOT NULL,
email VARCHAR NOT NULL UNIQUE,
phone VARCHAR,
user_id UUID REFERENCES users(id),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- The index on event_id has been removed since the 'attendees' table does not contain


```

