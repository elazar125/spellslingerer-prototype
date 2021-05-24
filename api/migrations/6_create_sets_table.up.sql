CREATE TABLE IF NOT EXISTS public.sets
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT sets_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.sets
    OWNER to postgres;
