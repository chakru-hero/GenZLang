package main
import (
	"os"
	"fmt"
	"unicode"
	"bytes"
)
//type TokenType struct {
//	int_lit string := "int_lit",
//	_return string := "return",
//	semi string := "semi"
//}
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

func token_to_asm(tokens []Token) string{
	var output string = "global _start\n_start:\n";
	for i:=0;i<len(tokens);i++{
		 token := tokens[i];
		if token.TokenType == "return" {
			if i+1 < len(tokens) && tokens[i+1].TokenType == "int_lit" {
				if i+2 < len(tokens) && tokens[i+2].TokenType == "semi" {
					output += "    mov rax, 60\n";
					output += "    mov rdi, "  + tokens[i+1].value +"\n";
					output += "    syscall";
				}
			}
		}
	}
	return output;
}

func main (){
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect usage. Correct usage is ...");
		fmt.Println("Correct usage is gz <input.genz>");
		os.Exit(1);
	}

	dat, err := os.ReadFile(args[1]);
	check(err);
	tokens := tokenize(string(dat));
	file, err := os.Create("out.asm");
	check(err);
	file.Write([]byte(token_to_asm(tokens)));
	os.Exit(0);


}

