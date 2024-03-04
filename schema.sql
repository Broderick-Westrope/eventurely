-- users Table
CREATE TABLE IF NOT EXISTS "users"
(
    id                   SERIAL PRIMARY KEY,
    email                TEXT NOT NULL UNIQUE,
    password_hash        TEXT NOT NULL,
    first_name           TEXT,
    last_name            TEXT,
    phone_number         TEXT UNIQUE,
    is_phone_verified    BOOLEAN                  DEFAULT FALSE,
    is_email_verified    BOOLEAN                  DEFAULT FALSE,
    dietary_preferences  TEXT,
    special_requirements TEXT,
    joined_at            TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- events Table
CREATE TABLE IF NOT EXISTS "events"
(
    id              SERIAL PRIMARY KEY,
    owner_id        INTEGER                                               NOT NULL,
    title           TEXT                                                  NOT NULL,
    description     TEXT,
    starts_at       TIMESTAMP WITH TIME ZONE                              NOT NULL,
    ends_at         TIMESTAMP WITH TIME ZONE,
    location        TEXT,
    unique_link     TEXT UNIQUE,
    privacy_setting TEXT CHECK (privacy_setting IN ('PUBLIC', 'PRIVATE')) NOT NULL DEFAULT 'PRIVATE',
    created_at      TIMESTAMP WITH TIME ZONE                                       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE                                       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES "users" (id)
);

-- invitations Table
CREATE TABLE IF NOT EXISTS "invitations"
(
    id           SERIAL PRIMARY KEY,
    event_id     INTEGER                                               NOT NULL,
    user_id      INTEGER                                               NOT NULL,
    sent_at      TIMESTAMP WITH TIME ZONE                                       DEFAULT CURRENT_TIMESTAMP,
    responded_at TIMESTAMP WITH TIME ZONE,
    status       TEXT CHECK (status IN ('SENT', 'YES', 'NO', 'MAYBE')) NOT NULL DEFAULT 'SENT',
    FOREIGN KEY (event_id) REFERENCES "events" (id),
    FOREIGN KEY (user_id) REFERENCES "users" (id)
);

-- Indexes
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON "users" (email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_phone_number ON "users" (phone_number);
CREATE INDEX IF NOT EXISTS idx_event_owner_id ON "events" (owner_id);
CREATE INDEX IF NOT EXISTS idx_invitation_event_id ON "invitations" (event_id);
CREATE INDEX IF NOT EXISTS idx_invitation_user_id ON "invitations" (user_id);
CREATE INDEX IF NOT EXISTS idx_invitation_status ON "invitations" (status);
CREATE INDEX IF NOT EXISTS idx_invitation_event_id_status ON "invitations" (event_id, status);
