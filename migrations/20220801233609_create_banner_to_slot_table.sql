-- +goose Up
-- +goose StatementBegin
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banner_to_slot;
-- +goose StatementEnd
