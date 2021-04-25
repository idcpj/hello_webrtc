'use strict';


class Socket {

    /**
     * @type {WebSocket}
     */
    ws=null

    /**
     *
     * @type {App}
     */
    app;

    /**
     * @type {String}
     */
    uid;


    /**
     *
     * @param url {String}
     * @param uid {String}
     * @param app {App}
     */
    constructor(url, uid,app) {
        this.uid = uid
        this.app = app

        this.ws = new WebSocket(url + "?uid=" + uid);

    }


    /**
     * @param type  {String} type=[open,message,error,close]
     * @param func {Function}
     */
    bindEvent(type,func){

        /**
         * @param resp {CloseEvent|Event|MessageEvent}
         */
        let handle = (resp)=>{


            // open,error,close
            if (resp.type!=="message") {
                // open,error,close 没有 data
                func(resp)
                return;
            }

            const data = JSON.parse(resp.data)

            // message 单独处理
            switch (data.type) {
                case PEER_OFFER:
                    // 过滤自己
                    if (data.uid===this.app.uid){
                        console.log("过滤自己")
                        return
                    }
                    console.log(PEER_OFFER);
                    this.app.peer.setRemoteDescription(data.data);
                    this.app.peer.createAnswer();
                    break;

                case PEER_ANSWER:
                    // 过滤自己
                    if (data.uid===this.app.uid){
                        console.log("过滤自己")
                        return
                    }
                    console.log(PEER_ANSWER);
                    this.app.peer.setRemoteDescription(data.data);
                    break;

                case PEER_CANDIDATE:
                    // 过滤自己
                    if (data.uid===this.app.uid){
                        console.log("过滤自己")
                        return
                    }
                    this.app.peer.addIceCandidate(data.data);
                    break;
                case SEND_MSG:
                    //todo
                    break;
                default:
                    func(data.type,data.roomid,data.uid,data.data)
            }
        };

        this.ws.addEventListener(type,handle)
    }

    /**
     *
     * @param data {request}
     */
    send(data ){
        let data1 = JSON.stringify(data);
        // console.log(data1);
        this.ws.send(data1)
    }

    close( ){
        this.ws.close()
    }


}



