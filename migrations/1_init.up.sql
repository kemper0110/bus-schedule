BEGIN;

CREATE TYPE role AS ENUM ('customer', 'admin', 'carrier');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login varchar(50) UNIQUE,
    password varchar(100),
    role role DEFAULT 'customer'
);

COMMIT;