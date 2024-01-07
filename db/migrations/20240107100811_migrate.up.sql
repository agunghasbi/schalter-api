CREATE TYPE "sale_status" AS ENUM (
  'pending',
  'paid',
  'canceled'
);

CREATE TYPE "gender" AS ENUM (
  'male',
  'female'
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "organizer_id" bigint,
  "name" varchar,
  "date_start" date,
  "date_end" date,
  "time_start" time,
  "time_end" time,
  "location_name" varchar,
  "location_maps" varchar,
  "description" text,
  "banner" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "organizers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "event_id" bigint,
  "name" varchar,
  "price" bigint,
  "amount" bigint,
  "description" text,
  "sale_start" timestamp,
  "sale_end" timestamp,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "tikets_sale" (
  "id" bigserial PRIMARY KEY,
  "ticket_id" bigint,
  "user_id" bigint,
  "amount" bigint,
  "status" sale_status,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "phone_number" varchar,
  "birth_date" date,
  "gender" gender,
  "email" varchar,
  "password" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "events" ("name");

CREATE INDEX ON "organizers" ("name");

CREATE INDEX ON "tickets" ("name");

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

ALTER TABLE "events" ADD FOREIGN KEY ("organizer_id") REFERENCES "organizers" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "tikets_sale" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");

ALTER TABLE "tikets_sale" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
