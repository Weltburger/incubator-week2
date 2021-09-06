ALTER TABLE "public"."comments"
    DROP CONSTRAINT fk_post_id;

ALTER TABLE "public"."posts"
    DROP CONSTRAINT fk_user_id;
