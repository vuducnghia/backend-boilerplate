CREATE TABLE users (
       id           VARCHAR PRIMARY KEY,
       username     VARCHAR UNIQUE NOT NULL,
       created_at   TIMESTAMPTZ NULL,
       modified_at  TIMESTAMPTZ NULL,
       deleted_at   TIMESTAMPTZ NULL,
       first_name   VARCHAR NOT NULL,
       last_name    VARCHAR NOT NULL,
       phone_number VARCHAR NOT NULL,
       password     VARCHAR NOT NULL
);
