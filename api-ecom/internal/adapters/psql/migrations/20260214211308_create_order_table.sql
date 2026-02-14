-- +goose Up
-- +goose StatementBegin
create table if not exists orders (
    id BIGSERIAL PRIMARY KEY,
    user_id bigint not NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

create table if not exists order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGSERIAL not NULL,
    product_id BIGSERIAL not NULL,
    price INTEGER not NULL,
    quantity INTEGER not NULL,
    constraint fk_order foreign key (order_id) references orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists order_items;

DROP table if exists orders;

-- +goose StatementEnd
