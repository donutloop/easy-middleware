package easy_middleware

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSqlDb(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		if rv := r.Context().Value("db"); rv != nil {
			if _, ok := rv.(*sql.DB); !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	Open = func(driver, dsn string) (*sql.DB, error) {
		return nil, nil
	}
	Close = func(db *sql.DB) error {
		return nil
	}
	test := httptest.NewServer(SqlDb("test", "foo")(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL)
	if err != nil {
		t.Errorf("sql db middleware request: %s", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Error("Sql db middleware request: Unexpected bad request")
	}
}
