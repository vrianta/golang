package structparser

import (
	"os"
	"testing"
)

// ----------------------------------------------------------------------

// TestConfig is a comprehensive struct designed to test all tag behaviors.
type TestConfig struct {
	Host         string `env:"TEST_HOST" default:"localhost"`
	Port         string `env:"TEST_PORT"`
	ApiKey       string `env:"TEST_KEY" required:"true"`
	AppMode      string `default:"production"`
	Debug        bool   `arg:"debug" default:"false"`
	Retries      int    `arg:"retries" default:"3"`
	RequiredTest string `required:"true"`
}

// PrivateConfig is used to test the unexported/private field error rule.
type PrivateConfig struct {
	secret string `env:"SECRET_KEY"`
}

// ----------------------------------------------------------------------
// Test Suite 1: Core Functionality (Strongly Typed to *TestConfig)
// ----------------------------------------------------------------------

func TestParse_CoreLogic(t *testing.T) {
	// Notice how input is now strictly *TestConfig, enabling perfect compiler inference!
	tests := []struct {
		name          string
		input         *TestConfig
		setupEnv      map[string]string
		setupArgs     []string
		expectedCheck func(t *testing.T, cfg *TestConfig) // Strongly typed callback!
		wantErr       bool
		expectedErr   string
	}{
		{
			name: "Error: Required field is missing (zero value)",
			input: &TestConfig{
				RequiredTest: "t",
			},
			setupEnv: map[string]string{
				"TEST_HOST": "127.0.0.1",
			},
			wantErr:     true,
			expectedErr: "field ApiKey is required but has no value",
		},
		{
			name: "Success: Fallback to default when env is missing",
			input: &TestConfig{
				RequiredTest: "t",
			},
			setupEnv: map[string]string{
				"TEST_KEY": "secure-token-123", // satisfies the required constraint
			},
			wantErr: false,
			expectedCheck: func(t *testing.T, cfg *TestConfig) {
				if cfg.Host != "localhost" {
					t.Errorf("expected Host to fall back to default 'localhost', got %q", cfg.Host)
				}
				if cfg.AppMode != "production" {
					t.Errorf("expected AppMode default 'production', got %q", cfg.AppMode)
				}
				if cfg.ApiKey != "secure-token-123" {
					t.Errorf("expected ApiKey to match env string, got %q", cfg.ApiKey)
				}
			},
		},
		{
			name: "Success: Env takes priority over default",
			input: &TestConfig{
				RequiredTest: "t",
			},
			setupEnv: map[string]string{
				"TEST_HOST": "custom-domain.com",
				"TEST_KEY":  "token-xyz",
			},
			wantErr: false,
			expectedCheck: func(t *testing.T, cfg *TestConfig) {
				if cfg.Host != "custom-domain.com" {
					t.Errorf("expected Host to override default and be 'custom-domain.com', got %q", cfg.Host)
				}
				if cfg.ApiKey != "token-xyz" {
					t.Errorf("expected ApiKey to be 'token-xyz', got %q", cfg.ApiKey)
				}
			},
		},
		{
			name: "Success: CLI arg sets bool field",
			input: &TestConfig{
				RequiredTest: "t",
			},
			setupEnv: map[string]string{
				"TEST_KEY": "secure-token-123",
			},
			setupArgs: []string{"cmd", "--debug"},
			wantErr:   false,
			expectedCheck: func(t *testing.T, cfg *TestConfig) {
				if !cfg.Debug {
					t.Errorf("expected Debug to be true when --debug is provided")
				}
				if cfg.AppMode != "production" {
					t.Errorf("expected AppMode default 'production', got %q", cfg.AppMode)
				}
			},
		},
		{
			name: "Success: CLI arg overrides default value",
			input: &TestConfig{
				RequiredTest: "t",
			},
			setupEnv: map[string]string{
				"TEST_KEY": "secure-token-123",
			},
			setupArgs: []string{"cmd", "--retries=5"},
			wantErr:   false,
			expectedCheck: func(t *testing.T, cfg *TestConfig) {
				if cfg.Retries != 5 {
					t.Errorf("expected Retries to be 5 when --retries=5 is provided, got %d", cfg.Retries)
				}
			},
		},
		{
			name:        "Error: Required field data is missing entirely",
			input:       &TestConfig{}, // ApiKey & RequiredTest left at zero-value
			setupEnv:    nil,           // Explicitly no environment variables set
			wantErr:     true,
			expectedErr: "field ApiKey is required but has no value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origArgs := os.Args
			if tt.setupArgs != nil {
				os.Args = tt.setupArgs
			}
			defer func() {
				os.Args = origArgs
				for k := range tt.setupEnv {
					os.Unsetenv(k)
				}
			}()

			// 1. Set environment variables
			for k, v := range tt.setupEnv {
				os.Setenv(k, v)
			}
			_, err := Parse(tt.input, false)

			// 3. Assert on expected errors
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err.Error() != tt.expectedErr {
				t.Errorf("Parse() error message = %q, expectedErr %q", err.Error(), tt.expectedErr)
			}

			// 4. Assert on struct states
			if !tt.wantErr && tt.expectedCheck != nil {
				tt.expectedCheck(t, tt.input)
			}
		})
	}
}

// ----------------------------------------------------------------------
// Test Suite 2: Edge Cases and Type Validation
// ----------------------------------------------------------------------

func TestParse_EdgeCases(t *testing.T) {
	t.Run("Error: Pass nil pointer", func(t *testing.T) {
		_, err := Parse((*TestConfig)(nil), false)

		expectedErr := "must pass a non-nil pointer to a struct"
		if err == nil || err.Error() != expectedErr {
			t.Errorf("expected error %q, got: %v", expectedErr, err)
		}
	})

	t.Run("Error: Pass pointer to non-struct (int)", func(t *testing.T) {
		_, err := Parse(new(int), false)

		expectedErr := "provided value is not a pointer to a struct"
		if err == nil || err.Error() != expectedErr {
			t.Errorf("expected error %q, got: %v", expectedErr, err)
		}
	})

	t.Run("Error: Struct contains unexported private fields", func(t *testing.T) {
		_, err := Parse(&PrivateConfig{}, false)

		expectedErr := "field secret is not public"
		if err == nil || err.Error() != expectedErr {
			t.Errorf("expected error %q, got: %v", expectedErr, err)
		}
	})
}
