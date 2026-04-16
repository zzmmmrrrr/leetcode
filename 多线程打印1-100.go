package main
import(
	"fmt"
	"sync"
)

func main(){
	cnt:=1
	chan1,chan2,chan3:=make(chan struct{},1),make(chan struct{},1),make(chan struct{},1)

	var wg sync.WaitGroup
	wg.Add(3)
	
	go func(){
		defer wg.Done()
		for {
			<-chan1
			if cnt>100{
				chan2<-struct{}{}
				return
			}
			fmt.Println(cnt)
			cnt++
			chan2<-struct{}{}
		}
	}()

	go func(){
		defer wg.Done()
		for {
			<-chan2
			if cnt>100{
				chan3<-struct{}{}
				return
			}
			fmt.Println(cnt)
			cnt++
			chan3<-struct{}{}
		}
	}()

	go func(){
		defer wg.Done()
		for {
			<-chan3
			if cnt>100{
				chan1<-struct{}{}
				return
			}
			fmt.Println(cnt)
			cnt++
			chan1<-struct{}{}
		}
	}()	
	chan1<- struct{}{}

	wg.Wait()
	fmt.Println("finished")
}