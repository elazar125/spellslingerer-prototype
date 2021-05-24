CREATE TABLE IF NOT EXISTS public.sub_types
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT sub_types_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.sub_types
    OWNER to postgres;
