CREATE TABLE transactions
(
    id TEXT NOT NULL,
    account_id TEXT NOT NULL,
    amount REAL NOT NULL,
    status TEXT NOT NULL,
    error_message TEXT NOT NULL,
    create_at TEXT NOT NULL,
    update_at TEXT NOT NULL
);