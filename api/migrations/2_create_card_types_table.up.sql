CREATE TABLE IF NOT EXISTS public.card_types
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT card_types_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.card_types
    OWNER to postgres;
