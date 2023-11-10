// Cho một mảng gồm các phần tử 0 và 1. Cần đưa ra tổng khoảng cách của các kết nối giữa các cặp phần tử của mảng thỏa mãn các điều kiện sau:

// Chỉ có thể nối phần tử 0 và 1 với nhau
// Tất cả các phần tử đều có ít nhất một kết nối
// Một phần tử có thể nối với 1 hoặc nhiều phần tử khác
// Tổng độ dài (theo chỉ số mảng) kết nối là nhỏ nhất
// Dữ liệu nhập
// Dòng đầu tiên chứa số nguyên 
// �
// (
// 1
// ≤
// �
// ≤
// 1
// 0
// )
// t(1≤t≤10) - số lượng test.

// Với mỗi bộ test, dòng đầu tiên gồm số nguyên 
// �
// (
// 2
// ≤
// �
// ≤
// 2
// ×
// 1
// 0
// 3
// )
// n(2≤n≤2×10
// ​3
// ​​ ) - độ dài của mảng. Dòng thứ hai gồm 
// �
// n số nguyên chứa các số 0 hoặc 1. Dữ liệu đảm bảo hai số 0 và 1 xuất hiện ít nhất 1 lần trong mảng.

// Dữ liệu xuất
// Gồm 
// �
// T dòng cho 
// �
// T test cases, mỗi dòng in ra tổng độ dài để kết nối và thỏa mãn điều kiện là nhỏ nhất.

// Ví dụ
// input
// 2
// 9
// 0 1 1 1 0 0 1 0 0
// 3
// 0 1 0
// output
// 8
// 2
// Ở test đầu tiên, ta có các cặp kết nối tối ưu là : (0, 1); (0,2); (3,4); (5, 6); (6, 7); (6, 8). Tổng độ dài kết nối là 8.

#include <iostream>
#include <cmath>
using namespace std;
int main()
{
    int t;
   cin >> t;
    while (t--)
    {
         int n;
         cin >> n;
         int a[n];
         for (int i = 0; i < n; i++)
         {
              cin >> a[i];
         }
         int dem = 0;
         for (int i = 0; i < n; i++)
         {
              if (a[i] == 0)
              {
                dem++;
              }
         }
         cout << min(dem, n - dem) << endl;
    }
    return 0;
    
}
//output 4 1 false
//write



#include <iostream>
#include <cmath>
using namespace std;
int main()
{
    int t;
   cin >> t;
    while (t--)
    {
         int n;
         cin >> n;
         int a[n];
         for (int i = 0; i < n; i++)
         {
              cin >> a[i];
         }
         int dem = 0;
         for (int i = 0; i < n; i++)
         {
              if (a[i] == 0)
              {
                dem++;
              }
         }
         cout << min(dem, n - dem) << endl;
    }
    return 0;
    
}