CREATE TABLE IF NOT EXISTS public.rarities
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT rarities_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.rarities
    OWNER to postgres;
