CREATE TABLE manufacturers (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
---- create above / drop below ----
DROP TABLE manufacturers;