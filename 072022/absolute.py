# <8 kyu> Return Negative
# by Deantwo

def make_negative(number):
    print(number)
    if number == 0:
        return 0
    if number > 0:
        number = number * -1
    return number
