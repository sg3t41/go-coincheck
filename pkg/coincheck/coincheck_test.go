package coincheck

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		opts    []Option
		wantErr bool
		check   func(*Coincheck) bool
	}{
		"正常系: Optionでcredentialsがセットされる": {
			opts: []Option{
				func(c *Coincheck) error {
					c.credentials = credentials{key: "k", secret: "s"}
					return nil
				},
			},
			wantErr: false,
			check: func(c *Coincheck) bool {
				return c.credentials.key == "k" && c.credentials.secret == "s"
			},
		},
		"異常系: Optionがエラーを返す": {
			opts: []Option{
				func(c *Coincheck) error {
					return errors.New("option error")
				},
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			cc, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("エラー期待値: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if err != nil {
				return
			}
			if tt.check != nil && !tt.check(cc) {
				t.Errorf("Option適用結果不一致: %+v", cc)
			}
		})
	}
}
