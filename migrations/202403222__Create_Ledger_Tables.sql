-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ref_ledger_types (
	ledger_type_id VARCHAR(4),
	first_grade    VARCHAR(1) NOT NULL,
	second_grade   VARCHAR(2) NOT NULL,
	third_grade    VARCHAR(3) NOT NULL,
	type_name      TEXT NOT NULL,
	type_name_zh   TEXT,
	description_en TEXT,
	description_zh TEXT,
	PRIMARY KEY (ledger_type_id)
);
COMMENT ON TABLE ref_ledger_types IS 'To generate the balance sheet and the income statement sheet, predefined by the system to generate the accounting reports. REF[IFRS 16]: https://gcis.nat.gov.tw/mainNew/matterAction.do?method=showFile&fileNo=t70215_p';
COMMENT ON COLUMN ref_ledger_types.ledger_type_id IS 'Fourth grade, e.g., A111 Salary revenue';
COMMENT ON COLUMN ref_ledger_types.first_grade IS 'Ref: ref_ledger_first_grade';
COMMENT ON COLUMN ref_ledger_types.second_grade IS 'Ref: ref_ledger_second_grade';
COMMENT ON COLUMN ref_ledger_types.third_grade IS 'Ref: ref_ledger_third_grade';

CREATE TABLE IF NOT EXISTS ref_ledger_first_grade (
	first_grade    VARCHAR(1), 
	type_name      TEXT NOT NULL,
	type_name_zh   TEXT NOT NULL,
	description_en TEXT,
	description_zh TEXT,
	PRIMARY KEY (first_grade)
);

CREATE TABLE IF NOT EXISTS ref_ledger_second_grade (
	first_grade    VARCHAR(1) NOT NULL,
	second_grade   VARCHAR(2),
	type_name      TEXT NOT NULL,
	type_name_zh   TEXT,
	description_en TEXT,
	description_zh TEXT,
	PRIMARY KEY (second_grade)
);

CREATE TABLE IF NOT EXISTS ref_ledger_third_grade (
	first_grade    VARCHAR(1) NOT NULL,
	second_grade   VARCHAR(2) NOT NULL,
	third_grade    VARCHAR(3),
	type_name      TEXT NOT NULL,
	type_name_zh   TEXT,
	description_en TEXT,
	description_zh TEXT,
	PRIMARY KEY (third_grade)
);

CREATE TABLE IF NOT EXISTS user_ledgers (
	ledger_id       TEXT,
	ledger_owner    TEXT NOT NULL,
	ledger_name     TEXT NOT NULL,
	ledger_type_id  VARCHAR(4) NOT NULL, -- e.g., A111
	currency        VARCHAR(3) NOT NULL,
	balance         NUMERIC(20, 6),
	ledger_tags     TEXT[], -- List of string
	ledger_rules    JSONB,
	ledger_status   INT2, -- 0: Inactive 1: Active
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(ledger_id)
);
DROP TRIGGER IF EXISTS user_ledgers_updated_at ON user_ledgers;
CREATE TRIGGER user_ledgers_updated_at BEFORE UPDATE ON user_ledgers FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();
COMMENT ON TABLE user_ledgers IS 'Defined by users, where the balance sheet is generated from. The app should provide a template for general types';
COMMENT ON COLUMN user_ledgers.ledger_id IS 'e.g., walker088_A111_01';
COMMENT ON COLUMN user_ledgers.ledger_owner IS 'e.g., walker088';
COMMENT ON COLUMN user_ledgers.ledger_name IS 'Under the ledger_type_name, e.g., Operating expenses, Cash, Tax, LivingExpense';
COMMENT ON COLUMN user_ledgers.ledger_type_id IS 'Related to ledger_types, e.g., A111';

CREATE TABLE IF NOT EXISTS user_ledger_journal (
	user_id        TEXT,
	transac_id     INT8,
	transac_date   DATE NOT NULL,
	credit_account TEXT NOT NULL,
	debit_account  TEXT NOT NULL,
	amount         NUMERIC(20, 6) NOT NULL,
	photo_addr     TEXT[],
	description    VARCHAR(100),
	created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	PRIMARY KEY(transac_id, user_id)
);
COMMENT ON COLUMN user_ledger_journal.transac_id IS 'Should be interpreted as in the frontend like TX00000001';
COMMENT ON COLUMN user_ledger_journal.transac_date IS 'transac date could be different from the creation date';

CREATE OR REPLACE FUNCTION flush_new_balance_to_user_ledger()   
RETURNS TRIGGER AS $$
BEGIN
	UPDATE user_ledgers
	SET balance = NEW.balance
	WHERE
		ledger_id = NEW.ledger_id;
    RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TABLE IF NOT EXISTS user_ledger_transactions (
	ledger_id  TEXT,
	transac_id INT8,
	balance    NUMERIC(20, 6) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	PRIMARY KEY(ledger_id, transac_id)
);
DROP TRIGGER IF EXISTS flush_new_balance_to_user_ledger ON user_ledger_transactions;
CREATE TRIGGER flush_new_balance_to_user_ledger AFTER INSERT ON user_ledger_transactions FOR EACH ROW EXECUTE PROCEDURE flush_new_balance_to_user_ledger();
COMMENT ON TABLE user_ledger_transactions IS 'Used to generate the income statement and trail balance for a given period.';
COMMENT ON COLUMN user_ledger_transactions.balance IS 'The balance in the ledger currency';


CREATE TABLE IF NOT EXISTS user_stocks_us (
	ledger_id      TEXT,
	stock_id       TEXT,
	orden          INT4,
	operation      INT2,
	unit_price     NUMERIC(12, 6),
	amt_shares     INT4,
	description    TEXT,
	acc_num_shares INT4,
	avg_share_cost NUMERIC(12, 6),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	PRIMARY KEY(ledger_id, stock_id, orden)
);
COMMENT ON COLUMN user_stocks_us.ledger_id IS 'e.g., walker088_1321_01 Firstrade.US stock Investment';
COMMENT ON COLUMN user_stocks_us.stock_id IS 'e.g., TSLA';
COMMENT ON COLUMN user_stocks_us.operation IS '0: buy, 1: sell';

CREATE TABLE IF NOT EXISTS user_stocks_tw (
	ledger_id   TEXT,
	stock_id    TEXT,
	orden       INT4,
	operation   INT2,
	unit_price  NUMERIC(12, 6),
	tx_fee NUMERIC(12, 6),
	tx_tax NUMERIC(12, 6),
	amt_shares  INT4,
	description TEXT,
	acc_num_shares INT4,
	avg_share_cost NUMERIC(12, 6),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	PRIMARY KEY(ledger_id, stock_id, orden)
);
COMMENT ON COLUMN user_stocks_tw.ledger_id IS 'e.g., walker088_1321_02 Capital.TW stock Investment';
COMMENT ON COLUMN user_stocks_tw.stock_id IS 'e.g., TSMC';
COMMENT ON COLUMN user_stocks_tw.operation IS '0: buy, 1: sell';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ref_ledger_types;
DROP TABLE ref_ledger_first_grade;
DROP TABLE ref_ledger_second_grade;
DROP TABLE ref_ledger_third_grade;
DROP TABLE user_ledgers;
DROP TABLE user_ledger_journal;
DROP FUNCTION IF EXISTS flush_new_balance_to_user_ledger() CASCADE;
DROP TABLE user_ledger_transactions;
DROP TABLE user_stocks_us;
DROP TABLE user_stocks_tw;
-- +goose StatementEnd