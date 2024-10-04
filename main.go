package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
	"errors"
	//	"slices"
	"strings"
)

// var word_weight_dictionary = map[string]uint{
// 	" of ": 650000000,
// 	" a ": 600000000,
// 	" the ": 650000000,
// 	"oo": 600000000,
// 	" is ":600000000,
// 	"th":600000000,
// 	"ing ":700000000,
// }
// var char_weight_dictionary = map[string]uint{
// 	"e": 529117365,
// 	"t": 390965105,
// 	"a": 374061888,
// 	"o": 326627740,
// 	"i": 320410057,
// 	"n": 313720540,
// 	"s": 294300210,
// 	"r": 277000841,
// 	"h": 216768975,
// 	"l": 183996130,
// 	"d": 169330528,
// 	"c": 138416451,
// 	"u": 117295780,
// 	"m": 110504544,
// 	"f": 95422055,
// 	"g": 91258980,
// 	"p": 90376747,
// 	"w": 79843664,
// 	"y": 75294515,
// 	"b": 70195826,
// 	"v": 46337161,
// 	"k": 35373464,
// 	"j": 9613410,
// 	"x": 8369915,
// 	"z": 4975847,
// 	"q": 4550166,
// }


var word_weight_dictionary=map[string]uint{
        "of":142,
        "a":131,
        "the":142,
        "oo":131,
        "is":131,
        "th":131,
        "ing":200,
	"and":200,
	"in":160,
	"tion":150,
	"to":170,
}
var char_weight_dictionary=map[string]uint{
        "e":116,
        "t":85,
        "a":82,
        "o":71,
        "i":70,
        "n":68,
        "s":64,
        "r":60,
        "h":47,
        "l":40,
        "d":37,
        "c":30,
        "u":25,
        "m":24,
        "f":20,
        "g":20,
        "p":19,
        "w":17,
        "y":16,
        "b":15,
        "v":10,
        "k":7,
        "j":2,
        "x":1,
        "z":1,
        "q":1,
}



const hex_encoded_string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func hex_to_base64(input string) string{
	bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var output string = base64.StdEncoding.EncodeToString(bytes)
	return output
}

func fixed_xor(input1 string, input2 string) (string, error) {
	bytes1, err1 := hex.DecodeString(input1)
	if err1 != nil {
		return "", err1
	}

	bytes2, err2 := hex.DecodeString(input2)
	if err2 != nil {
		return "", err2
	}

	if len(bytes1) != len(bytes2){
		return "", errors.New("Strings must have the same lenght")
	}
	
	var output []byte = make([]byte, len(bytes1))
	for i := range bytes1{
		output[i] = bytes1[i] ^ bytes2[i]
	}
	return hex.EncodeToString(output), nil
}

func char_xor(input string, char int) string{
	bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var output []byte = make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++{
		output[i] = bytes[i] ^ byte(char)
	}
	return hex.EncodeToString(output)
}


func break_single_byte_XOR_cypher(input string) (string, byte){
	//var str_len = len(input)
	var scores = make([]uint, 256)
	var max uint = 0
	var message string
	var cypher byte
	for i := 0; i < 256; i++{
		var bytes, _ = hex.DecodeString(char_xor(input, i))
		var str = string(bytes)
		for index := range str{
			scores[i] = 0
			scores[i] += char_weight_dictionary[strings.ToLower(string(str[index]))]
		}
		for word, key := range word_weight_dictionary{
			if strings.Contains(str, word){
				scores[i] += key
			}
		}
		if scores[i] > max {
			fmt.Println("Max:", max, "Score:", scores[i], "Letter:", string(i))
			max = scores[i]
			message = str
			cypher = byte(i)
		}
	}
	//fmt.Println("Message:", message, "      ", "Cypher:", string(cypher))
	return message, byte(cypher)
}



func main(){
	// Convert hex to base 64 ====================================================
	fmt.Println("Convert hex to base 64:")
	var input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(hex_to_base64(input), "\n===========================\n\n")



	
	// Fixed XOR ================================================================
	fmt.Println("Fixed XOR")
	var input1 = "1c0111001f010100061a024b53535009181c"
	var input2 = "686974207468652062756c6c277320657965"
	var output, err = fixed_xor(input1, input2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output, "\n===========================\n\n")




	
	// Single Byte XOR cypher =====================================================
	fmt.Println("Single Byte XOR Cypher")
	message, cypher_key := break_single_byte_XOR_cypher(hex_encoded_string)	
	fmt.Println("Message:", message, "      ", "Cypher:", string(cypher_key))




	

	//
}