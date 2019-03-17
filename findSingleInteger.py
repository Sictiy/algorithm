# encoding:utf-8
''' 
找出整形数组里个数不为3个数字，只存在一个
每位对3取余 00->01->10->11(00) 
'''
def find(array):
    one = 0
    two = 0
    for i in range(0, len(array), 1):
        # 加array[i]以后的取余结果
        two = two | one & array[i]
        one = one ^ array[i]
        # 把11变成00
        three = one & two
        two = two ^ three
        one = one ^ three
    print(one)
    print(two)
    print(one|two)

array = [1,1,1,3,3,3,4,4,4,6,6]
find(array)