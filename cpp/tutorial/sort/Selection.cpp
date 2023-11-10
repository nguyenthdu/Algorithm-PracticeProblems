#include <iostream>
#include <vector>
void Selection(std::vector<int>& arr){
	for(int i =0; i< arr.size()-1; i++){
		int min = i;
		for(int j=i+1; j<arr.size();j++)
			if (arr[min]>arr[j]) min =j;
		if (min !=i) std::swap(arr[min],arr[i]);
		
	}
}
int main() {
    std::vector<int> arr = {64, 34, 25, 12, 22, 11, 90};
    std::cout << "Danh sách trước khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    Selection(arr);
    std::cout << "Danh sách sau khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;

    return 0;
}
