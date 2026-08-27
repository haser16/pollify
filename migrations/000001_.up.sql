CREATE SCHEMA IF NOT EXISTS pollify;

CREATE TABLE pollify.users
(
    id            SERIAL PRIMARY KEY,
    full_name     VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 2 AND 100),
    email         VARCHAR(100) NOT NULL UNIQUE,
    phone_number  VARCHAR(15)  CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
),
    password      CHAR(120) NOT NULL
);

CREATE TABLE pollify.polls
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 2 AND 100) UNIQUE,
    description VARCHAR(1000)         CHECK (char_length(description) BETWEEN 1 AND 1000),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at  TIMESTAMPTZ NOT NULL,
    is_expire   BOOLEAN NOT NULL DEFAULT FALSE,
    owner_id    INT NOT NULL REFERENCES pollify.users(id) ON DELETE CASCADE
);

CREATE TABLE pollify.questions
(
    id            SERIAL PRIMARY KEY,
    poll_id       INT NOT NULL REFERENCES pollify.polls(id) ON DELETE CASCADE,
    question_text VARCHAR(1000) NOT NULL,
    is_multiple   BOOLEAN DEFAULT FALSE
);

CREATE TABLE pollify.options
(
    id          SERIAL PRIMARY KEY,
    question_id INT NOT NULL REFERENCES pollify.questions(id) ON DELETE CASCADE,
    option_text VARCHAR(500) NOT NULL
);

CREATE TABLE pollify.votes
(
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL REFERENCES pollify.users(id) ON DELETE CASCADE,
    question_id INT NOT NULL REFERENCES pollify.questions(id) ON DELETE CASCADE,
    option_id   INT NOT NULL REFERENCES pollify.options(id) ON DELETE CASCADE,
    voted_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_user_option UNIQUE (user_id, option_id)
);
