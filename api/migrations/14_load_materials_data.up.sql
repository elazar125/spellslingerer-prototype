INSERT INTO public.materials(name)
VALUES
('Mox Pearl Shard'),
('Mox Sapphire Shard'),
('Mox Jet Shard'),
('Mox Ruby Shard'),
('Mox Emerald Shard'),
('Mox Diamond Shard'),
('White Anima'),
('Blue Anima'),
('Black Anima'),
('Red Anima'),
('Green Anima'),
('Prismatic Anima'),
('White Essence'),
('Blue Essence'),
('Black Essence'),
('Red Essence'),
('Green Essence'),
('Prismatic Essence'),
('Chromatic Gilded Lotus')
ON CONFLICT (name) DO NOTHING;
