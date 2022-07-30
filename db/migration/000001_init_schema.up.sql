CREATE TYPE "role_enum" AS ENUM (
  'ADMIN',
  'MOD',
  'USER'
);

CREATE TYPE "car_type_enum" AS ENUM (
  'Hatchback',
  'Sedan',
  'SUV'
);

CREATE TYPE "car_seats_enum" AS ENUM (
  '5',
  '7'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone_number" int,
  "identical_id" int NOT NULL,
  "address" varchar,
  "balance" bigint NOT NULL,
  "account_number" varchar NOT NULL,
  "car_rental_id" int NOT NULL,
  "car_owner_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "user_role" (
  "user_id" bigserial NOT NULL,
  "role_id" bigserial NOT NULL
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "role" role_enum DEFAULT 'USER'
);

CREATE TABLE "car" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "pricing_id" int NOT NULL,
  "type_id" int NOT NULL,
  "brand" varchar NOT NULL,
  "image" varchar,
  "time_available_from" timestamptz,
  "time_available_to" timestamptz,
  "location_id" int NOT NULL,
  "created at" varchar NOT NULL
);

CREATE TABLE "car_type" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "type" car_type_enum NOT NULL DEFAULT 'Sedan',
  "seat" car_seats_enum NOT NULL DEFAULT '5'
);

CREATE TABLE "car_pricing" (
  "id" bigserial PRIMARY KEY,
  "payable_amount" bigint NOT NULL,
  "revenue" bigint NOT NULL,
  "created at" varchar NOT NULL
);

CREATE TABLE "car_location" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "lat" bigint,
  "long" bigint
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "car" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("car_rental_id") REFERENCES "car" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("car_owner_id") REFERENCES "car" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "car" ADD FOREIGN KEY ("pricing_id") REFERENCES "car_pricing" ("id");

ALTER TABLE "car" ADD FOREIGN KEY ("type_id") REFERENCES "car_type" ("id");

ALTER TABLE "car" ADD FOREIGN KEY ("location_id") REFERENCES "car_location" ("id");
