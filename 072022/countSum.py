# <8 kyu> Count of positives / sum of negatives
# by Dentzil

# Given an array of integers.

# Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

# If the input is an empty array or is null, return an empty array.

# Example

# For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].


# Thought Process: If positive increment positive variable, if negative += value to running total held in negative variable.
# Edge case if passed array arr is empty then we simply return arr. Would be a great use case for error handling, but challenge requires
# empty to be returned only.

def count_positives_sum_negatives(arr):
    if arr == []:
        return arr
    pos = 0
    neg = 0
    for x in range(len(arr)):
        if arr[x] > 0:
            pos += 1
        elif arr[x] < 0:
            neg += arr[x]
    newArr = [pos, neg]

    return newArr
