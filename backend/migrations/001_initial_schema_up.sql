-- Create makes table
CREATE TABLE makes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create models table
CREATE TABLE models (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    make_id INTEGER REFERENCES makes(id) ON DELETE CASCADE,
    year_range VARCHAR(50), 
    generation VARCHAR(100),
    location VARCHAR(50) DEFAULT 'Global',
    codename VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_models_make_name_gen_year UNIQUE(make_id, name, year_range, generation)
);

-- Create challenges table
CREATE TABLE challenges (
    id SERIAL PRIMARY KEY,
    date DATE UNIQUE NOT NULL,
    image_url TEXT NOT NULL,
    solution_make_id INTEGER REFERENCES makes(id),
    solution_model_id INTEGER REFERENCES models(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    anonymous_id UUID UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create submissions table
CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    challenge_id INTEGER REFERENCES challenges(id) ON DELETE CASCADE,
    make_id INTEGER,
    model_id INTEGER,
    is_correct BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_models_make_id ON models(make_id);
CREATE INDEX idx_challenges_date ON challenges(date);
CREATE INDEX idx_submissions_user_challenge ON submissions(user_id, challenge_id);

-- Insert makes from new data
INSERT INTO makes (name) VALUES
('Opel'),
('Volkswagen'),
('Audi'),
('BMW'),
('Mercedes-Benz'),
('Volvo'),
('Renault'),
('Peugeot'),
('Fiat'),
('Alfa Romeo'),
('Porsche'),
('Ferrari'),
('Lamborghini'),
('Aston Martin'),
('Bentley'),
('Rolls-Royce'),
('Maserati'),
('Bugatti'),
('Toyota'),
('Lexus'),
('Honda'),
('Acura'),
('Nissan'),
('Infiniti'),
('Mazda'),
('Subaru'),
('Mitsubishi'),
('Suzuki'),
('Hyundai'),
('Kia'),
('Genesis'),
('Ford'),
('Chevrolet'),
('GMC'),
('Cadillac'),
('Dodge'),
('Tesla'),
('Rivian'),
('Lucid');

-- Insert models from new data
INSERT INTO models (name, make_id, year_range, generation, location) VALUES
-- Opel models
('Astra', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', 'L', 'Global'),
('Corsa', (SELECT id FROM makes WHERE name = 'Opel'), '2020-2024', 'F', 'Global'),
('Grandland', (SELECT id FROM makes WHERE name = 'Opel'), '2018-2024', '1st Gen', 'Global'),
('Mokka', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', '2nd Gen', 'Global'),
('Insignia', (SELECT id FROM makes WHERE name = 'Opel'), '2017-2023', 'B', 'Global'),

-- Volkswagen models
('Golf', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2020-2024', 'Mk8', 'Global'),
('Passat', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2019-2023', 'B8', 'Global'),
('Tiguan', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', '2nd Gen', 'Global'),
('T-Roc', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', '1st Gen', 'Global'),
('Touareg', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', '3rd Gen', 'Global'),

-- Audi models
('A3', (SELECT id FROM makes WHERE name = 'Audi'), '2020-2024', '8Y', 'Global'),
('A4', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', 'B9', 'Global'),
('A6', (SELECT id FROM makes WHERE name = 'Audi'), '2019-2024', 'C8', 'Global'),
('Q5', (SELECT id FROM makes WHERE name = 'Audi'), '2018-2024', '2nd Gen', 'Global'),
('Q7', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', '2nd Gen', 'Global'),

-- BMW models
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G20', 'Global'),
('5 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2017-2023', 'G30', 'Global'),
('X3', (SELECT id FROM makes WHERE name = 'BMW'), '2018-2024', 'G01', 'Global'),
('X5', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G05', 'Global'),
('M3', (SELECT id FROM makes WHERE name = 'BMW'), '2021-present', 'G80', 'Global'),
('M4', (SELECT id FROM makes WHERE name = 'BMW'), '2021-2024', 'G82', 'Global'),

-- Mercedes-Benz models
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2007-2014', 'W204', 'Global'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W206', 'Global'),
('E-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2017-2023', 'W213', 'Global'),
('S-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W223', 'Global'),
('GLA', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2020-2024', 'H247', 'Global'),
('G-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2018-2024', 'W463', 'Global'),

-- Volvo models
('S60', (SELECT id FROM makes WHERE name = 'Volvo'), '2019-2024', '3rd Gen', 'Global'),
('S90', (SELECT id FROM makes WHERE name = 'Volvo'), '2017-2024', '2nd Gen', 'Global'),
('XC40', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', '1st Gen', 'Global'),
('XC60', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', '2nd Gen', 'Global'),
('XC90', (SELECT id FROM makes WHERE name = 'Volvo'), '2016-2024', '2nd Gen', 'Global'),

-- Renault models
('Clio', (SELECT id FROM makes WHERE name = 'Renault'), '2019-2024', '5th Gen', 'Global'),
('Megane', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2023', '4th Gen', 'Global'),
('Captur', (SELECT id FROM makes WHERE name = 'Renault'), '2020-2024', '2nd Gen', 'Global'),
('Kadjar', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2022', '1st Gen', 'Global'),
('Duster', (SELECT id FROM makes WHERE name = 'Renault'), '2018-2024', '2nd Gen', 'Global'),

-- Peugeot models
('208', (SELECT id FROM makes WHERE name = 'Peugeot'), '2019-2024', '2nd Gen', 'Global'),
('308', (SELECT id FROM makes WHERE name = 'Peugeot'), '2021-2024', '3rd Gen', 'Global'),
('508', (SELECT id FROM makes WHERE name = 'Peugeot'), '2018-2023', '2nd Gen', 'Global'),
('2008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2020-2024', '2nd Gen', 'Global'),
('3008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2017-2024', '2nd Gen', 'Global'),

-- Fiat models
('500', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', '3rd Gen', 'Global'),
('500X', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2023', '1st Gen', 'Global'),
('Panda', (SELECT id FROM makes WHERE name = 'Fiat'), '2012-2024', '3rd Gen', 'Global'),
('Tipo', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', '2nd Gen', 'Global'),

-- Alfa Romeo models
('Giulia', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2016-2024', '952', 'Global'),
('Stelvio', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2017-2024', '949', 'Global'),
('Giulietta', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2014-2021', '940', 'Global'),

-- Porsche models
('911', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', '992', 'Global'),
('718 Cayman', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global'),
('718 Boxster', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global'),
('Panamera', (SELECT id FROM makes WHERE name = 'Porsche'), '2017-2024', '971', 'Global'),
('Cayenne', (SELECT id FROM makes WHERE name = 'Porsche'), '2018-2024', '3rd Gen', 'Global'),
('Macan', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', '1st Gen', 'Global'),
('Taycan', (SELECT id FROM makes WHERE name = 'Porsche'), '2020-2024', '1st Gen', 'Global'),

-- Ferrari models
('488 GTB', (SELECT id FROM makes WHERE name = 'Ferrari'), '2015-2019', 'F142', 'Global'),
('F8 Tributo', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2023', 'F142M', 'Global'),
('Roma', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F169', 'Global'),
('SF90 Stradale', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F173', 'Global'),
('812 Superfast', (SELECT id FROM makes WHERE name = 'Ferrari'), '2017-2024', 'F152M', 'Global'),
('Portofino', (SELECT id FROM makes WHERE name = 'Ferrari'), '2018-2023', 'F164', 'Global'),

-- Lamborghini models
('Huracán', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2015-2024', '1st Gen', 'Global'),
('Aventador', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2012-2022', '1st Gen', 'Global'),
('Urus', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2018-2024', '1st Gen', 'Global'),
('Revuelto', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2024', '1st Gen', 'Global'),

-- Aston Martin models
('DB11', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2017-2023', '1st Gen', 'Global'),
('DBS Superleggera', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2019-2023', '1st Gen', 'Global'),
('Vantage', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2018-2024', '2nd Gen', 'Global'),
('DBX', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2020-2024', '1st Gen', 'Global'),

-- Bentley models
('Continental GT', (SELECT id FROM makes WHERE name = 'Bentley'), '2018-2024', '3rd Gen', 'Global'),
('Flying Spur', (SELECT id FROM makes WHERE name = 'Bentley'), '2020-2024', '3rd Gen', 'Global'),
('Bentayga', (SELECT id FROM makes WHERE name = 'Bentley'), '2016-2024', '1st Gen', 'Global'),

-- Rolls-Royce models
('Phantom', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2018-2024', 'VIII', 'Global'),
('Ghost', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2021-2024', '2nd Gen', 'Global'),
('Wraith', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2014-2023', '1st Gen', 'Global'),
('Cullinan', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2019-2024', '1st Gen', 'Global'),

-- Maserati models
('Ghibli', (SELECT id FROM makes WHERE name = 'Maserati'), '2014-2023', 'M156', 'Global'),
('Quattroporte', (SELECT id FROM makes WHERE name = 'Maserati'), '2013-2023', 'M156', 'Global'),
('Levante', (SELECT id FROM makes WHERE name = 'Maserati'), '2016-2024', '1st Gen', 'Global'),
('MC20', (SELECT id FROM makes WHERE name = 'Maserati'), '2021-2024', '1st Gen', 'Global'),

-- Bugatti models
('Veyron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2005-2015', '1st Gen', 'Global'),
('Chiron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2016-2024', '1st Gen', 'Global'),
('Mistral', (SELECT id FROM makes WHERE name = 'Bugatti'), '2024', '1st Gen', 'Global'),

-- Toyota models
('Corolla', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'E210', 'Global'),
('Camry', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'XV70', 'Global'),
('Land Cruiser', (SELECT id FROM makes WHERE name = 'Toyota'), '2021-2024', 'J300', 'Global'),
('RAV4', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'XA50', 'Global'),
('Hilux', (SELECT id FROM makes WHERE name = 'Toyota'), '2016-2024', 'AN120', 'Global'),
('Prius', (SELECT id FROM makes WHERE name = 'Toyota'), '2023-2024', '5th Gen', 'Global'),
('Supra', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'A90', 'Global'),

-- Lexus models
('IS', (SELECT id FROM makes WHERE name = 'Lexus'), '2021-2024', '3rd Gen', 'Global'),
('ES', (SELECT id FROM makes WHERE name = 'Lexus'), '2019-2024', '7th Gen', 'Global'),
('RX', (SELECT id FROM makes WHERE name = 'Lexus'), '2023-2024', '5th Gen', 'Global'),
('NX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', '2nd Gen', 'Global'),
('LX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', '3rd Gen', 'Global'),

-- Honda models
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', '11th Gen', 'Global'),
('Accord', (SELECT id FROM makes WHERE name = 'Honda'), '2018-2023', '10th Gen', 'Global'),
('CR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2023-2024', '6th Gen', 'Global'),
('HR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', '3rd Gen', 'Global'),
('City', (SELECT id FROM makes WHERE name = 'Honda'), '2020-2024', '7th Gen', 'Global'),

-- Acura models
('ILX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2022', '1st Gen', 'Global'),
('TLX', (SELECT id FROM makes WHERE name = 'Acura'), '2021-2024', '2nd Gen', 'Global'),
('RDX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2024', '3rd Gen', 'Global'),
('MDX', (SELECT id FROM makes WHERE name = 'Acura'), '2022-2024', '4th Gen', 'Global'),

-- Nissan models
('Altima', (SELECT id FROM makes WHERE name = 'Nissan'), '2019-2024', '6th Gen', 'Global'),
('Sentra', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', '8th Gen', 'Global'),
('X-Trail', (SELECT id FROM makes WHERE name = 'Nissan'), '2022-2024', 'T33', 'Global'),
('Patrol', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', 'Y62', 'Global'),
('GT-R', (SELECT id FROM makes WHERE name = 'Nissan'), '2009-2024', 'R35', 'Global'),

-- Infiniti models
('Q50', (SELECT id FROM makes WHERE name = 'Infiniti'), '2015-2024', 'V37', 'Global'),
('Q60', (SELECT id FROM makes WHERE name = 'Infiniti'), '2017-2022', 'V37', 'Global'),
('QX50', (SELECT id FROM makes WHERE name = 'Infiniti'), '2019-2024', '2nd Gen', 'Global'),
('QX60', (SELECT id FROM makes WHERE name = 'Infiniti'), '2022-2024', '2nd Gen', 'Global'),

-- Mazda models
('Mazda3', (SELECT id FROM makes WHERE name = 'Mazda'), '2019-2024', '4th Gen', 'Global'),
('Mazda6', (SELECT id FROM makes WHERE name = 'Mazda'), '2018-2023', '3rd Gen', 'Global'),
('CX-5', (SELECT id FROM makes WHERE name = 'Mazda'), '2017-2024', '2nd Gen', 'Global'),
('CX-30', (SELECT id FROM makes WHERE name = 'Mazda'), '2020-2024', '1st Gen', 'Global'),
('MX-5', (SELECT id FROM makes WHERE name = 'Mazda'), '2016-2024', 'ND', 'Global'),

-- Subaru models
('Impreza', (SELECT id FROM makes WHERE name = 'Subaru'), '2017-2023', '5th Gen', 'Global'),
('WRX', (SELECT id FROM makes WHERE name = 'Subaru'), '2022-2024', 'VB', 'Global'),
('Outback', (SELECT id FROM makes WHERE name = 'Subaru'), '2020-2024', '6th Gen', 'Global'),
('Forester', (SELECT id FROM makes WHERE name = 'Subaru'), '2019-2024', '5th Gen', 'Global'),

-- Mitsubishi models
('Outlander', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2022-2024', '4th Gen', 'Global'),
('Pajero Sport', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2016-2024', '3rd Gen', 'Global'),
('Lancer', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2008-2016', '10th Gen', 'Global'),
('Eclipse Cross', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2018-2024', '1st Gen', 'Global'),

-- Suzuki models
('Swift', (SELECT id FROM makes WHERE name = 'Suzuki'), '2017-2024', '3rd Gen', 'Global'),
('Vitara', (SELECT id FROM makes WHERE name = 'Suzuki'), '2019-2024', '2nd Gen', 'Global'),
('Jimny', (SELECT id FROM makes WHERE name = 'Suzuki'), '2018-2024', '4th Gen', 'Global'),
('Baleno', (SELECT id FROM makes WHERE name = 'Suzuki'), '2016-2024', '2nd Gen', 'Global'),

-- Hyundai models
('Elantra', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', '7th Gen', 'Global'),
('Sonata', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', '8th Gen', 'Global'),
('Tucson', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', '4th Gen', 'Global'),
('Santa Fe', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', '4th Gen', 'Global'),
('Creta', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', '2nd Gen', 'Global'),
('Venue', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', '1st Gen', 'Global'),
('IONIQ 5', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', '1st Gen', 'Global'),

-- Kia models
('Rio', (SELECT id FROM makes WHERE name = 'Kia'), '2017-2023', '4th Gen', 'Global'),
('Cerato', (SELECT id FROM makes WHERE name = 'Kia'), '2019-2024', '4th Gen', 'Global'),
('Optima', (SELECT id FROM makes WHERE name = 'Kia'), '2016-2020', '4th Gen', 'Global'),
('K5', (SELECT id FROM makes WHERE name = 'Kia'), '2020-2024', '5th Gen', 'Global'),
('Sportage', (SELECT id FROM makes WHERE name = 'Kia'), '2022-2024', '5th Gen', 'Global'),
('Seltos', (SELECT id FROM makes WHERE name = 'Kia'), '2019-2024', '1st Gen', 'Global'),
('Sorento', (SELECT id FROM makes WHERE name = 'Kia'), '2021-2024', '4th Gen', 'Global'),
('EV6', (SELECT id FROM makes WHERE name = 'Kia'), '2021-2024', '1st Gen', 'Global'),

-- Genesis models
('G70', (SELECT id FROM makes WHERE name = 'Genesis'), '2019-2024', '1st Gen', 'Global'),
('G80', (SELECT id FROM makes WHERE name = 'Genesis'), '2021-2024', '2nd Gen', 'Global'),
('G90', (SELECT id FROM makes WHERE name = 'Genesis'), '2019-2024', '1st-2nd Gen', 'Global'),
('GV70', (SELECT id FROM makes WHERE name = 'Genesis'), '2021-2024', '1st Gen', 'Global'),
('GV80', (SELECT id FROM makes WHERE name = 'Genesis'), '2020-2024', '1st Gen', 'Global'),

-- Ford models
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2005-2009', 'S197', 'Global'),
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2023', 'S550', 'Global'),
('F-150', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2024', '13th-14th Gen', 'Global'),
('Ranger', (SELECT id FROM makes WHERE name = 'Ford'), '2019-2024', 'T6', 'Global'),
('Explorer', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', '6th Gen', 'Global'),
('Escape', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', '4th Gen', 'Global'),
('Bronco', (SELECT id FROM makes WHERE name = 'Ford'), '2021-2024', '6th Gen', 'Global'),
('Mustang Mach-E', (SELECT id FROM makes WHERE name = 'Ford'), '2021-2024', '1st Gen', 'Global'),

-- Chevrolet models
('Camaro', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2023', '6th Gen', 'Global'),
('Corvette', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2020-2024', 'C8', 'Global'),
('Silverado', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2019-2024', '4th Gen', 'Global'),
('Malibu', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2024', '9th Gen', 'Global'),
('Equinox', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2018-2024', '3rd Gen', 'Global'),
('Tahoe', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2021-2024', '5th Gen', 'Global'),

-- GMC models
('Sierra', (SELECT id FROM makes WHERE name = 'GMC'), '2019-2024', '5th Gen', 'Global'),
('Yukon', (SELECT id FROM makes WHERE name = 'GMC'), '2021-2024', '5th Gen', 'Global'),
('Acadia', (SELECT id FROM makes WHERE name = 'GMC'), '2017-2024', '2nd Gen', 'Global'),
('Terrain', (SELECT id FROM makes WHERE name = 'GMC'), '2018-2024', '2nd Gen', 'Global'),

-- Cadillac models
('CT4', (SELECT id FROM makes WHERE name = 'Cadillac'), '2020-2024', '1st Gen', 'Global'),
('CT5', (SELECT id FROM makes WHERE name = 'Cadillac'), '2020-2024', '1st Gen', 'Global'),
('Escalade', (SELECT id FROM makes WHERE name = 'Cadillac'), '2021-2024', '5th Gen', 'Global'),
('XT4', (SELECT id FROM makes WHERE name = 'Cadillac'), '2019-2024', '1st Gen', 'Global'),
('Lyriq', (SELECT id FROM makes WHERE name = 'Cadillac'), '2023-2024', '1st Gen', 'Global'),

-- Dodge models
('Charger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LD', 'Global'),
('Challenger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LC', 'Global'),
('Durango', (SELECT id FROM makes WHERE name = 'Dodge'), '2016-2024', '3rd Gen', 'Global'),

-- Tesla models
('Model S', (SELECT id FROM makes WHERE name = 'Tesla'), '2012-2024', '1st Gen', 'Global'),
('Model 3', (SELECT id FROM makes WHERE name = 'Tesla'), '2017-2024', '1st Gen', 'Global'),
('Model X', (SELECT id FROM makes WHERE name = 'Tesla'), '2016-2024', '1st Gen', 'Global'),
('Model Y', (SELECT id FROM makes WHERE name = 'Tesla'), '2020-2024', '1st Gen', 'Global'),
('Cybertruck', (SELECT id FROM makes WHERE name = 'Tesla'), '2024', '1st Gen', 'Global'),

-- Rivian models
('R1T', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', '1st Gen', 'Global'),
('R1S', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', '1st Gen', 'Global'),

-- Lucid models
('Air', (SELECT id FROM makes WHERE name = 'Lucid'), '2022-2024', '1st Gen', 'Global');

-- Insert challenges data
INSERT INTO challenges (date, image_url, solution_make_id, solution_model_id) VALUES
('2026-01-25', '/images/mustang_s550.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Ford')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Ford') AND LOWER(m.name) = LOWER('Mustang') AND m.year_range = '2015-2023')),
('2026-01-26', '/images/tesla_model3.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Tesla')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Tesla') AND LOWER(m.name) = LOWER('Model 3') AND m.year_range = '2017-2023')),
('2026-01-27', '/images/benz_c_w204.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Mercedes-Benz')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Mercedes-Benz') AND LOWER(m.name) = LOWER('C-Class') AND m.year_range = '2007-2014')),
('2026-01-28', '/images/mustang_s550.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Ford')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Ford') AND LOWER(m.name) = LOWER('Mustang') AND m.year_range = '2015-2023')),
('2026-01-29', '/images/tesla_model3.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Tesla')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Tesla') AND LOWER(m.name) = LOWER('Model 3') AND m.year_range = '2017-2023')),
('2026-01-30', '/images/benz_c_w204.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Mercedes-Benz')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Mercedes-Benz') AND LOWER(m.name) = LOWER('C-Class') AND m.year_range = '2007-2014')),
('2026-01-31', '/images/bmw_m3_g80.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('BMW')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('BMW') AND LOWER(m.name) = LOWER('M3') AND m.year_range = '2021-present')),
('2026-02-01', '/images/mustang_s197.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Ford')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Ford') AND LOWER(m.name) = LOWER('Mustang') AND m.year_range = '2005-2009')),
('2026-02-02', '/images/gtr_r35.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Nissan')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Nissan') AND LOWER(m.name) = LOWER('GT-R') AND m.year_range LIKE '%2009%')),
('2026-02-03', '/images/corvette_c8.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Chevrolet')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Chevrolet') AND LOWER(m.name) = LOWER('Corvette') AND m.year_range LIKE '%2020%')),
('2026-02-04', '/images/tucson_nx4.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Hyundai')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Hyundai') AND LOWER(m.name) = LOWER('Tucson') AND m.year_range LIKE '%2021%')),
('2026-02-05', '/images/passat_b8_estate.jpg',
    (SELECT id FROM makes WHERE LOWER(name) = LOWER('Volkswagen')),
    (SELECT m.id FROM models m JOIN makes ma ON m.make_id = ma.id WHERE LOWER(ma.name) = LOWER('Volkswagen') AND LOWER(m.name) = LOWER('Passat') AND m.year_range = '2015-2019'));