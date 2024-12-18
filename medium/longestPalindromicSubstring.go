package main

import "fmt"

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	maxLen := 0
	result := ""
	longest := func(l, r int) {
		for l >= 0 && r < length && s[l] == s[r] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				result = s[l : r+1]
			}
			l--
			r++
		}
	}
	for i := 0; i < length; i++ {
		longest(i, i)
		longest(i, i+1)
	}
	return result
}

func main() {
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Print(longestPalindrome(s))
}
