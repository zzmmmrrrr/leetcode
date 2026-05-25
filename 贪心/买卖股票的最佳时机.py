prices=[7,6,4,3,1,10]

def solve(prices:list[int])->int:
    ans=0
    curMin=prices[0]

    for p in prices:
        curMin=min(curMin,p)
        ans=max(ans,p-curMin)

    return ans

print(solve(prices))