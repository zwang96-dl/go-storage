def wiggleSort(nums):
    temp = nums[:]
    temp.sort()

    for i in range(1, len(nums), 2):
        nums[i] = temp.pop()
    
    for i in range(0, len(nums), 2):
        nums[i] = temp.pop()