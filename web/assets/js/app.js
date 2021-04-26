'use strict';


class App {
    /** @type {Socket} */
    websocket=null;

    /** @type {string} */
    uid

    /** @type {Peer} */
    peer=null

    /** @type {String} */
    roomid

    constructor() {
    }

    /**
     * @param type  {string}  example:open,message,error,close,
     * @param func {Function}
     */
    bindSocketEvent(type, func) {
        this.websocket.bindEvent(type,func)
    }


    /**
     *
     * @param type
     * @param data {Object}
     */
    send(type,data) {
        this.websocket.send(sendMessage(type,this.uid,this.roomid,data))
    }

    heart() {
        this.send(SOCKET_HEART)
    }

    async createOffer(){
        await this.peer.createOffer()
    }


    /**
     *
     * @param url {string}
     * @param uid {string}
     * @param roomid {string}
     * @constructor
     */
    Socket(url, uid, roomid) {
        this.uid = uid
        this.roomid=roomid
        this.websocket = new Socket(url, uid, this);

    }

    Join(){
        this.send(ROOM_JOIN)
    }


    /**
     *
     * @param VideoDOM {HTMLVideoElement}
     * @param remoteVideo {HTMLVideoElement}
     * @constructor
     */
     RunPeer(VideoDOM, remoteVideo){
         this.peer=new Peer(this,VideoDOM,remoteVideo);
    }

    close(){
        this.websocket.close()
        this.peer.close()
    }

}
