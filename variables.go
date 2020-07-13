package gorean

/*
@Title
	korean token splitter
@Description
	ruby 'korean-string(https://github.com/bhumphreys/korean-string)' is transported with go
@Reference
- main reference
	- https://www.bada-ie.com/board/view/?page=9&uid=1782&category_code=&code=all
- sub reference
	1. http://www.w3c.or.kr/i18n/hangul-i18n/ko-code.html
*/

var (
	koreanDoubleJaumMap = map[string]bool{
		"ㄲ": true,
		"ㄸ": true,
		"ㅃ": true,
		"ㅆ": true,
		"ㅉ": true,
	}
	koreanCoupleJaumMap = map[string]string{
		"ㄱㄱ": "ㄲ",
		"ㄷㄷ": "ㄸ",
		"ㅂㅂ": "ㅃ",
		"ㅈㅈ": "ㅉ",
		"ㅅㅅ": "ㅆ",
		"ㄱㅅ": "ㄳ",
		"ㄴㅈ": "ㄵ",
		"ㄴㅎ": "ㄶ",
		"ㄹㄱ": "ㄺ",
		"ㄹㅁ": "ㄻ",
		"ㄹㅂ": "ㄼ",
		"ㄹㅅ": "ㄽ",
		"ㄹㅌ": "ㄾ",
		"ㄹㅍ": "ㄿ",
		"ㄹㅎ": "ㅀ",
		"ㅂㅅ": "ㅄ",
	}
	koreanCoupleJaumFirstMap = map[string]string{
		"ㄲ": "ㄱ",
		"ㄸ": "ㄷ",
		"ㅃ": "ㅂ",
		"ㅉ": "ㅈ",
		"ㅆ": "ㅅ",
		"ㄳ": "ㄱ",
		"ㄵ": "ㄴ",
		"ㄶ": "ㄴ",
		"ㄺ": "ㄹ",
		"ㄻ": "ㄹ",
		"ㄼ": "ㄹ",
		"ㄽ": "ㄹ",
		"ㄾ": "ㄹ",
		"ㄿ": "ㄹ",
		"ㅀ": "ㄹ",
		"ㅄ": "ㅂ",
	}
	koreanCoupleMoumFirstMap = map[string]string{
		"ㅢ": "ㅡ",
		"ㅝ": "ㅜ",
		"ㅟ": "ㅜ",
		"ㅞ": "ㅜ",
		"ㅙ": "ㅗ",
		"ㅚ": "ㅗ",
		"ㅘ": "ㅗ",
	}

	koranChosung = []string{
		"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ",
		"ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ",
		"ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ",
		"ㅋ", "ㅌ", "ㅍ", "ㅎ",
	}
	koreanChosungRuneToIndexMap = map[int64]int64{
		'ㄱ': 0,
		'ㄲ': 1,
		'ㄴ': 2,
		'ㄷ': 3,
		'ㄸ': 4,
		'ㄹ': 5,
		'ㅁ': 6,
		'ㅂ': 7,
		'ㅃ': 8,
		'ㅅ': 9,
		'ㅆ': 10,
		'ㅇ': 11,
		'ㅈ': 12,
		'ㅉ': 13,
		'ㅊ': 14,
		'ㅋ': 15,
		'ㅌ': 16,
		'ㅍ': 17,
		'ㅎ': 18,
	}
	koreanChosungHexToRuneMap = map[int64]rune{
		0x3131: 'ㄱ',
		0x3132: 'ㄲ',
		0x3134: 'ㄴ',
		0x3137: 'ㄷ',
		0x3138: 'ㄸ',
		0x3139: 'ㄹ',
		0x3141: 'ㅁ',
		0x3142: 'ㅂ',
		0x3143: 'ㅃ',
		0x3145: 'ㅅ',
		0x3146: 'ㅆ',
		0x3147: 'ㅇ',
		0x3148: 'ㅈ',
		0x3149: 'ㅉ',
		0x314a: 'ㅊ',
		0x314b: 'ㅋ',
		0x314c: 'ㅌ',
		0x314d: 'ㅍ',
		0x314e: 'ㅎ',
	}

	koreanJungsung = []string{
		"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ",
		"ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ",
		"ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ",
		"ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ",
		"ㅣ",
	}
	koreanJungsungRuneToIndexMap = map[int64]int64{
		'ㅏ': 0,
		'ㅐ': 1,
		'ㅑ': 2,
		'ㅒ': 3,
		'ㅓ': 4,
		'ㅔ': 5,
		'ㅕ': 6,
		'ㅖ': 7,
		'ㅗ': 8,
		'ㅘ': 9,
		'ㅙ': 10,
		'ㅚ': 11,
		'ㅛ': 12,
		'ㅜ': 13,
		'ㅝ': 14,
		'ㅞ': 15,
		'ㅟ': 16,
		'ㅠ': 17,
		'ㅡ': 18,
		'ㅢ': 19,
		'ㅣ': 20,
	}
	koreanJungsungHexToRuneMap = map[int64]rune{
		0x314f: 'ㅏ',
		0x3150: 'ㅐ',
		0x3151: 'ㅑ',
		0x3152: 'ㅒ',
		0x3153: 'ㅓ',
		0x3154: 'ㅔ',
		0x3155: 'ㅕ',
		0x3156: 'ㅖ',
		0x3157: 'ㅗ',
		0x3158: 'ㅘ',
		0x3159: 'ㅙ',
		0x315a: 'ㅚ',
		0x315b: 'ㅛ',
		0x315c: 'ㅜ',
		0x315d: 'ㅝ',
		0x315e: 'ㅞ',
		0x315f: 'ㅟ',
		0x3160: 'ㅠ',
		0x3161: 'ㅡ',
		0x3162: 'ㅢ',
		0x3163: 'ㅣ',
	}

	koreanJongsung = []string{
		"", "ㄱ", "ㄲ", "ㄳ", "ㄴ",
		"ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ",
		"ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ",
		"ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ",
		"ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ",
		"ㅌ", "ㅍ", "ㅎ",
	}
	koreanJongsungRuneToIndexMap = map[int64]int64{
		' ': 0,
		'ㄱ': 1,
		'ㄲ': 2,
		'ㄳ': 3,
		'ㄴ': 4,
		'ㄵ': 5,
		'ㄶ': 6,
		'ㄷ': 7,
		'ㄹ': 8,
		'ㄺ': 9,
		'ㄻ': 10,
		'ㄼ': 11,
		'ㄽ': 12,
		'ㄾ': 13,
		'ㄿ': 14,
		'ㅀ': 15,
		'ㅁ': 16,
		'ㅂ': 17,
		'ㅄ': 18,
		'ㅅ': 19,
		'ㅆ': 20,
		'ㅇ': 21,
		'ㅈ': 22,
		'ㅊ': 23,
		'ㅋ': 24,
		'ㅌ': 25,
		'ㅍ': 26,
		'ㅎ': 27,
	}
	koreanJongsungHexToRuneMap = map[int64]rune{
		0x0000: ' ',
		0x313a: 'ㄺ',
		0x313b: 'ㄻ',
		0x313c: 'ㄼ',
		0x313d: 'ㄽ',
		0x313e: 'ㄾ',
		0x313f: 'ㄿ',
		0x314a: 'ㅊ',
		0x314b: 'ㅋ',
		0x314c: 'ㅌ',
		0x314d: 'ㅍ',
		0x314e: 'ㅎ',
		0x3131: 'ㄱ',
		0x3132: 'ㄲ',
		0x3133: 'ㄳ',
		0x3134: 'ㄴ',
		0x3135: 'ㄵ',
		0x3136: 'ㄶ',
		0x3137: 'ㄷ',
		0x3139: 'ㄹ',
		0x3140: 'ㅀ',
		0x3141: 'ㅁ',
		0x3142: 'ㅂ',
		0x3144: 'ㅄ',
		0x3145: 'ㅅ',
		0x3146: 'ㅆ',
		0x3147: 'ㅇ',
		0x3148: 'ㅈ',
	}

	//koreanChosungLen  = int64(19) // int64(len(koreanCh))
	koreanJungsungLen = int64(21) // int64(len(JUNGSUNG))
	koreanJongsungLen = int64(28) // int64(len(JONGSUNG))
)

const (
	koreanStartHexCode int64 = 0xAC00 // 가
	koreanEndHexCode   int64 = 0xD7A3 // 힣

	chosungIndex  int = 0
	jungsungIndex int = 1
	jongsungIndex int = 2
)