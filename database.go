package main

import (
	_ "github.com/Shopify/sarama"
	_ "github.com/aerospike/aerospike-client-go"
	_ "github.com/couchbase/gocb"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/elastic/go-elasticsearch/v6"
	_ "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocql/gocql"
	_ "github.com/influxdata/influxdb-client-go"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/olivere/elastic"
	_ "github.com/rethinkdb/rethinkdb-go"
	_ "github.com/zegl/goriak"
	_ "go.mongodb.org/mongo-driver/mongo"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jmoiron/sqlx"
)
