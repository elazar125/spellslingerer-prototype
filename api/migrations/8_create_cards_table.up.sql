CREATE TABLE IF NOT EXISTS public.cards
(
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text COLLATE pg_catalog."default",
    cost bigint,
    power bigint,
    health bigint,
    ability text COLLATE pg_catalog."default",
    chance bigint,
    rarity_id bigint,
    card_type_id bigint,
    sub_type_id bigint,
    artist_id bigint,
    colour_id bigint,
    legendary boolean,
    collectible boolean,
    material_id bigint,
    mox_id bigint,
    set_id bigint,
    CONSTRAINT cards_name_key UNIQUE (name),
    CONSTRAINT fk_cards_artist FOREIGN KEY (artist_id)
        REFERENCES public.artists (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_card_type FOREIGN KEY (card_type_id)
        REFERENCES public.card_types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_colour FOREIGN KEY (colour_id)
        REFERENCES public.colours (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_material FOREIGN KEY (material_id)
        REFERENCES public.materials (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_mox FOREIGN KEY (mox_id)
        REFERENCES public.materials (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_rarity FOREIGN KEY (rarity_id)
        REFERENCES public.rarities (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_set FOREIGN KEY (set_id)
        REFERENCES public.sets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_cards_sub_type FOREIGN KEY (sub_type_id)
        REFERENCES public.sub_types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.cards
    OWNER to postgres;
