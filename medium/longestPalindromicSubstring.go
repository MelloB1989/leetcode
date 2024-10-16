package main

import "fmt"

func reverseString(s string) string {
	if len(s) == 0 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}

func checkPalindrom(s string) bool {
	return reverseString(s) == s
}

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	maxLen := 1
	maxStr := s[0:1]
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if checkPalindrom(s[i:j]) && j-i > maxLen {
				maxLen = j - i
				maxStr = s[i:j]
			}
		}
	}
	return maxStr
}

func main() {
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Print(longestPalindrome(s))
}
