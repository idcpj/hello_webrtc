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

    /** @type App */
    app


    pcConfig = {
        'iceServers': [{
            url: 'stun:stun.l.google.com:19302',
        }]
    }
    offerOptions = {
        offerToReceiveAudio: 1,
        offerToReceiveVideo: 1
    };

    constraints = {
        video: {
            width: 640,
            height: 480
        },
        audio: {
            echoCancellation: true,
            noiseSupperssion: true,
            autoGainControl: true
        }
    };

    /**
     *
     * @param localVideo {HTMLVideoElement}
     * @param remoteVideo {HTMLVideoElement}
     * @param app {App}
     */
    constructor(app, localVideo, remoteVideo) {
        this.app = app

        this.localVideo = localVideo
        this.remoteVideo = remoteVideo

        this._initPeer()

    }

    _initPeer(){
        console.log("init localPeer");
        this._createPeer()


        // navigator.mediaDevices.getUserMedia(this.constraints)
        navigator.mediaDevices.getDisplayMedia(this.constraints)
            .then((stream) =>  this._getMediaStream(stream))
            .then(()=>{
                this._addTracks()
                this.localPeer.addEventListener("icecandidate", event => this._icecandidate(event))
                this.localPeer.addEventListener("track", event=>this._track(event))

            }).catch(e => {
                console.log("getUserMedia is error:", e)
            });
    }

    close() {
        this.localPeer.close()
        this.remoteVideo.pause()
    }

    setRemoteDescription(data) {
        console.log("setRemoteDescription");
        this.localPeer.setRemoteDescription(new RTCSessionDescription(data)).catch(e => console.log(e))
    }

    createAnswer() {

        this.localPeer.createAnswer().then( (sdp) =>{

            console.log("get answer");
            this.localPeer.setLocalDescription(sdp).catch(e => console.log(e))
            this.app.send(sdp.type, sdp)

        }).catch(e => console.log(e))
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
            this.app.send(PEER_CANDIDATE, event.candidate)
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


    _addTracks() {
        console.log("add localPeer tracks");
        this.localStream.getTracks().forEach((track) => {
            this.localPeer.addTrack(track, this.localStream);
        });
    }

    _createPeer() {
        console.log("create peer");
        this.localPeer = new RTCPeerConnection()
    }

    async createOffer() {
        console.log("create Offer");
        /**
         *
         * @type {RTCSessionDescriptionInit}
         */
        await this.localPeer.createOffer(this.offerOptions).then(sdp=>{
            this.localPeer.setLocalDescription(sdp).catch(e => console.log(e))
            // 通知远程
            this.app.send(sdp.type, sdp)
        })


    }


}