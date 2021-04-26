'use strict'

// work 与 main的交互
// type= heart
const workData={
    type:"",
    msg:"",
    status:1, // 1=success,2=failed
    data:{},
}


setInterval(function () {
    var req =  workData
    req.type="heart"
    self.postMessage(req)

},10000)

self.addEventListener("message",function (event) {
    console.log(event);
})

self.addEventListener("messageerror",function (event) {
    console.log("error",event);
})