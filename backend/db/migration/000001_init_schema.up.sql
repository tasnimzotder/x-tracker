CREATE TABLE "devices" (
  "id" bigserial PRIMARY KEY,
  "device_name" varchar(255) NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "status" varchar(255) NOT NULL DEFAULT 'offline'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) NOT NULL UNIQUE,
  "hashed_password" varchar(255) NOT NULL UNIQUE,
  "email" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
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
