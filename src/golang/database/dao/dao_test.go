package dao

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

//var pool *database.PgPool

func TestGetByGithubId(t *testing.T) {
	//userInfo :=
	fmt.Println("helloooo")
}

/*
	func upsertTestUser(u *pojo.UserInfo) error {
		upsert := `INSERT INTO user_info (
			user_id, user_id_gh, user_name, user_mail, user_type,
			main_country, main_language, main_currency
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
		ON CONFLICT (user_id)
		DO UPDATE SET
			user_id_gh = EXCLUDED.user_id_gh,
			user_name = EXCLUDED.user_name,
			user_mail = EXCLUDED.user_mail,
			user_type = EXCLUDED.user_type,
			main_country = EXCLUDED.main_country,
			main_language = EXCLUDED.main_language,
			main_currency = EXCLUDED.main_currency;`
		_, err := pool.GetPool().Exec(
			context.Background(),
			upsert,
			u.Id, u.IdGh, u.Name, u.Mail, u.TypeCode,
			u.MainCountry, u.MainLanguage, u.MainCurrency,
		)
		return err
	}
*/
func cleanUpTestData() {
	fmt.Println("cleanUpTestData")
}

func TestMain(m *testing.M) {
	// BeforeTest
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println(basepath)

	/*
			pool, err := database.New(config.GetPgConfig())
			if err != nil {
				fmt.Printf("Init DB Conn Pool error: %s", err)
				os.Exit(1)
			}
			defer pool.ShutDownPool()


		testUser := pojo.UserInfo{
			Id:           "TEST",
			IdGh:         "19999999",
			Name:         "GO TESTER",
			Mail:         "gotester@walker088.tw",
			TypeCode:     0,
			MainCountry:  "TW",
			MainLanguage: "EN",
			MainCurrency: "USD",
		}
		if err := upsertTestUser(&testUser); err != nil {
			fmt.Printf("Failed to Create Testing User: %s", err)
			os.Exit(1)
		}
	*/
	// ExecTest
	code := m.Run()

	// Exit
	cleanUpTestData()
	os.Exit(code)
}
