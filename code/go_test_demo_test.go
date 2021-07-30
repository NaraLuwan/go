package main

import "testing"

/**
go test会为指定的源码文件生成一个虚拟代码包——“command-line-arguments”
而array_test.go引用了其他包中的数据并不属于代码包“command-line-arguments”，直接执行编译会不通过
需要在执行go test时指定引用的包：go test -v go_test_demo_test.go go_test_demo.go
*/

func Test_checkEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "hqcadmin@163.com",
			},
			want: true,
		},
		{
			name: "valid email",
			args: args{
				email: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckEmail(tt.args.email); got != tt.want {
				t.Errorf("checkEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
