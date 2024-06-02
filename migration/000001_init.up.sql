CREATE TABLE IF NOT EXISTS orderr
(
    order_uid varchar(255) not null unique,
    track_number varchar(255) not null,
    entry varchar(255) not null,
    locale varchar(255) not null,
    internal_signature varchar(255) not null,
    customer_id varchar(255) not null,
    delivery_service varchar(255) not null,
    shardkey varchar(255) not null,
    sm_id int not null unique,
    date_created varchar(255) not null,
    oof_shard varchar(255) not null
);

CREATE TABLE IF NOT EXISTS delivery
(
    name varchar(255) not null,
    phone varchar(255) not null,
    zip int not null,
    city varchar(255) not null,
    address varchar(255) not null,
    region varchar(255) not null,
    email varchar(255) not null
);

CREATE TABLE IF NOT EXISTS payment
(
    transaction varchar(255) not null,
    request_id varchar(255) not null,
    currency varchar(255) not null,
    provider varchar(255) not null,
    amount int not null,
    payment_dt int not null,
    bank varchar(255) not null,
    delivery_cost int not null,
    goods_total int not null,
    custom_fee int not null
);

CREATE TABLE IF NOT EXISTS item
(
    chrt_id int not null,
    track_number varchar(255) not null,
    price int not null,
    rid varchar(255) not null,
    name varchar(255) not null,
    sale int,
    size int,
    total_price int,
    nm_id varchar(255) not null,
    brand varchar(255) not null,
    status int
);

