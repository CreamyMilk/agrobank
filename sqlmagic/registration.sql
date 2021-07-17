DROP TABLE IF EXISTS user_registration;
CREATE TABLE user_registration (
    userid            int NOT NULL AUTO_INCREMENT,
    fname             VARCHAR(100),
    mname             VARCHAR(100),
    lname             VARCHAR(100),
    idnumber          VARCHAR(100),
    checkoutRequestID VARCHAR(100),
    photo_url         VARCHAR(100),
    phonenumber       VARCHAR(12) UNIQUE,
    email             VARCHAR(100),
    passwordHash      VARCHAR(65),
    informal_address  VARCHAR(100),
    xCords            VARCHAR(100),
    yCords            VARCHAR(100),
    role              VARCHAR(100),
    stage             VARCHAR(100),
    createdAt        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userid)
);

DROP TABLE IF EXISTS registration_limbo;
CREATE TABLE registration_limbo (
    registerID      int NOT NULL AUTO_INCREMENT,
    fname           VARCHAR(100),
    mname           VARCHAR(100),
    lname           VARCHAR(100),
    idnumber        VARCHAR(100),
    phonenumber     VARCHAR(12),
    fcmToken        VARCHAR(500),
    checkoutRequestID       VARCHAR(100),
    photo_url       VARCHAR(100),
    email           VARCHAR(1000),
    passwordHash    VARCHAR(65) DEFAULT "0x",
    informal_address VARCHAR(1000),
    xCords          VARCHAR(100),
    yCords          VARCHAR(100),
    role            VARCHAR(100),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(registerID)
);

INSERT registration_limbo (idnumber,phonenumber,fname,mname,lname,fcmToken,checkoutRequestID,photo_url,passwordHash,email,informal_address,xCords,yCords,role) VALUES
("144433434","254797678251","F","M","L","eg32","PPPZ","https://pic","$2a$04$xLHL53ke/GFU.LsG/1KOUOOy8zjWNYvyzSGM0vkoM0kKx.SDcpKtm","cool@example.com","kiambu south","10.00","10.00","Farmer");


SELECT registerID,phonenumber,checkoutRequestID,passwordHash,role FROM registration_limbo;
