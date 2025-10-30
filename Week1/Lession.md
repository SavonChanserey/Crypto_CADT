Introduction to Cryptography 101

- Terminology

• Encrypted/Decrypt: the process of converting plaintext into ciphertext using an algorithm + key.
• Encoding/Decoding: Transforming data into another format
• Cipher Text (C): The unreadable message after encryption
• Plaintext (P): The original readable message before encryption
• Key (K): A secret value used by an algorithm to encrypt or decrypt data.
• Private Key (SK) / Public Key (PK): A key for using in key exchange process
• PKI (Public Key Infrastructure): manages public keys and certificates for secure communication
• Asymmetric / Symmetric: Modern cryptography system. Asym(key pair). Sym(Same Key)
• Hash : A one-way function that converts data into a fixed-length value
• Digital Signature: Use hash and SK to prove authenticity and integrity.
• Cryptography: The science of designing algorithms to secure communication (focuses on building).
• Cryptanalysis: The practice of breaking cryptographic systems or finding weaknesses.
• Cryptology: It is the the broad study of codes or secret writing and included both cryptography and cryptanalysis.

- DAD/CIA Traids

These are the 3 pillars of information security:
• Confidentiality: no unauthorized
• Integrity: ensure the data no one change
• Availability: online 24h 

These are the opposite (what attackers try to do):
• Disclosure – breaking confidentiality (e.g., data leak, eavesdropping).
• Alteration – breaking integrity (e.g.,tampering with files or messages).
• Denial – breaking availability (e.g., DDOS attacks, shutting down a server)


- The era of Cryptography

- The Classical Era (Ancient - WWII)
    Marked by manual ciphers like the Caesar Cipher and mechanical devices like the Enigma machine. Security relied on the secrecy of the method
    itself.

- The Modern Era (1970s - Present)
    The digital revolution brought public-key cryptography (RSA), data encryption standards (DES, AES), and the secure internet. Security
    shifted to the secrecy of keys, not algorithms.

- The Quantum Era (Today & Tomorrow)
    The rise of quantum computers threatens to break modern encryption.This era is defined by the race to develop Post-Quantum Cryptography (PQC) and explore concepts like homomorphic encryption.

- Ancient Grace 5th (Sparta)
• A Scytale, an early device for encryption.

- 6th B.C
• Atbash Cipher

- 1st B.C Rome Empire (Julius Caesar)

• Caesar Cipher
Encryption:
C = (P + K) mod 26
Decryption:
P = (C - K) mod 26

- 15 - 16th
• Vigenere cipher(Blause de Vigenere)

- Before WW2
• Signal Code
    . Morse Code (encoding)

- Alan Turing is the person that breaking the Enigma.

- Japanese Purple Cipher (Type B Cipher Machine) Used in WW2
• Type B Enigma (Enigma I) Unbreakable Machine
• American cryptanalysts cracked PURPLE.

- Type of Cryptography System

    There are 3 types of cryptographic systems: Symmetric, Asymmetric,
    and Hashing. Officially, only 2 are encryption systems (Symmetric &
    Asymmetric), but in practice we often include Hashing as the third type.

• Hashing: md5, SHA, HMAC, etc...
• Symmetric: AES, DES, RC4, ChaCha20 (use 1 key)
• Asymmetric: RSA, ECC, Dh, PK, etc... (use 2 key) (public and private)


- Entropy?

    In cryptography, entropy is a measure of uncertainty,
    randomness, or unpredictability.
    • High entropy = secure secrets
    • Low entropy = weak, guessable secrets
    • Demonstration: flip the coin


- Cryptanalyst Basic

    Frequency Analysis
    Guess and Analysis
    Protocol and logic attack
    Brute-force attack
    Etc..


- Logic Gate

• And: 1 -> 1 = 1, otherwise = 0
• OR: 0 -> 0 = 0, otherwise = 1
• XOR: 1 -> 1 = 0, 0 -> 0 = 0, and 1 -> 0 = 1, 0 -> 1 = 1

