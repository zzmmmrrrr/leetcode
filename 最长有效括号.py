from collections import deque

def longestValid(s:str)->int:
    stack=[-1]
    ans=0

    for i,ch in enumerate(s):
        if ch=='(':
            stack.append(i)
        else:
            stack.pop()
            if  len(stack)==0:
                stack.append(i)
            else:
                ans=max(ans,i-stack[-1])
    return ans


s='()()))'
res=longestValid(s)
print(res)