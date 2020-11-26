CREATE TABLE IF NOT EXISTS wallet (
    id UUID NOT NULL DEFAULT uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
    owned_by UUID NOT NULL,
    status VARCHAR(128),
    balance NUMERIC DEFAULT 0,
    enabled_at TIMESTAMP,
    disabled_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS deposit (
    id UUID NOT NULL DEFAULT uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
    deposited_by UUID NOT NULL,
    reference_id TEXT UNIQUE,
    status VARCHAR(128),
    amount NUMERIC DEFAULT 0,
    deposited_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS withdrawal (
    id UUID NOT NULL DEFAULT uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
    withdrawn_by UUID NOT NULL,
    reference_id TEXT UNIQUE,
    status VARCHAR(128),
    amount NUMERIC DEFAULT 0,
    withdrawn_at TIMESTAMP NOT NULL
);