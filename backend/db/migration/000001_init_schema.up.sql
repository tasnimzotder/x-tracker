CREATE TABLE "devices" (
  "id" bigserial PRIMARY KEY,
  "device_key" uuid NOT NULL,
  "device_name" varchar(255) NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "status" varchar(255) NOT NULL DEFAULT 'offline',
  "last_updated_at" timestamptz NOT NULL DEFAULT now(),
  "user_group" varchar(255) NOT NULL DEFAULT 'default',
  "device_version" varchar(255) NOT NULL DEFAULT '0.0.1'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) NOT NULL UNIQUE,
  "hashed_password" varchar(255) NOT NULL UNIQUE,
  "email" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "last_updated_at" timestamptz NOT NULL DEFAULT now(),
  "phone_number" bigint,
  "country_code" int DEFAULT 91,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "postal_code" bigint
);

CREATE TABLE "device_access" (
  "id" bigserial PRIMARY KEY,
  "device_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "permission" varchar(255) NOT NULL DEFAULT 'view',
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "last_updated" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "device_activities" (
    "id" bigserial PRIMARY KEY,
    "device_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "panic" boolean NOT NULL DEFAULT false,
    "fall" boolean NOT NULL DEFAULT false
);

CREATE INDEX ON "device_access" ("device_id");
CREATE INDEX ON "device_access" ("user_id");

ALTER TABLE "device_access" ADD FOREIGN KEY ("device_id") REFERENCES "devices" ("id");
ALTER TABLE "device_access" ADD  FOREIGN KEY ("user_id") REFERENCES "users" ("id");
