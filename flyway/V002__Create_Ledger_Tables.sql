CREATE TABLE IF NOT EXISTS journal_book (
	transac_id INT8,
	user_id TEXT,
	orden INT4,
	transac_date DATE,
	debit_account TEXT,
	credit_account TEXT,
	amount NUMERIC(20, 4),
	currency VARCHAR(3),
	photo_addr TEXT,
	description VARCHAR(100),
	created_time TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(transac_id)
);
COMMENT ON COLUMN journal_book.transac_id IS 'Should be interpreted as in the frontend like TX00000001';
COMMENT ON COLUMN journal_book.orden IS 'To know how many transactions an user has made';
COMMENT ON COLUMN journal_book.transac_date IS 'transac date could be different from the creation date';
COMMENT ON COLUMN journal_book.currency IS 'ISO 4217 alphabetic_code_3, e.g, TWD as New Taiwn dollar';

CREATE TABLE IF NOT EXISTS ledger_balance (
	ledger_id TEXT,
	transac_id INT8,
	balance NUMERIC(20, 4),
	created_time TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(ledger_id, transac_id)
);
COMMENT ON TABLE ledger_balance IS 'Used to generate the trail balance for a given period, maintaine the current balance of each account';
COMMENT ON COLUMN ledger_balance.transac_id IS 'Used to generate income_statement';
COMMENT ON COLUMN ledger_balance.transac_id IS 'Used to generate balance_sheet';

CREATE TABLE IF NOT EXISTS user_ledgers (
	ledger_id TEXT,
	ledger_owner TEXT,
	ledger_name TEXT,
	ledger_type_id INT4, -- e.g., A111
	PRIMARY KEY(ledger_id)
);
COMMENT ON TABLE user_ledgers IS 'Defined by users, the app should provide a template for general types';
COMMENT ON COLUMN user_ledgers.ledger_id IS 'e.g., walker088_A111_01';
COMMENT ON COLUMN user_ledgers.ledger_owner IS 'e.g., walker088';
COMMENT ON COLUMN user_ledgers.ledger_name IS 'Under the ledger_type_name, e.g., Operating expenses, Cash, Tax, LivingExpense';
COMMENT ON COLUMN user_ledgers.ledger_type_id IS 'Related to ledger_types, e.g., A111';

CREATE TABLE IF NOT EXISTS user_stocks_us (
	ledger_id      TEXT,
	stock_id       TEXT,
	orden          INT4,
	operation      INT2,
	unit_price     NUMERIC(12, 4),
	amt_shares     INT4,
	description    TEXT,
	creation_date  TIMESTAMP WITH TIME ZONE,
	acc_num_shares INT4,
	avg_share_cost NUMERIC(12, 4),
	PRIMARY KEY(ledger_id, stock_id, orden)
);
COMMENT ON COLUMN user_stocks_us.ledger_id IS 'e.g., walker088_A111_01, Firstrade, IB, etc.';
COMMENT ON COLUMN user_stocks_us.stock_id IS 'e.g., TSLA';
COMMENT ON COLUMN user_stocks_us.operation IS '0: buy, 1: sell';

CREATE TABLE IF NOT EXISTS user_stocks_tw (
	ledger_id   TEXT,
	stock_id    TEXT,
	orden INT4,
	operation   INT2,
	unit_price  NUMERIC(12, 4),
	tx_fee NUMERIC(12, 4),
	tx_tax NUMERIC(12, 4),
	amt_shares  INT4,
	description TEXT,
	creation_date TIMESTAMP WITH TIME ZONE,
	acc_num_shares INT4,
	avg_share_cost NUMERIC(12, 4),
	PRIMARY KEY(ledger_id, stock_id, orden)
);
COMMENT ON COLUMN user_stocks_tw.ledger_id IS 'e.g., Capital, IB, etc.';
COMMENT ON COLUMN user_stocks_tw.stock_id IS 'e.g., TSMC';
COMMENT ON COLUMN user_stocks_tw.operation IS '0: buy, 1: sell';

CREATE TABLE IF NOT EXISTS ledger_types (
	ledger_type_id VARCHAR(4),
	first_grade    VARCHAR(1), 
	second_grade   VARCHAR(2),
	third_grade    VARCHAR(3),
	type_name TEXT,
	PRIMARY KEY (type_id)
);
COMMENT ON TABLE ledger_types IS 'To generate the balance sheet and the income statement sheet, predefined by the system to generate the accounting reports. REF[IFRS 16]: https://www.dgbas.gov.tw/News_Content.aspx?n=1961&s=17935';
COMMENT ON COLUMN ledger_types.first_grade IS 'Fourth grade, e.g., A111';
COMMENT ON COLUMN ledger_types.first_grade IS 'A: Assets, B: Liabilities, C: Owners equity, D: Self-Defined Operating income, E: Self-Defined expenses, F: Self-Defined Non‑operating income and expenses';
