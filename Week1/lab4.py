# Task1 - Binary -> ASCII 
binary = '01010011 01000101 01000011'
chars = binary.split(' ')
ascii_string = ''.join(chr(int(c, 2)) for c in chars)
print(f"Binary: {binary}, ASCII: {ascii_string}")
# Output: SEC

for c in chars:
    print(chr(int(c, 2)), end='')
print()

# Task2 - Hexadecimal -> ASCII
hex = '0x52 0x45 0x44'
chars = hex.split(' ')
ascii_string = ''.join(chr(int(c, 16)) for c in chars)
print(f"Hexadecimal: {hex}, ASCII: {ascii_string}")
# Output: RED

# Taks3 - XOR Decryption
ciphertext = '1110'
key = '1010'

for c, k in zip(ciphertext, key):
    if c != k:
        print('1',end='')
    else:
        print('0',end='')
print()

# Final secret phrase = SEC + RED + 0100 + Chanserey_Savon