CREATE TABLE auth (
      id                VARCHAR PRIMARY KEY,
      access_token      VARCHAR NULL,
      refresh_token     VARCHAR NULL,
      created_at        TIMESTAMPTZ NULL,
      modified_at       TIMESTAMPTZ NULL,
      deleted_at        TIMESTAMPTZ NULL,
      user_id           VARCHAR REFERENCES users (id)
);
CREATE UNIQUE INDEX auth_user_id_idx ON public.auth USING btree (user_id);





