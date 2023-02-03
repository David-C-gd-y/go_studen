package main

func main() {
	name := "dawei"
	power := 10000
	if name == "Leto" {
		print("the spice must flow")
	}

	if name == "dawei" {
		panic("hello dawei")
	}

	if (name == "dawei" && power > 9000) || (name == "xiaowei" && power > 10000) {
		panic("猛男")
	}

}
