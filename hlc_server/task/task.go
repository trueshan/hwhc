package task

import "github.com/hwhc/hlc_server/task/price"

func StartTask() {
	//go profileOverdue()
	//
	//go updateUserLevel1()
	//go updateUserLevel2()
	//go updateUserLevel3()
	//
	go price.Start()
	//
	//go datePrice()

	//go ConfigRate()

	//go Tbouns()

	//go partnerProducts()
}
