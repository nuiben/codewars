# <8 kyu> Sentence Smash
# by richardhsu

# Sentence Smash
# Write a function that takes an array of words and smashes them together into a sentence and returns the sentence. You can ignore any need to sanitize words or add punctuation, but you should add spaces between each word. Be careful, there shouldn't be a space at the beginning or the end of the sentence!

# Example
# ['hello', 'world', 'this', 'is', 'great']  =>  'hello world this is great'

# Thought Process:
# except for the last element of array, we add the element to a string variable with a space. If we are at the last element, do not include
# a space.

def smash(words):
    sentence = ""
    for x in range(len(words)):
        if x == len(words)-1:
            sentence = sentence + words[x]
        else:
            sentence = sentence + words[x] + " "
    return sentence
