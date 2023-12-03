class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        return sum([len(str(i)) % 2 == 0 for i in nums])
    # giải thích: đếm số lượng số có số chữ số chẵn
    # cách làm: chuyển số thành chuỗi, đếm số lượng kí tự trong chuỗi, nếu chia hết cho 2 thì là số chẵn
    # cách làm khác: dùng logarit, logarit cơ số 10 của số đó, nếu chia hết cho 2 thì là số chẵn
    # hiện thực:

    def findNumbers(self, nums: List[int]) -> int:
        return sum([int(math.log10(i)) % 2 == 1 for i in nums])
    # =1 có nghĩa là true
    # =0 có nghĩa là false


def findNumbers(self, nums: List[int]) -> int:
    result = 0
    for i in nums:
        if self.countN(i) % 2 == 0:
            result += 1
    return result


def countN(self, number: int) -> int:
    result = 0
    while number > 0:
        result += 1
        number = number // 10
    return result


class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        count = 0
        for num in nums:
            if num > 9 and num < 100 or num > 999 and num < 10000 or num == 100000:
                count += 1
        return count
