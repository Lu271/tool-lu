package crud

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func Init() cli.ActionFunc {
	return func(c *cli.Context) error {
		dbType := c.String("t")
		if dbType == "" {
			return fmt.Errorf("db type is empty")
		}

		dsn := c.String("n")
		if dsn == "" {
			return fmt.Errorf("dsn is empty")
		}

		queryPath := "internal/store/query"
		modelPath := "internal/store/model"
		folder := c.String("f")
		if folder != "" {
			queryPath = "internal/store/" + folder + "/query"
			modelPath = "internal/store/" + folder + "/model"
		}
		g := gen.NewGenerator(gen.Config{
			OutPath:      queryPath,
			ModelPkgPath: modelPath,
		})
		var db *gorm.DB
		if strings.ToLower(dbType) == "mysql " {
			db, _ = gorm.Open(mysql.Open(dsn))
		}
		g.UseDB(db)

		g.ApplyBasic(g.GenerateAllTable(
			gen.FieldGORMTag("created_at", "column:created_at;autoCreateTime"),
			gen.FieldGORMTag("updated_at", "column:updated_at;autoUpdateTime"),
			gen.FieldTypeReg("_at", "int64"),
		)...)
		g.Execute()
		return nil
	}
}
