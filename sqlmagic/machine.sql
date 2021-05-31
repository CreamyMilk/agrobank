DROP TABLE m_categories;
DROP TABLE machines;
CREATE TABLE m_categories (
    category_id       INT NOT NULL AUTO_INCREMENT,
    category_name     VARCHAR(100) UNIQUE,
    category_image    VARCHAR(1000),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY       (category_id)
);

CREATE TABLE machines (
    machineID          INT NOT NULL AUTO_INCREMENT,
    category_id         INT NOT NULL,
    owner_id            INT NOT NULL,
    machine_name        VARCHAR(100),
    machine_image       VARCHAR(1000),
    machine_image_large VARCHAR(1000),
    descriptions        VARCHAR(500),
    price               DECIMAL(10,4),
    stock               INT, 
    machine_packtype    VARCHAR(100),
    createdAt           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(machineID),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);



INSERT INTO m_categories (category_name,category_image) 
VALUES ("10 Ton Truck","https://foods.com");

INSERT INTO m_categories (category_name,category_image) 
VALUES ("3 Ton Truck","https://foods.com");


INSERT INTO m_categories (category_name,category_image) 
VALUES ("1 Ton Truck ","https://foods.com");


INSERT INTO m_categories (category_name,category_image) 
VALUES ("Motorbike","https://foods.com");



INSERT INTO machines (
category_id,owner_id,machine_name,machine_image,machine_image_large,
descriptions,price,stock,machine_packtype )
VALUES (1,3,"Machine","https://carotsimage.com","nolarge",
"Carrots are good for your eyes",100,5,"Boxes");



UPDATE machines
SET category_id=1,
    owner_id = 1,
    machine_name="A new Name",
    machine_image="image",
    machine_image_large="Large Image",
    descriptions="New Description",
    price=901.50,
    stock=12,
    machine_packtype="Boxes More Larger than normal"
WHERE machineID=99;
