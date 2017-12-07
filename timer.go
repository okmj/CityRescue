timer := time.NewTimer(time.Millisecond * 5000)
go func() {

  println("Timer expired")
}()
//stop := timer.Stop()
<-timer.C
