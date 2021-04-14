CREATE TABLE profile_images(
    id           int IS NOT NULL AUTO_INCREMENT,
    id_verification_url    VARCHAR(1000),
    id_card_image_url      VARCHAR(1000),
    PRIMARY KEY(id)
);

DROP TABLE user_registration;
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
    passwordHash      VARCHAR(60),
    informal_address  VARCHAR(100),
    xCords            VARCHAR(100),
    yCords            VARCHAR(100),
    role              VARCHAR(100),
    stage             VARCHAR(100),
    createdAt        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userid)
);

DROP TABLE registration_limbo;
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
    passwordHash    VARCHAR(60) DEFAULT "0x",
    informal_address VARCHAR(1000),
    xCords          VARCHAR(100),
    yCords          VARCHAR(100),
    role            VARCHAR(100),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(registerID)
);

INSERT registration_limbo (idnumber,phonenumber,fname,mname,lname,fcmToken,checkoutRequestID,photo_url,email,passwordHash,informal_address,xCords,yCords,role) VALUES
("144433434","0788787878","F","M","L","eg32","","https://pic","0xasdasd","cool@example.com","kiambu south","10.00","10.00","Farmer");


SELECT registerID,phonenumber,checkoutRequestID,passwordHash,role FROM registration_limbo;