package infra

import (
	"context"
	"log"
	"os"
	"rinha-backend/src/types"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Conn *pgxpool.Pool
)

func ConnectDB() types.IPgx {

	if Conn != nil {
		return Conn
	} else {
		log.Println("opening connections")
	}

	connectionString := os.Getenv("DATABASE_URL")

	if connectionString == "" {
		connectionString = "postgresql://postgres:@123456@localhost:5432/rinha"
	}

	maxConnEnv := os.Getenv("MAX_CONNECTIONS")
	if maxConnEnv == "" {
		maxConnEnv = "30"
	}

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	maxConn, _ := strconv.Atoi(maxConnEnv)

	config.MaxConns = int32(maxConn)
	config.MinConns = int32(maxConn)

	config.MaxConnIdleTime = time.Minute * 3

	Conn, err = pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Fatal(err)
	}

	err = Conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return Conn
}
