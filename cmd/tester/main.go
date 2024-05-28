package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	n := runtime.NumCPU()
	fmt.Println("numcpu")
	fmt.Println(n)

	g := []int{10}

	p := []float64{0.000001}

	P := []float64{0.001}

	r := []int{1}

	w := []int{1}

	m := make([]struct{}, runtime.GOMAXPROCS(0))

	if err := os.MkdirAll("./vkr/drwres", os.ModeDir); err != nil {
		fmt.Printf("ERROR mkdir: %s\n", err.Error())
		return
	}

	file, fileErr := os.OpenFile("./vkr/drwres/stats", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if fileErr != nil {
		fmt.Printf("ERROR stats: %s\n", fileErr.Error())
		return
	}
	defer file.Close()

	errs, fileErr := os.OpenFile("./vkr/drwres/errs", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if fileErr != nil {
		fmt.Printf("ERROR  errs: %s\n", fileErr.Error())
		return
	}
	defer errs.Close()

	coef := 30

	fmt.Println(coef*len(g)*len(p)*len(P)*len(r)*len(w)*len(m) + coef*len(g)*len(p)*len(r)*len(w)*len(m))
	fmt.Println(coef * len(g) * len(p) * len(P) * len(r) * len(w) * len(m))
	fmt.Println(coef * len(g) * len(p) * len(r) * len(w) * len(m))
	fmt.Println("")

	var counter = 0

	for range make([]struct{}, coef) {
		for _, gg := range g {
			for _, pp := range p {
				for _, rr := range r {
					for _, ww := range w {
						for mm := range m {
							xx := 0
							PP := float64(0)
							ii := 10
							if counter%100 == 0 {
								fmt.Println(time.Now())
								fmt.Println(counter)
							}
							counter++
							cmd := exec.Command("./runner",
								"-g", strconv.Itoa(gg),
								"-p", fmt.Sprintf("%.10f", pp),
								"-P", fmt.Sprintf("%.10f", PP),
								"-r", strconv.Itoa(rr),
								"-w", strconv.Itoa(ww),
								"-x", strconv.Itoa(xx))

							cmd.Env = []string{"GOMAXPROCS=" + strconv.Itoa(mm+1)}

							stdout := &bytes.Buffer{}
							stderr := &bytes.Buffer{}
							cmd.Stdout = stdout
							cmd.Stderr = stderr

							if err := cmd.Run(); err != nil {
								s := fmt.Sprintf("g=%v i=%v p=%v P=%v r=%v w=%v c=%v m=%v err=%v\n%v", gg, ii, pp, PP, rr, ww, n, mm, err.Error(), stderr.String())
								fmt.Print(s)
								if _, err := errs.WriteString(s); err != nil {
									fmt.Printf("ERROR:  err: %s\n", err.Error())
								}
								continue
							}

							if _, err := file.WriteString(stdout.String()); err != nil {
								fmt.Print(stdout.String())
								fmt.Printf("ERROR: file: %s\n", err.Error())
							}
						}
					}
				}
			}

		}
	}

	for range make([]struct{}, coef) {
		for _, gg := range g {
			for _, pp := range p {
				for _, PP := range P {
					for _, rr := range r {
						for _, ww := range w {
							for mm := range m {
								xx := 1
								ii := 10
								if counter%100 == 0 {
									fmt.Println(time.Now())
									fmt.Println(counter)
								}
								counter++
								cmd := exec.Command("./runner",
									"-g", strconv.Itoa(gg),
									"-p", fmt.Sprintf("%f", pp),
									"-P", fmt.Sprintf("%f", PP),
									"-r", strconv.Itoa(rr),
									"-w", strconv.Itoa(ww),
									"-x", strconv.Itoa(xx))

								cmd.Env = []string{"GOMAXPROCS=" + strconv.Itoa(mm+1)}

								stdout := &bytes.Buffer{}
								stderr := &bytes.Buffer{}
								cmd.Stdout = stdout
								cmd.Stderr = stderr

								if err := cmd.Run(); err != nil {
									s := fmt.Sprintf("g=%v i=%v p=%v P=%v r=%v w=%v c=%v m=%v err=%v\n%v", gg, ii, pp, PP, rr, ww, n, mm, err.Error(), stderr.String())
									fmt.Print(s)
									if _, err := errs.WriteString(s); err != nil {
										fmt.Printf("ERROR:  err: %s\n", err.Error())
									}
									continue
								}

								if _, err := file.WriteString(stdout.String()); err != nil {
									fmt.Print(stdout.String())
									fmt.Printf("ERROR: file: %s\n", err.Error())
								}
							}

						}
					}
				}
			}

		}
	}

	fmt.Println(time.Now())
	fmt.Println("DONE. Sleeping...")
	// for {
	// 	time.Sleep(30 * time.Minute)
	// }
}
