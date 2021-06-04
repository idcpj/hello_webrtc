'use strict';


class Socket {

    /** @type {WebSocket} */
    ws = null
    /** @type {String} */
    uid;
    /** @type {String} */
    roomid;

    /** @type {Peer} */
    peer


    /**
     *
     * @param url {String}
     * @param uid {String}
     * @param roomid {String}
     */
    constructor(url, uid,roomid) {
        this.uid = uid
        this.roomid=roomid
        this.ws = new WebSocket(url + "?uid=" + uid);

    }

    /**
     *
     * @param type {string}
     * @param func {Function} func(event,response,isYourSelf)
     */
    addEventListener(type, func) {
        const arr = ['error', 'open', 'close']
        if (arr.includes(type)) {
            this.ws.addEventListener(type, event => func(event))
        } else {

            this.ws.addEventListener("message", event => {
                let data = JSON.parse(event.data);
                if (type===data.type){
                    func(event,data,this.uid===data.uid)
                }
            })
        }
    }



    /**
     *
     * @param type {String}
     * @param roomid {String}
     * @param data {Object}
     */
    send(type,roomid,data=null) {

        this.ws.send(JSON.stringify({
            type: type,
            roomid: roomid,
            uid: this.uid,
            data: data,
        }))

    }

    sendDataChannel(type,data=null,roomid){
        this.peer.DataChannel.send(JSON.stringify({
            type: type,
            roomid: '',
            uid: this.uid,
            data: data,
        }))
    }

    close() {
        this.ws.close()
    }


}



