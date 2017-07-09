package main


// Error detection/recovery capacity.
//
// There are several levels of error detection/recovery capacity. Higher levels
// of error recovery are able to correct more errors, with the trade-off of
// increased symbol size.
type RecoveryLevel int

type dataEncoderType uint8

const (
	dataEncoderType1To9 dataEncoderType = iota
	dataEncoderType10To26
	dataEncoderType27To40
)

const (
	// Level M: 15% error recovery. Good default choice. 0b00
	Medium RecoveryLevel = iota
	// Level L: 7% error recovery. 0b01
	Low
	// Level H: 30% error recovery.0b10
	Highest
	// Level Q: 25% error recovery.0b11
	High
)

// qrCodeVersion describes the data length and encoding order of a single QR
// Code version. There are 40 versions numbers x 4 recovery levels == 160
// possible qrCodeVersion structures.
type qrCodeVersion struct {
	// Version number (1-40 inclusive).
	version          int

	// Error recovery level.
	level            RecoveryLevel

	dataEncoderType  dataEncoderType

	// Encoded data can be split into multiple blocks. Each block contains data
	// and error recovery bytes.
	//
	// Larger QR Codes contain more blocks.
	block            []block

	// Number of bits required to pad the combined data & error correction bit
	// stream up to the symbol's full capacity.
	numRemainderBits int
}

type block struct {
	numBlocks        int

	// Total codewords (numCodewords == numErrorCodewords+numDataCodewords).
	numCodewords     int

	// Number of data codewords.
	numDataCodewords int
}

var (
	Versions = []qrCodeVersion{
		{
			1,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					1,
					26,
					19,
				},
			},
			0,
		},
		{
			1,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					1,
					26,
					16,
				},
			},
			0,
		},
		{
			1,
			High,
			dataEncoderType1To9,
			[]block{
				{
					1,
					26,
					13,
				},
			},
			0,
		},
		{
			1,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					1,
					26,
					9,
				},
			},
			0,
		},
		{
			2,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					1,
					44,
					34,
				},
			},
			7,
		},
		{
			2,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					1,
					44,
					28,
				},
			},
			7,
		},
		{
			2,
			High,
			dataEncoderType1To9,
			[]block{
				{
					1,
					44,
					22,
				},
			},
			7,
		},
		{
			2,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					1,
					44,
					16,
				},
			},
			7,
		},
		{
			3,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					1,
					70,
					55,
				},
			},
			7,
		},
		{
			3,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					1,
					70,
					44,
				},
			},
			7,
		},
		{
			3,
			High,
			dataEncoderType1To9,
			[]block{
				{
					2,
					35,
					17,
				},
			},
			7,
		},
		{
			3,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					2,
					35,
					13,
				},
			},
			7,
		},
		{
			4,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					1,
					100,
					80,
				},
			},
			7,
		},
		{
			4,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					2,
					50,
					32,
				},
			},
			7,
		},
		{
			4,
			High,
			dataEncoderType1To9,
			[]block{
				{
					2,
					50,
					24,
				},
			},
			7,
		},
		{
			4,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					4,
					25,
					9,
				},
			},
			7,
		},
		{
			5,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					1,
					134,
					108,
				},
			},
			7,
		},
		{
			5,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					2,
					67,
					43,
				},
			},
			7,
		},
		{
			5,
			High,
			dataEncoderType1To9,
			[]block{
				{
					2,
					33,
					15,
				},
				{
					2,
					34,
					16,
				},
			},
			7,
		},
		{
			5,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					2,
					33,
					11,
				},
				{
					2,
					34,
					12,
				},
			},
			7,
		},
		{
			6,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					2,
					86,
					68,
				},
			},
			7,
		},
		{
			6,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					4,
					43,
					27,
				},
			},
			7,
		},
		{
			6,
			High,
			dataEncoderType1To9,
			[]block{
				{
					4,
					43,
					19,
				},
			},
			7,
		},
		{
			6,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					4,
					43,
					15,
				},
			},
			7,
		},
		{
			7,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					2,
					98,
					78,
				},
			},
			0,
		},
		{
			7,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					4,
					49,
					31,
				},
			},
			0,
		},
		{
			7,
			High,
			dataEncoderType1To9,
			[]block{
				{
					2,
					32,
					14,
				},
				{
					4,
					33,
					15,
				},
			},
			0,
		},
		{
			7,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					4,
					39,
					13,
				},
				{
					1,
					40,
					14,
				},
			},
			0,
		},
		{
			8,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					2,
					121,
					97,
				},
			},
			0,
		},
		{
			8,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					2,
					60,
					38,
				},
				{
					2,
					61,
					39,
				},
			},
			0,
		},
		{
			8,
			High,
			dataEncoderType1To9,
			[]block{
				{
					4,
					40,
					18,
				},
				{
					2,
					41,
					19,
				},
			},
			0,
		},
		{
			8,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					4,
					40,
					14,
				},
				{
					2,
					41,
					15,
				},
			},
			0,
		},
		{
			9,
			Low,
			dataEncoderType1To9,
			[]block{
				{
					2,
					146,
					116,
				},
			},
			0,
		},
		{
			9,
			Medium,
			dataEncoderType1To9,
			[]block{
				{
					3,
					58,
					36,
				},
				{
					2,
					59,
					37,
				},
			},
			0,
		},
		{
			9,
			High,
			dataEncoderType1To9,
			[]block{
				{
					4,
					36,
					16,
				},
				{
					4,
					37,
					17,
				},
			},
			0,
		},
		{
			9,
			Highest,
			dataEncoderType1To9,
			[]block{
				{
					4,
					36,
					12,
				},
				{
					4,
					37,
					13,
				},
			},
			0,
		},
		{
			10,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					2,
					86,
					68,
				},
				{
					2,
					87,
					69,
				},
			},
			0,
		},
		{
			10,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					4,
					69,
					43,
				},
				{
					1,
					70,
					44,
				},
			},
			0,
		},
		{
			10,
			High,
			dataEncoderType10To26,
			[]block{
				{
					6,
					43,
					19,
				},
				{
					2,
					44,
					20,
				},
			},
			0,
		},
		{
			10,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					6,
					43,
					15,
				},
				{
					2,
					44,
					16,
				},
			},
			0,
		},
		{
			11,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					4,
					101,
					81,
				},
			},
			0,
		},
		{
			11,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					1,
					80,
					50,
				},
				{
					4,
					81,
					51,
				},
			},
			0,
		},
		{
			11,
			High,
			dataEncoderType10To26,
			[]block{
				{
					4,
					50,
					22,
				},
				{
					4,
					51,
					23,
				},
			},
			0,
		},
		{
			11,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					3,
					36,
					12,
				},
				{
					8,
					37,
					13,
				},
			},
			0,
		},
		{
			12,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					2,
					116,
					92,
				},
				{
					2,
					117,
					93,
				},
			},
			0,
		},
		{
			12,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					6,
					58,
					36,
				},
				{
					2,
					59,
					37,
				},
			},
			0,
		},
		{
			12,
			High,
			dataEncoderType10To26,
			[]block{
				{
					4,
					46,
					20,
				},
				{
					6,
					47,
					21,
				},
			},
			0,
		},
		{
			12,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					7,
					42,
					14,
				},
				{
					4,
					43,
					15,
				},
			},
			0,
		},
		{
			13,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					4,
					133,
					107,
				},
			},
			0,
		},
		{
			13,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					8,
					59,
					37,
				},
				{
					1,
					60,
					38,
				},
			},
			0,
		},
		{
			13,
			High,
			dataEncoderType10To26,
			[]block{
				{
					8,
					44,
					20,
				},
				{
					4,
					45,
					21,
				},
			},
			0,
		},
		{
			13,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					12,
					33,
					11,
				},
				{
					4,
					34,
					12,
				},
			},
			0,
		},
		{
			14,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					3,
					145,
					115,
				},
				{
					1,
					146,
					116,
				},
			},
			3,
		},
		{
			14,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					4,
					64,
					40,
				},
				{
					5,
					65,
					41,
				},
			},
			3,
		},
		{
			14,
			High,
			dataEncoderType10To26,
			[]block{
				{
					11,
					36,
					16,
				},
				{
					5,
					37,
					17,
				},
			},
			3,
		},
		{
			14,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					11,
					36,
					12,
				},
				{
					5,
					37,
					13,
				},
			},
			3,
		},
		{
			15,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					5,
					109,
					87,
				},
				{
					1,
					110,
					88,
				},
			},
			3,
		},
		{
			15,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					5,
					65,
					41,
				},
				{
					5,
					66,
					42,
				},
			},
			3,
		},
		{
			15,
			High,
			dataEncoderType10To26,
			[]block{
				{
					5,
					54,
					24,
				},
				{
					7,
					55,
					25,
				},
			},
			3,
		},
		{
			15,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					11,
					36,
					12,
				},
				{
					7,
					37,
					13,
				},
			},
			3,
		},
		{
			16,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					5,
					122,
					98,
				},
				{
					1,
					123,
					99,
				},
			},
			3,
		},
		{
			16,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					7,
					73,
					45,
				},
				{
					3,
					74,
					46,
				},
			},
			3,
		},
		{
			16,
			High,
			dataEncoderType10To26,
			[]block{
				{
					15,
					43,
					19,
				},
				{
					2,
					44,
					20,
				},
			},
			3,
		},
		{
			16,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					3,
					45,
					15,
				},
				{
					13,
					46,
					16,
				},
			},
			3,
		},
		{
			17,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					1,
					135,
					107,
				},
				{
					5,
					136,
					108,
				},
			},
			3,
		},
		{
			17,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					10,
					74,
					46,
				},
				{
					1,
					75,
					47,
				},
			},
			3,
		},
		{
			17,
			High,
			dataEncoderType10To26,
			[]block{
				{
					1,
					50,
					22,
				},
				{
					15,
					51,
					23,
				},
			},
			3,
		},
		{
			17,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					2,
					42,
					14,
				},
				{
					17,
					43,
					15,
				},
			},
			3,
		},
		{
			18,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					5,
					150,
					120,
				},
				{
					1,
					151,
					121,
				},
			},
			3,
		},
		{
			18,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					9,
					69,
					43,
				},
				{
					4,
					70,
					44,
				},
			},
			3,
		},
		{
			18,
			High,
			dataEncoderType10To26,
			[]block{
				{
					17,
					50,
					22,
				},
				{
					1,
					51,
					23,
				},
			},
			3,
		},
		{
			18,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					2,
					42,
					14,
				},
				{
					19,
					43,
					15,
				},
			},
			3,
		},
		{
			19,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					3,
					141,
					113,
				},
				{
					4,
					142,
					114,
				},
			},
			3,
		},
		{
			19,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					3,
					70,
					44,
				},
				{
					11,
					71,
					45,
				},
			},
			3,
		},
		{
			19,
			High,
			dataEncoderType10To26,
			[]block{
				{
					17,
					47,
					21,
				},
				{
					4,
					48,
					22,
				},
			},
			3,
		},
		{
			19,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					9,
					39,
					13,
				},
				{
					16,
					40,
					14,
				},
			},
			3,
		},
		{
			20,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					3,
					135,
					107,
				},
				{
					5,
					136,
					108,
				},
			},
			3,
		},
		{
			20,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					3,
					67,
					41,
				},
				{
					13,
					68,
					42,
				},
			},
			3,
		},
		{
			20,
			High,
			dataEncoderType10To26,
			[]block{
				{
					15,
					54,
					24,
				},
				{
					5,
					55,
					25,
				},
			},
			3,
		},
		{
			20,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					15,
					43,
					15,
				},
				{
					10,
					44,
					16,
				},
			},
			3,
		},
		{
			21,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					4,
					144,
					116,
				},
				{
					4,
					145,
					117,
				},
			},
			4,
		},
		{
			21,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					17,
					68,
					42,
				},
			},
			4,
		},
		{
			21,
			High,
			dataEncoderType10To26,
			[]block{
				{
					17,
					50,
					22,
				},
				{
					6,
					51,
					23,
				},
			},
			4,
		},
		{
			21,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					19,
					46,
					16,
				},
				{
					6,
					47,
					17,
				},
			},
			4,
		},
		{
			22,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					2,
					139,
					111,
				},
				{
					7,
					140,
					112,
				},
			},
			4,
		},
		{
			22,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					17,
					74,
					46,
				},
			},
			4,
		},
		{
			22,
			High,
			dataEncoderType10To26,
			[]block{
				{
					7,
					54,
					24,
				},
				{
					16,
					55,
					25,
				},
			},
			4,
		},
		{
			22,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					34,
					37,
					13,
				},
			},
			4,
		},
		{
			23,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					4,
					151,
					121,
				},
				{
					5,
					152,
					122,
				},
			},
			4,
		},
		{
			23,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					4,
					75,
					47,
				},
				{
					14,
					76,
					48,
				},
			},
			4,
		},
		{
			23,
			High,
			dataEncoderType10To26,
			[]block{
				{
					11,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			4,
		},
		{
			23,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					16,
					45,
					15,
				},
				{
					14,
					46,
					16,
				},
			},
			4,
		},
		{
			24,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					6,
					147,
					117,
				},
				{
					4,
					148,
					118,
				},
			},
			4,
		},
		{
			24,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					6,
					73,
					45,
				},
				{
					14,
					74,
					46,
				},
			},
			4,
		},
		{
			24,
			High,
			dataEncoderType10To26,
			[]block{
				{
					11,
					54,
					24,
				},
				{
					16,
					55,
					25,
				},
			},
			4,
		},
		{
			24,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					30,
					46,
					16,
				},
				{
					2,
					47,
					17,
				},
			},
			4,
		},
		{
			25,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					8,
					132,
					106,
				},
				{
					4,
					133,
					107,
				},
			},
			4,
		},
		{
			25,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					8,
					75,
					47,
				},
				{
					13,
					76,
					48,
				},
			},
			4,
		},
		{
			25,
			High,
			dataEncoderType10To26,
			[]block{
				{
					7,
					54,
					24,
				},
				{
					22,
					55,
					25,
				},
			},
			4,
		},
		{
			25,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					22,
					45,
					15,
				},
				{
					13,
					46,
					16,
				},
			},
			4,
		},
		{
			26,
			Low,
			dataEncoderType10To26,
			[]block{
				{
					10,
					142,
					114,
				},
				{
					2,
					143,
					115,
				},
			},
			4,
		},
		{
			26,
			Medium,
			dataEncoderType10To26,
			[]block{
				{
					19,
					74,
					46,
				},
				{
					4,
					75,
					47,
				},
			},
			4,
		},
		{
			26,
			High,
			dataEncoderType10To26,
			[]block{
				{
					28,
					50,
					22,
				},
				{
					6,
					51,
					23,
				},
			},
			4,
		},
		{
			26,
			Highest,
			dataEncoderType10To26,
			[]block{
				{
					33,
					46,
					16,
				},
				{
					4,
					47,
					17,
				},
			},
			4,
		},
		{
			27,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					8,
					152,
					122,
				},
				{
					4,
					153,
					123,
				},
			},
			4,
		},
		{
			27,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					22,
					73,
					45,
				},
				{
					3,
					74,
					46,
				},
			},
			4,
		},
		{
			27,
			High,
			dataEncoderType27To40,
			[]block{
				{
					8,
					53,
					23,
				},
				{
					26,
					54,
					24,
				},
			},
			4,
		},
		{
			27,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					12,
					45,
					15,
				},
				{
					28,
					46,
					16,
				},
			},
			4,
		},
		{
			28,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					3,
					147,
					117,
				},
				{
					10,
					148,
					118,
				},
			},
			3,
		},
		{
			28,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					3,
					73,
					45,
				},
				{
					23,
					74,
					46,
				},
			},
			3,
		},
		{
			28,
			High,
			dataEncoderType27To40,
			[]block{
				{
					4,
					54,
					24,
				},
				{
					31,
					55,
					25,
				},
			},
			3,
		},
		{
			28,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					11,
					45,
					15,
				},
				{
					31,
					46,
					16,
				},
			},
			3,
		},
		{
			29,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					7,
					146,
					116,
				},
				{
					7,
					147,
					117,
				},
			},
			3,
		},
		{
			29,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					21,
					73,
					45,
				},
				{
					7,
					74,
					46,
				},
			},
			3,
		},
		{
			29,
			High,
			dataEncoderType27To40,
			[]block{
				{
					1,
					53,
					23,
				},
				{
					37,
					54,
					24,
				},
			},
			3,
		},
		{
			29,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					19,
					45,
					15,
				},
				{
					26,
					46,
					16,
				},
			},
			3,
		},
		{
			30,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					5,
					145,
					115,
				},
				{
					10,
					146,
					116,
				},
			},
			3,
		},
		{
			30,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					19,
					75,
					47,
				},
				{
					10,
					76,
					48,
				},
			},
			3,
		},
		{
			30,
			High,
			dataEncoderType27To40,
			[]block{
				{
					15,
					54,
					24,
				},
				{
					25,
					55,
					25,
				},
			},
			3,
		},
		{
			30,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					23,
					45,
					15,
				},
				{
					25,
					46,
					16,
				},
			},
			3,
		},
		{
			31,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					13,
					145,
					115,
				},
				{
					3,
					146,
					116,
				},
			},
			3,
		},
		{
			31,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					2,
					74,
					46,
				},
				{
					29,
					75,
					47,
				},
			},
			3,
		},
		{
			31,
			High,
			dataEncoderType27To40,
			[]block{
				{
					42,
					54,
					24,
				},
				{
					1,
					55,
					25,
				},
			},
			3,
		},
		{
			31,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					23,
					45,
					15,
				},
				{
					28,
					46,
					16,
				},
			},
			3,
		},
		{
			32,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					17,
					145,
					115,
				},
			},
			3,
		},
		{
			32,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					10,
					74,
					46,
				},
				{
					23,
					75,
					47,
				},
			},
			3,
		},
		{
			32,
			High,
			dataEncoderType27To40,
			[]block{
				{
					10,
					54,
					24,
				},
				{
					35,
					55,
					25,
				},
			},
			3,
		},
		{
			32,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					19,
					45,
					15,
				},
				{
					35,
					46,
					16,
				},
			},
			3,
		},
		{
			33,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					17,
					145,
					115,
				},
				{
					1,
					146,
					116,
				},
			},
			3,
		},
		{
			33,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					14,
					74,
					46,
				},
				{
					21,
					75,
					47,
				},
			},
			3,
		},
		{
			33,
			High,
			dataEncoderType27To40,
			[]block{
				{
					29,
					54,
					24,
				},
				{
					19,
					55,
					25,
				},
			},
			3,
		},
		{
			33,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					11,
					45,
					15,
				},
				{
					46,
					46,
					16,
				},
			},
			3,
		},
		{
			34,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					13,
					145,
					115,
				},
				{
					6,
					146,
					116,
				},
			},
			3,
		},
		{
			34,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					14,
					74,
					46,
				},
				{
					23,
					75,
					47,
				},
			},
			3,
		},
		{
			34,
			High,
			dataEncoderType27To40,
			[]block{
				{
					44,
					54,
					24,
				},
				{
					7,
					55,
					25,
				},
			},
			3,
		},
		{
			34,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					59,
					46,
					16,
				},
				{
					1,
					47,
					17,
				},
			},
			3,
		},
		{
			35,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					12,
					151,
					121,
				},
				{
					7,
					152,
					122,
				},
			},
			0,
		},
		{
			35,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					12,
					75,
					47,
				},
				{
					26,
					76,
					48,
				},
			},
			0,
		},
		{
			35,
			High,
			dataEncoderType27To40,
			[]block{
				{
					39,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			0,
		},
		{
			35,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					22,
					45,
					15,
				},
				{
					41,
					46,
					16,
				},
			},
			0,
		},
		{
			36,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					6,
					151,
					121,
				},
				{
					14,
					152,
					122,
				},
			},
			0,
		},
		{
			36,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					6,
					75,
					47,
				},
				{
					34,
					76,
					48,
				},
			},
			0,
		},
		{
			36,
			High,
			dataEncoderType27To40,
			[]block{
				{
					46,
					54,
					24,
				},
				{
					10,
					55,
					25,
				},
			},
			0,
		},
		{
			36,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					2,
					45,
					15,
				},
				{
					64,
					46,
					16,
				},
			},
			0,
		},
		{
			37,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					17,
					152,
					122,
				},
				{
					4,
					153,
					123,
				},
			},
			0,
		},
		{
			37,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					29,
					74,
					46,
				},
				{
					14,
					75,
					47,
				},
			},
			0,
		},
		{
			37,
			High,
			dataEncoderType27To40,
			[]block{
				{
					49,
					54,
					24,
				},
				{
					10,
					55,
					25,
				},
			},
			0,
		},
		{
			37,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					24,
					45,
					15,
				},
				{
					46,
					46,
					16,
				},
			},
			0,
		},
		{
			38,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					4,
					152,
					122,
				},
				{
					18,
					153,
					123,
				},
			},
			0,
		},
		{
			38,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					13,
					74,
					46,
				},
				{
					32,
					75,
					47,
				},
			},
			0,
		},
		{
			38,
			High,
			dataEncoderType27To40,
			[]block{
				{
					48,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			0,
		},
		{
			38,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					42,
					45,
					15,
				},
				{
					32,
					46,
					16,
				},
			},
			0,
		},
		{
			39,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					20,
					147,
					117,
				},
				{
					4,
					148,
					118,
				},
			},
			0,
		},
		{
			39,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					40,
					75,
					47,
				},
				{
					7,
					76,
					48,
				},
			},
			0,
		},
		{
			39,
			High,
			dataEncoderType27To40,
			[]block{
				{
					43,
					54,
					24,
				},
				{
					22,
					55,
					25,
				},
			},
			0,
		},
		{
			39,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					10,
					45,
					15,
				},
				{
					67,
					46,
					16,
				},
			},
			0,
		},
		{
			40,
			Low,
			dataEncoderType27To40,
			[]block{
				{
					19,
					148,
					118,
				},
				{
					6,
					149,
					119,
				},
			},
			0,
		},
		{
			40,
			Medium,
			dataEncoderType27To40,
			[]block{
				{
					18,
					75,
					47,
				},
				{
					31,
					76,
					48,
				},
			},
			0,
		},
		{
			40,
			High,
			dataEncoderType27To40,
			[]block{
				{
					34,
					54,
					24,
				},
				{
					34,
					55,
					25,
				},
			},
			0,
		},
		{
			40,
			Highest,
			dataEncoderType27To40,
			[]block{
				{
					20,
					45,
					15,
				},
				{
					61,
					46,
					16,
				},
			},
			0,
		},
	}
)



// A dataEncoder encodes data for a particular QR Code version.
type dataEncoder struct {
	// Minimum & maximum versions supported.
	minVersion                   int
	maxVersion                   int

	// Character count lengths.
	numNumericCharCountBits      int
	numAlphanumericCharCountBits int
	numByteCharCountBits         int
}

var dataEncoderTypeMap = map[dataEncoderType]*dataEncoder{
	dataEncoderType1To9: &dataEncoder{
		minVersion:                   1,
		maxVersion:                   9,
		numNumericCharCountBits:      10,
		numAlphanumericCharCountBits: 9,
		numByteCharCountBits:         8,
	},
	dataEncoderType10To26: &dataEncoder{
		minVersion:                   10,
		maxVersion:                   26,
		numNumericCharCountBits:      12,
		numAlphanumericCharCountBits: 11,
		numByteCharCountBits:         16,
	},
	dataEncoderType27To40:&dataEncoder{
		minVersion:                   27,
		maxVersion:                   40,
		numNumericCharCountBits:      14,
		numAlphanumericCharCountBits: 13,
		numByteCharCountBits:         16,
	},
}

func GetDataEncoder(version int)*dataEncoder{
	switch {
	case version >=1 && version <= 9:
		return dataEncoderTypeMap[dataEncoderType1To9]
	case version >=10 && version <=26:
		return dataEncoderTypeMap[dataEncoderType10To26]
	case version >=27 && version <=40:
		return dataEncoderTypeMap[dataEncoderType27To40]
	default:
		panic("version not found")
	}
}

func (de *dataEncoder)CharCountBits(format int)int{
	switch format{
	case 1:
		return de.numNumericCharCountBits
	case 2:
		return de.numAlphanumericCharCountBits
	case 4:
		return de.numByteCharCountBits
	default:
		logger.Println("format",format)
		panic("format not found")
	}
	return 0
}