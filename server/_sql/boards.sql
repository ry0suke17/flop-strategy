INSERT INTO card_numbers VALUES
  ('3f4601ff-54ae-4559-b5cb-767dcdc50fd3', '2', 2),
  ('30870ff7-c527-4ba8-97bd-857fb3ce1fde', '3', 3),
  ('dadf950e-2495-4204-bcb0-33a8c4163955', '4', 4),
  ('e7855916-c326-49b6-8515-5693936fe190', '5', 5);

INSERT INTO card_suits VALUES
  ('873848b2-a175-492b-95dd-95fd76d6f66f', 'c', 1),
  ('b167bcd7-e8fd-46e6-8ae6-d2887f13a76d', 'd', 2),
  ('7b388222-6e4a-458e-863b-c55283109e5c', 'h', 3),
  ('ee189460-1a30-4219-88b6-c7c9c02716db', 's', 4);

INSERT INTO cards VALUES
  -- スペードのカード
  ('8b8488b3-97b0-4954-8bf1-346d24643a86', '3f4601ff-54ae-4559-b5cb-767dcdc50fd3', 'ee189460-1a30-4219-88b6-c7c9c02716db'), -- 2s
  ('ae5ae17a-f186-41ff-878d-865f017be6e3', '30870ff7-c527-4ba8-97bd-857fb3ce1fde', 'ee189460-1a30-4219-88b6-c7c9c02716db'), -- 3s
  ('8c1ee6fd-8620-4f0e-a7f8-ab034fc66941', 'dadf950e-2495-4204-bcb0-33a8c4163955', 'ee189460-1a30-4219-88b6-c7c9c02716db'), -- 4s
  ('285ed38d-8e23-4122-bed1-b6ed4ecb4cf3', 'e7855916-c326-49b6-8515-5693936fe190', 'ee189460-1a30-4219-88b6-c7c9c02716db'), -- 5s
  -- ハートのカード
  ('23ae5887-289d-4bc7-be74-f4e92817a7d4', '3f4601ff-54ae-4559-b5cb-767dcdc50fd3', '7b388222-6e4a-458e-863b-c55283109e5c'), -- 2h
  ('77dac44f-8f89-4145-9511-3a9241614316', '30870ff7-c527-4ba8-97bd-857fb3ce1fde', '7b388222-6e4a-458e-863b-c55283109e5c'), -- 3h
  ('fbabd3aa-0d2d-4980-b7c4-6f79eb8747a4', 'dadf950e-2495-4204-bcb0-33a8c4163955', '7b388222-6e4a-458e-863b-c55283109e5c'), -- 4h
  ('fcf36f9c-0ae8-40ca-9eb0-d160b1844294', 'e7855916-c326-49b6-8515-5693936fe190', '7b388222-6e4a-458e-863b-c55283109e5c'), -- 5h
  -- ダイヤのカード
  ('5f56a061-a79a-4cc1-be3d-8e7c2e9a9e3a', 'e7855916-c326-49b6-8515-5693936fe190', 'b167bcd7-e8fd-46e6-8ae6-d2887f13a76d'); -- 5d

INSERT INTO boards VALUES
  -- 5s, 4s, 3s
  ('6a1e28e1-1d84-446a-a7a6-70e34ffd0d1c', '285ed38d-8e23-4122-bed1-b6ed4ecb4cf3', '8c1ee6fd-8620-4f0e-a7f8-ab034fc66941', 'ae5ae17a-f186-41ff-878d-865f017be6e3'),
  -- 5s, 5h, 3s
  ('2748ac73-67fe-4e5a-a556-f629ad1f92c8', '285ed38d-8e23-4122-bed1-b6ed4ecb4cf3', 'fcf36f9c-0ae8-40ca-9eb0-d160b1844294', 'ae5ae17a-f186-41ff-878d-865f017be6e3'),
  -- 5s, 5h, 5d
  ('f982fe1f-94b7-4dfc-a620-796e4fdff28c', '285ed38d-8e23-4122-bed1-b6ed4ecb4cf3', 'fcf36f9c-0ae8-40ca-9eb0-d160b1844294', '5f56a061-a79a-4cc1-be3d-8e7c2e9a9e3a');