-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stat;
-- +goose StatementEnd
