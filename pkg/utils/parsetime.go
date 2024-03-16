package utils

import "time"

func ParseTime(str string) (time.Time, error) {
	//str := jastipReq.Close_Order
	layout := "15.04"
	t, err := time.Parse(layout, str)
	if err != nil {
		//fmt.Println("Error parsing time:", err)
		return time.Time{}, err
	}
	
	return t, nil

}
