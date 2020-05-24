package testutil

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"task-api/src/interfaces"
	"testing"
)

// SetUpDB test用DBを作成、マイグレーション。dbDirPathはdbの設定ディレクトリのパスを受け取る。
func SetUpDB(dbDirPath string) {
	// test用DBを作成
	exec.Command(
		"mysql",
		"-h", os.Getenv("DB_HOST"),
		"-u", os.Getenv("DB_USER"),
		"-p"+os.Getenv("DB_PASSWORD"),
		"-e", "DROP DATABASE IF EXISTS "+os.Getenv("DB_NAME"),
	).Run()
	cmd := exec.Command(
		"mysql",
		"-h", os.Getenv("DB_HOST"),
		"-u", os.Getenv("DB_USER"),
		"-p"+os.Getenv("DB_PASSWORD"),
		"-e", "CREATE DATABASE "+os.Getenv("DB_NAME"),
	)
	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to create database error: ", err)
	}

	// migration実行
	cmd2 := exec.Command("goose", "-path", dbDirPath, "up")
	cmd2.Env = append(os.Environ(), "DB_NAME="+os.Getenv("DB_NAME"))
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd2.Run()
	if err != nil {
		log.Println("failed to migrate test db error: ", err)
	}
}

// ClearDB test用DBの削除
func ClearDB() {
	cmd := exec.Command(
		"mysql",
		"-h", os.Getenv("DB_HOST"),
		"-u", os.Getenv("DB_USER"),
		"-p"+os.Getenv("DB_PASSWORD"),
		"-e", "DROP DATABASE IF EXISTS "+os.Getenv("DB_NAME"),
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println("failed to clear database err: ", err)
	}
}

func InsertFromString(t *testing.T, query string) {
	cmd := exec.Command(
		"mysql",
		"-h", os.Getenv("DB_HOST"),
		"-u", os.Getenv("DB_USER"),
		"-p"+os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(query))
	stdin.Close()

	cmd.Run()
}

func ExecSchema(tx interfaces.SQLHandler, fpath string) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("schema reading error: %v", err)
	}

	queries := strings.Split(string(b), ";")

	for _, query := range queries[:len(queries)-1] {
		_, err := tx.Exec(query)
		if err != nil {
			log.Fatalf("exec scheme error: %v", err)
		}
	}
}
