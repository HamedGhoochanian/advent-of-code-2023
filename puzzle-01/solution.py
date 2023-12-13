with open('input2.txt','r') as f:
    lines = f.readlines()
lines = [l[:len(l)-1] for l in lines]
print(lines[-10:])

def check_for_words(line,i,reverse=False):
    if reverse:
        c3 = line[i-2:i+1]
        c4 = line[i-3:i+1]
        c5 = line[i-4:i+1]
    else:
        c3 = line[i:i+3]
        c4 = line[i:i+4]
        c5 = line[i:i+5]
    print(c3,c4,c5)
    if c3 == 'one':
        return 1
    if c3 == 'two':
        return 2
    if c3=='six':
        return 6
    if c4=='four':
        return 4
    if c4=='five':
        return 5
    if c4=='nine':
        return 9
    if c5=='three':
        return 3
    if c5=='seven':
        return 7
    if c5=='eight':
        return 8
    return None
    


values = []

for l in lines:
    a,b = 0,0
    for i,char in enumerate(l):
        if char in ['0','1','2','3','4','5','6','7','8','9']:
            a = int(char)
            break
        if char in ['o','t','f','s','e','n']:
            res = check_for_words(l,i)
            if res is not None:
                a = res
                break
    
    for i,char in enumerate(l[::-1]):
        if char in ['0','1','2','3','4','5','6','7','8','9']:
            b = int(char)
            break

        if char in ['e','o','r','x','n','t']:
            res = check_for_words(l,i,reverse=True)
            if res is not None:
                b = res
                break
    values.append(a*10+b)

print(values)

print(sum(values))