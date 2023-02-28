CREATE OR REPLACE FUNCTION flush_updated_at_time()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS user_info (
	user_id         TEXT,
	user_id_gh      TEXT,
	user_id_gg      TEXT,
	user_name       VARCHAR(20) NOT NULL,
	user_mail       TEXT NOT NULL,
	user_type       INT2 NOT NULL,
	main_country    VARCHAR(2) NOT NULL,
	main_language   VARCHAR(2) NOT NULL,
	main_currency   VARCHAR(3) NOT NULL,
	created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at      TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(user_id)
);
DROP TRIGGER IF EXISTS user_info_updated_at ON user_info;
CREATE TRIGGER user_info_updated_at BEFORE UPDATE ON user_info FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();
COMMENT ON COLUMN user_info.user_type IS '0: Personal, 1: Company, 2: Pending for approval';
COMMENT ON COLUMN user_info.main_country IS 'ISO 3166-1 alphabetic_code_2, e.g, TW as Taiwan';
COMMENT ON COLUMN user_info.main_language IS 'ISO 639-1 alphabetic_code_2, e.g., ZH as Chinese';
COMMENT ON COLUMN user_info.main_currency IS 'ISO 4217 alphabetic_code_3, e.g, TWD as New Taiwn dollar';

CREATE TABLE IF NOT EXISTS ref_countries_iso3166_1 (
	alphabetic_code_2   VARCHAR(2),
	alphabetic_code_3   VARCHAR(3),
	numeric_code        INT2,
	country_name        TEXT,
	official_state_name TEXT,
	sovereignty         TEXT,
	top_domain          TEXT,
	PRIMARY KEY(alphabetic_code_2)
);
COMMENT ON TABLE ref_countries_iso3166_1 IS 'ISO 3166-1 codes list, REF: https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes';

CREATE TABLE IF NOT EXISTS ref_currencies_iso4217 (
	alphabetic_code VARCHAR(3),
	numeric_code    INT2,
	minor_unit      INT2,
	currency_name   TEXT,
	PRIMARY KEY(alphabetic_code)
);
COMMENT ON TABLE ref_currencies_iso4217 IS 'ISO 4217 codes list, REF: https://en.wikipedia.org/wiki/ISO_4217';

CREATE TABLE IF NOT EXISTS ref_languages_iso639_1 (
	alphabetic_code VARCHAR(2),
	language_name   TEXT,
	language_name_en TEXT,
	PRIMARY KEY(alphabetic_code)
);
COMMENT ON TABLE ref_languages_iso639_1 IS 'ISO 639-1 codes list, REF: https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes';

CREATE TABLE IF NOT EXISTS ref_exchange_rates (
	from_code VARCHAR(3),
	to_code VARCHAR(10),
	exchange_rate NUMERIC(20, 6),
	updated_at TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(from_code, to_code)
);
DROP TRIGGER IF EXISTS ref_exchange_rates_updated_at ON ref_exchange_rates;
CREATE TRIGGER ref_exchange_rates_updated_at BEFORE UPDATE ON ref_exchange_rates FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();
COMMENT ON TABLE ref_exchange_rates IS 'USD based exchange rate table';

CREATE TABLE IF NOT EXISTS ref_stock_prices (
	stock_id   TEXT,
	currency   VARCHAR(3),
	updated_at TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY(stock_id)
);
DROP TRIGGER IF EXISTS ref_stock_prices_updated_at ON ref_stock_prices;
CREATE TRIGGER ref_stock_prices_updated_at BEFORE UPDATE ON ref_stock_prices FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();

CREATE TABLE IF NOT EXISTS generated_reports_annual (
	report_year INT4,
    report_owner TEXT,
	report_type INT2,
	report_content JSONB,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY (report_year, report_owner, report_type)
);
DROP TRIGGER IF EXISTS generated_reports_annual_updated_at ON generated_reports_annual;
CREATE TRIGGER generated_reports_annual_updated_at BEFORE UPDATE ON generated_reports_annual FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();
COMMENT ON COLUMN generated_reports_annual.report_type IS '0: balance_sheet, 1: income_statement, 3: Investment_sheet';
COMMENT ON COLUMN generated_reports_annual.report_content IS 'Applications should implement the corresponding parser for each report type';

CREATE TABLE IF NOT EXISTS generated_reports_seasonal (
	report_year INT4,
	report_season INT2,
	report_type INT2,
	report_owner TEXT,
	report_content JSONB,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE,
	PRIMARY KEY (report_year, report_season, report_owner, report_type)
);
DROP TRIGGER IF EXISTS generated_reports_seasonal_updated_at ON generated_reports_seasonal;
CREATE TRIGGER generated_reports_seasonal_updated_at BEFORE UPDATE ON generated_reports_seasonal FOR EACH ROW EXECUTE PROCEDURE flush_updated_at_time();
COMMENT ON COLUMN generated_reports_seasonal.report_season IS '1: [jan, mar], 2: [apr, jun], 3: [jul, sep], 4: [oct, dic]';
COMMENT ON COLUMN generated_reports_seasonal.report_type IS '0: balance_sheet, 1: income_statement, 3: Investment_sheet';
COMMENT ON COLUMN generated_reports_seasonal.report_content IS 'Applications should implement the corresponding parser for each report type';
