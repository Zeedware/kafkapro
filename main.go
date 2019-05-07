package main

import (
	"fmt"
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	configurationFactory := NewConfigurationFactory()
	cfg := configurationFactory.NewConfiguration()
	kafkaCfg := configurationFactory.NewKafkaConfiguration()

	kafkaProducer := NewKafkaProducer(kafkaCfg)

	data, err := ioutil.ReadFile(cfg.inputFileName)
	checkErr(err)
	for i := 0; i < cfg.Loop(); i++ {
		status := "success"
		fmt.Println("message:", string(data))
		err = kafkaProducer.sendMessage(data)
		if err != nil {
			fmt.Println("error:", err)
			status = "failed"
		}
		fmt.Println("status: ", status)
	}

}
