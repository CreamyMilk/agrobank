DROP TABLE escrowInvoices;
CREATE TABLE escrowInvoices(
    eid                int NOT NULL AUTO_INCREMENT,
    reconciliationcode VARCHAR(20),
    senderWalletName   VARCHAR(100),
    receiverWalletName VARCHAR(100),
    prodcutID          INT,
    amount             BIGINT,
    CreatedAt          DATETIME,
    Deadline           DATETIME,
    CompletedAT        DATETIME,
    TransactionCode    VARCHAR(100),
    PRIMARY KEY(eid)
);