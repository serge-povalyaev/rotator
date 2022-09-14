BEGIN;

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

CREATE TABLE banner_to_slot
(
    slot_id   bigint NOT NULL,
    banner_id bigint NOT NULL
);

ALTER TABLE banner_to_slot
    ADD CONSTRAINT fk_banner_to_slot_slot_id FOREIGN KEY (slot_id)
        REFERENCES slot (slot_id);


ALTER TABLE banner_to_slot
    ADD CONSTRAINT fk_banner_to_slot_banner_id FOREIGN KEY (banner_id)
        REFERENCES banner (banner_id);

CREATE TABLE total_stat
(
    slot_id         bigint NOT NULL,
    banner_id       bigint NOT NULL,
    social_group_id bigint NOT NULL,
    shows           bigint,
    clicks          bigint,
    updated_at      date NOT NULL
);

ALTER TABLE total_stat
    ADD CONSTRAINT fk_total_stat_slot_id FOREIGN KEY (slot_id)
        REFERENCES slot (slot_id);

ALTER TABLE total_stat
    ADD CONSTRAINT fk_total_stat_banner_id FOREIGN KEY (banner_id)
        REFERENCES banner (banner_id);

ALTER TABLE total_stat
    ADD CONSTRAINT fk_total_stat_social_group_id FOREIGN KEY (social_group_id)
        REFERENCES social_group (social_group_id);

CREATE TABLE stat
(
    slot_id         bigint NOT NULL,
    banner_id       bigint NOT NULL,
    social_group_id bigint NOT NULL,
    action_type     smallint,
    created_at      date NOT NULL
);

ALTER TABLE stat
    ADD CONSTRAINT fk_stat_slot_id FOREIGN KEY (slot_id)
        REFERENCES slot (slot_id);

ALTER TABLE stat
    ADD CONSTRAINT fk_stat_banner_id FOREIGN KEY (banner_id)
        REFERENCES banner (banner_id);

ALTER TABLE stat
    ADD CONSTRAINT fk_stat_social_group_id FOREIGN KEY (social_group_id)
        REFERENCES social_group (social_group_id);

END;