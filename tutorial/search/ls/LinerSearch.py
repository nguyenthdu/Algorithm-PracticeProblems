def LSearch(list,n:int,key:int):
	flag = 0;
	for i in range(n):
		if list[i]==key:
			print(f'found at position: ',i)
			flag=1;
			break;
	if flag==0:
		print('not found')

a=[4,6,7,8,9]
LSearch(a,5,6)