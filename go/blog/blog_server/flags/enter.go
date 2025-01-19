// flags/enter.go
package flags

import "flag"

type Options struct {
	  File    string
	  DB      bool
	  Version bool
}

var FlagOptions = new(Options)

func Parse() {
	  flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	  flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	  flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	  flag.Parse()
}

//需要进行数据库的迁移 -- 就是新建表
func Run() {
	  if FlagOptions.DB {
	  	MigrateDB()
	  }
}
