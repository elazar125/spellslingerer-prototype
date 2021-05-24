INSERT INTO public.colours(name)
VALUES
('White'),
('Blue'),
('Black'),
('Red'),
('Green'),
('Colourless')
ON CONFLICT (name) DO NOTHING;
