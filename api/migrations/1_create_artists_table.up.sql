CREATE TABLE IF NOT EXISTS public.artists
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    CONSTRAINT artists_name_key UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE public.artists
    OWNER to postgres;
