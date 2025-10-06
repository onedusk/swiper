# Page 18

## Text Content

```
-- =========================================
-- Migration 014: Create Registrations Table
-- =========================================
CREATE TABLE registrations (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
event_id UUID NOT NULL REFERENCES events(id),
attendee_id UUID NOT NULL REFERENCES attendees(id),
registration_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
status VARCHAR NOT NULL DEFAULT 'confirmed',
CONSTRAINT valid_status CHECK (status IN ('confirmed', 'cancelled', 'waitlisted'))
);
-- Indexes for Registrations
CREATE INDEX index_registrations_on_event_id ON registrations(event_id);
CREATE INDEX index_registrations_on_attendee_id ON registrations(attendee_id);
CREATE INDEX index_registrations_on_status ON registrations(status);

-- =========================================
-- Migration 015: Create Event Availability View
-- =========================================
CREATE OR REPLACE VIEW event_availability AS
SELECT
e.id,
e.name,
e.event_date,
e.start_time,
e.end_time,
e.total_seats,
i.name AS instructor_name,
i.id AS instructor_id,
COUNT(r.id) FILTER (WHERE r.status = 'confirmed') AS seats_taken,
e.total_seats - COUNT(r.id) FILTER (WHERE r.status = 'confirmed') AS seats_availabl
COUNT(r.id) FILTER (WHERE r.status = 'waitlisted') AS waitlisted_count
FROM events e
JOIN instructors i ON e.instructor_id = i.id
LEFT JOIN registrations r ON e.id = r.event_id
GROUP BY e.id, i.id;


```

