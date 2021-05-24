INSERT INTO public.card_types(name)
VALUES
('Creature'),
('Spell'),
('Trap'),
('Artifact'),
('Land')
ON CONFLICT (name) DO NOTHING;
