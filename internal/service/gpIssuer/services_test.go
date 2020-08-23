package gpIssuer

import (
	"massimple.com/wallet-controller/internal/models"
	"massimple.com/wallet-controller/internal/utils"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var baseAccount models.GpNewAccountInput
var id string
const emptyLegalAccount = "33436918"
func setup() {
	// We pick up the env setup
	utils.EnvInit("TEST")
	// we now connect
	GPInit() // todo no debería estar usando este método
	// we generate an account to use as test
	rand.Seed(time.Now().UnixNano())
	id = strconv.Itoa(20000000 + rand.Intn(30000000))
	baseAccount = models.GpNewAccountInput{
		Name:           "Emilio",
		Lastname:       "Basualdo Cibils",
		DocumentNumber: id,
		BirthDate:      "1997-06-16",
		Cellphone:      "005491133071114",
		ExternalId:     models.ID(id),
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestAltaDeCuenta_creates(t *testing.T) {
	// SETUP
	// EXERCISE
	acc, err := AltaDeCuenta(baseAccount)
	// ASSERT
	t.Run("Creates correct account", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want %s", err, "nil")
		}
		if acc.ExternalId != baseAccount.ExternalId {
			t.Errorf("got %s, want %s", acc.ExternalId, baseAccount.ExternalId)
		}
	})
}

func TestCargaDeTarjeta_returnsNoSuchAccount(t *testing.T) {
	// SETUP
	carga := models.GpRecharge{
		Amount: 100,
	}
	// EXERCISE
	err := CargaDeTarjeta("12312312", carga)
	// ASSERT
	t.Run("No such account", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
	})
}

func TestCargaDeTarjeta_returnsOk(t *testing.T) {
	// SETUP
	carga := models.GpRecharge{
		Amount: 100,
	}
	// EXERCISE
	err := CargaDeTarjeta(emptyLegalAccount, carga)
	// ASSERT
	t.Run("No error", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want %s", err, "nil")
		}
	})
}

func TestConsultaDeMovimientos_returnsNoSuchAccount(t *testing.T) {
	// SETUP
	// EXERCISE
	_, err := ConsultaDeMovimientos("12312312", 0)
	// ASSERT
	t.Run("No such account", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
	})
}

func TestConsultaDeMovimientos_returnsEmpty(t *testing.T) {
	// SETUP
	// EXERCISE
	resp, err := ConsultaDeMovimientos(emptyLegalAccount, 0)
	// ASSERT
	t.Run("Some movements", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want %s", err, "nil")
		}
		if resp.DateFrom.IsZero() {
			t.Errorf("got %t, want %t", resp.DateFrom.IsZero(), false)
		}
		if len(resp.Movements) == 0 {
			t.Errorf("got %d, want %s", len(resp.Movements), "non-zero")
		}
	})
}

func TestConsultaDeDisponileYSaldo_returnsNoSuchAccount(t *testing.T) {
	// SETUP
	// EXERCISE
	_, err := ConsultaDeDisponileYSaldo("12312312")
	// ASSERT
	t.Run("No such account", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
	})
}

func TestConsultaDeDisponileYSaldo_returnsNonZero(t *testing.T) {
	// SETUP
	// EXERCISE
	available, err := ConsultaDeDisponileYSaldo(emptyLegalAccount)
	// ASSERT
	t.Run("Returns non-zero", func(t *testing.T) {
		if err != nil {
			t.Errorf("got %s, want %s", err, "nil")
		}
		if available.LocalBalance != 0 {
			t.Errorf("got %f, want %f", available.LocalBalance, 0.0)
		}
		if available.LocalAvailableBuy == 0 {
			t.Errorf("got %f, want %s", available.LocalBalance, "non-zero")
		}
	})
}

func TestStillValidTokenRefreshed_tokenIsRefreshed(t *testing.T) {
	// SETUP
	oldToken := auth.AccessToken
	oldExpiry := auth.ExpiresAt
	// we change the token expiration time to force a refresh
	// The code is made so that the token is refreshed of the token will
	// expire in the next minute. So a token expiring in the next 30 seconds
	// must be refreshed.
	auth.ExpiresAt = time.Now().Add(time.Second * time.Duration(30))
	// EXERCISE
	_, err := ConsultaDeMovimientos("12312312", 0)
	// ASSERT
	t.Run("Token changed", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
		if auth.AccessToken == oldToken {
			t.Errorf("token seems the same")
		}
		if auth.ExpiresAt == oldExpiry {
			t.Errorf("Expiry seems the same")
		}
	})
}

func TestInvalidTokenRefreshed_tokenIsRefreshed(t *testing.T) {
	// SETUP
	oldToken := auth.AccessToken
	oldExpiry := auth.ExpiresAt
	// we change the token expiration time to force a refresh
	// The code is made so that the token is refreshed of the token will
	// expire in the next minute. So a token expired an hour ago
	// must be refreshed.
	auth.ExpiresAt = time.Now().Add(- time.Hour)
	// EXERCISE
	_, err := ConsultaDeMovimientos("12312312", 0)
	// ASSERT
	t.Run("Token changed", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
		if auth.AccessToken == oldToken {
			t.Errorf("token seems the same")
		}
		if auth.ExpiresAt == oldExpiry {
			t.Errorf("Expiry seems the same")
		}
	})
}

func TestValidTokenNotRefreshed_tokenIsNotRefreshed(t *testing.T) {
	// SETUP
	oldToken := auth.AccessToken
	oldExpiry := auth.ExpiresAt
	// If the token is still valid, it should not be changed
	// EXERCISE
	_, err := ConsultaDeMovimientos("12312312", 0)
	// ASSERT
	t.Run("Token not changed", func(t *testing.T) {
		if gpErr, ok := err.(*GPError); !ok && gpErr.Code != 3 {
			t.Errorf("got %s, want %s", err, &GPError{Code: 3})
		}
		if auth.AccessToken != oldToken {
			t.Errorf("token changed")
		}
		if auth.ExpiresAt != oldExpiry {
			t.Errorf("Expiry changed")
		}
	})
}
