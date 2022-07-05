# <8 kyu> Convert a string to an array
# by samjam48

# Write a function to split a string and convert it into an array of words.

# Examples (Input ==> Output):
# "Robin Singh" ==> ["Robin", "Singh"]

# "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

# Thought process: space is the key identifier where we split the string into substrings. Python has split built in
# but I only just discovered the simpler solution:

# def string_to_array(string):
#     return string.split(" ")

# My solution:

def string_to_array(s):
    word = ""
    arr = []
    for x in range(len(s)):
        if s[x].isspace():
            print(word)
            arr.append(word)
            word = ""
        else:
            word = word + s[x]
    arr.append(word)
    return arr
