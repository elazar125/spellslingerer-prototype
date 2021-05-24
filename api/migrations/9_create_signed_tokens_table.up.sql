CREATE TABLE IF NOT EXISTS public.signed_tokens
(
    token text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT signed_tokens_pkey PRIMARY KEY (token)
)

TABLESPACE pg_default;

ALTER TABLE public.signed_tokens
    OWNER to postgres;
