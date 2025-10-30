# Task 2.1 - Encode Your name

# My name is "Chanserey"
C = ord('C')
binary = bin(C)[2:]
hexadecimal = hex(C)[2:].upper()
print(f"Character: C, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: C, Binary: 1000011, Hexadecimal: 43

h = ord('h')
binary = bin(h)[2:]
hexadecimal = hex(h)[2:].upper()        
print(f"Character: h, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: h, Binary: 1101000, Hexadecimal: 68

a = ord('a')
binary = bin(a)[2:]
hexadecimal = hex(a)[2:].upper()
print(f"Character: a, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: a, Binary: 1100001, Hexadecimal: 61

n = ord('n')
binary = bin(n)[2:]
hexadecimal = hex(n)[2:].upper()
print(f"Character: n, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: n, Binary: 1101110, Hexadecimal: 6E

s = ord('s')
binary = bin(s)[2:]
hexadecimal = hex(s)[2:].upper()
print(f"Character: s, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: s, Binary: 1110011, Hexadecimal: 73    

e = ord('e')
binary = bin(e)[2:]
hexadecimal = hex(e)[2:].upper()
print(f"Character: e, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: e, Binary: 1100101, Hexadecimal: 65

r = ord('r')
binary = bin(r)[2:]
hexadecimal = hex(r)[2:].upper()
print(f"Character: r, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: r, Binary: 1110010, Hexadecimal: 72

e = ord('e')
binary = bin(e)[2:]
hexadecimal = hex(e)[2:].upper()
print(f"Character: e, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: e, Binary: 1100101, Hexadecimal: 65    

y = ord('y')
binary = bin(y)[2:]
hexadecimal = hex(y)[2:].upper()
print(f"Character: y, Binary: {binary}, Hexadecimal: {hexadecimal}")
# Output: Character: y, Binary: 1111001, Hexadecimal: 79


# Task 2.2 - Decode a Secret Message

# What does 0x41 represent in ASCII?
enc1 = chr(0x41)
print(enc1)
# Output: A

# Decode this binary: 01001000 01001001
enc2 = '01001000 01001001'
result = ''.join(chr(int(b, 2)) for b in enc2.split())
print(result)
# Output: HI    

