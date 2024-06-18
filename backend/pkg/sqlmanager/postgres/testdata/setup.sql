CREATE SCHEMA IF NOT EXISTS sqlmanagerpostgres;

SET search_path TO sqlmanagerpostgres;

CREATE TABLE users (
    id TEXT NOT NULL,
    age int NOT NULL,
    current_salary float NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    fullname TEXT GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED
);

CREATE TABLE users_with_identity (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    age int NOT NULL
);

CREATE TABLE unique_emails (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email text not null UNIQUE,
    username text not null,
    CONSTRAINT unique_email_username UNIQUE (email, username)
);

CREATE TABLE parent1 (id uuid NOT NULL DEFAULT gen_random_uuid(), CONSTRAINT pk_parent1_id PRIMARY KEY (id));
CREATE TABLE child1 (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    parent_id uuid NULL,
    CONSTRAINT pk_child1_id PRIMARY KEY (id),
    CONSTRAINT fk_child1_parent_id_parent1_id FOREIGN KEY (parent_id) REFERENCES parent1(id) ON
    DELETE
        CASCADE
);

-- testing basic circular deps
CREATE TABLE t1 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL,
    CONSTRAINT t1_b_fkey FOREIGN KEY (b) REFERENCES t1(a)
);
CREATE TABLE t2 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL
);
CREATE TABLE t3 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL
);
ALTER TABLE t2
ADD CONSTRAINT t2_b_fkey FOREIGN KEY (b) REFERENCES t3(a);
ALTER TABLE t3
ADD CONSTRAINT t3_b_fkey FOREIGN KEY (b) REFERENCES t2(a);

-- Testing composite keys
CREATE TABLE t4 (
    a int NOT NULL,
    b int NOT NULL,
    c int NULL,
    PRIMARY KEY (a, b)
);
CREATE TABLE t5 (
    x int NOT NULL,
    y int NOT NULL,
    z int NULL,
    CONSTRAINT t5_t4_fkey FOREIGN KEY (x, y) REFERENCES t4 (a, b)
);

CREATE SEQUENCE custom_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;
CREATE TYPE custom_type AS (
  part1 INTEGER,
  part2 TEXT
);
CREATE FUNCTION custom_function() RETURNS TRIGGER AS $$
BEGIN
  NEW.id := nextval('custom_seq');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TYPE custom_enum AS ENUM (
  'value1',
  'value2',
  'value3'
);
CREATE DOMAIN custom_domain AS TEXT
  CHECK (VALUE ~ '^[a-zA-Z]+$'); -- Only allows alphabetic characters

CREATE TABLE custom_table (
  id INTEGER NOT NULL DEFAULT nextval('custom_seq'),
  name custom_domain NOT NULL,
  data custom_type,
  status custom_enum NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
-- Adding a trigger to use the custom function for setting the 'id' field
CREATE TRIGGER set_custom_id
  BEFORE INSERT ON custom_table
  FOR EACH ROW
  EXECUTE FUNCTION custom_function();
CREATE INDEX idx_name ON custom_table(name);

CREATE TABLE tablewithcount (
    id TEXT NOT NULL
);
INSERT INTO tablewithcount(id) VALUES ('1'), ('2');
