package main

import "fmt"

func main(){
	var n, strength, speed, enemyStrength, enemySpeed int
	var i, defeated int

	fmt.Scan(&n)
	fmt.Scan(&strength, &speed)
	i = 1

	for i <= n && strength >= 3 && speed >= 3 {
		fmt.Scan(&enemyStrength, &enemySpeed)
		if (enemyStrength < strength && enemySpeed < speed) {
			if (strength > speed) {
				strength -= 6
				defeated++
			} else {
				speed -= 6
				defeated++
			}
		} else if (enemyStrength < strength && enemySpeed > speed) {
			strength -= 6
			defeated++
		} else if (enemyStrength >= strength && enemySpeed < speed) {
			speed -= 6
			defeated++
		}
		i++
	}
	fmt.Println(defeated, strength, speed)
}