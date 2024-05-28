package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./vkr/drwres/test.db")
	if err != nil {
		fmt.Printf("open db: %v", err.Error())
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("ping db: %v", err.Error())
		return
	}

	fields := []string{
		"g_by_proc",
		"p_of_write",
		"p_of_cpu_info",
		"w_num",
		"r_num",
		"cpu_count",
		"proc_count",
	}
	mus := []int{0, 1}

	wg := &sync.WaitGroup{}

	wg.Add(len(fields) * len(mus))
	for _, m := range mus {
		file, err := os.OpenFile("./vkr/drwres/"+strconv.Itoa(m), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("open file: %v", err.Error())
			return
		}
		defer file.Close()
		for _, f := range fields {
			m := m
			f := f
			func() {
				defer wg.Done()
				rows, err := db.Query(fmt.Sprintf(relationAvg, f), m)
				if err != nil {
					fmt.Printf("\n  q: %v\nerr: %v\n", f, err.Error())
					return
				}
				defer rows.Close()
				_, err = file.WriteString(f + "\n")
				if err != nil {
					fmt.Printf("\n  w: %v\nerr: %v\n", f, err.Error())
					return
				}
				var iTime, ff float64
				itimer := make([]float64, 0)
				ffr := make([]float64, 0)
				for rows.Next() {
					err = rows.Scan(&iTime, &ff)
					if err != nil {
						fmt.Printf("\n  s: %v\nerr: %v\n", f, err.Error())
						break
					}
					itimer = append(itimer, iTime)
					ffr = append(ffr, ff)
				}
				if rows.Err() != nil {
					fmt.Printf("\nrow: %v\nerr: %v\n", f, err.Error())
					return
				}
				_, err = file.WriteString(fmt.Sprintf("%s\n", k(itimer)))
				if err != nil {
					fmt.Printf("\n wi: %v\nerr: %v\n", f, err.Error())
				}

				_, err = file.WriteString(fmt.Sprintf("%s\n", k(ffr)))
				if err != nil {
					fmt.Printf("\n wf: %v\nerr: %v\n", f, err.Error())
				}
			}()
		}
	}
	wg.Wait()
}

type k []float64

func (kk k) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	for _, v := range kk {
		b.WriteString(fmt.Sprintf("%v,", v))
	}
	b.WriteString("]")
	return b.String()
}

var powerLineAvg = `
SELECT
	AVG(i_time_seconds), g_by_proc, p_of_write, p_of_cpu_info, w_num, r_num, proc_count
FROM
	bench
WHERE
    mu_is_drw=$1
GROUP BY
	g_by_proc, p_of_write, p_of_cpu_info, w_num, r_num, cpu_count, proc_count, mu_is_drw
ORDER BY
    AVG(i_time_seconds)
`

var relationAvg = `
SELECT
	AVG(i_time_seconds),
	%[1]v
FROM
    bench
WHERE
    mu_is_drw=$1
GROUP BY
    %[1]v, mu_is_drw
ORDER BY
    %[1]v
`
