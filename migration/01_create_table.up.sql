CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  

SELECT uuid_generate_v4();  

CREATE TABLE product(
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR,
    price NUMERIC
);

CREATE TABLE orders(
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    userID VARCHAR,
    productID VARCHAR,
    totalSum integer
);


INSERT INTO orders(userID,productID,totalSum) VALUES('a1758aa1-d582-4b92-abcd-5572f1738e4f','c8d328ed-154f-4008-a88e-390903145ec1',150);
INSERT INTO orders(userID,productID,totalSum) VALUES('11f8ffb4-78da-44b7-bed5-6d0e8e0f12a0','f269cf7d-5f39-4987-b20b-b43edfd7e32e',410);
INSERT INTO orders(userID,productID,totalSum) VALUES('0abbf2d6-1b37-4542-bfbb-b82458cc497c','c8d328ed-154f-4008-a88e-390903145ec1',320);
INSERT INTO orders(userID,productID,totalSum) VALUES('a1758aa1-d582-4b92-abcd-5572f1738e4f','f269cf7d-5f39-4987-b20b-b43edfd7e32e',250);

