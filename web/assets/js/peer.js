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

        this.localPeer =  new RTCPeerConnection(this.pcConfig)

    }

    /**
     *
     * @param mediaType {String}  example=video,desktop
     */
    initMedia(mediaType){

        let media;

        switch (mediaType) {
            case "video":
                 media = navigator.mediaDevices.getUserMedia(this.userMediaConfig);
                break;
            case "desktop":
                 media = navigator.mediaDevices.getDisplayMedia(this.displayMediaConfig)
                break;
            default:
                alert("未知媒体类型")
                return
        }

        media.then((stream) => {
            this._getMediaStream(stream)
        }).then(()=>{
            this.localStream.getTracks().forEach((track) => {
                console.log(track.getSettings());
                this.localPeer.addTrack(track, this.localStream);
            });

            this.localPeer.addEventListener("icecandidate", event => this._icecandidate(event))
            this.localPeer.addEventListener("track", event=>this._track(event))

        }).then(()=>[
        ])
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



    /**
     * @param stream {MediaStream}
     */
    _getMediaStream(stream) {
        console.log("get Media Stream");
        if (this.localStream) {
            stream.getTracks().forEach((track) => {
                this.localStream.addTrack(track);
                stream.removeTrack(track);
            })
        } else {
            this.localStream = stream
        }

        this.localVideo.srcObject = this.localStream
        // await this.localVideo.play();
    }


    /**
     *
     * @param event {RTCPeerConnectionIceEvent}
     */
    _icecandidate(event) {
        if (event.candidate) {
            this.ws.send(PEER_CANDIDATE,this.ws.roomid, event.candidate)
        }
    }

    /**
     *
     * @param event {RTCTrackEvent}
     */
    _track(event) {
        console.log("ontrack");
        this.remoteStream = event.streams[0]
        this.remoteVideo.srcObject =  event.streams[0]
        // this.remoteVideo.play().catch(e=>console.log(e));
        // this.once=true
    }



}