class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        return "".join(word1) == "".join(word2)

    # using string builder
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        return self.stringBuilder(word1) == self.stringBuilder(word2)
    
    def stringBuilder(self, word: List[str]) -> str:
        sb = []
        for c in word:
            sb.append(c)
        return "".join(sb)
        
    # fatser
    