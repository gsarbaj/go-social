CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    username varchar(255),
    email citext UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now()

)