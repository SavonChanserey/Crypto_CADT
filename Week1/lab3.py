# Task 3.1 - Simple XOR

ciphertext = '0110'
key = '1010'

plaintext = ''.join('1' if c != k else '0' for c, k in zip(ciphertext, key))
print(f"Ciphertext: {ciphertext}, Key: {key}, Plaintext: {plaintext}")
# Output: Ciphertext: 0110, Key: 1010, Plaintext: 1100

for c, k in zip(ciphertext, key):
    if c != k:
        print('1',end='')
    else:
        print('0',end='')
print()

# Task 3.2 - Longer Example

ciphertext = '110110'
key = '011011'

for c, k in zip(ciphertext, key):
    if c != k:
        print('1',end='')
    else:
        print('0',end='')
print()