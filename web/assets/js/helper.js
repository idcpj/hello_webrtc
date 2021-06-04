'use strict';

// peer
const PEER_CANDIDATE = "candidate";
const PEER_ANSWER    = "answer";
const PEER_OFFER     = "offer";

const PEER_READY = "peer_ready";

// socket
const SOCKET_HEART="heart"
const SOCKET_LOGIN="login"

// room
const ROOM_JOIN = "room_join";
const ROOM_QUIT = "room_quit";

// msg
const SEND_MSG = "send_msg";

// datachannel

const  CHANNEL_DATA="channel_data"
const  CHANNEL_MSG="channel_msg"
const  CHANNEL_MOUSE="channel_mouse"


// mouse

const MOUSE_MOVE="mouse_move"
const MOUSE_CLICK="mouse_click"
const MOUSE_DBCLICK="mouse_dbclick"

/**
 *
 * @type {{uid: string, data: {}, type: string, room: string}}
 */
const request = {
    type: "",
    roomid: "",
    uid: "",
    data: {},
}


const response = {
    status:"", // 1=success,==failed
    msg:"",
    type: "",
    roomid: "",
    uid: "",
    data: {},
}


// work 与 main的交互
// type= heart
const workData={
    type:"",
    msg:"",
    status:1, // 1=success,2=failed
    data:{},
}

