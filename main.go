package main

import "fmt";
 
func main() {

	//var str string = "hello World";
//	fmt.Println("hello World!");
//	fmt.Println(add());
//	fmt.Println(getIndexOfString("hello world!",3))
//	fmt.Println("Length of " + str + " is", getLength(str));
//	fmt.Println("FizzBuzz", fizzBuzz(5));

//	fmt.Scanf("%s", &input);
//	fmt.Printf(input);
	input := 5;
	fmt.Println("fibonacci", fibonacci(input));
}


func add() float32{
 return 1 + 0.1;
}

func getLength(str string) int{
	 
	return len(str);
}
func getIndexOfString(str string, i int)  string{
	return string(str[i]);
}



func isDivByThree(num int) bool {
	return num % 3 == 0;
}

func isDivByFive(num int) bool {
	return num % 5 == 0;
}
func fizzBuzz(number int) string{
	
	var result string;
	
	if(isDivByFive(number) && isDivByThree(number)){
		result = "foobar"
	}else if(isDivByFive(number)){
		result ="foo"
	}else if(isDivByThree(number)){
		result ="bar"
	}else {
		result = string(number);
	}

	return result;

}

func fibonacci(number int) int {
	 var (
		 i =  number
		 result = 1
	 );

	 for i > 0 {
		result = result * i;
		i--; 
	 }	
	 

	 return result;
}

