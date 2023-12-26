CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE user_role AS ENUM ('user', 'admin', 'superAdmin');
CREATE TABLE IF NOT EXISTS public.user (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar NOT NULL UNIQUE,
    "password" varchar NOT NULL,
    is_login boolean NOT NULL DEFAULT false,
    role user_role
);
INSERT INTO public."user"(username, password, role)
VALUES ('admin', 'admin', 'admin'),
    ('SuperAdmin', 'SuperAdmin', 'superAdmin'),
    ('user', 'user', 'user');