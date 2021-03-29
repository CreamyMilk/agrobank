
CREATE TABLE categories (
    category_id       INT NOT NULL AUTO_INCREMENT,
    category_name     VARCHAR(100) UNIQUE,
    category_image    VARCHAR(100),
    PRIMARY KEY       (category_id)
);

CREATE TABLE products (
    product_id          INT NOT NULL AUTO_INCREMENT,
    category_id         INT NOT NULL,
    product_name        VARCHAR(100),
    product_image       VARCHAR(100),
    product_image_large VARCHAR(100),
    descriptions        VARCHAR(500),
    amount              INT,
    stock               INT, 
    product_packtype    VARCHAR(100),
    PRIMARY KEY(product_id),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);



INSERT INTO categories (category_name,category_image) 
VALUES ("FoodStaff","https://foods.com");


INSERT INTO products (
category_id,product_name,product_image,product_image_large,
descriptions,amount,stock,product_packtype )
VALUES (1,"Carrots","https://carotsimage.com","nolarge",
"Carrots are good for your eyes",100,5,"Boxes");