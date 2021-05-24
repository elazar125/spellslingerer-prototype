CREATE TABLE IF NOT EXISTS public.colours
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT colours_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.colours
    OWNER to postgres;
