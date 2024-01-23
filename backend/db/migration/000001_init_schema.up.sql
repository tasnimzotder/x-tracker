CREATE TABLE "devices" (
    "id" bigserial PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "device_key" uuid NOT NULL,
    "device_name" varchar(255) NOT NULL UNIQUE,
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "last_seen" timestamptz NOT NULL DEFAULT now(),
    "device_type" VARCHAR(255) NOT NULL DEFAULT 'unknown',
    "device_version" VARCHAR(255) NOT NULL DEFAULT 'unknown'
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "hashed_password" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "role" VARCHAR(255) NOT NULL DEFAULT 'user',
    "phone_number" BIGINT,
    "country_code" INT DEFAULT 91,
    "first_name" VARCHAR(255),
    "last_name" VARCHAR(255),
    "postal_code" VARCHAR(255)
);

CREATE TABLE "device_activities" (
    "id" bigserial PRIMARY KEY,
    "device_id" BIGINT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "activity_type" VARCHAR(255) NOT NULL,
    "activity_data" JSONB NOT NULL
);

CREATE INDEX ON "device_activities" ("device_id");

ALTER TABLE "devices" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "device_activities" ADD FOREIGN KEY ("device_id") REFERENCES "devices" ("id");
