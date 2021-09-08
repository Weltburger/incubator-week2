CREATE TABLE ticket (
    "id"                int8 NOT NULL PRIMARY KEY,
    "passenger_id"      int8 NOT NULL,
    "flight_id"         int8 NOT NULL,
    "class_id"          int8 NOT NULL,
    "lunch_id"          int8 NOT NULL,
    "price"             numeric(15,6) not null default 0::numeric
);

CREATE TABLE passenger (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL,
    "age"               int2 NOT NULL,
    "email"             VARCHAR (100) NOT NULL,
    "phone_number"      VARCHAR (18)
);

CREATE TABLE flight (
    "id"                 int8 NOT NULL PRIMARY KEY,
    "from"               int8 NOT NULL,
    "to"                 int8 NOT NULL,
        "date"             date not null
);

CREATE TABLE country (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE city (
    "id"                int8 NOT NULL PRIMARY KEY,
        "country_id"        int8 NOT NULL,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE "class" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE lunch (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

ALTER TABLE ticket
    ADD CONSTRAINT fk_ticket_passenger
        FOREIGN KEY ("passenger_id")
            REFERENCES "passenger" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE ticket
    ADD CONSTRAINT fk_ticket_lunch
        FOREIGN KEY ("lunch_id")
            REFERENCES "lunch" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE ticket
    ADD CONSTRAINT fk_ticket_class
        FOREIGN KEY ("class_id")
            REFERENCES "class" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE ticket
    ADD CONSTRAINT fk_ticket_flight
        FOREIGN KEY ("flight_id")
            REFERENCES "flight" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE city
    ADD CONSTRAINT fk_city_country
        FOREIGN KEY ("country_id")
            REFERENCES "country" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE flight
    ADD CONSTRAINT fk_flight_city_from
        FOREIGN KEY ("from")
            REFERENCES "city" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE flight
    ADD CONSTRAINT fk_flight_city_to
        FOREIGN KEY ("to")
            REFERENCES "city" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
