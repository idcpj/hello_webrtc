class Peer {
    /** @type  RTCPeerConnection */
    localPeer
    /** @type MediaStream */
    localStream
    /** @type HTMLVideoElement */
    localVideo;

    /** @type MediaStream */
    remoteStream
    /** @type HTMLVideoElement */
    remoteVideo;
    /** @type RTCDataChannel */
    DataChannel;
    /** @type  Object*/
    DataChannelMap={};

    /** @type Socket */
    ws;

    pcConfig = {
        'iceServers': [{
            url: 'stun:stun.l.google.com:19302',
        }]
    }
    offerOptions = {
        offerToReceiveAudio: 1,
        offerToReceiveVideo: 1
    };

    userMediaConfig = {
        // width: { min: 640, ideal: 1920, max: 1920 },
        // height: { min: 400, ideal: 1080 },
        aspectRatio: 1.777777778,
        facingMode: {exact: "user"},
        audio: {
            echoCancellation: false,
            noiseSuppression: false,
            autoGainControl: false,
        },
    }
    displayMediaConfig = {
        // width: { min: 640, ideal: 1920, max: 1920 },
        // height: { min: 400, ideal: 1080 },
        aspectRatio: 1.777777778,
        facingMode: { exact: "user" },
    };

    /**
     *
     * @param localVideo {HTMLVideoElement}
     * @param remoteVideo {HTMLVideoElement}
     * @param ws {Socket}
     */
    constructor(localVideo, remoteVideo,ws) {

        this.localVideo = localVideo
        this.remoteVideo = remoteVideo
        this.ws=ws
        ws.peer=this

        this.localPeer =  new RTCPeerConnection(this.pcConfig)


        this.localPeer.addEventListener("icecandidate", event => {
            if (event.candidate) {
                this.ws.send(PEER_CANDIDATE,this.ws.roomid, event.candidate)
            }
        })

        this.localPeer.addEventListener("track", event=>{
            this.remoteStream=new MediaStream();
            this.remoteStream.addTrack(event.track)
            this.remoteVideo.srcObject =  event.streams[0]
        })

        this.localPeer.addEventListener("connectionstatechange",event=>{
            if (this.localPeer.connectionState === 'connected') {
                console.log("Peers connected finish!");
            }
        })

        this.localPeer.addEventListener("datachannel",event=>{
            const dataChannel = event.channel;
            console.log("=========dataChannel=========",dataChannel);

        })

        this.DataChannel = this.localPeer.createDataChannel(this.ws.uid, {negotiated: true, id: 0})

        this.DataChannel.addEventListener("open",event=>{
            console.log("data channel open",event);
        })
        this.DataChannel.addEventListener("message",event=>{
            console.log("data channel message",event);
            let data =JSON.parse(event.data)
            this.DataChannelMap[data.type](event,data,false)

        })
        this.DataChannel.addEventListener("close",event=>{
            console.log("data channel close",event);

        })
        this.DataChannel.addEventListener("error",event=>{
            console.error("data channel error",event);
        })

    }

    addEventListener(type,func){
        this.DataChannelMap[type]=func
    }



    /**
     *
     * @param mediaType {String}  example=video,desktop
     */
    initMedia(mediaType){

        let media=new MediaStream();

        switch (mediaType) {
            case "video":
                media = navigator.mediaDevices.getUserMedia(this.userMediaConfig);
                break;
            case "desktop":
                media = navigator.mediaDevices.getDisplayMedia(this.displayMediaConfig)
                break;
            case "":
                console.log("空类型,不进行本地媒体获取");
                return
            default:
                alert("未知媒体类型")
                return
        }

        media.then((stream) => {

            console.log("get Media Stream");

            if (this.localStream) {
                stream.getTracks().forEach((track) => {
                    this.localStream.addTrack(track);
                    stream.removeTrack(track);
                })
            } else {
                this.localStream = stream
            }

            this.localVideo!=null?this.localVideo.srcObject = this.localStream:"";

        }).then(()=>{
            this.localStream.getTracks().forEach((track) => {
                console.log(track.getSettings());
                this.localPeer.addTrack(track, this.localStream);
            });
        }).catch(e=>{
            console.log(e);
        })
    }

    close() {
        this.localPeer.close()
        this.remoteVideo.pause()
    }

    setRemoteDescription(data) {
        this.localPeer.setRemoteDescription(new RTCSessionDescription(data)).catch(e => console.log(e))
    }

    createAnswer() {

        this.localPeer.createAnswer().then( (sdp) =>{

            this.localPeer.setLocalDescription(sdp).catch(e => console.log(e))
            this.ws.send(sdp.type,this.ws.roomid,sdp)

        }).catch(e => console.log(e))
    }


    async createOffer() {
        await this.localPeer.createOffer(this.offerOptions).then(sdp=>{
            this.localPeer.setLocalDescription(sdp).catch(e => console.log(e))
            // 通知远程
            this.ws.send(sdp.type,this.ws.roomid,sdp)
        })
    }

    /**
     * @param data  {RTCIceCandidateInit}
     */
    addIceCandidate(data) {
        this.localPeer.addIceCandidate(new RTCIceCandidate(data)).catch(e => console.log(e))
    }





}