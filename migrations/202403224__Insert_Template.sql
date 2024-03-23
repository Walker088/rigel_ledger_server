-- +goose Up
-- +goose StatementBegin
TRUNCATE ref_countries_iso3166_1;
INSERT INTO ref_countries_iso3166_1 (
	alphabetic_code_2, alphabetic_code_3, 
	numeric_code, country_name, 
	official_state_name, sovereignty, top_domain
)
VALUES 
('TW', 'TWN', 158, 'Taiwan', 'The Republic of China', 'Disputed', '.tw'),
('PY', 'PRY', 600, 'Paraguay', 'The Republic of Paraguay', 'UN member state', '.py'),
('AU', 'AUS', 36, 'Australia', 'The Commonwealth of Australia', 'UN member state', '.au'),
('US', 'USA', 840, 'United States of America (the)', 'The United States of America', 'UN member state', '.us'),
('TR', 'TUR', 792, 'Türkiye', 'The Republic of Türkiye', 'UN member state', '.tr');

TRUNCATE ref_currencies_iso4217;
INSERT INTO ref_currencies_iso4217 (
	alphabetic_code, numeric_code, minor_unit, currency_name
)
VALUES 
('TWD', 901, 2, 'New Taiwan dollar'),
('PYG', 600, 0, 'Paraguayan guaraní'),
('AUD', 36, 2, 'Australian dollar'),
('USD', 840, 2, 'United States dollar'),
('TRY', 949, 2, 'Turkish lira');

TRUNCATE user_info;
INSERT INTO user_info (
	user_id, user_id_gh, user_name, user_mail, user_type,
	main_country, main_language, main_currency
)
VALUES (
	'walker088',
	'19402938',
	'CHUN WEI CHEN',
	'cwwalker088@gmail.com',
	0,
	'TW',
	'EN',
	'USD'
);

TRUNCATE ref_exchange_rates;
INSERT INTO ref_exchange_rates (from_code, to_code, exchange_rate)
VALUES
('USD', 'TWD', 30.508798),
('USD', 'PYG', 7339.425186),
('USD', 'AUD', 1.459748),
('USD', 'TRY', 18.962906);

TRUNCATE user_ledgers;
INSERT INTO user_ledgers (ledger_id, ledger_owner, ledger_name, ledger_type_id, currency, balance)
VALUES
-- Assets
('walker088_1111_01', 'walker088', 'USD Cash', '1111', 'USD', 2635),
('walker088_1111_02', 'walker088', 'PYG Cash', '1111', 'PYG', 2598000),
('walker088_1111_03', 'walker088', 'TWD Cash', '1111', 'TWD', 3520),
('walker088_1111_04', 'walker088', 'TRY Cash', '1111', 'TRY', 0),

('walker088_1113_01', 'walker088', 'MegaBank USD Cash', '1113', 'USD', 4582.43),
('walker088_1113_02', 'walker088', 'MegaBank TWD Cash', '1113', 'TWD', 25100),
('walker088_1113_03', 'walker088', 'CathayBank TWD Cash', '1113', 'TWD', 6352),
('walker088_1113_04', 'walker088', 'BancoContiental USD Cash', '1113', 'USD', 1025.24),
('walker088_1113_05', 'walker088', 'BancoContiental PYG Cash', '1113', 'PYG', 4587550),
('walker088_1113_06', 'walker088', 'Wise Cash USD', '1113', 'USD', 150),
('walker088_1113_07', 'walker088', 'Wise Cash AUD', '1113', 'AUD', 850),
('walker088_1114_01', 'walker088', 'MegaBank USD Deposit', '1114', 'USD', 24589),
('walker088_1131_01', 'walker088', 'Recevable USD', '1131', 'USD', 10000),
('walker088_1131_02', 'walker088', 'Recevable PYG', '1131', 'PYG', 20550000),
('walker088_1131_03', 'walker088', 'Recevable TWD', '1131', 'TWD', 1547),
('walker088_1139_01', 'walker088', 'Uncollectible Recevable USD', '1139', 'USD', 0),
('walker088_1139_02', 'walker088', 'Uncollectible Recevable PYG', '1139', 'PYG', 0),
('walker088_1139_03', 'walker088', 'Uncollectible Recevable TWD', '1139', 'TWD', 0),

('walker088_1314_01', 'walker088', 'TW_Gov labor insurance', '1314', 'TWD', 0),
('walker088_1321_01', 'walker088', 'Firstrade.US stock Investment', '1321', 'USD', 26000),
('walker088_1321_02', 'walker088', 'Capital.TW stock Investment', '1321', 'TWD', 1000000),
('walker088_1345_01', 'walker088', 'NanShanInsur. 多美福 2020_2025', '1345', 'USD', 16000),
('walker088_1441_01', 'walker088', 'Vehicle', '1441', 'USD', 26000),

-- Liabilities
('walker088_2112_01', 'walker088', 'CathayBank Credit Card', '2112', 'TWD', 0),
('walker088_2131_01', 'walker088', 'Payable USD', '2131', 'USD', 0),
('walker088_2131_02', 'walker088', 'Payable PYG', '2131', 'PYG', 0),
('walker088_2131_03', 'walker088', 'Payable TWD', '2131', 'TWD', 0),
('walker088_2131_04', 'walker088', 'Payable AUD', '2131', 'AUD', 0),

-- Owner's equity
('walker088_3111_01', 'walker088', 'Initail Capital', '3111', 'USD', 110981.67),
('walker088_3111_02', 'walker088', 'Initail Capital', '3111', 'PYG', 28023550),
('walker088_3111_03', 'walker088', 'Initail Capital', '3111', 'TWD', 1036519),
('walker088_3111_04', 'walker088', 'Initail Capital', '3111', 'AUD', 850),
('walker088_3111_05', 'walker088', 'Initail Capital', '3111', 'TRY', 0),

('walker088_3351_01', 'walker088', 'Accumulated profit or loss', '3351', 'USD', 0),
('walker088_3351_02', 'walker088', 'Accumulated profit or loss', '3351', 'PYG', 0),
('walker088_3351_03', 'walker088', 'Accumulated profit or loss', '3351', 'TWD', 0),
('walker088_3351_04', 'walker088', 'Accumulated profit or loss', '3351', 'AUD', 0),
('walker088_3351_05', 'walker088', 'Accumulated profit or loss', '3351', 'TRY', 0),

-- Revenue
('walker088_A111_01', 'walker088', 'Base Salary (ICDF)', 'A111', 'TWD', 0),
('walker088_A111_02', 'walker088', 'Hardship allowance (ICDF)', 'A111', 'USD', 0),
('walker088_A111_03', 'walker088', 'Housing allowance (ICDF)', 'A111', 'USD', 0),
('walker088_A111_04', 'walker088', 'Other allowance (ICDF)', 'A111', 'USD', 0),

('walker088_A211_01', 'walker088', 'Interest revenue/income USD', 'A211', 'USD', 0),
('walker088_A211_02', 'walker088', 'Interest revenue/income TWD', 'A211', 'TWD', 0),
('walker088_A211_03', 'walker088', 'Interest revenue/income PYG', 'A211', 'PYG', 0),
('walker088_A211_04', 'walker088', 'Firstrade.US stock dividend', 'A211', 'USD', 0),
('walker088_A211_05', 'walker088', 'Capital.TW stock dividend', 'A211', 'TWD', 0),
('walker088_A212_01', 'walker088', 'Firstrade.US stock gain on disposal', 'A212', 'USD', 0),
('walker088_A212_02', 'walker088', 'Capital.TW stock gain on disposal', 'A212', 'TWD', 0),
('walker088_A213_01', 'walker088', 'Gain on disposal of assets USD', 'A213', 'USD', 0),
('walker088_A215_01', 'walker088', 'Firstrade.US stock gain on valuation', 'A215', 'USD', 0),
('walker088_A215_02', 'walker088', 'Capital.TW stock gain on valuation', 'A215', 'TWD', 0),

('walker088_A812_01', 'walker088', 'Revenue from sale of scraps PYG', 'A812', 'PYG', 0),
('walker088_A813_01', 'walker088', 'Gain on cash inventory surplus PYG', 'A813', 'PYG', 0),
('walker088_A814_01', 'walker088', 'Gain on reversal of bad debts USD', 'A814', 'USD', 0),
('walker088_A814_02', 'walker088', 'Gain on reversal of bad debts PYG', 'A814', 'PYG', 0),
('walker088_A814_03', 'walker088', 'Gain on reversal of bad debts TWD', 'A814', 'TWD', 0),
('walker088_A814_04', 'walker088', 'Gain on reversal of bad debts AUD', 'A814', 'AUD', 0),

('walker088_A815_01', 'walker088', 'Other income', 'A815', 'USD', 0),
('walker088_A815_02', 'walker088', 'Other income', 'A815', 'PYG', 0),
('walker088_A815_03', 'walker088', 'Other income', 'A815', 'TWD', 0),
('walker088_A815_04', 'walker088', 'Other income', 'A815', 'AUD', 0),

-- Expenses
('walker088_B111_01', 'walker088', 'Meal USD', 'B111', 'USD', 0),
('walker088_B111_02', 'walker088', 'Meal PYG', 'B111', 'PYG', 0),
('walker088_B111_03', 'walker088', 'Meal TWD', 'B111', 'TWD', 0),
('walker088_B111_04', 'walker088', 'Meal AUD', 'B111', 'AUD', 0),

('walker088_B112_01', 'walker088', 'Transportation USD', 'B112', 'USD', 0),
('walker088_B112_02', 'walker088', 'Transportation PYG', 'B112', 'PYG', 0),
('walker088_B112_03', 'walker088', 'Transportation TWD', 'B112', 'TWD', 0),
('walker088_B112_04', 'walker088', 'Transportation AUD', 'B112', 'AUD', 0),

('walker088_B113_01', 'walker088', 'Education USD', 'B113', 'USD', 0),
('walker088_B113_02', 'walker088', 'Education PYG', 'B113', 'PYG', 0),
('walker088_B113_03', 'walker088', 'Education TWD', 'B113', 'TWD', 0),
('walker088_B113_04', 'walker088', 'Education AUD', 'B113', 'AUD', 0),

('walker088_B114_01', 'walker088', 'Health/medical USD', 'B114', 'USD', 0),
('walker088_B114_02', 'walker088', 'Health/medical PYG', 'B114', 'PYG', 0),
('walker088_B114_03', 'walker088', 'Health/medical TWD', 'B114', 'TWD', 0),
('walker088_B114_04', 'walker088', 'Health/medical AUD', 'B114', 'AUD', 0),

('walker088_B115_01', 'walker088', 'Beauty USD', 'B115', 'USD', 0),
('walker088_B115_02', 'walker088', 'Beauty PYG', 'B115', 'PYG', 0),
('walker088_B115_03', 'walker088', 'Beauty TWD', 'B115', 'TWD', 0),
('walker088_B115_04', 'walker088', 'Beauty AUD', 'B115', 'AUD', 0),

('walker088_B116_01', 'walker088', 'Rent USD', 'B116', 'USD', 0),
('walker088_B116_02', 'walker088', 'Rent PYG', 'B116', 'PYG', 0),
('walker088_B116_03', 'walker088', 'Rent TWD', 'B116', 'TWD', 0),
('walker088_B116_04', 'walker088', 'Rent AUD', 'B116', 'AUD', 0),

('walker088_B117_01', 'walker088', 'Commission USD', 'B117', 'USD', 0),
('walker088_B117_02', 'walker088', 'Commission PYG', 'B117', 'PYG', 0),
('walker088_B117_03', 'walker088', 'Commission TWD', 'B117', 'TWD', 0),
('walker088_B117_04', 'walker088', 'Commission AUD', 'B117', 'AUD', 0),

('walker088_B118_01', 'walker088', 'Repair/maintenance USD', 'B118', 'USD', 0),
('walker088_B118_02', 'walker088', 'Repair/maintenance PYG', 'B118', 'PYG', 0),
('walker088_B118_03', 'walker088', 'Repair/maintenance TWD', 'B118', 'TWD', 0),
('walker088_B118_04', 'walker088', 'Repair/maintenance AUD', 'B118', 'AUD', 0),

('walker088_B119_01', 'walker088', 'Utilities USD', 'B119', 'USD', 0),
('walker088_B119_02', 'walker088', 'Utilities PYG', 'B119', 'PYG', 0),
('walker088_B119_03', 'walker088', 'Utilities TWD', 'B119', 'TWD', 0),
('walker088_B119_04', 'walker088', 'Utilities AUD', 'B119', 'AUD', 0),

('walker088_B121_01', 'walker088', 'Insurance fee USD', 'B121', 'USD', 0),
('walker088_B121_02', 'walker088', 'Insurance fee PYG', 'B121', 'PYG', 0),
('walker088_B121_03', 'walker088', 'Insurance fee TWD', 'B121', 'TWD', 0),
('walker088_B121_04', 'walker088', 'Insurance fee AUD', 'B121', 'AUD', 0),

('walker088_B122_01', 'walker088', 'Supplies expense USD', 'B122', 'USD', 0),
('walker088_B122_02', 'walker088', 'Supplies expense PYG', 'B122', 'PYG', 0),
('walker088_B122_03', 'walker088', 'Supplies expense TWD', 'B122', 'TWD', 0),
('walker088_B122_04', 'walker088', 'Supplies expense AUD', 'B122', 'AUD', 0),

('walker088_B123_01', 'walker088', 'Taxes.US', 'B123', 'USD', 0),
('walker088_B123_02', 'walker088', 'Taxes.PY', 'B123', 'PYG', 0),
('walker088_B123_03', 'walker088', 'Taxes.TW', 'B123', 'TWD', 0),
('walker088_B123_04', 'walker088', 'Taxes.AU', 'B123', 'AUD', 0),

('walker088_B151_01', 'walker088', 'Apparel USD', 'B151', 'USD', 0),
('walker088_B151_02', 'walker088', 'Apparel PYG', 'B151', 'PYG', 0),
('walker088_B151_03', 'walker088', 'Apparel TWD', 'B151', 'TWD', 0),
('walker088_B151_04', 'walker088', 'Apparel AUD', 'B151', 'AUD', 0),

('walker088_B152_01', 'walker088', 'Travelling USD', 'B152', 'USD', 0),
('walker088_B152_02', 'walker088', 'Travelling PYG', 'B152', 'PYG', 0),
('walker088_B152_03', 'walker088', 'Travelling TWD', 'B152', 'TWD', 0),
('walker088_B152_04', 'walker088', 'Travelling AUD', 'B152', 'AUD', 0),

('walker088_B153_01', 'walker088', 'Gift USD', 'B153', 'USD', 0),
('walker088_B153_02', 'walker088', 'Gift PYG', 'B153', 'PYG', 0),
('walker088_B153_03', 'walker088', 'Gift TWD', 'B153', 'TWD', 0),
('walker088_B153_04', 'walker088', 'Gift AUD', 'B153', 'AUD', 0),

('walker088_B154_01', 'walker088', 'Entertainment/Social USD', 'B154', 'USD', 0),
('walker088_B154_02', 'walker088', 'Entertainment/Social PYG', 'B154', 'PYG', 0),
('walker088_B154_03', 'walker088', 'Entertainment/Social TWD', 'B154', 'TWD', 0),
('walker088_B154_04', 'walker088', 'Entertainment/Social AUD', 'B154', 'AUD', 0),

('walker088_B155_01', 'walker088', 'Online service subscription USD', 'B155', 'USD', 0),
('walker088_B155_02', 'walker088', 'Online service subscription PYG', 'B155', 'PYG', 0),
('walker088_B155_03', 'walker088', 'Online service subscription TWD', 'B155', 'TWD', 0),
('walker088_B155_04', 'walker088', 'Online service subscription AUD', 'B155', 'AUD', 0),
('walker088_B155_05', 'walker088', 'Online service subscription TRY', 'B155', 'TRY', 0),

('walker088_B156_01', 'walker088', 'Pets USD', 'B156', 'USD', 0),
('walker088_B156_02', 'walker088', 'Pets PYG', 'B156', 'PYG', 0),
('walker088_B156_03', 'walker088', 'Pets TWD', 'B156', 'TWD', 0),
('walker088_B156_04', 'walker088', 'Pets AUD', 'B156', 'AUD', 0),

('walker088_B211_01', 'walker088', 'Interest expense TWD', 'B211', 'TWD', 0),
('walker088_B212_01', 'walker088', 'Firstrade.US stock loss on disposal', 'B212', 'USD', 0),
('walker088_B212_02', 'walker088', 'Capital.TW stock loss on disposal', 'B212', 'TWD', 0),
('walker088_B213_01', 'walker088', 'Loss on disposal of assets USD', 'B213', 'USD', 0),
('walker088_B214_01', 'walker088', 'Firstrade.US stock loss on valuation', 'B214', 'USD', 0),
('walker088_B214_02', 'walker088', 'Capital.TW stock loss on valuation', 'B214', 'TWD', 0),

('walker088_B811_01', 'walker088', 'Donation USD', 'B811', 'USD', 0),
('walker088_B811_02', 'walker088', 'Donation PYG', 'B811', 'PYG', 0),
('walker088_B811_03', 'walker088', 'Donation TWD', 'B811', 'TWD', 0),
('walker088_B811_04', 'walker088', 'Donation AUD', 'B811', 'AUD', 0),

('walker088_B812_01', 'walker088', 'Loss on uncollectible accounts USD', 'B812', 'USD', 0),
('walker088_B812_02', 'walker088', 'Loss on uncollectible accounts PYG', 'B812', 'PYG', 0),
('walker088_B812_03', 'walker088', 'Loss on uncollectible accounts TWD', 'B812', 'TWD', 0),
('walker088_B812_04', 'walker088', 'Loss on uncollectible accounts AUD', 'B812', 'AUD', 0),

('walker088_B818_01', 'walker088', 'Other expense USD', 'B818', 'USD', 0),
('walker088_B818_02', 'walker088', 'Other expense PYG', 'B818', 'PYG', 0),
('walker088_B818_03', 'walker088', 'Other expense TWD', 'B818', 'TWD', 0),
('walker088_B818_04', 'walker088', 'Other expense AUD', 'B818', 'AUD', 0);

TRUNCATE user_ledger_journal;
TRUNCATE user_ledger_transactions;
UPDATE user_ledgers
	SET balance = 0
	WHERE 
		ledger_owner = 'walker088'
		AND ledger_type_id = '3351';

-- 1. Income
INSERT INTO user_ledger_journal (user_id, transac_id, transac_date, credit_account, debit_account, amount, description)
SELECT
	'walker088', 
	COUNT(1) + 1,
	CURRENT_DATE, 
	'walker088_1113_02', -- [Asset]   MegaBank TWD Cash
	'walker088_A111_01', -- [Revenue] Base Salary (ICDF)
	60000, 
	'ICDF Salary Feb 2023'
FROM
	user_ledger_journal
WHERE user_id = 'walker088'
RETURNING *;

INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
SELECT
	ledger_id,
	1,
	balance + 60000
FROM
	user_ledgers
WHERE
	ledger_id = 'walker088_1113_02' OR ledger_id = 'walker088_A111_01';

UPDATE user_ledgers
SET balance = balance + 60000
WHERE
	ledger_id = (SELECT ledger_id FROM user_ledgers WHERE ledger_owner = 'walker088' AND ledger_type_id = '3351' AND currency = 'TWD');

INSERT INTO user_ledger_journal (user_id, transac_id, transac_date, credit_account, debit_account, amount, description)
SELECT
	'walker088', 
	COUNT(1) + 1,
	CURRENT_DATE, 
	'walker088_1113_01', -- [Asset]   MegaBank USD Cash
	'walker088_A111_02', -- [Revenue] Hardship allowance (ICDF)
	4100, 
	'ICDF Hardship allowance Feb 2023'
FROM
	user_ledger_journal
WHERE user_id = 'walker088';

INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
SELECT
	ledger_id,
	2,
	balance + 4100
FROM
	user_ledgers
WHERE
	ledger_id = 'walker088_1113_01' OR ledger_id = 'walker088_A111_02';

UPDATE user_ledgers
SET balance = balance + 4100
WHERE 
	ledger_id = (SELECT ledger_id FROM user_ledgers WHERE ledger_owner = 'walker088' AND ledger_type_id = '3351' AND currency = 'USD');

-- Expense
INSERT INTO user_ledger_journal (user_id, transac_id, transac_date, credit_account, debit_account, amount, description)
SELECT
	'walker088', 
	COUNT(1) + 1,
	CURRENT_DATE, 
	'walker088_B111_03', -- [Expense]   Meal TWD
	'walker088_2112_01', -- [Liability] CathayBank Credit Card
	215, 
	'Bellini, Lunch'
FROM
	user_ledger_journal
WHERE user_id = 'walker088'
RETURNING *;

INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
SELECT
	ledger_id,
	3,
	balance + 215
FROM
	user_ledgers
WHERE
	ledger_id = 'walker088_B111_03' OR ledger_id = 'walker088_2112_01';

UPDATE user_ledgers
SET balance = balance - 215
WHERE ledger_id = (SELECT ledger_id FROM user_ledgers WHERE ledger_owner = 'walker088' AND ledger_type_id = '3351' AND currency = 'TWD');

-- Asset/Liability/Owner's equity
INSERT INTO user_ledger_journal (user_id, transac_id, transac_date, credit_account, debit_account, amount, description)
SELECT
	'walker088', 
	COUNT(1) + 1,
	CURRENT_DATE, 
	'walker088_2112_01', -- [Liability] CathayBank Credit Card
	'walker088_1113_03', -- [Asset]     CathayBank TWD Cash
	215, 
	'Pay Credit Card Bill, Feb 2023'
FROM
	user_ledger_journal
WHERE user_id = 'walker088'
RETURNING *;

INSERT INTO user_ledger_transactions (ledger_id, transac_id, balance)
SELECT
	ledger_id,
	4,
	balance - 215
FROM
	user_ledgers
WHERE
	ledger_id = 'walker088_2112_01' OR ledger_id = 'walker088_1113_03';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE ref_countries_iso3166_1;
TRUNCATE ref_currencies_iso4217;
TRUNCATE user_info;
TRUNCATE ref_exchange_rates;
TRUNCATE user_ledgers;
TRUNCATE user_ledger_journal;
TRUNCATE user_ledger_transactions;
-- +goose StatementEnd