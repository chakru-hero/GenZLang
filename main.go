package main
import (
	"os"
	"fmt"
	"unicode"
	"bytes"
)

type Token struct{
	TokenType string;
	value string;
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func tokenize(str string) []Token {
	var tokenArray []Token;

	var b bytes.Buffer;
	for i:=0; i<len(str); i++ {
		c :=str[i];
		if unicode.IsLetter(rune(c)) {
			b.WriteByte(c);
			i++;
			for unicode.IsPrint(rune(str[i])) && !unicode.IsSpace(rune(str[i])) {
				b.WriteByte(str[i]);
 				i++;
			}
			i--;
			if("return" == b.String()){
				tokenArray = append(tokenArray,Token{TokenType: b.String(), value: ""} );
				b.Reset();
			} else{
				println(b.String());
				fmt.Println("Compilation Error : While Creating AlphaNum Token");
				os.Exit(0);
			}
		} else if unicode.IsSpace(rune(c)){

		} else if unicode.IsDigit(rune(c)) {
			b.WriteByte(c);
			i++;
			for unicode.IsDigit(rune(str[i])){
				b.WriteByte(str[i]);
				i++;
			}
			i--;
			tokenArray = append(tokenArray, Token{TokenType: "int_lit", value: b.String()});
			b.Reset();
		} else if string(c) == ";" {
			tokenArray = append(tokenArray, Token{TokenType: "semi", value: ""});
		} else {
			fmt.Println("Compilation Error : While creating Tokens");
		}

	}
	return tokenArray; 
}

func main (){
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect usage. Correct usage is ...");
		fmt.Println("Correct usage is gz <input.gz>");
		os.Exit(1);
	}

	dat, err := os.ReadFile(args[1]);
	check(err);
	tokens := tokenize(string(dat));
	for i:=0; i< len(tokens);i++ {
		fmt.Println(tokens[i].TokenType, tokens[i].value);
	}

	os.Exit(0);


}

