ALTER TABLE "public"."posts"
    ADD CONSTRAINT fk_user_id
        FOREIGN KEY("user_id")
            REFERENCES "public"."users"(id)
            ON UPDATE CASCADE
            ON DELETE RESTRICT;

ALTER TABLE "public"."comments"
    ADD CONSTRAINT fk_post_id
        FOREIGN KEY("post_id")
            REFERENCES "public"."posts"(id)
            ON UPDATE CASCADE
            ON DELETE RESTRICT;
