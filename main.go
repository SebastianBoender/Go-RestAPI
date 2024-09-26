package main

import (
	"net/http"
	"database/sql"
	"fmt"
    "log"
    "os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type recentSales struct {
    ID     int64
    Offer_name  string
    Credits float64
    Country  string
    DateTime string

}

func main() {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "",
        AllowNativePasswords: true,
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    router := gin.Default()

	router.GET("/recentSales/:id",getRecentSales)
	router.Run("localhost:3600")
}

func getRecentSales(c *gin.Context) {
    uid := c.Param("id")
    saleInfo,err := getRecentSalesByID(uid)

     if err != nil {
        	panic(err)
        }

	if saleInfo == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, saleInfo)
	}
}

func getRecentSalesByID(uid string) ([]recentSales, error) {
    var recentSale []recentSales


    rows, err := db.Query("SELECT `id`, `offer_name`, `credits`, `country`, `date` FROM offer_process WHERE `status` = 1 AND `uid` = ? ORDER BY `date` ASC LIMIT 4", uid)
    if err != nil {
        return nil, fmt.Errorf("recentSale %q: %v", uid, err)
    }

    defer rows.Close()
    
    for rows.Next() {
        var rec recentSales
        if err := rows.Scan(&rec.ID, &rec.Offer_name, &rec.Credits, &rec.Country, &rec.DateTime); err != nil {
            return nil, fmt.Errorf("recentSale %q: %v", uid, err)
        }
        recentSale = append(recentSale, rec)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("recentSale %q: %v", uid, err)
    }

    return recentSale, nil
}




