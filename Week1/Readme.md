1. Decimal to Hexadecimal

To convert a decimal to hexdecimal, repeatedly divide the number by 16, record the quotient and remainder (0-9 or A-F for 10-15), and collect the remainders in reverse order.

Example: Convert 42 to hexadecimal
- 42 divide 16 = 2 remainder 10 (10 = A in hex)
- 2 divide 16 = 0 remainder 2
- stope when quotient is 0
- Read remainders from bottom to top: 2A

2. Binary to Hexadecimal 

To convert binary to hexadecimal, group the binary digits into sets of 4 (starting from the right), convert eac h group to its hex equivalent (0-9, A-F)

Example: Convert binary 101010 to hexadecimal
- Group binary: 0010 1010 (pad with leading zeros if needed)
- Convert each group:
    . 0010 = 2 
    . 1010 = 10 = A

- Combine: 2A

So ,binary 101010 is 2A in hexadecimal

ASCII Table
a : 097, z: 122 (+25)
A : 065, Z: 090 (+25)
But we can find this value by using ord. Example: C = ord('C')


XOR is one of the simplest building blocks of encryption. It's also reversible, which makes it powerful.

- The zip () function pairs corresponding characters from the ciphertext and key into tuples.

Example:
- ciphertext = '0110' -> characters: 0, 1, 1, 0
- key = '1010' -> characters: 1, 0, 1, 0
- zip creates pairs: ('0', '1'), ('1', '0'), ('1', '1'), ('0', '0')

Note: if ciphertext = 1010 want to convert to string use str(), if don't want to use just put this ''

- The split() use for:
    hex = '0x52 0x45 0x44'
    chars = hex.split(' ')
    change to ["0x52", "0x45", "0x44"]