package main

import (
    "database/postgresql"
    "log"
    "os"
)

type Person struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.Default()

    db, err := sql.Open("postgresql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router.GET("/persons", func(c *gin.Context) {
        var persons []Person
        rows, err := db.Query("SELECT id, name, age FROM persons")
        if err != nil {
            log.Printf("Failed to execute query: %v", err)
            c.JSON(500, gin.H{"error": "Failed to execute query"})
            return
        }
        defer rows.Close()
        for rows.Next() {
            var p Person
            if err := rows.Scan(&p.ID, &p.Name, &p.Age); err != nil {
                log.Printf("Failed to scan row: %v", err)
                c.JSON(500, gin.H{"error": "Failed to scan row"})
                return
            }
            persons = append(persons, p)
        }
        if err := rows.Err(); err != nil {
            log.Printf("Failed to iterate over rows: %v", err)
            c.JSON(500, gin.H{"error": "Failed to iterate over rows"})
            return
        }
        c.JSON(200, persons)
    })

    router.Run(":8080")
}
