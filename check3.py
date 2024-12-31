char = input("Enter a single character: ")

if len(char) != 1:
    print("Error: Please enter exactly one character.")
else:
    ascii_value = ord(char)
    print(f"The ASCII value of '{char}' is {ascii_value}")
