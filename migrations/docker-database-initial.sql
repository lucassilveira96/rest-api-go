CREATE SEQUENCE product_categories_id_seq;

-- Table: public.product_categories

-- DROP TABLE IF EXISTS public.product_categories;

CREATE TABLE IF NOT EXISTS public.product_categories
(
    id bigint NOT NULL DEFAULT nextval('product_categories_id_seq'::regclass),
    description text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT product_categories_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.product_categories
    OWNER to root;

ALTER SEQUENCE product_categories_id_seq
OWNED BY product_categories.id;
-- Table: public.users

-- DROP TABLE IF EXISTS public.users;
CREATE SEQUENCE users_id_seq;
CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username text COLLATE pg_catalog."default",
    age bigint,
    email text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to root;

ALTER SEQUENCE users_id_seq
OWNED BY users.id;
-- Table: public.products

-- DROP TABLE IF EXISTS public.products;
CREATE SEQUENCE products_id_seq;
CREATE TABLE IF NOT EXISTS public.products
(
    id bigint NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    description text COLLATE pg_catalog."default",
    price numeric,
    product_categories_id bigint,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT fk_products_product_category FOREIGN KEY (product_categories_id)
        REFERENCES public.product_categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.products
    OWNER to root;

ALTER SEQUENCE products_id_seq
OWNED BY products.id;
-- Index: idx_products_deleted_at

-- DROP INDEX IF EXISTS public.idx_products_deleted_at;

CREATE INDEX IF NOT EXISTS idx_products_deleted_at
    ON public.products USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;