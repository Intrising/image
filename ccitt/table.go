// generated by "go run gen.go". DO NOT EDIT.

package ccitt

// Each table is represented by an array of [2]int16's: a binary tree. Each
// array element (other than element 0, which means invalid) is a branch node
// in that tree. The root node is always element 1 (the second element).
//
// To walk the tree, look at the next bit in the bit stream, using it to select
// the first or second element of the [2]int16. If that int16 is 0, we have an
// invalid code. If it is positive, go to that branch node. If it is negative,
// then we have a leaf node, whose value is the bitwise complement (the ^
// operator) of that int16.
//
// Comments above each table also show the same structure visually. The "b123"
// lines show the 123'rd branch node. The "=XXXXX" lines show an invalid code.
// The "=v1234" lines show a leaf node with value 1234. When reading the bit
// stream, a 0 or 1 bit means to go up or down, as you move left to right.
//
// For example, in modeTable, branch node b005 is three steps up from the root
// node, meaning that we have already seen "000". If the next bit is "0" then
// we move to branch node b006. Otherwise, the next bit is "1", and we move to
// the leaf node v0000 (also known as the modePass constant). Indeed, the bits
// that encode modePass are "0001".
//
// Tables 1, 2 and 3 come from the "ITU-T Recommendation T.6: FACSIMILE CODING
// SCHEMES AND CODING CONTROL FUNCTIONS FOR GROUP 4 FACSIMILE APPARATUS"
// specification:
//
// https://www.itu.int/rec/dologin_pub.asp?lang=e&id=T-REC-T.6-198811-I!!PDF-E&type=items

// modeTable represents Table 1 and the End-of-Line code.
//
//                              +=XXXXX
// b015                       +-+
//                            | +=v0010
// b014                     +-+
//                          | +=XXXXX
// b013                   +-+
//                        | +=XXXXX
// b012                 +-+
//                      | +=XXXXX
// b011               +-+
//                    | +=XXXXX
// b009             +-+
//                  | +=v0009
// b007           +-+
//                | | +=v0008
// b010           | +-+
//                |   +=v0005
// b006         +-+
//              | | +=v0007
// b008         | +-+
//              |   +=v0004
// b005       +-+
//            | +=v0000
// b003     +-+
//          | +=v0001
// b002   +-+
//        | | +=v0006
// b004   | +-+
//        |   +=v0003
// b001 +-+
//        +=v0002
var modeTable = [...][2]int16{
	0:  {0, 0},
	1:  {2, ^2},
	2:  {3, 4},
	3:  {5, ^1},
	4:  {^6, ^3},
	5:  {6, ^0},
	6:  {7, 8},
	7:  {9, 10},
	8:  {^7, ^4},
	9:  {11, ^9},
	10: {^8, ^5},
	11: {12, 0},
	12: {13, 0},
	13: {14, 0},
	14: {15, 0},
	15: {0, ^10},
}

// whiteTable represents Tables 2 and 3 for a white run.
//
//                      +=XXXXX
// b059               +-+
//                    | |     +=v1792
// b096               | |   +-+
//                    | |   | | +=v1984
// b100               | |   | +-+
//                    | |   |   +=v2048
// b094               | | +-+
//                    | | | |   +=v2112
// b101               | | | | +-+
//                    | | | | | +=v2176
// b097               | | | +-+
//                    | | |   | +=v2240
// b102               | | |   +-+
//                    | | |     +=v2304
// b085               | +-+
//                    |   |   +=v1856
// b098               |   | +-+
//                    |   | | +=v1920
// b095               |   +-+
//                    |     |   +=v2368
// b103               |     | +-+
//                    |     | | +=v2432
// b099               |     +-+
//                    |       | +=v2496
// b104               |       +-+
//                    |         +=v2560
// b040             +-+
//                  | | +=v0029
// b060             | +-+
//                  |   +=v0030
// b026           +-+
//                | |   +=v0045
// b061           | | +-+
//                | | | +=v0046
// b041           | +-+
//                |   +=v0022
// b016         +-+
//              | |   +=v0023
// b042         | | +-+
//              | | | | +=v0047
// b062         | | | +-+
//              | | |   +=v0048
// b027         | +-+
//              |   +=v0013
// b008       +-+
//            | |     +=v0020
// b043       | |   +-+
//            | |   | | +=v0033
// b063       | |   | +-+
//            | |   |   +=v0034
// b028       | | +-+
//            | | | |   +=v0035
// b064       | | | | +-+
//            | | | | | +=v0036
// b044       | | | +-+
//            | | |   | +=v0037
// b065       | | |   +-+
//            | | |     +=v0038
// b017       | +-+
//            |   |   +=v0019
// b045       |   | +-+
//            |   | | | +=v0031
// b066       |   | | +-+
//            |   | |   +=v0032
// b029       |   +-+
//            |     +=v0001
// b004     +-+
//          | |     +=v0012
// b030     | |   +-+
//          | |   | |   +=v0053
// b067     | |   | | +-+
//          | |   | | | +=v0054
// b046     | |   | +-+
//          | |   |   +=v0026
// b018     | | +-+
//          | | | |     +=v0039
// b068     | | | |   +-+
//          | | | |   | +=v0040
// b047     | | | | +-+
//          | | | | | | +=v0041
// b069     | | | | | +-+
//          | | | | |   +=v0042
// b031     | | | +-+
//          | | |   |   +=v0043
// b070     | | |   | +-+
//          | | |   | | +=v0044
// b048     | | |   +-+
//          | | |     +=v0021
// b009     | +-+
//          |   |     +=v0028
// b049     |   |   +-+
//          |   |   | | +=v0061
// b071     |   |   | +-+
//          |   |   |   +=v0062
// b032     |   | +-+
//          |   | | |   +=v0063
// b072     |   | | | +-+
//          |   | | | | +=v0000
// b050     |   | | +-+
//          |   | |   | +=v0320
// b073     |   | |   +-+
//          |   | |     +=v0384
// b019     |   +-+
//          |     +=v0010
// b002   +-+
//        | |     +=v0011
// b020   | |   +-+
//        | |   | |   +=v0027
// b051   | |   | | +-+
//        | |   | | | | +=v0059
// b074   | |   | | | +-+
//        | |   | | |   +=v0060
// b033   | |   | +-+
//        | |   |   |     +=v1472
// b086   | |   |   |   +-+
//        | |   |   |   | +=v1536
// b075   | |   |   | +-+
//        | |   |   | | | +=v1600
// b087   | |   |   | | +-+
//        | |   |   | |   +=v1728
// b052   | |   |   +-+
//        | |   |     +=v0018
// b010   | | +-+
//        | | | |     +=v0024
// b053   | | | |   +-+
//        | | | |   | | +=v0049
// b076   | | | |   | +-+
//        | | | |   |   +=v0050
// b034   | | | | +-+
//        | | | | | |   +=v0051
// b077   | | | | | | +-+
//        | | | | | | | +=v0052
// b054   | | | | | +-+
//        | | | | |   +=v0025
// b021   | | | +-+
//        | | |   |     +=v0055
// b078   | | |   |   +-+
//        | | |   |   | +=v0056
// b055   | | |   | +-+
//        | | |   | | | +=v0057
// b079   | | |   | | +-+
//        | | |   | |   +=v0058
// b035   | | |   +-+
//        | | |     +=v0192
// b005   | +-+
//        |   |     +=v1664
// b036   |   |   +-+
//        |   |   | |   +=v0448
// b080   |   |   | | +-+
//        |   |   | | | +=v0512
// b056   |   |   | +-+
//        |   |   |   |   +=v0704
// b088   |   |   |   | +-+
//        |   |   |   | | +=v0768
// b081   |   |   |   +-+
//        |   |   |     +=v0640
// b022   |   | +-+
//        |   | | |     +=v0576
// b082   |   | | |   +-+
//        |   | | |   | | +=v0832
// b089   |   | | |   | +-+
//        |   | | |   |   +=v0896
// b057   |   | | | +-+
//        |   | | | | |   +=v0960
// b090   |   | | | | | +-+
//        |   | | | | | | +=v1024
// b083   |   | | | | +-+
//        |   | | | |   | +=v1088
// b091   |   | | | |   +-+
//        |   | | | |     +=v1152
// b037   |   | | +-+
//        |   | |   |     +=v1216
// b092   |   | |   |   +-+
//        |   | |   |   | +=v1280
// b084   |   | |   | +-+
//        |   | |   | | | +=v1344
// b093   |   | |   | | +-+
//        |   | |   | |   +=v1408
// b058   |   | |   +-+
//        |   | |     +=v0256
// b011   |   +-+
//        |     +=v0002
// b001 +-+
//        |     +=v0003
// b012   |   +-+
//        |   | | +=v0128
// b023   |   | +-+
//        |   |   +=v0008
// b006   | +-+
//        | | |   +=v0009
// b024   | | | +-+
//        | | | | | +=v0016
// b038   | | | | +-+
//        | | | |   +=v0017
// b013   | | +-+
//        | |   +=v0004
// b003   +-+
//          |   +=v0005
// b014     | +-+
//          | | |   +=v0014
// b039     | | | +-+
//          | | | | +=v0015
// b025     | | +-+
//          | |   +=v0064
// b007     +-+
//            | +=v0006
// b015       +-+
//              +=v0007
var whiteTable = [...][2]int16{
	0:   {0, 0},
	1:   {2, 3},
	2:   {4, 5},
	3:   {6, 7},
	4:   {8, 9},
	5:   {10, 11},
	6:   {12, 13},
	7:   {14, 15},
	8:   {16, 17},
	9:   {18, 19},
	10:  {20, 21},
	11:  {22, ^2},
	12:  {^3, 23},
	13:  {24, ^4},
	14:  {^5, 25},
	15:  {^6, ^7},
	16:  {26, 27},
	17:  {28, 29},
	18:  {30, 31},
	19:  {32, ^10},
	20:  {^11, 33},
	21:  {34, 35},
	22:  {36, 37},
	23:  {^128, ^8},
	24:  {^9, 38},
	25:  {39, ^64},
	26:  {40, 41},
	27:  {42, ^13},
	28:  {43, 44},
	29:  {45, ^1},
	30:  {^12, 46},
	31:  {47, 48},
	32:  {49, 50},
	33:  {51, 52},
	34:  {53, 54},
	35:  {55, ^192},
	36:  {^1664, 56},
	37:  {57, 58},
	38:  {^16, ^17},
	39:  {^14, ^15},
	40:  {59, 60},
	41:  {61, ^22},
	42:  {^23, 62},
	43:  {^20, 63},
	44:  {64, 65},
	45:  {^19, 66},
	46:  {67, ^26},
	47:  {68, 69},
	48:  {70, ^21},
	49:  {^28, 71},
	50:  {72, 73},
	51:  {^27, 74},
	52:  {75, ^18},
	53:  {^24, 76},
	54:  {77, ^25},
	55:  {78, 79},
	56:  {80, 81},
	57:  {82, 83},
	58:  {84, ^256},
	59:  {0, 85},
	60:  {^29, ^30},
	61:  {^45, ^46},
	62:  {^47, ^48},
	63:  {^33, ^34},
	64:  {^35, ^36},
	65:  {^37, ^38},
	66:  {^31, ^32},
	67:  {^53, ^54},
	68:  {^39, ^40},
	69:  {^41, ^42},
	70:  {^43, ^44},
	71:  {^61, ^62},
	72:  {^63, ^0},
	73:  {^320, ^384},
	74:  {^59, ^60},
	75:  {86, 87},
	76:  {^49, ^50},
	77:  {^51, ^52},
	78:  {^55, ^56},
	79:  {^57, ^58},
	80:  {^448, ^512},
	81:  {88, ^640},
	82:  {^576, 89},
	83:  {90, 91},
	84:  {92, 93},
	85:  {94, 95},
	86:  {^1472, ^1536},
	87:  {^1600, ^1728},
	88:  {^704, ^768},
	89:  {^832, ^896},
	90:  {^960, ^1024},
	91:  {^1088, ^1152},
	92:  {^1216, ^1280},
	93:  {^1344, ^1408},
	94:  {96, 97},
	95:  {98, 99},
	96:  {^1792, 100},
	97:  {101, 102},
	98:  {^1856, ^1920},
	99:  {103, 104},
	100: {^1984, ^2048},
	101: {^2112, ^2176},
	102: {^2240, ^2304},
	103: {^2368, ^2432},
	104: {^2496, ^2560},
}

// blackTable represents Tables 2 and 3 for a black run.
//
//                      +=XXXXX
// b017               +-+
//                    | |     +=v1792
// b042               | |   +-+
//                    | |   | | +=v1984
// b063               | |   | +-+
//                    | |   |   +=v2048
// b029               | | +-+
//                    | | | |   +=v2112
// b064               | | | | +-+
//                    | | | | | +=v2176
// b043               | | | +-+
//                    | | |   | +=v2240
// b065               | | |   +-+
//                    | | |     +=v2304
// b022               | +-+
//                    |   |   +=v1856
// b044               |   | +-+
//                    |   | | +=v1920
// b030               |   +-+
//                    |     |   +=v2368
// b066               |     | +-+
//                    |     | | +=v2432
// b045               |     +-+
//                    |       | +=v2496
// b067               |       +-+
//                    |         +=v2560
// b013             +-+
//                  | |     +=v0018
// b031             | |   +-+
//                  | |   | |   +=v0052
// b068             | |   | | +-+
//                  | |   | | | | +=v0640
// b095             | |   | | | +-+
//                  | |   | | |   +=v0704
// b046             | |   | +-+
//                  | |   |   |   +=v0768
// b096             | |   |   | +-+
//                  | |   |   | | +=v0832
// b069             | |   |   +-+
//                  | |   |     +=v0055
// b023             | | +-+
//                  | | | |     +=v0056
// b070             | | | |   +-+
//                  | | | |   | | +=v1280
// b097             | | | |   | +-+
//                  | | | |   |   +=v1344
// b047             | | | | +-+
//                  | | | | | |   +=v1408
// b098             | | | | | | +-+
//                  | | | | | | | +=v1472
// b071             | | | | | +-+
//                  | | | | |   +=v0059
// b032             | | | +-+
//                  | | |   |   +=v0060
// b072             | | |   | +-+
//                  | | |   | | | +=v1536
// b099             | | |   | | +-+
//                  | | |   | |   +=v1600
// b048             | | |   +-+
//                  | | |     +=v0024
// b018             | +-+
//                  |   |     +=v0025
// b049             |   |   +-+
//                  |   |   | |   +=v1664
// b100             |   |   | | +-+
//                  |   |   | | | +=v1728
// b073             |   |   | +-+
//                  |   |   |   +=v0320
// b033             |   | +-+
//                  |   | | |   +=v0384
// b074             |   | | | +-+
//                  |   | | | | +=v0448
// b050             |   | | +-+
//                  |   | |   |   +=v0512
// b101             |   | |   | +-+
//                  |   | |   | | +=v0576
// b075             |   | |   +-+
//                  |   | |     +=v0053
// b024             |   +-+
//                  |     |     +=v0054
// b076             |     |   +-+
//                  |     |   | | +=v0896
// b102             |     |   | +-+
//                  |     |   |   +=v0960
// b051             |     | +-+
//                  |     | | |   +=v1024
// b103             |     | | | +-+
//                  |     | | | | +=v1088
// b077             |     | | +-+
//                  |     | |   | +=v1152
// b104             |     | |   +-+
//                  |     | |     +=v1216
// b034             |     +-+
//                  |       +=v0064
// b010           +-+
//                | |   +=v0013
// b019           | | +-+
//                | | | |     +=v0023
// b052           | | | |   +-+
//                | | | |   | | +=v0050
// b078           | | | |   | +-+
//                | | | |   |   +=v0051
// b035           | | | | +-+
//                | | | | | |   +=v0044
// b079           | | | | | | +-+
//                | | | | | | | +=v0045
// b053           | | | | | +-+
//                | | | | |   | +=v0046
// b080           | | | | |   +-+
//                | | | | |     +=v0047
// b025           | | | +-+
//                | | |   |     +=v0057
// b081           | | |   |   +-+
//                | | |   |   | +=v0058
// b054           | | |   | +-+
//                | | |   | | | +=v0061
// b082           | | |   | | +-+
//                | | |   | |   +=v0256
// b036           | | |   +-+
//                | | |     +=v0016
// b014           | +-+
//                |   |     +=v0017
// b037           |   |   +-+
//                |   |   | |   +=v0048
// b083           |   |   | | +-+
//                |   |   | | | +=v0049
// b055           |   |   | +-+
//                |   |   |   | +=v0062
// b084           |   |   |   +-+
//                |   |   |     +=v0063
// b026           |   | +-+
//                |   | | |     +=v0030
// b085           |   | | |   +-+
//                |   | | |   | +=v0031
// b056           |   | | | +-+
//                |   | | | | | +=v0032
// b086           |   | | | | +-+
//                |   | | | |   +=v0033
// b038           |   | | +-+
//                |   | |   |   +=v0040
// b087           |   | |   | +-+
//                |   | |   | | +=v0041
// b057           |   | |   +-+
//                |   | |     +=v0022
// b020           |   +-+
//                |     +=v0014
// b008         +-+
//              | |   +=v0010
// b015         | | +-+
//              | | | +=v0011
// b011         | +-+
//              |   |     +=v0015
// b027         |   |   +-+
//              |   |   | |     +=v0128
// b088         |   |   | |   +-+
//              |   |   | |   | +=v0192
// b058         |   |   | | +-+
//              |   |   | | | | +=v0026
// b089         |   |   | | | +-+
//              |   |   | | |   +=v0027
// b039         |   |   | +-+
//              |   |   |   |   +=v0028
// b090         |   |   |   | +-+
//              |   |   |   | | +=v0029
// b059         |   |   |   +-+
//              |   |   |     +=v0019
// b021         |   | +-+
//              |   | | |     +=v0020
// b060         |   | | |   +-+
//              |   | | |   | | +=v0034
// b091         |   | | |   | +-+
//              |   | | |   |   +=v0035
// b040         |   | | | +-+
//              |   | | | | |   +=v0036
// b092         |   | | | | | +-+
//              |   | | | | | | +=v0037
// b061         |   | | | | +-+
//              |   | | | |   | +=v0038
// b093         |   | | | |   +-+
//              |   | | | |     +=v0039
// b028         |   | | +-+
//              |   | |   |   +=v0021
// b062         |   | |   | +-+
//              |   | |   | | | +=v0042
// b094         |   | |   | | +-+
//              |   | |   | |   +=v0043
// b041         |   | |   +-+
//              |   | |     +=v0000
// b016         |   +-+
//              |     +=v0012
// b006       +-+
//            | |   +=v0009
// b012       | | +-+
//            | | | +=v0008
// b009       | +-+
//            |   +=v0007
// b004     +-+
//          | | +=v0006
// b007     | +-+
//          |   +=v0005
// b002   +-+
//        | | +=v0001
// b005   | +-+
//        |   +=v0004
// b001 +-+
//        | +=v0003
// b003   +-+
//          +=v0002
var blackTable = [...][2]int16{
	0:   {0, 0},
	1:   {2, 3},
	2:   {4, 5},
	3:   {^3, ^2},
	4:   {6, 7},
	5:   {^1, ^4},
	6:   {8, 9},
	7:   {^6, ^5},
	8:   {10, 11},
	9:   {12, ^7},
	10:  {13, 14},
	11:  {15, 16},
	12:  {^9, ^8},
	13:  {17, 18},
	14:  {19, 20},
	15:  {^10, ^11},
	16:  {21, ^12},
	17:  {0, 22},
	18:  {23, 24},
	19:  {^13, 25},
	20:  {26, ^14},
	21:  {27, 28},
	22:  {29, 30},
	23:  {31, 32},
	24:  {33, 34},
	25:  {35, 36},
	26:  {37, 38},
	27:  {^15, 39},
	28:  {40, 41},
	29:  {42, 43},
	30:  {44, 45},
	31:  {^18, 46},
	32:  {47, 48},
	33:  {49, 50},
	34:  {51, ^64},
	35:  {52, 53},
	36:  {54, ^16},
	37:  {^17, 55},
	38:  {56, 57},
	39:  {58, 59},
	40:  {60, 61},
	41:  {62, ^0},
	42:  {^1792, 63},
	43:  {64, 65},
	44:  {^1856, ^1920},
	45:  {66, 67},
	46:  {68, 69},
	47:  {70, 71},
	48:  {72, ^24},
	49:  {^25, 73},
	50:  {74, 75},
	51:  {76, 77},
	52:  {^23, 78},
	53:  {79, 80},
	54:  {81, 82},
	55:  {83, 84},
	56:  {85, 86},
	57:  {87, ^22},
	58:  {88, 89},
	59:  {90, ^19},
	60:  {^20, 91},
	61:  {92, 93},
	62:  {^21, 94},
	63:  {^1984, ^2048},
	64:  {^2112, ^2176},
	65:  {^2240, ^2304},
	66:  {^2368, ^2432},
	67:  {^2496, ^2560},
	68:  {^52, 95},
	69:  {96, ^55},
	70:  {^56, 97},
	71:  {98, ^59},
	72:  {^60, 99},
	73:  {100, ^320},
	74:  {^384, ^448},
	75:  {101, ^53},
	76:  {^54, 102},
	77:  {103, 104},
	78:  {^50, ^51},
	79:  {^44, ^45},
	80:  {^46, ^47},
	81:  {^57, ^58},
	82:  {^61, ^256},
	83:  {^48, ^49},
	84:  {^62, ^63},
	85:  {^30, ^31},
	86:  {^32, ^33},
	87:  {^40, ^41},
	88:  {^128, ^192},
	89:  {^26, ^27},
	90:  {^28, ^29},
	91:  {^34, ^35},
	92:  {^36, ^37},
	93:  {^38, ^39},
	94:  {^42, ^43},
	95:  {^640, ^704},
	96:  {^768, ^832},
	97:  {^1280, ^1344},
	98:  {^1408, ^1472},
	99:  {^1536, ^1600},
	100: {^1664, ^1728},
	101: {^512, ^576},
	102: {^896, ^960},
	103: {^1024, ^1088},
	104: {^1152, ^1216},
}

// COPY PASTE table.go BEGIN

const (
	modePass = iota // Pass
	modeH           // Horizontal
	modeV0          // Vertical-0
	modeVR1         // Vertical-Right-1
	modeVR2         // Vertical-Right-2
	modeVR3         // Vertical-Right-3
	modeVL1         // Vertical-Left-1
	modeVL2         // Vertical-Left-2
	modeVL3         // Vertical-Left-3
	modeExt         // Extension
	modeEOL         // End-of-Line
)

// COPY PASTE table.go END
