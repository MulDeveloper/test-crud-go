CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS cars (
    id UUID NOT NULL,
    brand VARCHAR(150) NOT NULL,
    model VARCHAR(150) NOT NULL,
    price NUMERIC NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_cars PRIMARY KEY(id)
);