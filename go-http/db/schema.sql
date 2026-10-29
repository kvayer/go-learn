CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    user_name VARCHAR(50) NOT NULL,
    user_email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    event_name VARCHAR(255) NOT NULL,
    start_at TIMESTAMPTZ NOT NULL,
    limit_participants INTEGER NOT NULL CHECK (limit_participants > 0),
    organizer_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT
);

CREATE TABLE user_events (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    user_id UUID NOT NULL REFERENCES users(id),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, event_id)
);