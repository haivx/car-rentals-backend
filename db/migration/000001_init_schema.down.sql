-- DROP TABLE schema_migrations;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS role CASCADE;
DROP TABLE IF EXISTS car;
DROP TABLE IF EXISTS car_type;
DROP TABLE IF EXISTS car_location;
DROP TABLE IF EXISTS car_pricing;
DROP type IF exists role_enum;
DROP type IF exists car_type_enum;
DROP type IF exists car_seats_enum;