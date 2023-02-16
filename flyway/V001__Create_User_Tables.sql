CREATE TABLE IF NOT EXISTS users (
	user_id TEXT,
	user_name VARCHAR(20),
	user_mail TEXT,
	user_type INT2,
	main_country    VARCHAR(2),
	main_currency   VARCHAR(3),
	main_language   VARCHAR(2),
	PRIMARY KEY(user_id)
);
COMMENT ON COLUMN users.user_type IS '0: Personal, 1: Company, 2: Pending for approval';
COMMENT ON COLUMN users.main_country IS 'ISO 3166-1 alphabetic_code_2, e.g, TW as Taiwan';
COMMENT ON COLUMN users.main_currency IS 'ISO 4217 alphabetic_code_3, e.g, TWD as New Taiwn dollar';
COMMENT ON COLUMN users.main_language IS 'ISO 639-1 alphabetic_code_2, e.g., ZH as Chinese';

CREATE TABLE IF NOT EXISTS user_currency (
	user_id TEXT,
	currency   VARCHAR(3)
);
COMMENT ON COLUMN user_currency.currency IS 'ISO 4217 alphabetic_code_3, e.g, TWD as New Taiwn dollar';

CREATE TABLE IF NOT EXISTS countries_iso3166_1 (
	alphabetic_code_2   VARCHAR(2),
	alphabetic_code_3   VARCHAR(3),
	numeric_code        INT2,
	country_name        TEXT,
	official_state_name TEXT,
	sovereignty         TEXT,
	top_domain          TEXT,
	PRIMARY KEY(alphabetic_code_2)
);
COMMENT ON TABLE countries_iso3166_1 IS 'ISO 3166-1 codes list, REF: https://datahub.io/core/country-list';

CREATE TABLE IF NOT EXISTS currencies_iso4217 (
	alphabetic_code VARCHAR(3),
	numeric_code    INT2,
	minor_unit      INT2,
	currency_name   TEXT,
	PRIMARY KEY(alphabetic_code)
);
COMMENT ON TABLE currencies_iso4217 IS 'ISO 4217 codes list, REF: https://datahub.io/core/currency-codes';

CREATE TABLE IF NOT EXISTS languages_iso639_1 (
	alphabetic_code VARCHAR(2),
	language_name   TEXT,
	PRIMARY KEY(alphabetic_code)
);
COMMENT ON TABLE languages_iso639_1 IS 'ISO 639-1 codes list, REF: https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes';

CREATE TABLE IF NOT EXISTS generated_reports_annual (
	report_year INT4,
    report_owner TEXT,
	report_type INT2,
	report_content JSONB,
	PRIMARY KEY (report_year, report_owner, report_type)
);
COMMENT ON COLUMN generated_reports_annual.report_type IS '0: balance_sheet, 1: income_statement, 3: Investment_sheet';
COMMENT ON COLUMN generated_reports_annual.report_content IS 'Applications should implement the corresponding parser for each report type';

CREATE TABLE IF NOT EXISTS generated_reports_seasonal (
	report_year INT4,
	report_season INT2,
	report_type INT2,
	report_owner TEXT,
	report_content JSONB,
	PRIMARY KEY (report_year, report_season, report_owner, report_type)
);
COMMENT ON COLUMN generated_reports_seasonal.report_season IS '1: [jan, mar], 2: [apr, jun], 3: [jul, sep], 4: [oct, dic]';
COMMENT ON COLUMN generated_reports_seasonal.report_type IS '0: balance_sheet, 1: income_statement, 3: Investment_sheet';
COMMENT ON COLUMN generated_reports_seasonal.report_content IS 'Applications should implement the corresponding parser for each report type';
