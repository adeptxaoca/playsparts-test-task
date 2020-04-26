CREATE TABLE parts (
   id serial PRIMARY KEY,
   manufacturer_id int REFERENCES manufacturers,
   name varchar(255) NOT NULL,
   vendor_code varchar(100) NOT NULL UNIQUE,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
   deleted_at TIMESTAMP WITH TIME ZONE
);
---- create above / drop below ----
DROP TABLE parts;