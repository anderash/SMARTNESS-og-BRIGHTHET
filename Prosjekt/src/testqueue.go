package main

import (
	"./queue"
	"encoding/json"
	"fmt"
	"time"
)

type ElevInfo struct {
	IPADDR     string
	F_NEW_INFO bool

	F_DEAD_ELEV bool
	F_BUTTONPRESS bool

	POSITION    int
	DIRECTION   int
	DESTINATION int

	ButtonType  int
	ButtonFloor int
}

func main() {
	my_ipaddr := "111.11.111"
	other_ipaddr := "999.99.999"
	c_to_queueManager := make(chan []byte)

	queue.InitQueuemanager(my_ipaddr, c_to_queueManager)
	// queue.AppendElevator(other_ipaddr)

	// position := 3
	// direction := 1
	// destination_pos := 3
	// queue.SetElevator(other_ipaddr, position, direction, destination_pos)
	
	// queue.PrintActiveElevators()

	// // button_type := 0
	// button_floor := 2
	// // queue.AppendOrder(button_type, button_floor)

	// fmt.Println("Cost:", queue.CostFunction(my_ipaddr, button_floor, "up"))
	// fmt.Println("Cost:", queue.CostFunction(other_ipaddr, button_floor, "up"))

	// queue.PrintActiveElevators()


	position := 3
	direction := 1
	destination_floor := 3
	elev_info := ElevInfo{other_ipaddr, true, false, false, position,direction,destination_floor,0,0}
	encoded_elev_info, err := json.Marshal(elev_info)
	if err != nil{
		fmt.Println("error: ", err)
	}

	// time.Sleep(1000 * time.Millisecond)
	c_to_queueManager <- encoded_elev_info
	time.Sleep(10 * time.Millisecond)
	queue.PrintActiveElevators()

	button_type := 0
	button_floor := 2

	elev_info = ElevInfo{other_ipaddr, true, false, true, position,direction,destination_floor,button_type,button_floor}
	encoded_elev_info, err2 := json.Marshal(elev_info)
	if err2 != nil{
		fmt.Println("error: ", err2)
	}
	time.Sleep(1000 * time.Millisecond)
	c_to_queueManager <- encoded_elev_info
	time.Sleep(10 * time.Millisecond)
	queue.PrintActiveElevators()

	// queue.AppendOrder(button_type, button_floor)


}
