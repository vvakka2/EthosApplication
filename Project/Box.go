package main
import (
"ethos/syscall"
ethosLog "ethos/log"
"ethos/efmt"
"ethos/ethos"
"math"
)
func main() {
me := syscall.GetUser()
path := "/user/"+me+"/myDir/"
status := ethosLog.RedirectToLog("box")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr,"Error opening %v: %v\n",path,status)
syscall.Exit(syscall.StatusOk)
}
data := BoxType {1,1,3,3}
efmt.Println("Values x1:",data.x1,"y1:",data.y1)
efmt.Println("Second Values X2: ",data.x2," Y2: ",data.y2)
distance := math.Sqrt(math.Pow((data.x2-data.x1),2)+math.Pow((data.y2-data.y1),2))
efmt.Println("Distance: ",distance)
fd,status := ethos.OpenDirectoryPath(path)
data.Write(fd)
efmt.Fprint(syscall.Stderr,"Success")
}
