CREATE TABLE IF NOT EXISTS todos (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    is_completed boolean NOT NULL DEFAULT false,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);