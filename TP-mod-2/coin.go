package main

import "fmt"

func main(){
	var shard, galactum, nebula, stellaris int
	var lumin int

	fmt.Scan(&lumin)
	shard = lumin / 1200
	lumin = lumin % 1200
	galactum = lumin / 120
	lumin = lumin % 120
	nebula = lumin / 40
	lumin = lumin % 40 
	stellaris = lumin / 20
	lumin = lumin % 20
	fmt.Println(shard, galactum, nebula, stellaris, lumin)

}

/* 
1 shard = 10 galactum
1 galactum = 3 nebula
1 nebula = 2 stellaris
1 stellaris = 20 lumin

1 shard = 1200 lumin
1 galactum = 120 lumin
1 nebula = 40 lumin
1 stellaris = 20 lumin
*/