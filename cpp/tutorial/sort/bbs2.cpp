#include <iostream>

using namespace std;
void BubleSort(int a[], int n){
	for(int i =0; i<n; i++){
		bool isSorted = true;
		for (int j=0; j<n-i-1; j++){
			if(a[j]>a[j+1]){
				isSorted = false;
				int temp = a[j];
				a[j] = a[j+1];
				a[j+1] = temp;
			}

		}
		if(isSorted){
				break;
			}
	}
} 

int main(int argc, char const *argv[])
{
	int  a[5]={3,4,1,6,2};
	BubleSort(a,5);
 	for (int i=0; i<5; i++){
 		cout <<a[i]<<endl;
 	}
}