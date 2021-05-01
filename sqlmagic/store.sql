DROP TABLE categories;
DROP TABLE products;
CREATE TABLE categories (
    category_id       INT NOT NULL AUTO_INCREMENT,
    category_name     VARCHAR(100) UNIQUE,
    category_image    VARCHAR(100),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY       (category_id)
);

CREATE TABLE products (
    product_id          INT NOT NULL AUTO_INCREMENT,
    category_id         INT NOT NULL,
    owner_id            INT NOT NULL,
    product_name        VARCHAR(100),
    product_image       VARCHAR(100),
    product_image_large VARCHAR(100),
    descriptions        VARCHAR(500),
    price              DECIMAL,
    stock               DECIMAL, 
    product_packtype    VARCHAR(100),
    createdAt       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(product_id),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);



INSERT INTO categories (category_name,category_image) 
VALUES ("FoodStaff","https://foods.com");


INSERT INTO products (
category_id,owner_id,product_name,product_image,product_image_large,
descriptions,price,stock,product_packtype )
VALUES (1,1,"Carrots","https://carotsimage.com","nolarge",
"Carrots are good for your eyes",100,5,"Boxes");



UPDATE products
SET category_id=1,
    owner_id = 1,
    product_name="A new Name",
    product_image="image",
    product_image_large="Large Image",
    descriptions="New Description",
    price=901.50,
    stock=12,
    product_packtype="Boxes More Larger than normal"
WHERE product_id=99;
