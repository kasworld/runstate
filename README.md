# runstate
manage grouped goruntine running state

running state manage when multi gorutine is running

0 : runnging and can run all goroutine

bit 0 : need stop all , somebody want stop

bit 1 ~ 63 : use for each go routine

can use max 63 goroutine