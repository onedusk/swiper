# Page 16

## Text Content

```
-- =========================================
-- Migration 009: Create Locations Table
-- =========================================
CREATE TABLE locations (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR NOT NULL,
address VARCHAR NOT NULL,
city VARCHAR NOT NULL,
state VARCHAR,
country VARCHAR NOT NULL,
postal_code VARCHAR,
capacity INTEGER,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================================
-- Migration 010: Create Instructors Table
-- =========================================
CREATE TABLE instructors (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR NOT NULL,
email VARCHAR NOT NULL UNIQUE,
bio TEXT,
phone VARCHAR,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================================
-- Migration 011: Create Events Table
-- =========================================
CREATE TABLE events (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR NOT NULL,
description TEXT,
event_date DATE NOT NULL,
start_time TIME WITH TIME ZONE NOT NULL,
end_time TIME WITH TIME ZONE NOT NULL,


```

