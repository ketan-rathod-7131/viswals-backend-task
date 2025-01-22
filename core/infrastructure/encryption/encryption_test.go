package encryption_test

import (
	"testing"

	"github.com/viswals/core/infrastructure/encryption"
)

func TestEncryption(t *testing.T) {
	const key = "abcdefghabcdefghabcdefghabcdefgh"

	tests := []struct {
		name       string
		data       string
		shouldFail bool
	}{
		{name: "valid email", data: "ketan.rathod@bacancy.com", shouldFail: false},
		{name: "empty string", data: "", shouldFail: false},
		{name: "special characters", data: "!@#$%^&*()_+=-`~", shouldFail: false},
		{name: "long string", data: "this is a very long string to test the encryption and decryption mechanisms in detail", shouldFail: false},
	}

	em, err := encryption.New([]byte(key))
	if err != nil {
		t.Fatalf("failed to initialize encryption manager: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// encrypt the data
			encryptedData, err := em.Encrypt(tt.data)
			if (err != nil) != tt.shouldFail {
				t.Fatalf("encrypt() unexpected error: %v", err)
			}

			t.Logf("encrypted data for '%s': %s", tt.data, encryptedData)

			// decrypt the data if encryption succeeded
			if !tt.shouldFail {
				decryptedData, err := em.Decrypt(encryptedData)
				if err != nil {
					t.Fatalf("decrypt() failed: %v", err)
				}

				t.Logf("decrypted data: %s", decryptedData)

				// ensure the decrypted data matches the original
				if decryptedData != tt.data {
					t.Errorf("decrypted data does not match original. got: %s, want: %s", decryptedData, tt.data)
				}
			}
		})
	}
}
