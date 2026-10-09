package failover

import (
	"math"
	"strings"
	"testing"
	"time"
)

func f64(v float64) *float64 { return &v }

func TestOptionsValidate(t *testing.T) {
	tests := []struct {
		name        string
		opts        *Options
		wantErr     bool
		errContains []string
	}{
		{
			name:    "nil options",
			opts:    nil,
			wantErr: false,
		},
		{
			name:    "valid roundRobin",
			opts:    &Options{Strategy: RoundRobin, RetryableMethods: []string{"GET", "post"}, MaxRetries: 3},
			wantErr: false,
		},
		{
			name:    "empty strategy is allowed",
			opts:    &Options{},
			wantErr: false,
		},
		{
			name:        "invalid strategy",
			opts:        &Options{Strategy: "random"},
			wantErr:     true,
			errContains: []string{`invalid failover strategy "random", supported values are "none", "roundRobin" or an empty value`},
		},
		{
			name:        "invalid retryableMethods entry",
			opts:        &Options{Strategy: RoundRobin, RetryableMethods: []string{"G3T"}},
			wantErr:     true,
			errContains: []string{`invalid failover retryableMethods entry "G3T"`},
		},
		{
			name:        "negative maxRetries",
			opts:        &Options{Strategy: RoundRobin, MaxRetries: -1},
			wantErr:     true,
			errContains: []string{"failover maxRetries must be >= 0, got -1"},
		},
		{
			name:        "negative initialInterval",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{InitialInterval: -time.Second}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval must be >= 0, got -1s"},
		},
		{
			name:        "negative maxInterval",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{MaxInterval: -time.Second}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.maxInterval must be >= 0, got -1s"},
		},
		{
			name:        "zero multiplier",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{Multiplier: f64(0)}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got 0"},
		},
		{
			name:        "NaN multiplier",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{Multiplier: f64(math.NaN())}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got NaN"},
		},
		{
			name:        "infinite multiplier",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{Multiplier: f64(math.Inf(1))}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got +Inf"},
		},
		{
			name:        "randomizationFactor above 1",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{RandomizationFactor: f64(1.5)}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.randomizationFactor must be between 0 and 1, got 1.5"},
		},
		{
			name:        "initialInterval above explicit maxInterval",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{InitialInterval: 500 * time.Millisecond, MaxInterval: 100 * time.Millisecond}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval (500ms) must not be greater than maxInterval (100ms)"},
		},
		{
			name:        "initialInterval above default maxInterval",
			opts:        &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{InitialInterval: 2 * time.Minute}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval (2m0s) must not be greater than maxInterval (1m0s)"},
		},
		{
			name:    "valid backoff",
			opts:    &Options{Strategy: RoundRobin, ExponentialBackoff: &ExponentialBackoffOptions{InitialInterval: time.Second, MaxInterval: 10 * time.Second, Multiplier: f64(1.5), RandomizationFactor: f64(0.5)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.Validate()
			if tt.wantErr && err == nil {
				t.Fatalf("expected an error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			for _, want := range tt.errContains {
				if err == nil || !strings.Contains(err.Error(), want) {
					t.Errorf("expected error to contain %q, got: %v", want, err)
				}
			}
		})
	}
}
