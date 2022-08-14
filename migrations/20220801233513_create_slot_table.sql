-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS slot
(
    slot_id serial NOT NULL,
    name    text,
    CONSTRAINT "slot_pk" PRIMARY KEY (slot_id)
);

INSERT INTO slot (name)
VALUES ('Шапка'),
       ('Слева'),
       ('Справа'),
       ('Футер');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS slot;
-- +goose StatementEnd
