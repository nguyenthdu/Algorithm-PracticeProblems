class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        return max([sum(i) for i in accounts])

    def maximumWealth(self, accounts: List[List[int]]) -> int:#best
        return max(map(sum, accounts))
        #giải thích:
        #hàm max sẽ tìm giá trị lớn nhất trong list
        #hàm map sẽ áp dụng hàm sum vào từng phần tử của accounts
        #hàm sum sẽ tính tổng các phần tử trong list
        #ví dụ: sum([1,2,3]) = 6
        

    def maximumWealth(self, accounts: List[List[int]]) -> int:
        return max(map(lambda x: sum(x), accounts))

    def maximumWealth(self, accounts: List[List[int]]) -> int:
        return max(map(lambda x: sum(x), accounts))

    def maximumWealth(self, accounts: List[List[int]]) -> int:
        max = 0
        for i in accounts:
            if sum(i) > max:
                max = sum(i)
        return max
