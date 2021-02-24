CREATE TABLE wallets_store (
    wid int NOT NULL AUTO_INCREMENT,
    wallet_name VARCHAR(100) UNIQUE,
	balance BIGINT,
    PRIMARY KEY (wid)
);



INSERT INTO wallets_store (wallet_name,balance) VALUES("JOB",1000);
SELECT * FROM wallets_store;


CREATE TABLE [IF NOT EXISTS] ledger(
    event_id serial PRIMARY KEY,
    from string
    to string
    type int
)

