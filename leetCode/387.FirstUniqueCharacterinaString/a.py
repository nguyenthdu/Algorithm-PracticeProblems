class Solution:
    def firstUniqChar(self, s: str) -> int:
        char_dict = {}
        for i in range(len(s)):
            if s[i] in char_dict:
                char_dict[s[i]] = -1#-1 means duplicate
            else:
                char_dict[s[i]] = i#i means index
        for i in range(len(s)):
            if char_dict[s[i]] != -1:
                return char_dict[s[i]]
        return -1
