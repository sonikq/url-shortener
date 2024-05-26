package app

import (
	"github.com/sonikq/url-shortener/internal/app/pkg/utils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	cfg := Config{}
	ParseConfig(&cfg)
	require.NotEmpty(t, cfg)
}

func Test_getEnvInt(t *testing.T) {
	type args struct {
		key           string
		argumentValue *int
	}

	tests := []struct {
		name     string
		initFunc func()
		args     args
		want     int
	}{
		{
			name:     "with valid os env",
			initFunc: func() { os.Setenv("NUMBER", "250") },
			args: args{
				key:           "NUMBER",
				argumentValue: utils.Ptr(125),
			},
			want: 250,
		},
		{
			name:     "with invalid os env",
			initFunc: func() { os.Setenv("NUMBER", "invalid") },
			args: args{
				key:           "NUMBER",
				argumentValue: utils.Ptr(125),
			},
			want: 125,
		},
		{
			name:     "with empty os env",
			initFunc: func() { os.Setenv("NUMBER", "") },
			args: args{
				key:           "NUMBER",
				argumentValue: utils.Ptr(125),
			},
			want: 125,
		},
		{
			name:     "without os env",
			initFunc: func() { os.Unsetenv("NUMBER") },
			args: args{
				key:           "NUMBER",
				argumentValue: utils.Ptr(125),
			},
			want: 125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initFunc()
			if got := getEnvInt(tt.args.key, tt.args.argumentValue); got != tt.want {
				t.Errorf("getEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvString(t *testing.T) {
	type args struct {
		key           string
		argumentValue *string
	}

	tests := []struct {
		name     string
		initFunc func()
		args     args
		want     string
	}{
		{
			name:     "with normal os env",
			initFunc: func() { os.Setenv("TEXT", "normal") },
			args: args{
				key:           "TEXT",
				argumentValue: utils.Ptr("argument"),
			},
			want: "normal",
		},
		{
			name:     "with empty os env",
			initFunc: func() { os.Setenv("TEXT", "") },
			args: args{
				key:           "TEXT",
				argumentValue: utils.Ptr("argument"),
			},
			want: "",
		},
		{
			name:     "without os env",
			initFunc: func() { os.Unsetenv("TEXT") },
			args: args{
				key:           "TEXT",
				argumentValue: utils.Ptr("argument"),
			},
			want: "argument",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initFunc()
			if got := getEnvString(tt.args.key, tt.args.argumentValue); got != tt.want {
				t.Errorf("getEnvString() = %v, want %v", got, tt.want)
			}
		})
	}
}
