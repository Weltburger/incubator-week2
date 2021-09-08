CREATE TABLE book (
    "id"                int8 NOT NULL PRIMARY KEY,
    "type_id"           int8 NOT NULL,
        "author_id"         int8 NOT NULL,
        "publisher_id"      int8 NOT NULL,
        "cover_id"          int2 NOT NULL,
    "title"             VARCHAR (100) NOT NULL,
    "pages"             int4 NOT NULL,
    "weight"            float4 NOT NULL,
    "price"             numeric(15,6) not null default 0::numeric,
    "count"             int8 NOT NULL
);

CREATE TABLE "order" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "customer_id"       int8 NOT NULL,
    "sum"               numeric(15,6) not null default 0::numeric,
        "info"              VARCHAR (500) NOT NULL,
    "date_made"         TIMESTAMP (0) DEFAULT now() NOT NULL,
    "date_paid"         TIMESTAMP (0) DEFAULT now() NOT NULL,
    "date_done"         TIMESTAMP (0)
);

CREATE TABLE order_info (
    "order_id"          int8 NOT NULL,
    "book_id"           int8 NOT NULL,
        "count"             int8 NOT NULL
);

CREATE TABLE customer (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL,
    "email"             VARCHAR (100) NOT NULL,
    "address"           VARCHAR (100),
        "phone"             VARCHAR (18)
);

CREATE TABLE author (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE cover (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE "type" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "type"              VARCHAR (100) NOT NULL
);

CREATE TABLE publisher (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

ALTER TABLE book
    ADD CONSTRAINT fk_book_author
        FOREIGN KEY ("author_id")
            REFERENCES "author" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE book
    ADD CONSTRAINT fk_book_type
        FOREIGN KEY ("type_id")
            REFERENCES "type" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE book
    ADD CONSTRAINT fk_book_publisher
        FOREIGN KEY ("publisher_id")
            REFERENCES "publisher" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
                        
ALTER TABLE book
    ADD CONSTRAINT fk_book_cover
        FOREIGN KEY ("cover_id")
            REFERENCES "cover" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "order"
    ADD CONSTRAINT fk_order_customer
        FOREIGN KEY ("customer_id")
            REFERENCES "customer" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE order_info
    ADD CONSTRAINT fk_order_info_order_id
        FOREIGN KEY ("order_id")
            REFERENCES "order" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE order_info
    ADD CONSTRAINT fk_order_info_book_id
        FOREIGN KEY ("book_id")
            REFERENCES "book" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
