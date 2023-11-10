#include <iostream>

#include <vector>

void quickSort(std::vector<int>& arr){
	int left = 0;
	int right = arr.size()-1;

	int pivot = arr[(left+right)/2];
	int i=left, j=right;
	while(i<j){
		while(arr[i]<arr[pivot]){
			i++;
		}
		while(arr[i]>arr[pivot]){
			j--;
		}
		if (i<=j){
			std::swap(arr[i],arr[j]);
			i++;
			j--;
		}
	}
	


}
int main() {
    std::vector<int> arr = {64, 34, 25, 12, 22, 11, 90};
    std::cout << "Danh sách trước khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    quickSort(arr);
    std::cout << "Danh sách sau khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;

    return 0;
}