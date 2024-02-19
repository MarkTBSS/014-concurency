package main

import (
	"fmt"
	"time"
)

func worker(id int, jobsChannel <-chan string, resultsChannel chan<- string) {
	// หนึ่ง Job ใช้เวลา 2 วินาที
	for paper := range jobsChannel {
		fmt.Println("Worker resultsChannelParam <- paper Start ", id)
		fmt.Println(id, "Start Job", paper)
		time.Sleep(2 * time.Second) // สมมติว่าหนึ่งงานใช้เวลา 2 วินาที
		fmt.Println(id, "End   Job", paper)
		resultsChannel <- paper
		fmt.Println("Worker resultsChannelParam <- paper Finished ", id)
	}
}

func main() {
	start := time.Now()
	// ขนาดของช่องทาง (buffered channel)
	const numberOfJobs = 4
	jobsChannel := make(chan string, numberOfJobs)    // Create a buffered channel for jobs
	resultsChannel := make(chan string, numberOfJobs) // Create a buffered channel for results

	/* สร้างใบงาน (จากตัวอย่างคือ X ใบงาน)
	for j := 1; j <= numberOfJobs; j++ {
		jobsChannel <- fmt.Sprintf("job-%d", j)
	}
	close(jobsChannel) // ปิดช่องทาง jobs เมื่อใส่งานครบแล้ว */

	// สร้างใบงาน (จากตัวอย่างคือ X ใบงาน)
	// ทำส่วนนี้เป็นส่วนแรก
	papers := []string{"Paper-A", "Paper--B", "Paper---C", "Paper-----D"}
	for id, paper := range papers {
		jobsChannel <- paper
		fmt.Println("jobsChannel <- paper Finished", id)
	}
	close(jobsChannel) // ปิดช่องทาง jobs เมื่อใส่งานครบแล้ว

	/* // เริ่มการทำงานที่จุดนี้โดยใช้ go routine (Asynconous)
	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobsChannel, resultsChannel)
	} */
	go worker(1, jobsChannel, resultsChannel)
	go worker(2, jobsChannel, resultsChannel)

	// Collect results from the worker goroutines
	for a := 1; a <= numberOfJobs; a++ {
		fmt.Println("result := <-resultsChannel Start", a)
		result := <-resultsChannel // Receive result from results channel
		fmt.Println("result := <-resultsChannel Finished", a, " : ", result)
	}
	fmt.Println("Total Run Time : ", time.Since(start))
}
