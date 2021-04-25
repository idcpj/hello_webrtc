'use strict';

// peer
const PEER_CANDIDATE = "candidate";
const PEER_ANSWER    = "answer";
const PEER_OFFER     = "offer";

const PEER_READY = "peer_ready";


// room
const ROOM_JOIN = "room_join";
const ROOM_QUIT = "room_quit";

// msg
const SEND_MSG = "send_msg";

/**
 *
 * @type {{uid: string, data: {}, type: string, room: string}}
 */
let request = {
    type: "",
    roomid: "",
    uid: "",
    data: {},
}

/**
 *
 * @param type {String}
 * @param uid {String}
 * @param roomid {String}
 * @param data {Object}
 * @return request
 */
function sendMessage(type = '', uid = '', roomid = '', data = {}) {
    return {
        type: type,
        roomid: roomid,
        uid: uid,
        data: data,
    }
}

