def solution():
    f = file.open("input.txt", "r")
    prev_line = 0
    counter = -1
    for line in f:
        if line > prev_line:
            counter +=1
        
        prev_line = line
    
    return counter

if __name__ == "__main__":
    solution()
