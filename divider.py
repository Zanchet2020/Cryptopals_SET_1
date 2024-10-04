

string = """
var word_weight_dictionary = map[string]uint{
	" of ": 650000000,
	" a ": 600000000,
	" the ": 650000000,
	"oo": 600000000,
	" is ": 600000000,
	"th": 600000000,
	"ing ": 700000000,
}
var char_weight_dictionary = map[string]uint{
	"e": 529117365,
	"t": 390965105,
	"a": 374061888,
	"o": 326627740,
	"i": 320410057,
	"n": 313720540,
	"s": 294300210,
	"r": 277000841,
	"h": 216768975,
	"l": 183996130,
	"d": 169330528,
	"c": 138416451,
	"u": 117295780,
	"m": 110504544,
	"f": 95422055,
	"g": 91258980,
	"p": 90376747,
	"w": 79843664,
	"y": 75294515,
	"b": 70195826,
	"v": 46337161,
	"k": 35373464,
	"j": 9613410,
	"x": 8369915,
	"z": 4975847,
	"q": 4550166,
}
"""


for line in string.splitlines():
    for word in line.split(' '):
        if(word.replace(',', '').isdigit()):
            print(int(word.replace(',', '')) // 4550166, end='')
        else:
            print(word, end='')
    print(',\n', end='')
