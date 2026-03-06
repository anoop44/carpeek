-- Insert makes from new data
INSERT INTO makes (name) VALUES
('Opel'), ('Volkswagen'), ('Audi'), ('BMW'), ('Mercedes-Benz'), ('Volvo'), ('Renault'), ('Peugeot'), ('Fiat'), ('Alfa Romeo'),
('Porsche'), ('Ferrari'), ('Lamborghini'), ('Aston Martin'), ('Bentley'), ('Rolls-Royce'), ('Maserati'), ('Bugatti'),
('Toyota'), ('Lexus'), ('Honda'), ('Acura'), ('Nissan'), ('Infiniti'), ('Mazda'), ('Subaru'), ('Mitsubishi'), ('Suzuki'),
('Hyundai'), ('Kia'), ('Genesis'), ('Ford'), ('Chevrolet'), ('GMC'), ('Cadillac'), ('Dodge'), ('Tesla'), ('Rivian'), ('Lucid'), ('Jaguar'), ('Land Rover'), ('Jeep'), ('Mini');

-- Insert models from new data
INSERT INTO models (name, make_id, year_range, generation, location, image_url, known_for) VALUES
-- Opel models
('Astra', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', 'L', 'Global', 'opel/Opel_Astra_12_Turbo_Ultimate_(L)_–_f_13122024.jpg', 'Sharp styling and modern tech'),
('Corsa', (SELECT id FROM makes WHERE name = 'Opel'), '2020-2024', 'F', 'Global', 'opel/Opel_Corsa-e_at_IAA_2019_IMG_0738.jpg', 'Popular city car'),
('Grandland', (SELECT id FROM makes WHERE name = 'Opel'), '2018-2024', 'first', 'Global', 'opel/Opel_Grandland_Hybrid4_Automesse_Ludwigsburg_2022_1X7A5911.jpg', 'Compact SUV contender'),
('Mokka', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', 'second', 'Global', 'opel/Opel_Mokka-e_IMG_6111.jpg', 'Bold Vizor design'),
('Insignia', (SELECT id FROM makes WHERE name = 'Opel'), '2017-2023', 'B', 'Global', 'opel/Opel_Insignia_Sports_Tourer_15_DIT_Innovation_(B).jpg', 'Sleek executive tourer'),

-- Volkswagen models
('Golf', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2020-2024', 'Mk8', 'Global', 'volkswagen/2020_Volkswagen_Golf_MK8.jpg', 'The benchmark compact hatchback'),
('Golf', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2014-2019', 'Mk7', 'Global', 'volkswagen/2013_Volkswagen_Golf_MK7.jpg', 'The complete compact package'),
('Passat', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2019-2023', 'B8', 'Global', 'volkswagen/vw_passat_b8.jpg', 'Definitive family estate'),
('Tiguan', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'second', 'Global', 'volkswagen/VW_Tiguan_2nd.jpg', 'Best-selling refined SUV'),
('T-Roc', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'first', 'Global', 'volkswagen/VW_T-Roc_15_TSI_.jpg', 'Stylish compact crossover'),
('Touareg', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'third', 'Global', 'volkswagen/Volkswagen_Touareg.jpg', 'Premium luxury flagship SUV'),

-- Audi models
('A3', (SELECT id FROM makes WHERE name = 'Audi'), '2020-2024', '8Y', 'Global', 'audi/audi_s3_8y.jpg', 'Premium compact superiority'),
('A4', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', 'B9', 'Global', 'audi/Audi_A4_B9_sedan.jpg', 'Executive sedan staple'),
('A6', (SELECT id FROM makes WHERE name = 'Audi'), '2019-2024', 'C8', 'Global', 'audi/audi_a6_c8.jpg', 'High-tech business class'),
('Q5', (SELECT id FROM makes WHERE name = 'Audi'), '2018-2024', 'second', 'Global', 'audi/Audi_Q5_2nd.jpg', 'Best-selling luxury SUV'),
('Q7', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', 'second', 'Global', 'audi/Audi_Q7_Second.jpg', 'Seven-seat luxury family hauler'),
('TT', (SELECT id FROM makes WHERE name = 'Audi'), '2015-2023', '8S', 'Global', 'audi/2019_Audi_TT_8S.jpg', 'Design icon'),
('R8', (SELECT id FROM makes WHERE name = 'Audi'), '2015-2023', '4S', 'Global', 'audi/2018_Audi_R8_4S.jpg', 'V10 mid-engine masterpiece'),


-- BMW models
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G20', 'Global', 'bmw/2019_BMW_318d_G20.jpg', 'The ultimate sports sedan'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1982-1994', 'E30', 'Global', 'bmw/bmw_e30.jpg', 'The original M3 legend'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1998-2006', 'E46', 'Global', 'bmw/bmw_e46.jpg', 'The gold standard of sports sedans'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1990-1998', 'E36', 'Global', 'bmw/bmw_e36.jpg', 'Balanced driving perfection'),
('5 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2017-2023', 'G30', 'Global', 'bmw/2018_BMW_520d_G30.jpg', 'Dynamic business athlete'),
('5 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1995-2003', 'E39', 'Global', 'bmw/BMW_525i_(E39)_Executive_sedan.jpg', 'Often cited as the best 5 Series ever'),
('M4', (SELECT id FROM makes WHERE name = 'BMW'), '2014-2020', 'F82', 'Global', 'bmw/f82_m4.jpg', 'Swift twin-turbo coupe'),
('X1', (SELECT id FROM makes WHERE name = 'BMW'), '2009-2015', 'E84', 'Global', 'bmw/BMW_X1_xDrive18d_(E84)_.jpg', 'Compact agile utility'),
('X3', (SELECT id FROM makes WHERE name = 'BMW'), '2018-2024', 'G01', 'Global', 'bmw/bmw_x3_g01.jpeg', 'SAV versatility with sportiness'),
('X5', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G05', 'Global', 'bmw/BMW-X5-G05.png', 'The boss of luxury SUVs'),
('M3', (SELECT id FROM makes WHERE name = 'BMW'), '2021-present', 'G80', 'Global', 'bmw/bmw_m3_g80.jpg', 'Polarizing grille, undeniable performance'),
('M4', (SELECT id FROM makes WHERE name = 'BMW'), '2021-2024', 'G82', 'Global', 'bmw/bmw-m4-g82.jpg', 'Aggressive coupe dominator'),

-- Mercedes-Benz models
('AMG GT', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2017-2020', 'C190', 'Global', 'merc/Mercedes_AMG_GT.jpg', 'Track focused beast'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2007-2014', 'W204', 'Global', 'merc/Mercedes-Benz_C300_w204.jpg', 'Angular styling reliability'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W206', 'Global', 'merc/Mercede-Benz-C_W206.jpg', 'Baby S-Class technology'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2018-2021', 'W205', 'Global', 'merc/Mercedes-Benz_C_200_Avantgarde_W_205.jpg', 'Executive comfort and tech'),
('E-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2017-2023', 'W213', 'Global', 'merc/Mercedes-Benz_E220d_W213.jpg', 'Smooth executive cruiser'),
('E-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '1976-1985', 'W123', 'Global', 'merc/Mercedes-Benz_W123.jpg', 'Undying mechanical durability'),
('S-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W223', 'Global', 'merc/Mercedes-Benz_S_W223.jpg', 'The standard of the world'),
('S-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '1979-1991', 'W126', 'Global', 'merc/Mercede-Benz-S_W126.jpg', 'Engineering peak of luxury'),
('GLA', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2020-2024', 'H247', 'Global', 'merc/Mercedes-Benz_GLA_H247.jpg', 'Compact urban luxury'),
('G-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2018-2024', 'W463', 'Global', 'merc/Mercedes-Benz_W463_G.jpg', 'The iconic Geländewagen'),

-- Volvo models
('S60', (SELECT id FROM makes WHERE name = 'Volvo'), '2019-2024', 'third', 'Global', 'volvo/Volvo_S60_3rd.jpg', 'Scandinavian sports sedan'),
('S90', (SELECT id FROM makes WHERE name = 'Volvo'), '2017-2024', 'second', 'Global', 'volvo/Volvo_S90_2nd.jpg', 'Elegant luxury flagship'),
('XC40', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', 'first', 'Global', 'volvo/Volvo_XC40_1st.jpg', 'Funky compact premium SUV'),
('XC60', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', 'second', 'Global', 'volvo/2018_Volvo_XC60_2nd.jpg', 'Nordic bestseller'),
('XC90', (SELECT id FROM makes WHERE name = 'Volvo'), '2016-2024', 'second', 'Global', 'volvo/Volvo_XC90_2nd.jpg', 'Thor hammer seven-seater'),

-- Renault models
('Clio', (SELECT id FROM makes WHERE name = 'Renault'), '2019-2024', 'fifth', 'Global', 'renault/Renault_Clio_5th.jpg', 'Chic supermini icon'),
('Megane', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2023', 'fourth', 'Global', 'renault/Renault_Megane_4th.jpg', 'Distinctive lighting signature'),
('Captur', (SELECT id FROM makes WHERE name = 'Renault'), '2020-2024', 'second', 'Global', 'renault/Renault_Captur_2nd.jpg', 'Versatile urban crossover'),
('Kadjar', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2022', 'first', 'Global', 'renault/Renault_Kadjar.jpg', 'Practical family SUV'),
('Duster', (SELECT id FROM makes WHERE name = 'Renault'), '2018-2024', 'second', 'Global', 'renault/duster_2nd.jpeg', 'Robust affordable adventurer'),

-- Peugeot models
('208', (SELECT id FROM makes WHERE name = 'Peugeot'), '2019-2024', 'second', 'Global', 'peugeot/Peugeot_208.jpg', 'Sabre-tooth styling'),
('308', (SELECT id FROM makes WHERE name = 'Peugeot'), '2021-2024', 'third', 'Global', 'peugeot/2022_-_Peugeot.jpg', 'New logo pioneer'),
('508', (SELECT id FROM makes WHERE name = 'Peugeot'), '2018-2023', 'second', 'Global', 'peugeot/2019_Peugeot_508.jpg', 'Radical fastback design'),
('2008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2020-2024', 'second', 'Global', 'peugeot/Peugeot_2008.jpg', 'Sharp-edged compact SUV'),
('3008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2017-2024', 'second', 'Global', 'peugeot/Peugeot_e-3008.jpg', 'Award-winning style'),

-- Fiat models
('500', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', 'third', 'Global', 'fiat/Fiat_500_3rd.jpeg', 'Timeless retro chic'),
('500X', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2023', 'first', 'Global', 'fiat/Fiat_500X.jpg', 'Crossover with heritage'),
('Panda', (SELECT id FROM makes WHERE name = 'Fiat'), '2012-2024', 'third', 'Global', 'fiat/2013_Fiat_Panda.jpg', 'Boxy practical hero'),
('Tipo', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', 'second', 'Global', 'fiat/Fiat_tipo_f.jpg', 'Value-focused practicality'),

-- Alfa Romeo models
('Giulia', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2016-2024', '952', 'Global', 'alfa_giulia.jpg', 'Emotional sports sedan'),
('Stelvio', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2017-2024', '949', 'Global', 'alfa_stelvio.jpg', 'The driving SUV'),
('Giulietta', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2014-2021', '940', 'Global', 'alfa_giulietta.jpg', 'Curvaceous hatchback'),

-- Porsche models
('911', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', '992', 'Global', 'porsche/992_911.jpg', 'The everyday supercar icon'),
('911', (SELECT id FROM makes WHERE name = 'Porsche'), '2006-2013', '997', 'Global', 'porsche/997_911.jpg', 'The usable daily supercar'),
('718 Cayman', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global', 'porsche/Porsche_718_Cayman_982.jpg', 'Mid-engine precision'),
('718 Boxster', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global', 'porsche/Porsche_718_boxster_982.jpg', 'Open-top purity'),
('Panamera', (SELECT id FROM makes WHERE name = 'Porsche'), '2017-2024', '971', 'Global', 'porsche/Porsche_972_Panamera.jpg', 'Four-door GT performance'),
('Cayenne', (SELECT id FROM makes WHERE name = 'Porsche'), '2018-2024', 'third', 'Global', 'porsche/Porsche_Cayenne_3rd.jpg', 'The sportiest SUV'),
('Macan', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', 'first', 'Global', 'porsche/porsche_macan.jpg', 'Compact handling benchmark'),
('Taycan', (SELECT id FROM makes WHERE name = 'Porsche'), '2020-2024', 'first', 'Global', 'porsche/porsche_taycan.jpg', 'Electric soul'),

-- Ferrari models
('488 GTB', (SELECT id FROM makes WHERE name = 'Ferrari'), '2015-2019', 'F142', 'Global', 'ferrari/488_gtb.jpg', 'Twin-turbo V8 fury'),
('F8 Tributo', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2023', 'F142M', 'Global', 'ferrari/f8_tributo.jpg', 'Evolution of excellence'),
('Roma', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F169', 'Global', 'ferrari/roma.jpg', 'La Nuova Dolce Vita'),
('SF90 Stradale', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F173', 'Global', 'ferrari/sf90_stradale.jpg', 'Hybrid hypercar performance'),
('812 Superfast', (SELECT id FROM makes WHERE name = 'Ferrari'), '2017-2024', 'F152M', 'Global', 'ferrari/812_Superfast.jpg', 'V12 grand tourer majesty'),
('Portofino', (SELECT id FROM makes WHERE name = 'Ferrari'), '2018-2023', 'F164', 'Global', 'ferrari/portofino.jpg', 'Convertible elegance'),
('458 Italia', (SELECT id FROM makes WHERE name = 'Ferrari'), '2010-2015', '458', 'Global', 'ferrari/458_italia.jpg', 'The last N/A V8'),
('F355', (SELECT id FROM makes WHERE name = 'Ferrari'), '1994-1999', 'F129', 'Global', 'ferrari/f355_berlinetta.jpg', 'Classic V8 mid-engine beauty'),


-- Lamborghini models
('Huracan', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2014-2024', 'LP610-4', 'Global', 'lamborghini/Lamborghini_Huracan_LP610.jpg', 'V10 mid-engine thrill'),
('Huracan', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2019-2023', 'EVO', 'Global', 'lamborghini/huracan_evoe.jpg', 'Enhanced aero and tech'),
('Aventador', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2012-2021', 'LP700-4', 'Global', 'lamborghini/aventador.jpg', 'The ultimate V12 poster car'),
('Murcielago', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2001-2010', 'LP640', 'Global', 'lamborghini/Lamborghini_Murcielago_LP-640.jpg', 'Dramatic scissor-door icon'),
('Urus', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2018-2024', 'first', 'Global', 'lamborghini/Lamborghini_Urus.jpg', 'Super SUV dominator'),
('Revuelto', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2024', 'first', 'Global', 'lamborghini/Lamborghini_Revuelto.jpg', 'The first HPEV hybrid'),

-- Aston Martin models
('DB11', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2017-2023', 'first', 'Global', 'aston_martin/DB11_V8.jpg', 'Grand touring sculpture'),
('DBS Superleggera', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2019-2023', 'first', 'Global', 'aston_martin/DBS_Superleggera.jpg', 'Brute in a suit'),
('Vantage', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2018-2024', 'second', 'Global', 'aston_martin/Vantage_V8.jpg', 'Hunter instinct'),
('DBX', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2020-2024', 'first', 'Global', 'aston_martin/DBX.jpg', 'Luxury capability'),

-- Bentley models
('Continental GT', (SELECT id FROM makes WHERE name = 'Bentley'), '2018-2024', 'third', 'Global', 'bentley/continental_GT_First_Edition.jpg', 'Deep-chested touring capability'),
('Flying Spur', (SELECT id FROM makes WHERE name = 'Bentley'), '2020-2024', 'third', 'Global', 'bentley/flying_spur_W12.jpg', 'Four-door super-luxury'),
('Bentayga', (SELECT id FROM makes WHERE name = 'Bentley'), '2016-2024', 'first', 'Global', 'bentley/bentayga_V8.jpg', 'Pinnacle SUV'),

-- Rolls-Royce models
('Phantom', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2018-2024', 'VIII', 'Global', 'rolls_royce/phantom_VIII.jpg', 'The best car in the world'),
('Ghost', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2021-2024', 'second', 'Global', 'rolls_royce/ghost.jpg', 'Post-opulence luxury'),
('Wraith', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2014-2023', 'first', 'Global', 'rolls_royce/wraith_V12.jpg', 'Ultimate gentleman GT'),
('Cullinan', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2019-2024', 'first', 'Global', 'rolls_royce/cullinan_V12.jpg', 'Effortless everywhere'),

-- Maserati models
('Ghibli', (SELECT id FROM makes WHERE name = 'Maserati'), '2014-2023', 'M156', 'Global', 'maserati/Maserati_Ghibli_M157.jpg', 'Italian sports sedan flare'),
('Quattroporte', (SELECT id FROM makes WHERE name = 'Maserati'), '2013-2023', 'M156', 'Global', 'maserati/Maserati_Quattroporte.jpg', 'The original race-bred saloon'),
('Levante', (SELECT id FROM makes WHERE name = 'Maserati'), '2016-2024', 'first', 'Global', 'maserati/Maserati_Levante_S_(01).jpg', 'The Maserati of SUVs'),
('MC20', (SELECT id FROM makes WHERE name = 'Maserati'), '2021-2024', 'first', 'Global', 'maserati/Maserati_MC20.jpg', 'Super sports car return'),
('GranTurismo', (SELECT id FROM makes WHERE name = 'Maserati'), '2007-2019', 'M145', 'Global', 'maserati/Maserati_GranTurismo.jpg', 'Italian vocal masterpiece'),


-- Bugatti models
('Veyron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2005-2015', 'first', 'Global', 'bugatti_veyron.jpg', 'The concord moment of cars'),
('Chiron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2016-2024', 'first', 'Global', 'bugatti_chiron.jpg', 'Breaking physics'),
('Mistral', (SELECT id FROM makes WHERE name = 'Bugatti'), '2024', 'first', 'Global', 'bugatti_mistral.jpg', 'The last W16 roadster'),

-- Toyota models
('Corolla', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'E210', 'Global', 'toyota/corolla_e210.jpg', 'Best selling car globally'),
('Camry', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'XV70', 'Global', 'toyota/Toyota_Camry_XV70.jpg', 'Reliable mid-size standard'),
('Land Cruiser', (SELECT id FROM makes WHERE name = 'Toyota'), '2021-2024', 'J300', 'Global', 'toyota/tlc_300.png', 'Unstoppable off-road legend'),
('Land Cruiser', (SELECT id FROM makes WHERE name = 'Toyota'), '2015-2021', 'J200', 'Global', 'toyota/tlc_200.jpg', 'Rugged luxury cruiser'),
('Tacoma', (SELECT id FROM makes WHERE name = 'Toyota'), '2016-present', 'N300', 'Global', 'toyota/Tacoma_N300.jpg', 'Reliable mid-size adventurer'),
('RAV4', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'XA50', 'Global', 'toyota/Toyota_RAV4_.jpg', 'The defining modern crossover'),
('Hilux', (SELECT id FROM makes WHERE name = 'Toyota'), '2016-2024', 'AN120', 'Global', 'toyota/Toyota_HiLux.jpg', 'Indestructible workhorse'),
('Prius', (SELECT id FROM makes WHERE name = 'Toyota'), '2023-2024', 'fifth', 'Global', 'toyota/Toyota_Prius_5th.jpg', 'From boring to beautiful'),
('Supra', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'A90', 'Global', 'toyota/supra_a90.jpg', 'German heart Japanese soul'),
('Supra', (SELECT id FROM makes WHERE name = 'Toyota'), '1993-2002', 'A80', 'Global', 'toyota/supra_a80.jpg', 'The 10-second car'),
('MR2', (SELECT id FROM makes WHERE name = 'Toyota'), '1990-1999', 'SW20', 'Global', 'toyota/ToyotaMR2.jpg', 'Mid-ship runabout'),
('86', (SELECT id FROM makes WHERE name = 'Toyota'), '2012-2020', 'ZN6', 'Global', 'toyota/Toyota_GR86.jpg', 'Lightweight handling purist'),

-- Lexus models
('IS', (SELECT id FROM makes WHERE name = 'Lexus'), '2021-2024', 'third', 'Global', 'lexus/Lexus_IS300.jpg', 'Sharp compact executive'),
('ES', (SELECT id FROM makes WHERE name = 'Lexus'), '2019-2024', 'seventh', 'Global', 'lexus/Lexus_ES_350(GSZ10).jpg', 'Serene comfort cruiser'),
('RX', (SELECT id FROM makes WHERE name = 'Lexus'), '2023-2024', 'fifth', 'Global', 'lexus/Lexus_RX_500h_F_SPORT.jpg', 'Pioneer of luxury SUVs'),
('NX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', 'second', 'Global', 'lexus/Lexus_NX_450h.jpg', 'Angular urban luxury'),
('LX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', 'third', 'Global', 'lexus/Lexus_LX_570.jpg', 'Land Cruiser in a tuxedo'),

-- Honda models
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', 'eleventh', 'Global', 'honda/Honda_Civic_11th.jpg', 'Mature sporty compact'),
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2006-2011', 'FD', 'Global', 'honda/honda_civic_fd.jpg', 'VTEC legend'),
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2016-2021', 'FC / FK', 'Global', 'honda/Honda_Civic_FC.jpg', 'Sophisticated global compact'),
('Accord', (SELECT id FROM makes WHERE name = 'Honda'), '2018-2023', 'tenth', 'Global', 'honda/Honda_Accord_10th.jpg', 'Sleek family sedan'),
('CR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2023-2024', 'sixth', 'Global', 'honda/Honda_CR-V_6th.jpg', 'Practical family hauler'),
('HR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', 'third', 'Global', 'honda/2023_Honda_HR-V.jpg', 'Stylish subcompact utility'),
('City', (SELECT id FROM makes WHERE name = 'Honda'), '2020-2024', 'seventh', 'Global', 'honda/honda_city_7th.jpg', 'Global compact sedan'),

-- Acura models
('ILX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2022', 'first', 'Global', 'acura/2019_Acura_ILX.jpg', 'Entry luxury sedan'),
('TLX', (SELECT id FROM makes WHERE name = 'Acura'), '2021-2024', 'second', 'Global', 'acura/2016_Acura_TLX_V6.jpg', 'Precision crafted performance'),
('RDX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2024', 'third', 'Global', 'acura/2019_Acura_RDX.jpg', 'Sharp-handling SUV'),
('MDX', (SELECT id FROM makes WHERE name = 'Acura'), '2022-2024', 'fourth', 'Global', 'acura/2022_Acura_MDX_.jpg', 'Flagship three-row SUV'),

-- Nissan models
('Altima', (SELECT id FROM makes WHERE name = 'Nissan'), '2019-2024', 'sixth', 'Global', 'nissan/2024_Nissan_Altima_SR.jpg', 'Efficient mid-size sedan'),
('Sentra', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', 'eighth', 'Global', 'nissan/2021_Nissan_Sentra_SR.jpg', 'Sharp compact value'),
('X-Trail', (SELECT id FROM makes WHERE name = 'Nissan'), '2022-2024', 'T33', 'Global', 'nissan/Nissan_X-Trail_(T33).jpg', 'Family adventure ready'),
('Patrol', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', 'Y62', 'Global', 'nissan/NISSAN_PATROL_Y62.jpg', 'Desert conquering giant'),
('GT-R', (SELECT id FROM makes WHERE name = 'Nissan'), '2009-2024', 'R35', 'Global', 'nissan/r35.jpg', 'The supercar killer'),
('GT-R', (SELECT id FROM makes WHERE name = 'Nissan'), '1999-2002', 'R34', 'Global', 'nissan/r34.jpg', 'The legendary Skyline'),
('370Z', (SELECT id FROM makes WHERE name = 'Nissan'), '2009-2020', 'Z34', 'Global', 'nissan/z34_370z.jpg', 'V6 sports coupe'),
('350Z', (SELECT id FROM makes WHERE name = 'Nissan'), '2002-2009', 'Z33', 'Global', 'nissan/z33_350Z.jpg', 'V6 sports coupe icon'),


-- Infiniti models
('Q50', (SELECT id FROM makes WHERE name = 'Infiniti'), '2015-2024', 'V37', 'Global', 'infiniti_q50.jpg', 'Twin-turbo sports sedan'),
('Q60', (SELECT id FROM makes WHERE name = 'Infiniti'), '2017-2022', 'V37', 'Global', 'infiniti_q60.jpg', 'Stunning luxury coupe'),
('QX50', (SELECT id FROM makes WHERE name = 'Infiniti'), '2019-2024', 'second', 'Global', 'infiniti_qx50.jpg', 'Variable compression innovation'),
('QX60', (SELECT id FROM makes WHERE name = 'Infiniti'), '2022-2024', 'second', 'Global', 'infiniti_qx60.jpg', 'Serene family luxury'),

-- Mazda models
('Mazda3', (SELECT id FROM makes WHERE name = 'Mazda'), '2019-2024', 'fourth', 'Global', 'mazda3_bp.jpg', 'Artistic compact design'),
('Mazda6', (SELECT id FROM makes WHERE name = 'Mazda'), '2018-2023', 'third', 'Global', 'mazda6_gj.jpg', 'Kodo design elegance'),
('CX-5', (SELECT id FROM makes WHERE name = 'Mazda'), '2017-2024', 'second', 'Global', 'mazda_cx5.jpg', 'Driver-focused crossover'),
('CX-30', (SELECT id FROM makes WHERE name = 'Mazda'), '2020-2024', 'first', 'Global', 'mazda_cx30.jpg', 'Rugged sweeping lines'),
('MX-5', (SELECT id FROM makes WHERE name = 'Mazda'), '2016-2024', 'ND', 'Global', 'mazda_mx5_nd.jpg', 'Jinba Ittai roadster'),

-- Subaru models
('Impreza', (SELECT id FROM makes WHERE name = 'Subaru'), '2017-2023', 'fifth', 'Global', 'subaru_impreza.jpg', 'Symmetrical AWD compact'),
('WRX', (SELECT id FROM makes WHERE name = 'Subaru'), '2022-2024', 'VB', 'Global', 'subaru_wrx_vb.jpg', 'Rally bred sedan'),
('Outback', (SELECT id FROM makes WHERE name = 'Subaru'), '2020-2024', 'sixth', 'Global', 'subaru_outback.jpg', 'The original crossover wagon'),
('Forester', (SELECT id FROM makes WHERE name = 'Subaru'), '2019-2024', 'fifth', 'Global', 'subaru_forester.jpg', 'Boxy visibility champion'),

-- Mitsubishi models
('Outlander', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2022-2024', 'fourth', 'Global', 'mitsubishi/2025_Mitsubishi_Outlander.jpg', 'Bold dynamic shield'),
('Pajero Sport', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2016-2024', 'third', 'Global', 'mitsubishi/Mitsubishi_Pajero_Sport.jpg', 'Robust off-road SUV'),
('Lancer', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2008-2016', 'tenth', 'Global', 'mitsubishi/2012_Mitsubishi_Lancer.jpg', 'Evo heritage sedan'),
('Eclipse Cross', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2018-2024', 'first', 'Global', 'mitsubishi/2023_Mitsubishi_Eclipse_Cross_SEL_Touring.jpg', 'Coupe-style SUV'),
('Lancer Evolution', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2003-2006', 'Evo 8', 'Global', 'mitsubishi/2008_Mitsubishi_Lancer_Evolution_X.jpg', 'Rally legend for the road'),


-- Suzuki models
('Swift', (SELECT id FROM makes WHERE name = 'Suzuki'), '2017-2024', 'third', 'Global', 'suzuki/Suzuki_Swift_(2024)_hybrid.jpg', 'Lightweight fun hatch'),
('Vitara', (SELECT id FROM makes WHERE name = 'Suzuki'), '2019-2024', 'second', 'Global', 'suzuki/2024_Suzuki_Vitara_(4th_generation).jpg', 'Compact capable SUV'),
('Jimny', (SELECT id FROM makes WHERE name = 'Suzuki'), '2018-2024', 'fourth', 'Global', 'suzuki/2019_Suzuki_Jimny_SZ5_4X4.jpg', 'Tiny off-road giant'),
('Baleno', (SELECT id FROM makes WHERE name = 'Suzuki'), '2016-2024', 'second', 'Global', 'suzuki/2017_Suzuki_Baleno_SZ3_Dualjet.jpg', 'Practical spacious hatch'),

-- Hyundai models
('Elantra', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'seventh', 'Global', 'hyundai/Hyundai_Avante_CN7_white_(10).jpg', 'Parametric dynamics design'),
('Sonata', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', 'eighth', 'Global', 'hyundai/2024_Hyundai_Sonata_SEL.jpg', 'Sensuous sportiness'),
('Tucson', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'fourth', 'Global', 'hyundai/2022_Hyundai_Tucson.jpg', 'Revolutionary lighting face'),
('Santa Fe', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', 'fourth', 'Global', 'hyundai/2024_Hyundai_Santa_Fe_Luxury_AWD.jpg', 'Bold family SUV'),
('Creta', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', 'second', 'Global', 'hyundai/2022_Hyundai_Creta_16.jpg', 'Compact global hit'),
('Venue', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', 'first', 'Global', 'hyundai/2022_Hyundai_Venue.jpg', 'Connected compact SUV'),
('IONIQ 5', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'first', 'Global', 'hyundai/Hyundai_Ioniq_5_AWD.jpg', 'Retro-modern EV'),

-- Kia models
('Rio', (SELECT id FROM makes WHERE name = 'Kia'), '2017-2023', 'fourth', 'Global', 'kia_rio.jpg', 'Solid subcompact value'),
('Cerato', (SELECT id FROM makes WHERE name = 'Kia'), '2019-2024', 'fourth', 'Global', 'kia_cerato.jpg', 'Stylish compact sedan'),
('Optima', (SELECT id FROM makes WHERE name = 'Kia'), '2016-2020', 'fourth', 'Global', 'kia_optima.jpg', 'Value-packed sedan'),
('K5', (SELECT id FROM makes WHERE name = 'Kia'), '2020-2024', 'fifth', 'Global', 'kia_k5.jpg', 'Tiger nose evolution'),
('Sportage', (SELECT id FROM makes WHERE name = 'Kia'), '2022-2024', 'fifth', 'Global', 'kia_sportage_nq5.jpg', 'Futuristic organic SUV'),
('Seltos', (SELECT id FROM makes WHERE name = 'Kia'), '2019-2024', 'first', 'Global', 'kia_seltos.jpg', 'Sophisticated compact SUV'),
('Sorento', (SELECT id FROM makes WHERE name = 'Kia'), '2021-2024', 'fourth', 'Global', 'kia_sorento_mq4.jpg', 'Refined family hauler'),
('EV6', (SELECT id FROM makes WHERE name = 'Kia'), '2021-2024', 'first', 'Global', 'kia_ev6.jpg', 'Performance EV crossover'),

-- Genesis models
('G70', (SELECT id FROM makes WHERE name = 'Genesis'), '2019-2024', 'first', 'Global', 'genesis_g70.jpg', 'Athletic luxury sedan'),
('G80', (SELECT id FROM makes WHERE name = 'Genesis'), '2021-2024', 'second', 'Global', 'genesis_g80.jpg', 'Athletic elegance defined'),
('G90', (SELECT id FROM makes WHERE name = 'Genesis'), '2019-2024', 'first/second', 'Global', 'genesis_g90.jpg', 'Flagship Korean luxury'),
('GV70', (SELECT id FROM makes WHERE name = 'Genesis'), '2021-2024', 'first', 'Global', 'genesis_gv70.jpg', 'Dynamic luxury SUV'),
('GV80', (SELECT id FROM makes WHERE name = 'Genesis'), '2020-2024', 'first', 'Global', 'genesis_gv80.jpg', 'Distinctive grandeur'),

-- Ford models
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2005-2009', 'S197', 'Global', 'ford/mustang_s197.jpg', 'Retro-futuristic muscle car'),
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2023', 'S550', 'Global', 'ford/2019_Ford_Mustang_S550.jpg', 'Modern global muscle'),
('F-150', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2024', 'thirteenth', 'Global', 'ford/2016_Ford_F-150_13_14.jpg', 'America''s best seller'),
('Ranger', (SELECT id FROM makes WHERE name = 'Ford'), '2019-2024', 'T6', 'Global', 'ford/Ford_Ranger_T6.jpg', 'ford/Built Ford Tough global ute'),
('Explorer', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', 'sixth', 'Global', 'ford/Ford_Explorer_(sixth_generation).jpg', 'Modern family adventure'),
('Escape', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', 'fourth', 'Global', 'ford/2021_Ford_Escape_.jpg', 'Curvy compact SUV'),
('Bronco', (SELECT id FROM makes WHERE name = 'Ford'), '2021-present', 'sixth', 'Global', 'ford/Ford_Bronco_(6th).jpg', 'Capable retro-modern explorer'),
('Mustang Mach-E', (SELECT id FROM makes WHERE name = 'Ford'), '2021-2024', 'first', 'Global', 'ford/2021_Ford_Mustang_Mach-E_.jpg', 'Electric pony SUV'),

-- Chevrolet models
('Camaro', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2023', 'sixth', 'Global', 'chevrolet/camaro_sixth.jpg', 'Sharply styled muscle'),
('Corvette', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2020-2024', 'C8', 'Global', 'chevrolet/Corvette-C8.jpg', 'Mid-engine revolution'),
('Silverado', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2019-2024', 'fourth', 'Global', 'chevrolet/2022_Chevrolet_Silverado_2500HD.jpg', 'Heavy duty capability'),
('Malibu', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2024', 'ninth', 'Global', 'chevrolet/Chevrolet_Malibu_LT_(IX,_Facelift).jpg', 'Comfortable mid-size sedan'),
('Equinox', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2018-2024', 'third', 'Global', 'chevrolet/Chevrolet_Equinox_LT_(III,_Facelift).jpg', 'Popular family crossover'),
('Tahoe', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2021-2024', 'fifth', 'Global', 'chevrolet/2022_Chevrolet_Tahoe.jpg', 'Full-size dominance'),

-- GMC models
('Sierra', (SELECT id FROM makes WHERE name = 'GMC'), '2019-2024', 'fifth', 'Global', 'gmc_sierra.jpg', 'Professional grade truck'),
('Yukon', (SELECT id FROM makes WHERE name = 'GMC'), '2021-2024', 'fifth', 'Global', 'gmc_yukon.jpg', 'Premium full-size SUV'),
('Acadia', (SELECT id FROM makes WHERE name = 'GMC'), '2017-2024', 'second', 'Global', 'gmc_acadia.jpg', 'Versatile mid-size crossover'),
('Terrain', (SELECT id FROM makes WHERE name = 'GMC'), '2018-2024', 'second', 'Global', 'gmc_terrain.jpg', 'Stylized compact SUV'),

-- Cadillac models
('CT4', (SELECT id FROM makes WHERE name = 'Cadillac'), '2020-2024', 'first', 'Global', 'cadillac_ct4.jpg', 'Compact luxury sport'),
('CT5', (SELECT id FROM makes WHERE name = 'Cadillac'), '2020-2024', 'first', 'Global', 'cadillac_ct5.jpg', 'Stylish mid-size luxury'),
('Escalade', (SELECT id FROM makes WHERE name = 'Cadillac'), '2021-2024', 'fifth', 'Global', 'cadillac_escalade.jpg', 'The standard of bling'),
('XT4', (SELECT id FROM makes WHERE name = 'Cadillac'), '2019-2024', 'first', 'Global', 'cadillac_xt4.jpg', 'Tailored urban crossover'),
('Lyriq', (SELECT id FROM makes WHERE name = 'Cadillac'), '2023-2024', 'first', 'Global', 'cadillac_lyriq.jpg', 'Electric luxury future'),

-- Dodge models
('Charger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LD', 'Global', 'dodge/dodge_charger.jpg', 'Four-door muscle car'),
('Challenger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LC', 'Global', 'dodge/2017_Dodge_Challenger_R_T_Scat_Pack.jpg', 'Classic muscle tribute'),
('Challenger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-present', 'LA', 'Global', 'dodge/dodge_challenger_la.jpg', 'High-performance muscle'),
('Durango', (SELECT id FROM makes WHERE name = 'Dodge'), '2016-2024', 'third', 'Global', 'dodge/2021-dodge-durango-srt-hellcat.jpg', 'Muscle SUV for seven'),

-- Tesla models
('Model S', (SELECT id FROM makes WHERE name = 'Tesla'), '2012-present', 'first', 'Global', 'tesla/model_s.jpg', 'The EV game changer'),
('Model 3', (SELECT id FROM makes WHERE name = 'Tesla'), '2017-2024', 'first', 'Global', 'tesla/model_3.jpg', 'Mass market electric revolution'),
('Model X', (SELECT id FROM makes WHERE name = 'Tesla'), '2016-2024', 'first', 'Global', 'tesla/model_x.jpg', 'Falcon wing spectacle'),
('Model Y', (SELECT id FROM makes WHERE name = 'Tesla'), '2020-2024', 'first', 'Global', 'tesla/model_y.jpg', 'Best selling electric crossover'),
('Cybertruck', (SELECT id FROM makes WHERE name = 'Tesla'), '2024', 'first', 'Global', 'tesla/cybertruck.jpg', 'Polygonal exoskeleton truck'),

-- Rivian models
('R1T', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', 'first', 'Global', 'rivian_r1t.jpg', 'Electric adventure truck'),
('R1S', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', 'first', 'Global', 'rivian_r1s.jpg', 'Electric adventure SUV'),

-- Lucid models
('Air', (SELECT id FROM makes WHERE name = 'Lucid'), '2022-2024', 'first', 'Global', 'lucid_air.jpg', 'Luxury efficiency champion'),


-- Jaguar models
('F-Type', (SELECT id FROM makes WHERE name = 'Jaguar'), '2017-2019', 'X152', 'Global', 'jaguar/2017_Jaguar_F-Type_V6_R-Dynamic.jpg', 'British sports car excellence'),
('XE', (SELECT id FROM makes WHERE name = 'Jaguar'), '2015-2020', 'X760', 'Global', 'jaguar/2019_Jaguar_XE_S_Automatic.jpg', 'Compact executive with style'),
('XF', (SELECT id FROM makes WHERE name = 'Jaguar'), '2016-2020', 'X260', 'Global', 'jaguar/2018_Jaguar_XF_V6_S_Diesel_Automatic_30.jpg', 'Mid-size luxury with flair'),
('XF', (SELECT id FROM makes WHERE name = 'Jaguar'), '2007-2015', 'X250', 'Global', 'jaguar/2010_Jaguar_XF_Premium_Luxury_V6_Automatic_30.jpg', 'Modern Jaguar design revolution'),
('XJ', (SELECT id FROM makes WHERE name = 'Jaguar'), '2010-2019', 'X351', 'Global', 'jaguar/Jaguar_XJ.jpg', 'Flagship luxury sedan'),
('E-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2018-2024', 'first', 'Global', 'jaguar/2019-jaguar-e-pace-epace-cp-0.jpg', 'Compact luxury SUV'),
('F-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2017-2024', 'first', 'Global', 'jaguar/f-pace.jpg', 'Sporty luxury SUV'),
('I-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2018-2024', 'first', 'Global', 'jaguar/2018_Jaguar_I-Pace_EV400_AWD.jpg', 'Electric performance SUV'),

-- Land Rover models
('Range Rover', (SELECT id FROM makes WHERE name = 'Land Rover'), '2018-2024', 'fifth', 'Global', 'landrover/2022_Land_Rover_Range_Rover_SE_P440e_AWD_Automatic_30.jpg', 'Luxury off-road icon'),
('Discovery', (SELECT id FROM makes WHERE name = 'Land Rover'), '2017-2024', 'fifth', 'Global', 'landrover/2018_Land_Rover_Discovery_Luxury_HSE_TD6_30.jpg', 'Versatile family SUV'),
('Range Rover Sport', (SELECT id FROM makes WHERE name = 'Land Rover'), '2005-2013', 'L320', 'Global', 'landrover/First_Range_Rover_Sport_l320.jpg', 'Luxury off-road capability'),
('Range Rover Velar', (SELECT id FROM makes WHERE name = 'Land Rover'), '2017-present', 'L560', 'Global', 'landrover/range_rover_velar.jpeg', 'Sleek reductionist design'),
('Defender', (SELECT id FROM makes WHERE name = 'Land Rover'), '2020-present', 'L663', 'Global', 'landrover/defender_l663.jpg', 'Reinvented off-road icon'),
('Defender', (SELECT id FROM makes WHERE name = 'Land Rover'), '1983-2016', 'L316', 'Global', 'landrover/defender_l316.jpg', 'Original mechanical hero'),

-- Jeep models
('Wrangler', (SELECT id FROM makes WHERE name = 'Jeep'), '2018-2024', 'JL', 'Global', 'jeep/Jeep_Wrangler_Unlimited_(JL)_PHEV.jpg', 'The ultimate off-road freedom'),
('Wrangler', (SELECT id FROM makes WHERE name = 'Jeep'), '2007-2018', 'JK', 'Global', 'jeep/Jeep_JK_Wrangler_Sahara_2-Door_Convertible.jpg', 'Ultimate off-road freedom'),
('Grand Cherokee', (SELECT id FROM makes WHERE name = 'Jeep'), '2011-2021', 'WK2', 'Global', 'jeep/2012_Jeep_Grand_Cherokee_(WK2_MY12)_Laredo_CRD_4WD.jpg', 'Luxury off-road capability'),
('Cherokee', (SELECT id FROM makes WHERE name = 'Jeep'), '2014-2024', 'KL', 'Global', 'jeep/2019_Jeep_Cherokee_Latitude.jpg', 'Stylish compact SUV'),
('Renegade', (SELECT id FROM makes WHERE name = 'Jeep'), '2015-2024', 'BU', 'Global', 'jeep/Jeep_Renegade_16_MultiJet_2WD.jpg', 'Subcompact urban adventurer'),

-- Mini models
('Cooper', (SELECT id FROM makes WHERE name = 'Mini'), '2001-2006', 'R50', 'Global', 'mini/cooper_r50.jpg', 'Retro hatchback fun'),
('Cooper S', (SELECT id FROM makes WHERE name = 'Mini'), '2002-2006', 'R53', 'Global', 'mini/cooper_r53.jpg', 'Supercharged pocket rocket'),
('Clubman', (SELECT id FROM makes WHERE name = 'Mini'), '2015-2023', 'first', 'Global', 'mini/mini-clubman.jpg', 'Quirky six-door hatch'),
('Cooper', (SELECT id FROM makes WHERE name = 'Mini'), '2014-present', 'F56', 'Global', 'mini/mini-cooper-s-f56.jpg', 'Go-kart handling fun'),
('Countryman', (SELECT id FROM makes WHERE name = 'Mini'), '2010-2016', 'R60', 'Global', 'mini/2018_Mini_Countryman_Cooper_S_R60.jpg', 'The original Mini SUV'),
('Paceman', (SELECT id FROM makes WHERE name = 'Mini'), '2013-2016', 'R61', 'Global', 'mini/paceman_r61.jpeg', 'Coupe-style crossover');