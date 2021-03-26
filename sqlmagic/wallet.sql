CREATE DATABASE agrodb;
USE agrodb;
CREATE TABLE wallets_store (
    wid int NOT NULL AUTO_INCREMENT,
    wallet_name VARCHAR(100) UNIQUE,
	balance BIGINT,
    PRIMARY KEY (wid)
);

CREATE TABLE transaction_costs(
    rate_id int NOT NULL AUTO_INCREMENT,
    upper_limit BIGINT,
    cost BIGINT,
    PRIMARY KEY (rate_id)
);

INSERT INTO transaction_costs (upper_limit,cost) VALUES (100,1);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (1000,5);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (10000,10);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (100000,15);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (1000000,20);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (10000000,25);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (100000000,100000000);


SELECT cost FROM transaction_costs WHERE upper_limit >=10000 LIMIT 1;


INSERT INTO wallets_store (wallet_name,balance) VALUES("JOB",1000);
SELECT * FROM wallets_store;


CREATE TABLE transactions_list{
    transID       VARCHAR(100) UNIQUE,
    sender_name   VARCHAR(100) UNIQUE,
    receiver_name VARCHAR(100) UNIQUE
}

CREATE TABLE [IF NOT EXISTS] ledger(
    event_id serial PRIMARY KEY,
    from string
    to string
    type int
)

