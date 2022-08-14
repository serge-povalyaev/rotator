-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banner
(
    banner_id serial NOT NULL,
    name      text,
    CONSTRAINT "banner_pk" PRIMARY KEY (banner_id)
);

INSERT INTO banner (name)
VALUES ('Реклама курсов'),
       ('Реклама магазинов'),
       ('Реклама автомобилей'),
       ('Реклама техники'),
       ('Реклама еды');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banner;
-- +goose StatementEnd
