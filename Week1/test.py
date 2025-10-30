encrpyt = 0xed
key = 0xad
decrypt = encrpyt ^ key
print(f"Encrypted: {hex(encrpyt)}, Key: {hex(key)}, Decrypted: {hex(decrypt)}")