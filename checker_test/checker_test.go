package checker_test

import (
	"gowebapp/checker"
	"strconv"
	"testing"
)

func TestLengthCheck(t *testing.T) {
	/*
		input 7文字
		condition 8文字
	*/
	// input 7文字
	value1 := "abcdefg"
	condition1 := 8
	// expected
	expected1 := false
	// verify
	actual1 := checker.LengthCheck(value1, condition1)
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		input 8文字
		condition 8文字
	*/
	// input
	value2 := "abcdefgh"
	condition2 := 8
	// expected
	expected2 := true
	// exercise
	actual2 := checker.LengthCheck(value2, condition2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/**
	input 9文字
	condition 8文字
	*/
	// input
	value3 := "abcdefghi"
	condition3 := 8
	// expected
	expected3 := true
	// exercise
	actual3 := checker.LengthCheck(value3, condition3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/**
	input 9文字
	condition 10文字
	*/
	// input
	value4 := "abcdefghi"
	condition4 := 10
	// expected
	expected4 := false
	// exercise
	actual4 := checker.LengthCheck(value4, condition4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		input 0文字
		condition 8文字
	*/
	// input
	value5 := ""
	condition5 := 8
	// expected
	expected5 := false
	// exercise
	actual5 := checker.LengthCheck(value5, condition5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		input 1文字
		condition 0文字
	*/
	// input
	value6 := "a"
	condition6 := 0
	// expected
	expected6 := true
	// exercise
	actual6 := checker.LengthCheck(value6, condition6)
	// verfiy
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		input 0文字
		condition 0文字
	*/
	// input
	value7 := ""
	condition7 := 0
	// expected
	expected7 := true
	// exercise
	actual7 := checker.LengthCheck(value7, condition7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}
}

func TestComboUpperLowerCase(t *testing.T) {
	/*
		小文字のみ
	*/
	// input
	value1 := "abcdefg"
	// expected
	expected1 := false
	// exercise
	actual1 := checker.ComboUpperLowerCase(value1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		大文字のみ
	*/
	// input
	value2 := "ABCDEFG"
	// expected
	expected2 := false
	// exercise
	actual2 := checker.ComboUpperLowerCase(value2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		大文字小文字混合
	*/
	// input
	value3 := "abcdABCD"
	// expected
	expected3 := true
	// exercise
	actual3 := checker.ComboUpperLowerCase(value3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		全角英字小文字大文字混合
	*/
	// input
	value4 := "ａｂｃｄＡＢＣＤ"
	// expected
	expected4 := false
	// exercise
	actual4 := checker.ComboUpperLowerCase(value4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		半角数字
	*/
	// input
	value5 := "1234567890"
	// expected
	expected5 := false
	// exercise
	actual5 := checker.ComboUpperLowerCase(value5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		半角記号
	*/
	// input
	value6 := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	// expected
	expected6 := false
	// exercise
	actual6 := checker.ComboUpperLowerCase(value6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		半角英字小文字大文字、半角数字、半角記号混合
	*/
	// input
	value7 := "abcABC01234!.:@;"
	// expected
	expected7 := true
	// exercise
	actual7 := checker.ComboUpperLowerCase(value7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		半角英字小文字大文字、全角英字混合
	*/
	// input
	value8 := "abcABCａｂｃＡＢＣ"
	// expected
	expected8 := true
	// exercise
	actual8 := checker.ComboUpperLowerCase(value8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		未入力
	*/
	// input
	value9 := ""
	// expected
	expected9 := false
	// exercise
	actual9 := checker.ComboUpperLowerCase(value9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}
}

func TestComboCharType(t *testing.T) {
	/*
		半角英字、半角数字
		condition 2
	*/
	// input
	value1 := "abcde12345"
	condition1 := 2
	// expected
	expected1 := true
	// exercise
	actual1 := checker.ComboCharaType(value1, condition1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		半角英字、半角記号
		condition 2
	*/
	// input
	value2 := "abcde@:[]"
	condition2 := 2
	// expected
	expected2 := true
	// exercise
	actual2 := checker.ComboCharaType(value2, condition2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		半角記号、半角数字
		condition 2
	*/
	// input
	value3 := "[@p:;1234567890"
	condition3 := 2
	// expected
	expected3 := true
	// exercise
	actual3 := checker.ComboCharaType(value3, condition3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		半角英字、半角数字、半角記号
		condition 2
	*/
	// input
	value4 := "abcABC0123456789[@];,'()"
	condition4 := 2
	// expected
	expected4 := true
	// exercise
	actual4 := checker.ComboCharaType(value4, condition4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		半角英字、半角カナ
		condition 2
	*/
	// input
	value5 := "abcdeｱｲｳｴｵﾜｦﾝ"
	condition5 := 2
	// expected
	expected5 := false
	// exercise
	actual5 := checker.ComboCharaType(value5, condition5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		半角英字、半角数字、半角記号、半角カナ
		condition 2
	*/
	// input
	value6 := "abcdeABCDE1234567890:;@'(&%ｱｲｳｴｵﾜｦﾝ"
	condition6 := 2
	// expected
	expected6 := true
	// exercise
	actual6 := checker.ComboCharaType(value6, condition6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		半角英字
		condition 2
	*/
	// input
	value7 := "abcdeABCDE"
	condition7 := 2
	// expected
	expected7 := false
	// exercise
	actual7 := checker.ComboCharaType(value7, condition7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		半角数字
		condition 2
	*/
	// input
	value8 := "0123456789"
	condition8 := 2
	// expected
	expected8 := false
	// exercise
	actual8 := checker.ComboCharaType(value8, condition8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		半角記号
		condition 2
	*/
	// input
	value9 := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	condition9 := 2
	// expected
	expected9 := false
	// exercise
	actual9 := checker.ComboCharaType(value9, condition9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}

	/*
		半角英字、半角数字
		condition 3
	*/
	// input
	value10 := "abcde12345"
	condition10 := 3
	// expected
	expected10 := false
	// exercise
	actual10 := checker.ComboCharaType(value10, condition10)
	// verify
	if actual10 != expected10 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual10), strconv.FormatBool(expected10))
	}

	/*
		入力無し
		condition 2
	*/
	// input
	value11 := ""
	condition11 := 2
	// expected
	expected11 := false
	// exercise
	actual11 := checker.ComboCharaType(value11, condition11)
	// verify
	if actual11 != expected11 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual11), strconv.FormatBool(expected11))
	}

	/*
		半角英字
		condition 0
	*/
	// input
	value12 := "abcXYZ"
	condition12 := 0
	// expected
	expected12 := true
	// exercise
	actual12 := checker.ComboCharaType(value12, condition12)
	// verify
	if actual12 != expected12 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual12), strconv.FormatBool(expected12))
	}

	/*
		入力無し
		condition 0
	*/
	// input
	value13 := ""
	condition13 := 0
	// expected
	expected13 := true
	// exercise
	actual13 := checker.ComboCharaType(value13, condition13)
	// verify
	if actual13 != expected13 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual13), strconv.FormatBool(expected13))
	}
}

func TestContinuousChar(t *testing.T) {
	/*
		先頭半角英字3文字連続
		condition 3
	*/
	// input
	value1 := "aaabcdefji[@;(=695"
	condition1 := 3
	// expected
	expected1 := false
	// exercise
	actual1 := checker.ContinuousChar(value1, condition1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		中間半角英字3文字連続
		condition 3
	*/
	// input
	value2 := "abcccdjilacb1235@')&]"
	condition2 := 3
	// expected
	expected2 := false
	//exercise
	actual2 := checker.ContinuousChar(value2, condition2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		末尾半角英字3文字連続
		condition 3
	*/
	// input
	value3 := "abcd1235@]'=&ccc"
	condition3 := 3
	// expected
	expected3 := false
	// exercise
	actual3 := checker.ContinuousChar(value3, condition3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		半角数字3文字連続
		condition 3
	*/
	// input
	value4 := "ajil1380;@333:ab"
	condition4 := 3
	// expected
	expected4 := false
	// exercise
	actual4 := checker.ContinuousChar(value4, condition4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		半角記号3文字連続
		condition 3
	*/
	// input
	value5 := "abcde12356:(&=@@@ba"
	condition5 := 3
	// expected
	expected5 := false
	// exercise
	actual5 := checker.ContinuousChar(value5, condition5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		全角英字3文字連続
		condition 3
	*/
	// input
	value6 := "ａｂｃＸＹＺＱＱＱＺＹＡ"
	condition6 := 3
	// expected
	expected6 := false
	// exercise
	actual6 := checker.ContinuousChar(value6, condition6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		半角英字4文字連続
		condition 3
	*/
	// input
	value7 := "abcddddxyzZ:1356"
	condition7 := 3
	// expected
	expected7 := false
	// exercise
	actual7 := checker.ContinuousChar(value7, condition7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		半角英字2文字連続
		condition 3
	*/
	// input
	value8 := "aacab3jil:3"
	condition8 := 3
	// expected
	expected8 := true
	// exercise
	actual8 := checker.ContinuousChar(value8, condition8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		半角英字3文字出現非連続
		condition 3
	*/
	// input
	value9 := "abcega,70a1]"
	condition9 := 3
	// expected
	expected9 := true
	// exercise
	actual9 := checker.ContinuousChar(value9, condition9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}

	/*
		半角英字2文字連続
		condition 2
	*/
	// input
	value10 := "aacab3jil:3"
	condition10 := 2
	// expected
	expected10 := false
	// exercise
	actual10 := checker.ContinuousChar(value10, condition10)
	// verify
	if actual10 != expected10 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual10), strconv.FormatBool(expected10))
	}

	/*
		未入力
		condition 3
	*/
	// input
	value11 := ""
	condition11 := 3
	// expected
	expected11 := true
	// exercise
	actual11 := checker.ContinuousChar(value11, condition11)
	// verify
	if actual11 != expected11 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual11), strconv.FormatBool(expected11))
	}

	/*
		入力あり
		condition 0
	*/
	// input
	value12 := "abc60ji]@"
	condition12 := 0
	// expected
	expected12 := false
	// exercise
	actual12 := checker.ContinuousChar(value12, condition12)
	// verify
	if actual12 != expected12 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual12), strconv.FormatBool(expected12))
	}

	/*
		入力無し
		condition 0
	*/
	// input
	value13 := ""
	condition13 := 0
	// expected
	expected13 := true
	// exercise
	actual13 := checker.ContinuousChar(value13, condition13)
	// verify
	if actual13 != expected13 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual13), strconv.FormatBool(expected13))
	}

	/*
		半角英字3文字連続2回
		condition 3
	*/
	// input
	value14 := "abcccjiLLL"
	condition14 := 3
	// expected
	expected14 := false
	// exercise
	actual14 := checker.ContinuousChar(value14, condition14)
	// verify
	if actual14 != expected14 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual14), strconv.FormatBool(expected14))
	}
}

func TestCommonWords(t *testing.T) {
	/*
		辞書に登録した単語と完全一致
		password
	*/
	// input
	value1 := "password"
	// expected
	expected1 := false
	// exercise
	actual1 := checker.CommonWords(value1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		辞書に登録した単語を先頭に含む
		password
	*/
	value2 := "password123:jiAbo@["
	// expected
	expected2 := false
	// exercise
	actual2 := checker.CommonWords(value2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		辞書に登録した単語を中間に含む
		password
	*/
	// input
	value3 := "132:ajipassword[]()="
	// expected
	expected3 := false
	// exercise
	actual3 := checker.CommonWords(value3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		辞書に登録した単語を末尾に含む
		football
	*/
	// input
	value4 := "ji:123]o()football"
	// expected
	expected4 := false
	// exercise
	actual4 := checker.CommonWords(value4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		辞書に登録された単語（大文字）を含む
		password
	*/
	// input
	value5 := "PASSWORDbaji*AC7-"
	// expected
	expected5 := false
	// exercise
	actual5 := checker.CommonWords(value5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		辞書に登録された単語（小文字大文字混合）を含む
		password
	*/
	// input
	value6 := ",ajiAPaSsWorDji]132"
	// expected
	expected6 := false
	// exercise
	actual6 := checker.CommonWords(value6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		辞書に登録されていないワードを含む
	*/
	// input
	value7 := "ab,jpy6-][@EabC"
	// expected
	expected7 := true
	// exercise
	actual7 := checker.CommonWords(value7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		辞書に登録された単語の一部を含む
		passwor
	*/
	// input
	value8 := "Dba123passwor7-jkoD:,"
	// expected
	expected8 := true
	// exercsie
	actual8 := checker.CommonWords(value8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		未入力
	*/
	// input
	value9 := ""
	// expected
	expected9 := true
	// exercise
	actual9 := checker.CommonWords(value9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}

	/*
		辞書登録済みの単語2回
	*/
	// input
	value10 := "password,ji;acPaSSWORD"
	// expected
	expected10 := false
	// exercise
	actual10 := checker.CommonWords(value10)
	// verify
	if actual10 != expected10 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual10), strconv.FormatBool(expected10))
	}
}

func TestGetSatisfiedCondition(t *testing.T) {
	/*
		充足：〇 不足：×
		〇：LengthCheck
		〇：ComboUpperLowerCase
		〇：ComboCharaType
		〇：ContinuousChar
		〇：CommonWords
	*/
	// input
	value1 := "bajiA:c]70213ACk"
	// expected
	expectedInt1 := 5
	expectedMsgLen1 := 0
	// exercise
	actualInt1, actualMsg1 := checker.GetSatisfiedCondition(value1)
	// verify
	if actualInt1 != expectedInt1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(actualInt1), strconv.Itoa(expectedInt1))
	}
	if len(actualMsg1) != expectedMsgLen1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(len(actualMsg1)), strconv.Itoa(expectedMsgLen1))
	}

	/*
		充足：〇 不足：×
		×：LengthCheck
		×：ComboUpperLowerCase
		×：ComboCharaType
		〇：ContinuousChar
		×：CommonWords
	*/
	// input
	value2 := "welcome"
	// expected
	expectedInt2 := 1
	expectedMsgLen2 := 4
	expectedMsg2_1 := "8文字以上にしてください"
	expectedMsg2_2 := "大文字と小文字の組み合わせにしてください"
	expectedMsg2_3 := "文字種（英字、数字、記号等）を組み合わせましょう"
	expectedMsg2_4 := "一般的な単語を使用するのはやめましょう"
	// exercise
	actualInt2, actualMsg2 := checker.GetSatisfiedCondition(value2)
	// verify
	if actualInt2 != expectedInt2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(actualInt2), strconv.Itoa(expectedInt2))
	}
	if len(actualMsg2) != expectedMsgLen2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(len(actualMsg2)), strconv.Itoa(expectedMsgLen2))
	}
	if actualMsg2[0] != expectedMsg2_1 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2[0], expectedMsg2_1)
	}
	if actualMsg2[1] != expectedMsg2_2 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2[1], expectedMsg2_2)
	}
	if actualMsg2[2] != expectedMsg2_3 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2[2], expectedMsg2_3)
	}
	if actualMsg2[3] != expectedMsg2_4 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2[3], expectedMsg2_4)
	}

	/*
		充足：〇 不足：×
		〇：LengthCheck
		〇：ComboUpperLowerCase
		〇：ComboCharaType
		×：ContinuousChar
		〇：CommonWords
	*/
	// input
	value3 := "aaacJI:346CC"
	// expected
	expectedInt3 := 4
	expectedMsgLen3 := 1
	expectedMsg3 := "同じ文字を連続して使用するのはやめましょう"
	// exercise
	actualInt3, actualMsg3 := checker.GetSatisfiedCondition(value3)
	// verify
	if actualInt3 != expectedInt3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(actualInt3), strconv.Itoa(expectedInt3))
	}
	if len(actualMsg3) != expectedMsgLen3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.Itoa(len(actualMsg3)), strconv.Itoa(expectedMsgLen3))
	}
	if actualMsg3[0] != expectedMsg3 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg3[0], expectedMsg3)
	}
}
