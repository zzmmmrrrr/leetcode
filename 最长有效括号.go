package main
import (
	"fmt"
)
func main(){
	s:=")(()()))"
	fmt.Println(dp_method(s))
	fmt.Println(stack_method(s))
}
func dp_method(s string) int{
	n:=len(s)
	dp:=make([]int,n)
	ans:=0

	for i:=1;i<n;i++{
		if s[i]==')'{

			if s[i-1]=='('{//前一个字符是左括号
				dp[i]=2
				if i>=2{
					dp[i]+=dp[i-2]
				}
			} else if dp[i-1]>0{//前一个字符是右括号
				j:=i-dp[i-1]-1
				if j>=0&&s[j]=='('{
					dp[i]=dp[i-1]+2
					if j>0{
						dp[i]+=dp[j-1]
					}
				}
			}
			ans=max(dp[i],ans)

		}
	}
	return ans
}

//dp[i]，以s[i]结尾的最长有效括号串长度
//dp[i]
//()()()()
//(((()))
//0123456


func stack_method(s string) int{
	ans:=0
	stk:=[]int{-1}
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			stk=append(stk,i)
		} else {
			stk=stk[:len(stk)-1]
			if len(stk)==0{
				stk=append(stk,i)
			} else{
				ans=max(ans,i-stk[len(stk)-1])
			}
		}
	}
	return ans
}