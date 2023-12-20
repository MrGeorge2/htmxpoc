CREATE TABLE authors (
    id uuid NOT NULL CONSTRAINT pk_authors PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL
);