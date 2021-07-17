DROP     DATABASE IF EXISTS agrodb;
CREATE   DATABASE agrodb;
USE      agrodb;

CREATE TABLE wallets_store (
    wid          int NOT NULL AUTO_INCREMENT,
    wallet_name  VARCHAR(100) UNIQUE,
	balance      BIGINT,
    PRIMARY KEY  (wid)
);

CREATE TABLE transaction_costs(
    rate_id      int NOT NULL AUTO_INCREMENT,
    upper_limit  BIGINT,
    cost         BIGINT,
    PRIMARY KEY  (rate_id)
);

CREATE TABLE transactions_type(
    type INT NOT NULL UNIQUE,
    name VARCHAR(1000)
);

CREATE TABLE transactions_list(
    tid             int NOT NULL AUTO_INCREMENT,
    transuuid       VARCHAR(100) UNIQUE,
    sender_name     VARCHAR(100),
    receiver_name   VARCHAR(100),
    amount          BIGINT,
    charge          BIGINT,
    ttype           INT,
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY     (tid)
);

CREATE TABLE deposit_attempts(
    did               int NOT NULL AUTO_INCREMENT,
    checkoutRequestID VARCHAR(100) UNIQUE,
    walletname        VARCHAR(100),
    amount            BIGINT,
    ttype             INT DEFAULT 0,
    method            VARCHAR(100) DEFAULT "MPESA",
    mpesaID           VARCHAR(100),
    createdAt         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY       (did)
);


INSERT INTO transactions_type (type,name) VALUES (0, "Deposit");
INSERT INTO transactions_type (type,name) VALUES (1, "Withdraw");
INSERT INTO transactions_type (type,name) VALUES (2, "SendMoney");
INSERT INTO transactions_type (type,name) VALUES (3, "SendToMpesa");

INSERT INTO transaction_costs (upper_limit,cost) VALUES (100,1);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (1000,5);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (10000,10);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (100000,15);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (1000000,20);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (10000000,25);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (100000000,100);
INSERT INTO transaction_costs (upper_limit,cost) VALUES (999999999999,200);

SELECT  cost FROM transaction_costs WHERE upper_limit >=10000 LIMIT 1;
INSERT       INTO wallets_store (wallet_name,balance) VALUES("JOB",1000);
