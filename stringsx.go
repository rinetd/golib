package golib

// //字符串处理包，对标准库字符串的补充

// import (
// 	"bytes"
// 	"fmt"
// 	"math/rand"
// 	"strconv"
// 	"strings"
// 	"time"

// 	// iox "github.com/yudeguang/iox"
// )

// //返回第一次出现sep之后的字串符
// func After(s, sep string) string {
// 	if s == "" || sep == "" {
// 		return s
// 	}
// 	pos := strings.Index(s, sep)
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[pos+len(sep):]
// }

// //返回最后一次出现sep之后的字符串
// func AfterLast(s, sep string) string {
// 	if sep == "" || s == "" {
// 		return s
// 	}
// 	pos := strings.LastIndex(s, sep)
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[pos+len(sep):]
// }

// //返回第N次出现sep之后的字符串
// func AfterNSep(s, sep string, nTimes int) string {
// 	if sep == "" || s == "" || nTimes <= 0 {
// 		return s
// 	}
// 	f := iox.New(strings.NewReader(s))
// 	pos := int(f.IndexN(0, []byte(sep), nTimes))
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[pos+len(sep):]
// }

// //返回第一次出现sep之前的字符串
// func Before(s, sep string) string {
// 	if s == "" || sep == "" {
// 		return s
// 	}
// 	pos := strings.Index(s, sep)
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[:pos]
// }

// //返回最后一次出现sep之前的字符串
// func BeforeLast(s, sep string) string {
// 	if sep == "" || s == "" {
// 		return s
// 	}
// 	pos := strings.LastIndex(s, sep)
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[:pos]
// }

// //返回第N次出现sep之前的字符串
// func BeforeNSep(s, sep string, nTimes int) string {
// 	if sep == "" || s == "" || nTimes <= 0 {
// 		return s
// 	}
// 	f := iox.New(strings.NewReader(s))
// 	pos := int(f.IndexN(0, []byte(sep), nTimes))
// 	if pos == -1 {
// 		return ""
// 	}
// 	return s[:pos]
// }

// //返回第一次出现在两个字符串接之间的字符串
// func Between(s, begin, end string) string {
// 	if s == "" || begin == "" || end == "" {
// 		return ""
// 	}
// 	beginPos := strings.Index(s, begin)
// 	if beginPos != -1 {
// 		f := iox.New(strings.NewReader(s))
// 		endPos := int(f.IndexGen(int64(beginPos+len(begin)), int64(len(s)-1), []byte(end)))
// 		if endPos != -1 {
// 			return s[beginPos+len(begin) : endPos]
// 		}
// 	}
// 	return ""
// }

// //返回左侧N个字符
// func Left(s string, n int) string {
// 	if n <= 0 || s == "" {
// 		return ""
// 	}
// 	runes := []rune(s)
// 	if len(runes) <= n {
// 		return s
// 	}
// 	return string(runes[0:n])
// }

// //返回右侧N个字符
// func Right(s string, n int) string {
// 	if n <= 0 || s == "" {
// 		return ""
// 	}
// 	runes := []rune(s)
// 	if len(runes) <= n {
// 		return s
// 	}
// 	return string(runes[len(runes)-n:])
// }

// //用分隔符sep把若干个字符拼接在一起,实际为strings.Join的变体形式
// func JoinStrings(sep string, args ...string) string {
// 	return strings.Join(args, sep)
// }

// //用分隔符号把若干个数字排接在一起
// func JoinInts(sep string, args ...int) string {
// 	l := len(args)
// 	switch l {
// 	case 0:
// 		return ""
// 	case 1:
// 		return strconv.Itoa(args[0])
// 	case 2:
// 		return strconv.Itoa(args[0]) + sep + strconv.Itoa(args[1])
// 	case 3:
// 		return strconv.Itoa(args[0]) + sep + strconv.Itoa(args[1]) + sep + strconv.Itoa(args[2])
// 	}
// 	var buffer bytes.Buffer
// 	//前面若干条中间要加sep
// 	for i := 0; i < l-1; i++ {
// 		buffer.WriteString(strconv.Itoa(args[i]) + sep)
// 	}
// 	//最后次不加sep
// 	buffer.WriteString(strconv.Itoa(args[l-1]))
// 	return buffer.String()
// }

// //用分隔符sep把若干个字符或int,double等类型数据拼接在一起,实际为strings.Join的变体形式
// func JoinInterface(sep string, args ...interface{}) string {
// 	l := len(args)
// 	switch l {
// 	case 0:
// 		return ""
// 	case 1:
// 		return fmt.Sprint(args[0])
// 	case 2:
// 		return fmt.Sprint(args[0]) + sep + fmt.Sprint(args[1])
// 	case 3:
// 		return fmt.Sprint(args[0]) + sep + fmt.Sprint(args[1]) + sep + fmt.Sprint(args[2])
// 	}
// 	var buffer bytes.Buffer
// 	//前面若干条中间要加sep
// 	for i := 0; i < l-1; i++ {
// 		buffer.WriteString(fmt.Sprint(args[i]) + sep)
// 	}
// 	//最后次不加sep
// 	buffer.WriteString(fmt.Sprint(args[l-1]))
// 	return buffer.String()
// }

// //返回倒序字符串
// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
// 		runes[from], runes[to] = runes[to], runes[from]
// 	}
// 	return string(runes)
// }

// //返回把中间字符按一定规则替换后的字符串接
// //CenterPad("hello", 10, "*") => "he*****llo"
// func CenterPad(s string, length int, pad string) string {
// 	l := lenRune(s)
// 	if length <= l {
// 		return Left(s, length)
// 	}
// 	//对应mysql中是返回NULL
// 	if pad == "" && length >= l {
// 		return ""
// 	}
// 	//取得重复若干次pad之后剩余的文本"12312", Rightpad("hello", 10, "123") => "hello12312"
// 	pads := Right(Rightpad(s, length, pad), length-l)
// 	//判断s长度是单数还是双数
// 	remainder := l % 2
// 	quotient := l / 2
// 	if remainder == 0 {
// 		return Left(s, quotient) + pads + Right(s, quotient)
// 	} else {
// 		return Left(s, quotient) + pads + Right(s, quotient+1)
// 	}
// }

// //返回两侧字符按一定规则替换后的字符串接
// // LeftRightPad("hello", 4, " ")    => "hell"
// // LeftRightPad("hello", 10, " ")   => "  hello   "
// // LeftRightPad("hello", 10, "123") => "12hello123"
// func LeftRightPad(s string, length int, pad string) string {
// 	l := lenRune(s)
// 	if length <= l {
// 		return Left(s, length)
// 	}
// 	//对应mysql中是返回NULL
// 	if pad == "" && length >= l {
// 		return ""
// 	}
// 	//取得重复若干次pad之后剩余的文本"12312", Rightpad("hello", 10, "123") => "hello12312"
// 	pads := Right(Rightpad(s, length, pad), length-l)
// 	//判断pads长度是单数还是双数
// 	remainder := (length - l) % 2
// 	quotient := (length - l) / 2
// 	if remainder == 0 {
// 		return Left(pads, quotient) + s + Right(pads, quotient)
// 	} else {
// 		return Left(pads, quotient) + s + Right(pads, quotient+1)
// 	}
// }

// //返回字符串str，右面用字符串padstr填补直到str是len个字符长,此函数与mysql中RPAD()行为保持一致
// // Rightpad("hello", 4, " ")    => "hello"
// // Rightpad("hello", 10, " ")   => "hello     "
// // Rightpad("hello", 10, "123") => "hello12312"
// func Rightpad(s string, length int, pad string) string {
// 	l := lenRune(s)
// 	if length <= l {
// 		return Left(s, length)
// 	}
// 	//对应mysql中是返回NULL
// 	if pad == "" && length >= l {
// 		return ""
// 	}
// 	for {
// 		if lenRune(s) >= length {
// 			return Left(s, length)
// 		}
// 		s = s + pad
// 	}
// }

// //返回字符串str，左面用字符串padstr填补直到str是len个字符长,,此函数与mysql中LPAD()行为保持一致
// // LeftPad("hello", 4, " ")    => "hello"
// // LeftPad("hello", 10, " ")   => "     hello"
// // LeftPad("hello", 10, "123") => "12312hello"
// func LeftPad(s string, length int, pad string) string {
// 	l := lenRune(s)
// 	if length <= l {
// 		return Right(s, length)
// 	}
// 	//对应mysql中是返回NULL
// 	if pad == "" && length >= l {
// 		return ""
// 	}
// 	for {
// 		if lenRune(s) >= length {
// 			return Right(s, length)
// 		}
// 		s = pad + s
// 	}
// }

// //返回s中的字符数
// func lenRune(s string) int {
// 	return len([]rune(s))
// }

// //返回随机打乱后的字符串
// func Rand(s string) string {
// 	if s == "" {
// 		return s
// 	}
// 	runes := []rune(s)
// 	rand.Seed(time.Now().UnixNano())
// 	randedSlice := rand.Perm(len(runes))
// 	newRunes := make([]rune, 0, len(runes))
// 	for _, randedIndex := range randedSlice {
// 		newRunes = append(newRunes, runes[randedIndex])
// 	}
// 	return string(newRunes)
// }

// const numbersAndLetters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
// const commaAndNumbersAndLetters = `,abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`

// //只保留数字和英文字母,删除其它类型字母及标点符号
// func NumbersLettersLeft(s string) string {
// 	runes := []rune(s)
// 	newRunes := make([]rune, 0, len(runes))
// 	for _, r := range runes {
// 		if strings.Contains(numbersAndLetters, string(r)) {
// 			newRunes = append(newRunes, r)
// 		}
// 	}
// 	return string(newRunes)
// }

// //只保留逗号以及数字和英文字母，因为逗号一般用于分隔文本
// func CommaNumbersLettersLeft(s string) string {
// 	runes := []rune(s)
// 	newRunes := make([]rune, 0, len(runes))
// 	for _, r := range runes {
// 		if strings.Contains(commaAndNumbersAndLetters, string(r)) {
// 			newRunes = append(newRunes, r)
// 		}
// 	}
// 	return string(newRunes)
// }
