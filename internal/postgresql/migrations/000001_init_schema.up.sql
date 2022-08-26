create table if not exists schema_migrations
(
    version bigint  not null primary key,
    dirty   boolean not null
);



create table if not exists users
(
    id            bigserial primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    user_name     varchar(255) not null unique,
    email         varchar(100) not null unique,
    password_hash varchar(100) not null,
    profile_image text    default '' :: text,
    phone_number  text         not null unique,
    is_verified   boolean default false,
    fts           tsvector
);

-- create index on

create index on users (id, created_at) where is_verified = true;
create index on users (is_verified, id, created_at) where is_verified = false;



CREATE FUNCTION users_fts_trigger() RETURNS trigger AS
$$
BEGIN
    new.fts :=
                    setweight(to_tsvector('pg_catalog.english', new.user_name), 'A') ||
                    setweight(to_tsvector('pg_catalog.english', new.email), 'B') ||
                    setweight(to_tsvector('pg_catalog.english', new.phone_number), 'C');
    return new;
END
$$ LANGUAGE plpgsql;


CREATE TRIGGER tgr_search_idx_fts_update
    BEFORE INSERT OR UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION users_fts_trigger();



CREATE INDEX users_fts_idx ON users USING gist (fts);



create table if not exists payment_transactions
(
    id                       bigserial
        primary key,
    created_at               timestamp with time zone not null,
    updated_at               timestamp with time zone not null,
    transaction_ref          varchar(255)             not null,
    status                   boolean                  not null,
    transaction_complete     boolean                  not null,
    data                     jsonb                    not null,
    code                     varchar(255)             not null,
    payment_integration_type integer                  not null,
    payment_purpose          integer                  not null,
    amount                   double precision         not null,
    payment_mode             integer                  not null,
    currency                 varchar(255)             not null,
    phone_number             varchar(255)             not null,
    fts                      tsvector
);

CREATE FUNCTION payments_transaction_fts_trigger() RETURNS trigger AS
$$
BEGIN
    new.fts :=
                setweight(to_tsvector('pg_catalog.english', new.phone_number), 'A') ||
                setweight(to_tsvector('pg_catalog.english', new.code), 'B');
    return new;
END
$$ LANGUAGE plpgsql;


CREATE TRIGGER tgr_search_idx_fts_update
    BEFORE INSERT OR UPDATE
    ON payment_transactions
    FOR EACH ROW
EXECUTE FUNCTION payments_transaction_fts_trigger();

CREATE INDEX payments_transaction_fts_idx ON payment_transactions USING gist (fts);

CREATE INDEX idx_payments_transaction_data ON payment_transactions USING GIN (data jsonb_ops);

create index on payment_transactions (status, id, payment_mode, phone_number) where status = true and transaction_complete = true;;
create index on payment_transactions (status, id, payment_mode, phone_number);



create table if not exists category
(
    id         bigserial primary key,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null,
    icon       text
);
