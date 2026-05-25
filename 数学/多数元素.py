def moore_vote(nums:list[int])->int:
    candidate_num=0
    vote_cnt=0
    for x in nums:
        if vote_cnt==0:
            candidate_num=x
        if x==candidate_num:
            vote_cnt+=1
        else:
            vote_cnt-=1

    return candidate_num

def hashmap(nums:list[int])->int:
    cnt={}
    for x in nums:
        cnt[x]=cnt.get(x,0)+1
        if cnt[x]>len(nums)//2:
            return x
        

    return -1

def main():
    nums=[1,2,3,4,4,4,4,4,5,6,4,4,4]
    print(moore_vote(nums))
    print(hashmap(nums))


main()
