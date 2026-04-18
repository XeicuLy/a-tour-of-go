package main

import "fmt"

func numeric() {
	/*
		var (
			u8  uint8     = 255
			i8  int8      = 127
			f32 float32   = 0.2
			c64 complex64 = -5 + 12i
		)
		fmt.Println(u8, i8, f32, c64)
		fmt.Printf("type=%T value=%v\n", u8, u8)
		fmt.Printf("type=%T value=%v\n", i8, i8)
		fmt.Printf("type=%T value=%v\n", f32, f32)
		fmt.Printf("type=%T value=%v\n", c64, c64)
	*/

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("2 * 3 =", 2*3)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 % 3 =", 10%3)

	fmt.Println("1.5 + 2.5 =", 1.5+2.5)
	fmt.Println("1.5 - 0.5 =", 1.5-0.5)
	fmt.Println("1.5 * 2 =", 1.5*2)
	fmt.Println("1.5 / 0.5 =", 1.5/0.5)

	fmt.Println("complex(1, 2) + complex(3, 4) =", complex(1, 2)+complex(3, 4))
	fmt.Println("complex(1, 2) - complex(3, 4) =", complex(1, 2)-complex(3, 4))
	fmt.Println("complex(1, 2) * complex(3, 4) =", complex(1, 2)*complex(3, 4))
	fmt.Println("complex(1, 2) / complex(3, 4) =", complex(1, 2)/complex(3, 4))
}
