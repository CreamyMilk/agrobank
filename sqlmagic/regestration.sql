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
    phonenumber      INT,
    email            VARCHAR(100),
    informal_address VARCHAR(100)
    residense        VARCHAR(100),
    role             VARCHAR(100),
    stage            VARCHAR(100)
);
