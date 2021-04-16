package basegin

import (
	"testing"

	"github.com/cadicode/basegin/base"
	"github.com/cadicode/basegin/router"

	"github.com/gin-gonic/gin"
)

func TestCreateGin(t *testing.T) {
	type args struct {
		isProduct          bool
		isCors             bool
		logger             base.ILogger
		corsAllowOrigins   []string
		corsAllowHeaders   []string
		groupedControllers map[string][]router.IBaseController
	}
	tests := []struct {
		name    string
		args    args
		want    *gin.Engine
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				isProduct:          true,
				isCors:             false,
				logger:             nil,
				corsAllowOrigins:   nil,
				corsAllowHeaders:   nil,
				groupedControllers: nil,
			},
			want:    gin.Default(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateGin(tt.args.isProduct, tt.args.isCors, tt.args.logger, tt.args.corsAllowOrigins, tt.args.corsAllowHeaders, tt.args.groupedControllers)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
