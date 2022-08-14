-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS social_group
(
    social_group_id serial NOT NULL,
    name            text,
    CONSTRAINT "social_group_pk" PRIMARY KEY (social_group_id)
);

INSERT INTO social_group (name)
VALUES ('Айтишники'),
       ('Пенсионеры'),
       ('Студенты');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS social_group;
-- +goose StatementEnd
