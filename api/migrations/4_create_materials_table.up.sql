CREATE TABLE IF NOT EXISTS public.materials
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT materials_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.materials
    OWNER to postgres;
