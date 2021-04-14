CREATE TABLE profile_images(
    id           int IS NOT NULL AUTO_INCREMENT,
    id_verification_url    VARCHAR(1000),
    id_card_image_url      VARCHAR(1000),
    PRIMARY KEY(id)
);


CREATE TABLE user_registration (
    userid           int NOT NULL AUTO_INCREMENT,
    idnumber         int, 
    photo_url        VARCHAR(100),
    phonenumber      VARCHAR(100) UNIQUE,
    email            VARCHAR(100),
    informal_address VARCHAR(100)
    residense        VARCHAR(100),
    roles             VARCHAR(100),
    stage            VARCHAR(100),
    createdAt        TIMESTAMP IS NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userid)
);

DROP TABLE registration_limbo;
CREATE TABLE registration_limbo (
    registerID      int NOT NULL AUTO_INCREMENT,
    idnumber        VARCHAR(100),
    phonenumber     VARCHAR(12),
    fcmToken        VARCHAR(500),
    stkPushid       VARCHAR(100),
    photo_url       VARCHAR(100),
    email           VARCHAR(1000),
    informal_address VARCHAR(1000),
    xCords          VARCHAR(100),
    yCords          VARCHAR(100),
    role            VARCHAR(100),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(registerID)
);

INSERT registration_limbo (idnumber,phonenumber,fcmToken,stkPushid,photo_url,email,informal_address,xCords,yCords,role) VALUES
("144433434","0788787878","eg32","","https://pic","cool@example.com","kiambu south","10.00","10.00","Farmer");
