package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	type args struct {
		dividend uint16
		divisor  uint16
	}
	tests := []struct {
		args           args
		wantQuotient   uint16
		wantRemainder  uint16
		wantErr        bool
		wantErrMessage string
	}{
		{
			args: args{
				dividend: 3,
				divisor:  0,
			},
			wantErr:        true,
			wantErrMessage: "divisor must be not equals to 0",
		},
		{
			args: args{
				dividend: 3,
				divisor:  1,
			},
			wantQuotient:  3,
			wantRemainder: 0,
		},
		{
			args: args{
				dividend: 0,
				divisor:  3,
			},
			wantQuotient:  0,
			wantRemainder: 0,
		},
		{
			args: args{
				dividend: 3,
				divisor:  3,
			},
			wantQuotient:  1,
			wantRemainder: 0,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("successs: Div(%d,%d) retunrs (%d,%d)", tt.args.dividend, tt.args.divisor, tt.wantQuotient, tt.wantRemainder)
		if tt.wantErr {
			name = fmt.Sprintf("failure: Div(%d,%d) retunrs error `%s`", tt.args.dividend, tt.args.divisor, tt.wantErrMessage)
		}

		t.Run(name, func(t *testing.T) {
			gotQuotient, gotRemainder, err := Div(tt.args.dividend, tt.args.divisor)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantQuotient, gotQuotient, "quotient is wrong")
			assert.Equal(t, tt.wantRemainder, gotRemainder, "remainder is wrong")
		})
	}
}
