CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE wallets (
    wallet_id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner VARCHAR NOT NULL ,
    balance BIGINT NOT NULL ,
    currency VARCHAR NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entries (
    entrie_id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    wallet_id UUID NOT NULL ,
    amount BIGINT NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers (
    transfer_id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_wallet_id UUID NOT NULL ,
    to_wallet_id UUID NOT NULL ,
    amount BIGINT NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE entries ADD FOREIGN KEY (wallet_id) REFERENCES wallets(wallet_id);
ALTER TABLE transfers ADD FOREIGN KEY (from_wallet_id) REFERENCES wallets(wallet_id);
ALTER TABLE transfers ADD FOREIGN KEY (to_wallet_id) REFERENCES wallets(wallet_id);

CREATE INDEX IF NOT EXISTS wallets_owner_idx ON wallets (owner);
CREATE INDEX IF NOT EXISTS entries_wallet_id_idx ON entries(wallet_id);
CREATE INDEX IF NOT EXISTS transfers_from_wallet_id_idx ON transfers(from_wallet_id);
CREATE INDEX IF NOT EXISTS transfers_to_wallet_id_idx ON transfers(to_wallet_id);
CREATE INDEX IF NOT EXISTS transfers_from_wallet_id_to_wallet_id_idx ON transfers(from_wallet_id, to_wallet_id);
