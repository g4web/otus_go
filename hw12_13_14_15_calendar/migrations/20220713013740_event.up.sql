CREATE TABLE IF NOT EXISTS event
(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    title VARCHAR NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    notification_before int DEFAULT 0 NOT NULL
);
