package numUniqueEmails

import (
	"fmt"
	"strings"
)

/*
Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?



Example 1:

Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails



Note:

    1 <= emails[i].length <= 100
    1 <= emails.length <= 100
    Each emails[i] contains exactly one '@' character.
    All local and domain names are non-empty.
    Local names do not start with a '+' character.


*/

// Runtime: 8 ms, faster than 98.10% of Go online submissions for Unique Email Addresses.
// Memory Usage: 6 MB, less than 44.71% of Go online submissions for Unique Email Addresses.

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, email := range emails {
		var left string
		strs := strings.Split(email, "@")
		idx := strings.Index(strs[0], "+")
		if idx != -1 {
			left = strs[0][:idx]
		} else {
			left = strs[0]
		}
		res := strings.Join(strings.Split(left, "."), "") + "@" + strs[1]
		fmt.Println(res)
		m[res] = true
	}
	return len(m)
}
