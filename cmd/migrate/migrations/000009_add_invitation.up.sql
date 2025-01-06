CREATE TABLE IF NOT EXISTS user_invitations (
    t bytea PRIMARY KEY,
    user_id bigint NOT NULL
)