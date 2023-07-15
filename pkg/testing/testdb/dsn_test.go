package testdb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMySQL(t *testing.T) {
	expected := map[string]string{
		"user":   "bob",
		"host":   "1.2.3.4",
		"port":   "5432",
		"dbname": "mydb",
	}
	m, err := MySQLDSN("bob:secret@tcp(1.2.3.4:5432)/mydb")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expected, m); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}
}

func TestPostgres(t *testing.T) {
	expected := map[string]string{
		"user":    "bob",
		"host":    "1.2.3.4",
		"port":    "5432",
		"dbname":  "mydb",
		"sslmode": "verify-full",
	}
	m, err := PostgresDSN("postgres://bob:secret@1.2.3.4:5432/mydb?sslmode=verify-full")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expected, m); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}

	expected = map[string]string{
		"user":             "dog",
		"port":             "5433",
		"host":             "master-db-master-active.postgres.service.consul",
		"dbname":           "dogdatastaging",
		"application_name": "trace-api",
	}
	dsn := "password=zMWmQz26GORmgVVKEbEl dbname=dogdatastaging application_name=trace-api port=5433 host=master-db-master-active.postgres.service.consul user=dog"
	m, err = PostgresDSN(dsn)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(expected, m); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}
}
