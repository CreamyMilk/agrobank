CREATE TABLE profile_images(
    id           IS NOT NULL AUTO_INCREMENT,
    id_verification_url    VARCHAR(1000),
    id_card_image_url      VARCHAR(1000),
    PRIMARY KEY(id)
);


CREATE TABLE user_registration (
    userid           IS NOT NULL AUTO_INCREMENT,
    idnumber         INT, 
    photo_url        VARCHAR(100),
    phonenumber      VARCHAR(100) UNIQUE,
    email            VARCHAR(100),
    informal_address VARCHAR(100)
    residense        VARCHAR(100),
    role             VARCHAR(100),
    stage            VARCHAR(100),
    createdAt        TIMESTAMP IS NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(userid)
);


CREATE TABLE registration_limbo (
    registerID      IS NOT NULL AUTO_INCREMENT,
    phonenumber     VARCHAR(12),
    fcmToken        VARCHAR(500),
    stkPushid       VARCHAR(100),
    photo_url       VARCHAR(100),
    email           VARCHAR(1000),
    informal_ddress VARCHAR(1000),
    cordinates      VARCHAR(100),
    role            VARCHAR(100),
    createdAt       TIMESTAMP IS NOT NULL DEFAULT CURRENT_TIMESTAMP,

)
