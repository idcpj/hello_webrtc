'use strict';


class websocket {

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
        console.log(this.app);

        this.ws = new WebSocket(url + "?uid=" + uid);
    }


    /**
     * @param type  {String} type=[open,message,error,close]
     *
     * @param func {Function,String}
     */
    bindEvent(type,func){
        this.ws.addEventListener(type,func)
    }

    send(data ){
        this.ws.send(data)
    }

    close( ){
        this.ws.close()
    }


}



