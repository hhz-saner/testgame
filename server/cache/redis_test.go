package cache

import (
	"errors"
	"testing"
)

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value string
		ttl   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"test", "test", 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Set(tt.args.key, tt.args.value, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {

	type args struct {
		key   string
		value string
		ttl   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test", args{"test", "test", 10}, "test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if e := Set(tt.args.key, tt.args.value, tt.args.ttl); e != nil {
				t.Errorf("Get() error = %v, wantErr %v", e, tt.wantErr)
			}
			got, err := Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemember(t *testing.T) {
	type args struct {
		key string
		ttl int
		f   func() (string, error)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test", args{"test", 10, func() (s string, e error) {
			return "test", nil
		}}, "test", false},
		{"test", args{"test", 10, func() (s string, e error) {
			return "test1", nil
		}}, "test", false},
		{"testError", args{"testError", 10, func() (s string, e error) {
			return "testError", errors.New("something error")
		}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Remember(tt.args.key, tt.args.ttl, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Remember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ping(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"ping", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ping(); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
