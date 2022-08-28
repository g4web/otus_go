-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event
(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    title VARCHAR NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    notification_before int DEFAULT 0 NOT NULL,
    notification_is_sent BOOLEAN DEFAULT FALSE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event;
-- +goose StatementEnd
