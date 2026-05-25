package main
func find(x int) int{
	if parent[x]!=x{
		parent[x]=find(parent[x])
	}
	return parent[x]
}

func union(x,y int){
	rootx,rooty:=find(x),find(y)
	if rootx==rooty{
		return
	}
	if h[rootx]<h[rooty]{
		parent[rootx]=rooty
	} else if h[rootx]>h[rooty]{
		parent[rooty]=rootx
	} else{
		parent[rooty]=rootx
		h[rootx]++
	}
}
func main(){
	n:=10
	parent:=make([]int,n)
	h:=make([]int,n)
	for i:=0;i<n;i++{
		parent[i]=i
		h[i]=1
	}
	
}
