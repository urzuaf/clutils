package date_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/urzuaf/clutils/packages/go/date"
)

type TimezoneCase struct {
	Identifier   string `json:"identifier"`
	ExpectedIANA string `json:"expected_iana"`
	UTCTimestamp string `json:"utc_timestamp"`
	ExpectedHour int    `json:"expected_hour"`
	Description  string `json:"description"`
}

type TimezonesTestSuite struct {
	Timezones []TimezoneCase `json:"timezones"`
}

func loadTimezonesTestSuite(t *testing.T) TimezonesTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "date_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/date_formatting.json: %v", err)
	}

	var suite TimezonesTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de timezones: %v", err)
	}

	return suite
}

func TestTimezoneConstants(t *testing.T) {
	if date.TzSantiago != "America/Santiago" {
		t.Errorf("TzSantiago = %q; esperaba %q", date.TzSantiago, "America/Santiago")
	}
	if date.TzPuntaArenas != "America/Punta_Arenas" {
		t.Errorf("TzPuntaArenas = %q; esperaba %q", date.TzPuntaArenas, "America/Punta_Arenas")
	}
	if date.TzEaster != "Pacific/Easter" {
		t.Errorf("TzEaster = %q; esperaba %q", date.TzEaster, "Pacific/Easter")
	}
}

func TestTimezoneConversions(t *testing.T) {
	suite := loadTimezonesTestSuite(t)

	for _, tc := range suite.Timezones {
		t.Run(tc.Description, func(t *testing.T) {
			utcTime, err := time.Parse(time.RFC3339, tc.UTCTimestamp)
			if err != nil {
				t.Fatalf("error parseando timestamp UTC %q: %v", tc.UTCTimestamp, err)
			}

			var converted time.Time
			switch tc.Identifier {
			case "santiago":
				converted = date.ToChileTime(utcTime)
				if converted.Location().String() != tc.ExpectedIANA {
					t.Errorf("Location = %q; esperaba %q", converted.Location().String(), tc.ExpectedIANA)
				}
			case "punta_arenas":
				converted = date.ToMagallanesTime(utcTime)
				if converted.Location().String() != tc.ExpectedIANA {
					t.Errorf("Location = %q; esperaba %q", converted.Location().String(), tc.ExpectedIANA)
				}
			case "easter_island":
				converted = date.ToEasterIslandTime(utcTime)
				if converted.Location().String() != tc.ExpectedIANA {
					t.Errorf("Location = %q; esperaba %q", converted.Location().String(), tc.ExpectedIANA)
				}
			default:
				t.Fatalf("identificador de zona desconocido: %q", tc.Identifier)
			}

			if converted.Hour() != tc.ExpectedHour {
				t.Errorf("Hora convertida = %d; esperaba %d (time: %v)", converted.Hour(), tc.ExpectedHour, converted)
			}
		})
	}
}

func TestNowHelpers(t *testing.T) {
	nowSantiago := date.NowChile()
	if nowSantiago.Location().String() != date.TzSantiago {
		t.Errorf("NowChile().Location() = %q; esperaba %q", nowSantiago.Location().String(), date.TzSantiago)
	}

	nowPuntaArenas := date.NowMagallanes()
	if nowPuntaArenas.Location().String() != date.TzPuntaArenas {
		t.Errorf("NowMagallanes().Location() = %q; esperaba %q", nowPuntaArenas.Location().String(), date.TzPuntaArenas)
	}

	nowEaster := date.NowEasterIsland()
	if nowEaster.Location().String() != date.TzEaster {
		t.Errorf("NowEasterIsland().Location() = %q; esperaba %q", nowEaster.Location().String(), date.TzEaster)
	}
}
