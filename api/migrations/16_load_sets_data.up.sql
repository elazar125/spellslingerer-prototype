INSERT INTO public.sets(name)
VALUES
('Core Set'),
('Chandra'),
('Gideon'),
('Jace'),
('Liliana'),
('Nissa'),
('Ajani'),
('Teferi'),
('Kaya'),
('Nahiri'),
('Ashiok'),
('Ral'),
('Kiora'),
('Angrath'),
('Vraska'),
('Domri'),
('Opening Ceremony')
ON CONFLICT (name) DO NOTHING;
