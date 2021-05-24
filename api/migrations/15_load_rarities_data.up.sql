INSERT INTO public.rarities(name)
VALUES
('Core Set'),
('Signature'),
('Common'),
('Rare'),
('Epic'),
('Mythic')
ON CONFLICT (name) DO NOTHING;
