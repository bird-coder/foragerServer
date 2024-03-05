package dao

import (
	"fmt"
	"foragerServer/options"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func NewDB(c *options.MysqlConfig) *gorm.DB {
	mysqlConfig := c.Default
	dsn := fmt.Sprintf(initDsn(), mysqlConfig.User, mysqlConfig.Pass, mysqlConfig.Protocol,
		mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if c.Cluster {
		initCluster(db, c)
	}
	if err != nil {
		panic("init mysql failed")
	}
	return db
}

func initCluster(db *gorm.DB, c *options.MysqlConfig) {
	sourceConfigs := c.Sources
	replicaConfigs := c.Replicas

	if len(sourceConfigs) == 0 && len(replicaConfigs) == 0 {
		return
	}

	config := dbresolver.Config{}
	config.Sources = []gorm.Dialector{}
	config.Replicas = []gorm.Dialector{}

	dsnStr := initDsn()
	for _, sourceConfig := range sourceConfigs {
		dsn := fmt.Sprintf(dsnStr, sourceConfig.User, sourceConfig.Pass, sourceConfig.Protocol,
			sourceConfig.Host, sourceConfig.Port, sourceConfig.Database, sourceConfig.Charset)
		config.Sources = append(config.Sources, mysql.Open(dsn))
	}
	for _, replicaConfig := range replicaConfigs {
		dsn := fmt.Sprintf(dsnStr, replicaConfig.User, replicaConfig.Pass, replicaConfig.Protocol,
			replicaConfig.Host, replicaConfig.Port, replicaConfig.Database, replicaConfig.Charset)
		config.Replicas = append(config.Replicas, mysql.Open(dsn))
	}
	db.Use(dbresolver.Register(config))
}

func initDsn() string {
	return "%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
}
