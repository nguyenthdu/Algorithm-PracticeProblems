class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        for i in ransomNote:
            if i not in magazine:
                return False
            else:
                magazine = magazine.replace(i, '', 1)
        return True
    
    def canConstruct2(self, ransomNote: str, magazine: str) -> bool:
        # chuyển ransomNote  thành set để loại bỏ các ký tự trùng lặp tránh lặp lại 1 ký tự nhiều lần
        for i in set(ransomNote):
            # Kiểm tra số lần xuất hiện của ký tự i trong ransomNote có lớn hơn số lần xuất hiện của ký tự i trong magazine hay không
            if ransomNote.count(i) > magazine.count(i):#không có đủ ký tự trong magazine để tạo thành ransomNote
                return False
        return True
    
print(Solution().canConstruct("aa", "aab"))
print(Solution().canConstruct2("aa", "aab"))

