CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_number TEXT UNIQUE NOT NULL,
    address TEXT,
    pin TEXT NOT NULL,
    balance BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Add other tables like topups, payments, transfers if needed

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id),
    transaction_type TEXT NOT NULL,
    amount BIGINT NOT NULL,
    remarks TEXT,
    balance_before BIGINT,
    balance_after BIGINT,
    created_at TIMESTAMP DEFAULT NOW()
);