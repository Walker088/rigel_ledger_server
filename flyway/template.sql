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

INSERT INTO ref_exchange_rates (from_code, to_code, exchange_rate)
VALUES
('USD', 'TWD', 30.508798),
('USD', 'PYG', 7339.425186),
('USD', 'AUD', 1.459748);

INSERT INTO user_ledgers (ledger_id, ledger_owner, ledger_name, ledger_type_id)
VALUES
-- Assets
('walker088_1111_01', 'walker088', 'USD Cash', '1111'),
('walker088_1111_02', 'walker088', 'PYG Cash', '1111'),
('walker088_1111_03', 'walker088', 'TWD Cash', '1111'),
('walker088_1113_01', 'walker088', 'MegaBank Cash', '1113'),
('walker088_1113_02', 'walker088', 'CathayBank Cash', '1113'),
('walker088_1113_03', 'walker088', 'BancoContiental Cash', '1113'),
('walker088_1113_04', 'walker088', 'Wise Cash', '1113'),
('walker088_1131_01', 'walker088', 'Recevable', '1131'),
('walker088_1139_01', 'walker088', 'Uncollectible Recevable', '1139'),
('walker088_1314_01', 'walker088', 'TW_Gov labor insurance', '1314'),
('walker088_1321_01', 'walker088', 'Firstrade.US stock Investment', '1321'),
('walker088_1321_02', 'walker088', 'Capital.TW stock Investment', '1321'),
('walker088_1345_01', 'walker088', 'NanShanInsur. 多美福 2020_2025', '1345'),
('walker088_1441_01', 'walker088', 'Vehicle', '1441'),

-- Liabilities
('walker088_2112_01', 'walker088', 'CathayBank Credit Card', '2112'),
('walker088_2131_01', 'walker088', 'Payable', '2131'),

-- Owner's equity
('walker088_3111_01', 'walker088', 'Initail Capital', '3111'),
('walker088_3351_01', 'walker088', 'Accumulated profit or loss', '3351'),
('walker088_3353_01', 'walker088', 'Net income or loss for current season', '3351'),

-- Revenue
('walker088_A111_01', 'walker088', 'Base Salary (ICDF)', 'A111'),
('walker088_A111_02', 'walker088', 'Hardship allowance (ICDF)', 'A111'),
('walker088_A111_03', 'walker088', 'Housing allowance (ICDF)', 'A111'),
('walker088_A111_04', 'walker088', 'Other allowance (ICDF)', 'A111'),

('walker088_A211_01', 'walker088', 'Interest revenue/income USD', 'A211'),
('walker088_A211_02', 'walker088', 'Interest revenue/income TWD', 'A211'),
('walker088_A211_03', 'walker088', 'Interest revenue/income PYG', 'A211'),
('walker088_A211_04', 'walker088', 'Firstrade.US stock dividend', 'A211'),
('walker088_A211_05', 'walker088', 'Capital.TW stock dividend', 'A211'),
('walker088_A212_01', 'walker088', 'Firstrade.US stock gain on disposal', 'A212'),
('walker088_A212_02', 'walker088', 'Capital.TW stock gain on disposal', 'A212'),
('walker088_A213_01', 'walker088', 'Gain on disposal of assets', 'A213'),
('walker088_A215_01', 'walker088', 'Firstrade.US stock gain on valuation', 'A215'),
('walker088_A215_02', 'walker088', 'Capital.TW stock gain on valuation', 'A215'),

('walker088_A811_01', 'walker088', 'Donation income', 'A811'),
('walker088_A812_01', 'walker088', 'Revenue from sale of scraps', 'A812'),
('walker088_A813_01', 'walker088', 'Gain on cash inventory surplus', 'A813'),
('walker088_A814_01', 'walker088', 'Gain on reversal of bad debts', 'A814'),

('walker088_A815_01', 'walker088', 'Other income', 'A815'),

-- Expenses
('walker088_B111_01', 'walker088', 'Meal', 'B111'),
('walker088_B112_01', 'walker088', 'Transportation', 'B112'),
('walker088_B113_01', 'walker088', 'Education', 'B113'),
('walker088_B114_01', 'walker088', 'Health/medical', 'B114'),
('walker088_B115_01', 'walker088', 'Beauty', 'B115'),
('walker088_B116_01', 'walker088', 'Rent', 'B116'),
('walker088_B117_01', 'walker088', 'Commission', 'B117'),
('walker088_B118_01', 'walker088', 'Repair/maintenance', 'B118'),
('walker088_B119_01', 'walker088', 'Utilities', 'B119'),
('walker088_B121_01', 'walker088', 'Insurance fee', 'B121'),
('walker088_B122_01', 'walker088', 'Supplies expense', 'B122'),
('walker088_B123_01', 'walker088', 'Taxes.TW', 'B123'),

('walker088_B151_01', 'walker088', 'Apparel', 'B151'),
('walker088_B152_01', 'walker088', 'Travelling', 'B152'),
('walker088_B153_01', 'walker088', 'Gift', 'B153'),
('walker088_B154_01', 'walker088', 'Entertainment/Social', 'B154'),
('walker088_B155_01', 'walker088', 'Online service subscription', 'B155'),
('walker088_B156_01', 'walker088', 'Pets', 'B156'),

('walker088_B211_01', 'walker088', 'Interest expense', 'B211'),
('walker088_B212_01', 'walker088', 'Firstrade.US stock loss on disposal', 'B212'),
('walker088_B212_02', 'walker088', 'Capital.TW stock loss on disposal', 'B212'),
('walker088_B213_01', 'walker088', 'Loss on disposal of assets', 'B213'),
('walker088_B214_01', 'walker088', 'Firstrade.US stock loss on valuation', 'B214'),
('walker088_B214_02', 'walker088', 'Capital.TW stock loss on valuation', 'B214'),

('walker088_B811_01', 'walker088', 'Donation', 'B811'),
('walker088_B812_01', 'walker088', 'Loss on uncollectible accounts', 'B812'),
('walker088_B818_01', 'walker088', 'Other expense', 'B818'),
