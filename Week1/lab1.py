# Task 1.1 - Basic Conversions

# 42 (decimal) to binary and hexadecimal
decimal = 42
binary = bin(decimal)[2:]
print(binary)  
# Output: 101010

hexadecimal = hex(decimal)[2:].upper()
print(hexadecimal)
# Output: 2a

# 101101 (binary) to decimal and hexadecimal
binary = '101101'
decimal = int(binary, 2)
print(decimal)
# Output: 45

hex = hex(decimal)[2:].upper()
print(hex)
# Output: 2D

# 0x7F (hexadecimal) to decimal and binary
hex = '0x7F'
decimal = int(hex, 16)
print(decimal)
# Output: 127

binary = bin(decimal)[2:]
print(binary)
# Output: 1111111

# 2025 (decimal) to hexadecimal and binary
decimal = 2025
hexadecimal = hex(decimal)[2:].upper()
print(hexadecimal)
# Output: 7E9

binary = bin(decimal)[2:]
print(binary)
# Output: 111111001


# Task 1.2 - Fill The Blanks

Decimal = 65
binary = bin(Decimal)[2:]
hexadecimal = hex(Decimal)[2:].upper()
print(f"Decimal: {Decimal}, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Decimal: 65, Binary: 1000001, Hexadecimal: 41

Binary = '01010100'
Decimal = int(Binary, 2)
Hexadecimal = hex(Decimal)[2:].upper()
print(f"Binary: {Binary}, Decimal: {Decimal}, Hexadecimal: {Hexadecimal}")
# Output: Binary: 01010100, Decimal: 84, Hexadecimal:

Decimal = 155
binary = bin(Decimal)[2:]
hexadecimal = hex(Decimal)[2:].upper()
print(f"Decimal: {Decimal}, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Decimal: 155, Binary: 10011011, Hexadecimal: 9B

