-- Insert makes from new data
INSERT INTO makes (name) VALUES
('Opel'), ('Volkswagen'), ('Audi'), ('BMW'), ('Mercedes-Benz'), ('Volvo'), ('Renault'), ('Peugeot'), ('Fiat'), ('Alfa Romeo'),
('Porsche'), ('Ferrari'), ('Lamborghini'), ('Aston Martin'), ('Bentley'), ('Rolls-Royce'), ('Maserati'), ('Bugatti'),
('Toyota'), ('Lexus'), ('Honda'), ('Acura'), ('Nissan'), ('Infiniti'), ('Mazda'), ('Subaru'), ('Mitsubishi'), ('Suzuki'),
('Hyundai'), ('Kia'), ('Genesis'), ('Ford'), ('Chevrolet'), ('GMC'), ('Cadillac'), ('Dodge'), ('Tesla'), ('Rivian'), ('Lucid'), ('Jaguar'), ('Land Rover'), ('Jeep'), ('Mini');

-- Insert models from new data
INSERT INTO models (name, make_id, year_range, generation, location, image_url, known_for) VALUES
-- Opel models
('Astra', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', 'L', 'Global', 'opel_astra_l.jpg', 'Sharp styling and modern tech'),
('Corsa', (SELECT id FROM makes WHERE name = 'Opel'), '2020-2024', 'F', 'Global', 'opel_corsa_f.jpg', 'Popular city car'),
('Grandland', (SELECT id FROM makes WHERE name = 'Opel'), '2018-2024', 'first', 'Global', 'opel_grandland.jpg', 'Compact SUV contender'),
('Mokka', (SELECT id FROM makes WHERE name = 'Opel'), '2021-2024', 'second', 'Global', 'opel_mokka.jpg', 'Bold Vizor design'),
('Insignia', (SELECT id FROM makes WHERE name = 'Opel'), '2017-2023', 'B', 'Global', 'opel_insignia_b.jpg', 'Sleek executive tourer'),

-- Volkswagen models
('Golf', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2020-2024', 'Mk8', 'Global', 'vw_golf_mk8.jpg', 'The benchmark compact hatchback'),
('Golf', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2014-2019', 'Mk7', 'Global', 'Volkswagen_Golf.jpg', 'The complete compact package'),
('Passat', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2019-2023', 'B8', 'Global', 'vw_passat_b8.jpg', 'Definitive family estate'),
('Tiguan', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'second', 'Global', 'vw_tiguan_mk2.jpg', 'Best-selling refined SUV'),
('T-Roc', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'first', 'Global', 'vw_troc.jpg', 'Stylish compact crossover'),
('Touareg', (SELECT id FROM makes WHERE name = 'Volkswagen'), '2018-2024', 'third', 'Global', 'vw_touareg_mk3.jpg', 'Premium luxury flagship SUV'),

-- Audi models
('A3', (SELECT id FROM makes WHERE name = 'Audi'), '2020-2024', '8Y', 'Global', 'audi_a3_8y.jpg', 'Premium compact superiority'),
('A4', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', 'B9', 'Global', 'audi_a4_b9.jpg', 'Executive sedan staple'),
('A6', (SELECT id FROM makes WHERE name = 'Audi'), '2019-2024', 'C8', 'Global', 'audi_a6_c8.jpg', 'High-tech business class'),
('Q5', (SELECT id FROM makes WHERE name = 'Audi'), '2018-2024', 'second', 'Global', 'audi_q5_mk2.jpg', 'Best-selling luxury SUV'),
('Q7', (SELECT id FROM makes WHERE name = 'Audi'), '2016-2023', 'second', 'Global', 'audi_q7_mk2.jpg', 'Seven-seat luxury family hauler'),
('TT', (SELECT id FROM makes WHERE name = 'Audi'), '2015-2023', '8S', 'Global', 'tt_8s.jpg', 'Design icon'),
('R8', (SELECT id FROM makes WHERE name = 'Audi'), '2015-2023', '4S', 'Global', 'Audi_R8.jpg', 'V10 mid-engine masterpiece'),


-- BMW models
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G20', 'Global', 'bmw_3series_g20.jpg', 'The ultimate sports sedan'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1982-1994', 'E30', 'Global', 'bmw_e30_m3.jpg', 'The original M3 legend'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1998-2006', 'E46', 'Global', 'Bmw_3_Series.jpg', 'The gold standard of sports sedans'),
('3 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1990-1998', 'E36', 'Global', 'Bmw_3_Series_1.jpg', 'Balanced driving perfection'),
('5 Series', (SELECT id FROM makes WHERE name = 'BMW'), '2017-2023', 'G30', 'Global', 'bmw_5series_g30.jpg', 'Dynamic business athlete'),
('5 Series', (SELECT id FROM makes WHERE name = 'BMW'), '1995-2003', 'E39', 'Global', 'Bmw_5_Series.jpg', 'Often cited as the best 5 Series ever'),
('M4', (SELECT id FROM makes WHERE name = 'BMW'), '2014-2020', 'F82', 'Global', 'Bmw_M4.jpg', 'Swift twin-turbo coupe'),
('X1', (SELECT id FROM makes WHERE name = 'BMW'), '2009-2015', 'E84', 'Global', 'Bmw_X1.jpg', 'Compact agile utility'),
('X3', (SELECT id FROM makes WHERE name = 'BMW'), '2018-2024', 'G01', 'Global', 'bmw_x3_g01.jpg', 'SAV versatility with sportiness'),
('X5', (SELECT id FROM makes WHERE name = 'BMW'), '2019-2024', 'G05', 'Global', 'bmw_x5_g05.jpg', 'The boss of luxury SUVs'),
('M3', (SELECT id FROM makes WHERE name = 'BMW'), '2021-present', 'G80', 'Global', 'bmw_m3_g80.jpg', 'Polarizing grille, undeniable performance'),
('M4', (SELECT id FROM makes WHERE name = 'BMW'), '2021-2024', 'G82', 'Global', 'bmw_m4_g82.jpg', 'Aggressive coupe dominator'),

-- Mercedes-Benz models
('AMG GT', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2017-2020', 'C190', 'Global', 'amg_gt_r_c190.jpg', 'Track focused beast'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2007-2014', 'W204', 'Global', 'merc_c_w204.jpg', 'Angular styling reliability'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W206', 'Global', 'merc_c_w206.jpg', 'Baby S-Class technology'),
('C-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2018-2021', 'W205', 'Global', 'Mercedes-Benz_C-Class.jpg', 'Executive comfort and tech'),
('E-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2017-2023', 'W213', 'Global', 'merc_e_w213.jpg', 'Smooth executive cruiser'),
('E-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '1976-1985', 'W123', 'Global', 'Mercedes-Benz_W123.jpg', 'Undying mechanical durability'),
('S-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2021-2024', 'W223', 'Global', 'merc_s_w223.jpg', 'The standard of the world'),
('S-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '1979-1991', 'W126', 'Global', 'Mercedes-Benz_S-Class.jpg', 'Engineering peak of luxury'),
('GLA', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2020-2024', 'H247', 'Global', 'merc_gla_h247.jpg', 'Compact urban luxury'),
('G-Class', (SELECT id FROM makes WHERE name = 'Mercedes-Benz'), '2018-2024', 'W463', 'Global', 'merc_g_w463.jpg', 'The iconic Geländewagen'),

-- Volvo models
('S60', (SELECT id FROM makes WHERE name = 'Volvo'), '2019-2024', 'third', 'Global', 'volvo_s60_mk3.jpg', 'Scandinavian sports sedan'),
('S90', (SELECT id FROM makes WHERE name = 'Volvo'), '2017-2024', 'second', 'Global', 'volvo_s90_mk2.jpg', 'Elegant luxury flagship'),
('XC40', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', 'first', 'Global', 'volvo_xc40.jpg', 'Funky compact premium SUV'),
('XC60', (SELECT id FROM makes WHERE name = 'Volvo'), '2018-2024', 'second', 'Global', 'volvo_xc60_mk2.jpg', 'Nordic bestseller'),
('XC90', (SELECT id FROM makes WHERE name = 'Volvo'), '2016-2024', 'second', 'Global', 'volvo_xc90_mk2.jpg', 'Thor hammer seven-seater'),

-- Renault models
('Clio', (SELECT id FROM makes WHERE name = 'Renault'), '2019-2024', 'fifth', 'Global', 'renault_clio_v.jpg', 'Chic supermini icon'),
('Megane', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2023', 'fourth', 'Global', 'renault_megane_iv.jpg', 'Distinctive lighting signature'),
('Captur', (SELECT id FROM makes WHERE name = 'Renault'), '2020-2024', 'second', 'Global', 'renault_captur_ii.jpg', 'Versatile urban crossover'),
('Kadjar', (SELECT id FROM makes WHERE name = 'Renault'), '2016-2022', 'first', 'Global', 'renault_kadjar.jpg', 'Practical family SUV'),
('Duster', (SELECT id FROM makes WHERE name = 'Renault'), '2018-2024', 'second', 'Global', 'renault_duster_ii.jpg', 'Robust affordable adventurer'),

-- Peugeot models
('208', (SELECT id FROM makes WHERE name = 'Peugeot'), '2019-2024', 'second', 'Global', 'peugeot_208_ii.jpg', 'Sabre-tooth styling'),
('308', (SELECT id FROM makes WHERE name = 'Peugeot'), '2021-2024', 'third', 'Global', 'peugeot_308_iii.jpg', 'New logo pioneer'),
('508', (SELECT id FROM makes WHERE name = 'Peugeot'), '2018-2023', 'second', 'Global', 'peugeot_508_ii.jpg', 'Radical fastback design'),
('2008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2020-2024', 'second', 'Global', 'peugeot_2008_ii.jpg', 'Sharp-edged compact SUV'),
('3008', (SELECT id FROM makes WHERE name = 'Peugeot'), '2017-2024', 'second', 'Global', 'peugeot_3008_ii.jpg', 'Award-winning style'),

-- Fiat models
('500', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', 'third', 'Global', 'fiat_500.jpg', 'Timeless retro chic'),
('500X', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2023', 'first', 'Global', 'fiat_500x.jpg', 'Crossover with heritage'),
('Panda', (SELECT id FROM makes WHERE name = 'Fiat'), '2012-2024', 'third', 'Global', 'fiat_panda_iii.jpg', 'Boxy practical hero'),
('Tipo', (SELECT id FROM makes WHERE name = 'Fiat'), '2016-2024', 'second', 'Global', 'fiat_tipo.jpg', 'Value-focused practicality'),

-- Alfa Romeo models
('Giulia', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2016-2024', '952', 'Global', 'alfa_giulia.jpg', 'Emotional sports sedan'),
('Stelvio', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2017-2024', '949', 'Global', 'alfa_stelvio.jpg', 'The driving SUV'),
('Giulietta', (SELECT id FROM makes WHERE name = 'Alfa Romeo'), '2014-2021', '940', 'Global', 'alfa_giulietta.jpg', 'Curvaceous hatchback'),

-- Porsche models
('911', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', '992', 'Global', 'porsche_911_992.jpg', 'The everyday supercar icon'),
('911', (SELECT id FROM makes WHERE name = 'Porsche'), '2006-2013', '997', 'Global', 'Porsche_911_Turbo.jpg', 'The usable daily supercar'),
('718 Cayman', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global', 'porsche_718_cayman.jpg', 'Mid-engine precision'),
('718 Boxster', (SELECT id FROM makes WHERE name = 'Porsche'), '2016-2024', '982', 'Global', 'porsche_718_boxster.jpg', 'Open-top purity'),
('Panamera', (SELECT id FROM makes WHERE name = 'Porsche'), '2017-2024', '971', 'Global', 'porsche_panamera.jpg', 'Four-door GT performance'),
('Cayenne', (SELECT id FROM makes WHERE name = 'Porsche'), '2018-2024', 'third', 'Global', 'porsche_cayenne_mk3.jpg', 'The sportiest SUV'),
('Macan', (SELECT id FROM makes WHERE name = 'Porsche'), '2019-2024', 'first', 'Global', 'porsche_macan.jpg', 'Compact handling benchmark'),
('Taycan', (SELECT id FROM makes WHERE name = 'Porsche'), '2020-2024', 'first', 'Global', 'porsche_taycan.jpg', 'Electric soul'),

-- Ferrari models
('488 GTB', (SELECT id FROM makes WHERE name = 'Ferrari'), '2015-2019', 'F142', 'Global', 'ferrari_488.jpg', 'Twin-turbo V8 fury'),
('F8 Tributo', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2023', 'F142M', 'Global', 'ferrari_f8.jpg', 'Evolution of excellence'),
('Roma', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F169', 'Global', 'ferrari_roma.jpg', 'La Nuova Dolce Vita'),
('SF90 Stradale', (SELECT id FROM makes WHERE name = 'Ferrari'), '2020-2024', 'F173', 'Global', 'ferrari_sf90.jpg', 'Hybrid hypercar performance'),
('812 Superfast', (SELECT id FROM makes WHERE name = 'Ferrari'), '2017-2024', 'F152M', 'Global', 'ferrari_812.jpg', 'V12 grand tourer majesty'),
('Portofino', (SELECT id FROM makes WHERE name = 'Ferrari'), '2018-2023', 'F164', 'Global', 'ferrari_portofino.jpg', 'Convertible elegance'),
('458 Italia', (SELECT id FROM makes WHERE name = 'Ferrari'), '2010-2015', '458', 'Global', '458_italia.jpg', 'The last N/A V8'),
('F355', (SELECT id FROM makes WHERE name = 'Ferrari'), '1994-1999', 'F129', 'Global', 'ferrari_f355.jpg', 'Classic V8 mid-engine beauty'),


-- Lamborghini models
('Huracan', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2015-2024', 'first', 'Global', 'lambo_huracan.jpg', 'V10 screaming bull'),
('Huracan', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2014-2019', 'LP610-4', 'Global', 'Lamborghini_Huracan.jpg', 'V10 mid-engine thrill'),
('Huracan', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2019-2023', 'EVO', 'Global', 'Lamborghini_Huracan_1.jpg', 'Enhanced aero and tech'),
('Aventador', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2012-2021', 'LP700-4', 'Global', 'Lamborghini_Aventador.jpg', 'The ultimate V12 poster car'),
('Murcielago', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2001-2010', 'LP640', 'Global', 'murcielago_lp640.jpg', 'Dramatic scissor-door icon'),
('Urus', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2018-2024', 'first', 'Global', 'lambo_urus.jpg', 'Super SUV dominator'),
('Revuelto', (SELECT id FROM makes WHERE name = 'Lamborghini'), '2024', 'first', 'Global', 'lambo_revuelto.jpg', 'The first HPEV hybrid'),

-- Aston Martin models
('DB11', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2017-2023', 'first', 'Global', 'aston_db11.jpg', 'Grand touring sculpture'),
('DBS Superleggera', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2019-2023', 'first', 'Global', 'aston_dbs.jpg', 'Brute in a suit'),
('Vantage', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2018-2024', 'second', 'Global', 'aston_vantage.jpg', 'Hunter instinct'),
('DBX', (SELECT id FROM makes WHERE name = 'Aston Martin'), '2020-2024', 'first', 'Global', 'aston_dbx.jpg', 'Luxury capability'),

-- Bentley models
('Continental GT', (SELECT id FROM makes WHERE name = 'Bentley'), '2018-2024', 'third', 'Global', 'bentley_conti_gt.jpg', 'Deep-chested touring capability'),
('Flying Spur', (SELECT id FROM makes WHERE name = 'Bentley'), '2020-2024', 'third', 'Global', 'bentley_flying_spur.jpg', 'Four-door super-luxury'),
('Bentayga', (SELECT id FROM makes WHERE name = 'Bentley'), '2016-2024', 'first', 'Global', 'bentley_bentayga.jpg', 'Pinnacle SUV'),

-- Rolls-Royce models
('Phantom', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2018-2024', 'VIII', 'Global', 'rr_phantom.jpg', 'The best car in the world'),
('Ghost', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2021-2024', 'second', 'Global', 'rr_ghost.jpg', 'Post-opulence luxury'),
('Wraith', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2014-2023', 'first', 'Global', 'rr_wraith.jpg', 'Ultimate gentleman GT'),
('Cullinan', (SELECT id FROM makes WHERE name = 'Rolls-Royce'), '2019-2024', 'first', 'Global', 'rr_cullinan.jpg', 'Effortless everywhere'),

-- Maserati models
('Ghibli', (SELECT id FROM makes WHERE name = 'Maserati'), '2014-2023', 'M156', 'Global', 'maserati_ghibli.jpg', 'Italian sports sedan flare'),
('Quattroporte', (SELECT id FROM makes WHERE name = 'Maserati'), '2013-2023', 'M156', 'Global', 'maserati_quattroporte.jpg', 'The original race-bred saloon'),
('Levante', (SELECT id FROM makes WHERE name = 'Maserati'), '2016-2024', 'first', 'Global', 'maserati_levante.jpg', 'The Maserati of SUVs'),
('MC20', (SELECT id FROM makes WHERE name = 'Maserati'), '2021-2024', 'first', 'Global', 'maserati_mc20.jpg', 'Super sports car return'),
('GranTurismo', (SELECT id FROM makes WHERE name = 'Maserati'), '2007-2019', 'M145', 'Global', 'Maserati_Granturismo.jpg', 'Italian vocal masterpiece'),


-- Bugatti models
('Veyron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2005-2015', 'first', 'Global', 'bugatti_veyron.jpg', 'The concord moment of cars'),
('Chiron', (SELECT id FROM makes WHERE name = 'Bugatti'), '2016-2024', 'first', 'Global', 'bugatti_chiron.jpg', 'Breaking physics'),
('Mistral', (SELECT id FROM makes WHERE name = 'Bugatti'), '2024', 'first', 'Global', 'bugatti_mistral.jpg', 'The last W16 roadster'),

-- Toyota models
('Corolla', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'E210', 'Global', 'toyota_corolla_e210.jpg', 'Best selling car globally'),
('Camry', (SELECT id FROM makes WHERE name = 'Toyota'), '2018-2024', 'XV70', 'Global', 'toyota_camry_xv70.jpg', 'Reliable mid-size standard'),
('Land Cruiser', (SELECT id FROM makes WHERE name = 'Toyota'), '2021-2024', 'J300', 'Global', 'toyota_lc300.jpg', 'Unstoppable off-road legend'),
('Land Cruiser', (SELECT id FROM makes WHERE name = 'Toyota'), '2015-2021', 'J200', 'Global', 'Toyota_Land_Cruiser.jpg', 'Rugged luxury cruiser'),
('Tacoma', (SELECT id FROM makes WHERE name = 'Toyota'), '2016-present', 'N300', 'Global', 'Toyota_Tacoma.jpg', 'Reliable mid-size adventurer'),
('RAV4', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'XA50', 'Global', 'toyota_rav4.jpg', 'The defining modern crossover'),
('Hilux', (SELECT id FROM makes WHERE name = 'Toyota'), '2016-2024', 'AN120', 'Global', 'toyota_hilux.jpg', 'Indestructible workhorse'),
('Prius', (SELECT id FROM makes WHERE name = 'Toyota'), '2023-2024', 'fifth', 'Global', 'toyota_prius_v.jpg', 'From boring to beautiful'),
('Supra', (SELECT id FROM makes WHERE name = 'Toyota'), '2019-2024', 'A90', 'Global', 'toyota_supra_a90.jpg', 'German heart Japanese soul'),
('Supra', (SELECT id FROM makes WHERE name = 'Toyota'), '1993-2002', 'A80', 'Global', 'supra_mk4.jpg', 'The 10-second car'),
('MR2', (SELECT id FROM makes WHERE name = 'Toyota'), '1990-1999', 'SW20', 'Global', 'mr2_sw20.jpg', 'Mid-ship runabout'),
('86', (SELECT id FROM makes WHERE name = 'Toyota'), '2012-2020', 'ZN6', 'Global', 'Toyota_86.jpg', 'Lightweight handling purist'),

-- Lexus models
('IS', (SELECT id FROM makes WHERE name = 'Lexus'), '2021-2024', 'third', 'Global', 'lexus_is.jpg', 'Sharp compact executive'),
('ES', (SELECT id FROM makes WHERE name = 'Lexus'), '2019-2024', 'seventh', 'Global', 'lexus_es.jpg', 'Serene comfort cruiser'),
('RX', (SELECT id FROM makes WHERE name = 'Lexus'), '2023-2024', 'fifth', 'Global', 'lexus_rx.jpg', 'Pioneer of luxury SUVs'),
('NX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', 'second', 'Global', 'lexus_nx.jpg', 'Angular urban luxury'),
('LX', (SELECT id FROM makes WHERE name = 'Lexus'), '2022-2024', 'third', 'Global', 'lexus_lx.jpg', 'Land Cruiser in a tuxedo'),

-- Honda models
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', 'eleventh', 'Global', 'honda_civic_xi.jpg', 'Mature sporty compact'),
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2006-2011', 'FD', 'Global', 'fd_civic.jpg', 'VTEC legend'),
('Civic', (SELECT id FROM makes WHERE name = 'Honda'), '2016-2021', 'FC / FK', 'Global', 'Honda_Civic.jpg', 'Sophisticated global compact'),
('Accord', (SELECT id FROM makes WHERE name = 'Honda'), '2018-2023', 'tenth', 'Global', 'honda_accord_x.jpg', 'Sleek family sedan'),
('CR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2023-2024', 'sixth', 'Global', 'honda_crv_vi.jpg', 'Practical family hauler'),
('HR-V', (SELECT id FROM makes WHERE name = 'Honda'), '2022-2024', 'third', 'Global', 'honda_hrv.jpg', 'Stylish subcompact utility'),
('City', (SELECT id FROM makes WHERE name = 'Honda'), '2020-2024', 'seventh', 'Global', 'honda_city.jpg', 'Global compact sedan'),

-- Acura models
('ILX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2022', 'first', 'Global', 'acura_ilx.jpg', 'Entry luxury sedan'),
('TLX', (SELECT id FROM makes WHERE name = 'Acura'), '2021-2024', 'second', 'Global', 'acura_tlx.jpg', 'Precision crafted performance'),
('RDX', (SELECT id FROM makes WHERE name = 'Acura'), '2019-2024', 'third', 'Global', 'acura_rdx.jpg', 'Sharp-handling SUV'),
('MDX', (SELECT id FROM makes WHERE name = 'Acura'), '2022-2024', 'fourth', 'Global', 'acura_mdx.jpg', 'Flagship three-row SUV'),

-- Nissan models
('Altima', (SELECT id FROM makes WHERE name = 'Nissan'), '2019-2024', 'sixth', 'Global', 'nissan_altima.jpg', 'Efficient mid-size sedan'),
('Sentra', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', 'eighth', 'Global', 'nissan_sentra.jpg', 'Sharp compact value'),
('X-Trail', (SELECT id FROM makes WHERE name = 'Nissan'), '2022-2024', 'T33', 'Global', 'nissan_xtrail.jpg', 'Family adventure ready'),
('Patrol', (SELECT id FROM makes WHERE name = 'Nissan'), '2020-2024', 'Y62', 'Global', 'nissan_patrol.jpg', 'Desert conquering giant'),
('GT-R', (SELECT id FROM makes WHERE name = 'Nissan'), '2009-2024', 'R35', 'Global', 'nissan_gtr.jpg', 'The supercar killer'),
('GT-R', (SELECT id FROM makes WHERE name = 'Nissan'), '1999-2002', 'R34', 'Global', 'gtr_r34_skyline.jpg', 'The legendary Skyline'),
('370Z', (SELECT id FROM makes WHERE name = 'Nissan'), '2009-2020', 'Z34', 'Global', '370z_z34.jpg', 'V6 sports coupe'),
('350Z', (SELECT id FROM makes WHERE name = 'Nissan'), '2002-2009', 'Z33', 'Global', 'Nissan_350Z.jpg', 'V6 sports coupe icon'),


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
('Outlander', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2022-2024', 'fourth', 'Global', 'mitsubishi_outlander.jpg', 'Bold dynamic shield'),
('Pajero Sport', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2016-2024', 'third', 'Global', 'mitsubishi_pajero_sport.jpg', 'Robust off-road SUV'),
('Lancer', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2008-2016', 'tenth', 'Global', 'mitsubishi_lancer.jpg', 'Evo heritage sedan'),
('Eclipse Cross', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2018-2024', 'first', 'Global', 'mitsubishi_eclipse_cross.jpg', 'Coupe-style SUV'),
('Lancer Evolution', (SELECT id FROM makes WHERE name = 'Mitsubishi'), '2003-2006', 'Evo 8', 'Global', 'Mitsubishi_Lancer_Evolution.jpg', 'Rally legend for the road'),


-- Suzuki models
('Swift', (SELECT id FROM makes WHERE name = 'Suzuki'), '2017-2024', 'third', 'Global', 'suzuki_swift.jpg', 'Lightweight fun hatch'),
('Vitara', (SELECT id FROM makes WHERE name = 'Suzuki'), '2019-2024', 'second', 'Global', 'suzuki_vitara.jpg', 'Compact capable SUV'),
('Jimny', (SELECT id FROM makes WHERE name = 'Suzuki'), '2018-2024', 'fourth', 'Global', 'suzuki_jimny.jpg', 'Tiny off-road giant'),
('Baleno', (SELECT id FROM makes WHERE name = 'Suzuki'), '2016-2024', 'second', 'Global', 'suzuki_baleno.jpg', 'Practical spacious hatch'),

-- Hyundai models
('Elantra', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'seventh', 'Global', 'hyundai_elantra_cn7.jpg', 'Parametric dynamics design'),
('Sonata', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', 'eighth', 'Global', 'hyundai_sonata_dn8.jpg', 'Sensuous sportiness'),
('Tucson', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'fourth', 'Global', 'hyundai_tucson_nx4.jpg', 'Revolutionary lighting face'),
('Santa Fe', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', 'fourth', 'Global', 'hyundai_santafe.jpg', 'Bold family SUV'),
('Creta', (SELECT id FROM makes WHERE name = 'Hyundai'), '2020-2024', 'second', 'Global', 'hyundai_creta.jpg', 'Compact global hit'),
('Venue', (SELECT id FROM makes WHERE name = 'Hyundai'), '2019-2024', 'first', 'Global', 'hyundai_venue.jpg', 'Connected compact SUV'),
('IONIQ 5', (SELECT id FROM makes WHERE name = 'Hyundai'), '2021-2024', 'first', 'Global', 'hyundai_ioniq5.jpg', 'Retro-modern EV'),

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
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2005-2009', 'S197', 'Global', 'ford_mustang_s197.jpg', 'Retro-futuristic muscle car'),
('Mustang', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2023', 'S550', 'Global', 'ford_mustang_s550.jpg', 'Modern global muscle'),
('F-150', (SELECT id FROM makes WHERE name = 'Ford'), '2015-2024', 'thirteenth-fourteenth', 'Global', 'ford_f150.jpg', 'America''s best seller'),
('Ranger', (SELECT id FROM makes WHERE name = 'Ford'), '2019-2024', 'T6', 'Global', 'ford_ranger.jpg', 'Built Ford Tough global ute'),
('Explorer', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', 'sixth', 'Global', 'ford_explorer.jpg', 'Modern family adventure'),
('Escape', (SELECT id FROM makes WHERE name = 'Ford'), '2020-2024', 'fourth', 'Global', 'ford_escape.jpg', 'Curvy compact SUV'),
('Bronco', (SELECT id FROM makes WHERE name = 'Ford'), '2021-present', 'sixth', 'Global', 'Ford_Bronco.jpg', 'Capable retro-modern explorer'),
('Mustang Mach-E', (SELECT id FROM makes WHERE name = 'Ford'), '2021-2024', 'first', 'Global', 'ford_mach_e.jpg', 'Electric pony SUV'),

-- Chevrolet models
('Camaro', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2023', 'sixth', 'Global', 'chevy_camaro.jpg', 'Sharply styled muscle'),
('Corvette', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2020-2024', 'C8', 'Global', 'chevy_corvette_c8.jpg', 'Mid-engine revolution'),
('Silverado', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2019-2024', 'fourth', 'Global', 'chevy_silverado.jpg', 'Heavy duty capability'),
('Malibu', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2016-2024', 'ninth', 'Global', 'chevy_malibu.jpg', 'Comfortable mid-size sedan'),
('Equinox', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2018-2024', 'third', 'Global', 'chevy_equinox.jpg', 'Popular family crossover'),
('Tahoe', (SELECT id FROM makes WHERE name = 'Chevrolet'), '2021-2024', 'fifth', 'Global', 'chevy_tahoe.jpg', 'Full-size dominance'),

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
('Charger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LD', 'Global', 'dodge_charger.jpg', 'Four-door muscle car'),
('Challenger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-2023', 'LC', 'Global', 'dodge_challenger.jpg', 'Classic muscle tribute'),
('Challenger', (SELECT id FROM makes WHERE name = 'Dodge'), '2015-present', 'LA', 'Global', 'Dodge_Challenger.jpg', 'High-performance muscle'),
('Durango', (SELECT id FROM makes WHERE name = 'Dodge'), '2016-2024', 'third', 'Global', 'dodge_durango.jpg', 'Muscle SUV for seven'),

-- Tesla models
('Model S', (SELECT id FROM makes WHERE name = 'Tesla'), '2012-present', 'first', 'Global', 'Tesla_Model_S.jpg', 'The EV game changer'),
('Model 3', (SELECT id FROM makes WHERE name = 'Tesla'), '2017-2024', 'first', 'Global', 'tesla_model3.jpg', 'Mass market electric revolution'),
('Model X', (SELECT id FROM makes WHERE name = 'Tesla'), '2016-2024', 'first', 'Global', 'tesla_modelx.jpg', 'Falcon wing spectacle'),
('Model Y', (SELECT id FROM makes WHERE name = 'Tesla'), '2020-2024', 'first', 'Global', 'tesla_modely.jpg', 'Best selling electric crossover'),
('Cybertruck', (SELECT id FROM makes WHERE name = 'Tesla'), '2024', 'first', 'Global', 'tesla_cybertruck.jpg', 'Polygonal exoskeleton truck'),

-- Rivian models
('R1T', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', 'first', 'Global', 'rivian_r1t.jpg', 'Electric adventure truck'),
('R1S', (SELECT id FROM makes WHERE name = 'Rivian'), '2022-2024', 'first', 'Global', 'rivian_r1s.jpg', 'Electric adventure SUV'),

-- Lucid models
('Air', (SELECT id FROM makes WHERE name = 'Lucid'), '2022-2024', 'first', 'Global', 'lucid_air.jpg', 'Luxury efficiency champion'),


-- Jaguar models
('F-Type', (SELECT id FROM makes WHERE name = 'Jaguar'), '2017-2019', 'X152', 'Global', 'Jaguar_F-Type.jpg', 'British sports car excellence'),
('XE', (SELECT id FROM makes WHERE name = 'Jaguar'), '2015-2020', 'X760', 'Global', 'Jaguar_XE.jpg', 'Compact executive with style'),
('XF', (SELECT id FROM makes WHERE name = 'Jaguar'), '2016-2020', 'X260', 'Global', 'Jaguar_XF.jpg', 'Mid-size luxury with flair'),
('XF', (SELECT id FROM makes WHERE name = 'Jaguar'), '2007-2015', 'X250', 'Global', 'jaguar_xf_x250.jpg', 'Modern Jaguar design revolution'),
('XJ', (SELECT id FROM makes WHERE name = 'Jaguar'), '2010-2019', 'X351', 'Global', 'Jaguar_XJ.jpg', 'Flagship luxury sedan'),
('E-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2018-2024', 'first', 'Global', 'Jaguar_E-Pace.jpg', 'Compact luxury SUV'),
('F-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2017-2024', 'first', 'Global', 'Jaguar_F-Pace.jpg', 'Sporty luxury SUV'),
('I-Pace', (SELECT id FROM makes WHERE name = 'Jaguar'), '2018-2024', 'first', 'Global', 'Jaguar_I-Pace.jpg', 'Electric performance SUV'),

-- Land Rover models
('Range Rover', (SELECT id FROM makes WHERE name = 'Land Rover'), '2018-2024', 'fifth', 'Global', 'land_rover_range_rover.jpg', 'Luxury off-road icon'),
('Discovery', (SELECT id FROM makes WHERE name = 'Land Rover'), '2017-2024', 'fifth', 'Global', 'land_rover_discovery.jpg', 'Versatile family SUV'),
('Range Rover Sport', (SELECT id FROM makes WHERE name = 'Land Rover'), '2005-2013', 'L320', 'Global', 'Land_Rover_Range_Rover_Sport.jpg', 'Luxury off-road capability'),
('Range Rover Velar', (SELECT id FROM makes WHERE name = 'Land Rover'), '2017-present', 'L560', 'Global', 'Land_Rover_Range_Rover_Velar.jpg', 'Sleek reductionist design'),
('Defender', (SELECT id FROM makes WHERE name = 'Land Rover'), '2020-present', 'L663', 'Global', 'Land_Rover_Defender_110.jpg', 'Reinvented off-road icon'),
('Defender', (SELECT id FROM makes WHERE name = 'Land Rover'), '1983-2016', 'L316', 'Global', 'Land_Rover_Defender_2.jpg', 'Original mechanical hero'),

-- Jeep models
('Wrangler', (SELECT id FROM makes WHERE name = 'Jeep'), '2018-2024', 'JL', 'Global', 'Jeep_Wrangler_JL.jpg', 'The ultimate off-road freedom'),
('Wrangler', (SELECT id FROM makes WHERE name = 'Jeep'), '2007-2018', 'JK', 'Global', 'Jeep_Wrangler.jpg', 'Ultimate off-road freedom'),
('Grand Cherokee', (SELECT id FROM makes WHERE name = 'Jeep'), '2011-2021', 'WK2', 'Global', 'Jeep_Grand_Cherokee.jpg', 'Luxury off-road capability'),
('Cherokee', (SELECT id FROM makes WHERE name = 'Jeep'), '2014-2024', 'KL', 'Global', 'Jeep_Cherokee.jpg', 'Stylish compact SUV'),
('Renegade', (SELECT id FROM makes WHERE name = 'Jeep'), '2015-2024', 'BU', 'Global', 'Jeep_Renegade.jpg', 'Subcompact urban adventurer'),

-- Mini models
('Cooper', (SELECT id FROM makes WHERE name = 'Mini'), '2001-2006', 'R50', 'Global', 'Mini_Cooper_R50.jpg', 'Retro hatchback fun'),
('Cooper S', (SELECT id FROM makes WHERE name = 'Mini'), '2002-2006', 'R53', 'Global', 'Mini_Cooper_S_R53.jpg', 'Supercharged pocket rocket'),
('Clubman', (SELECT id FROM makes WHERE name = 'Mini'), '2015-2023', 'first', 'Global', 'Mini_Clubman.jpg', 'Quirky six-door hatch'),
('Cooper', (SELECT id FROM makes WHERE name = 'Mini'), '2014-present', 'F56', 'Global', 'Mini_Cooper.jpg', 'Go-kart handling fun'),
('Countryman', (SELECT id FROM makes WHERE name = 'Mini'), '2010-2016', 'R60', 'Global', 'Mini_Countryman_R60.jpg', 'The original Mini SUV'),
('Paceman', (SELECT id FROM makes WHERE name = 'Mini'), '2013-2016', 'R61', 'Global', 'Mini_Paceman_R61.jpg', 'Coupe-style crossover');