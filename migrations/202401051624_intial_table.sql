-- +goose Up

CREATE TABLE roles (
      id serial PRIMARY KEY,
      name varchar DEFAULT NULL,
      is_active boolean DEFAULT true,
      created_at timestamp DEFAULT NULL,
      updated_at timestamp DEFAULT NULL,
      deleted_at timestamp DEFAULT NULL
);

CREATE TABLE users (
      id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
      username varchar DEFAULT NULL,
      password varchar DEFAULT NULL,
      token varchar DEFAULT NULL,
      is_active boolean DEFAULT true,
      role_id integer,
      created_at timestamp DEFAULT NULL,
      updated_at timestamp DEFAULT NULL,
      deleted_at timestamp DEFAULT NULL
);

CREATE TABLE user_histories (
      user_id uuid,
      new_data text DEFAULT NULL,
      old_data text DEFAULT NULL,
      created_at timestamp DEFAULT NULL,
      updated_at timestamp DEFAULT NULL
);

ALTER TABLE users ADD CONSTRAINT fk_role_user FOREIGN KEY (role_id) REFERENCES roles(id);
ALTER TABLE user_histories ADD CONSTRAINT fk_user_histories1 FOREIGN KEY (user_id) REFERENCES users(id);

-- +goose Down

DROP TABLE roles;
DROP TABLE users;
DROP TABLE user_histories;