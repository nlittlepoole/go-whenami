package whenami

import (
	"database/sql"
	_ "github.com/shaxbee/go-spatialite"
	"io"
	"net/http"
	"os"
)

var db *sql.DB

const TIMEZONE_SQL string = `
	SELECT tz.tz_name
    FROM timezone AS tz
    WHERE tz.ROWID IN
    (
        SELECT ROWID
        FROM idx_timezone_geometry
        WHERE xmin < ? AND xmax > ? AND ymin < ? AND ymax > ?
    )
    AND Contains(geometry, MAKEPOINT( ?, ?));`

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	filepath := "tzgeo.sqlite"
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		out, err := os.Create(filepath)
		defer out.Close()
		resp, err := http.Get("https://github.com/judy2k/tzgeo/raw/master/tzgeo/tzgeo.sqlite")
		checkErr(err)
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		checkErr(err)
	}
	var err error
	db, err = sql.Open("spatialite", filepath)
	checkErr(err)
}

func WhenAmI(lat float64, lon float64) (string, error) {
	result := "Unknown"
	rows, err := db.Query(TIMEZONE_SQL, lon, lon, lat, lat, lon, lat)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
