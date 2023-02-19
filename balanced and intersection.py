## todo Task 2
# def is_balanced(expression):
#     all = []
#     open_brackets = ['(', '[', '{']
#     close_brackets = [')', ']', '}']
#     for char in expression:
#         if char in open_brackets:
#             all.append(char)
#         elif char in close_brackets:
#             if not all:
#                 return False
#             if open_brackets.index(all[-1]) == close_brackets.index(char):
#                 all.pop()
#             else:
#                 return False
#     return not all
#
#
# print(is_balanced("()[{}"))

## todo Task 3
def intersection(nums1, nums2):
    hash_table = {}
    result = []
    for num in nums1:
        hash_table[num] = 1
    for num in nums2:
        if num in hash_table and hash_table[num] == 1:
            result.append(num)
            hash_table[num] = 2
    return result


print(intersection([1, 2, 3, 2, 0], [5, 1, 2, 7, 3, 2]))