//Преобразуем yaml файлы в нужный нам формат
//находим число в бд
//создаем таблицу при первом запуске программы

package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"gopkg.in/yaml.v3"
)
type Config struct{
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Env      string `yaml:"env"`
}

// Преобразуем yaml файлы в нужный нам формат
func ReadConfigFromFile(filePath string) (*Config, error) {
	// Чтение содержимого файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Создание объекта Config
	config := &Config{}

	// Разбор YAML данных в объект Config
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// находим число в бд
func FindInSQLStorage(my_number int64)(answer string, count int64){
	
	config, err := ReadConfigFromFile("config/config.yaml")
    if err != nil {
        log.Fatal(err)
    }
    sqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)
    db, err := sql.Open("postgres", sqlconn)
	createTableIfNotExists(db)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Проверяем, есть ли запись с указанным "my_number"
    err = db.QueryRow(`SELECT "count" FROM "numbers" WHERE "num" = $1`, my_number).Scan(&count)
    if err != nil && err != sql.ErrNoRows {
        log.Fatal(err)
    }

	// Если запись найдена, увеличиваем значение "count" на 1
	if (err == nil || err == sql.ErrNoRows) && count!=0 {
		_, err = db.Exec(`UPDATE "numbers" SET "count" = $1 WHERE "num" = $2`, count+1, my_number)
        if err != nil {
            log.Fatal(err)
        }
		answer="yes"

	//если в таблице нет значений или они есть но число не совпало, то число добавляеться с count=0
    }else{
		insertDynStmt := `insert into "numbers"("num", "count") values($1, $2)`
    	_, err := db.Exec(insertDynStmt, my_number, count+1)
		if err != nil {
            log.Fatal(err)
        }
		answer="no"
    }
	
	return answer,count
}

//создаем таблицу при первом запуске программы
func createTableIfNotExists(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS numbers (
			num INT,
			count INT
		);
	`

	_, err := db.Exec(query)
	return err
}
