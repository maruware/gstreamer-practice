## RTP

### vp8

#### Receiver

```
gst-launch-1.0 -v udpsrc port=5000 caps = "application/x-rtp, media=(string)video, clock-rate=(int)90000, encoding-name=(string)VP8-DRAFT-IETF-01, payload=(int)96, ssrc=(uint)2990747501, clock-base=(uint)275641083, seqnum-base=(uint)34810" ! rtpvp8depay ! vp8dec ! videoconvert ! autovideosink
```

#### Publisher

```
gst-launch-1.0 -v videotestsrc ! video/x-raw,width=640,height=480 ! videoscale ! videoconvert ! vp8enc ! rtpvp8pay ! udpsink host=127.0.0.1 port=5000
```

### vp9

#### Receiver

```
gst-launch-1.0 -v udpsrc port=5000 caps = "application/x-rtp, media=(string)video, clock-rate=(int)90000, encoding-name=(string)VP9, payload=(int)96, ssrc=(uint)2990747501, clock-base=(uint)275641083, seqnum-base=(uint)34810" ! rtpvp9depay ! vp9dec ! videoconvert ! autovideosink
```

#### Publisher

```
gst-launch-1.0 -v videotestsrc ! video/x-raw,width=640,height=480 ! videoscale ! videoconvert ! vp9enc ! rtpvp9pay ! udpsink host=127.0.0.1 port=5000
```

### h264

#### Receiver

```
gst-launch-1.0 -v udpsrc port=5000 caps = "application/x-rtp, media=(string)video, clock-rate=(int)90000, encoding-name=(string)H264, payload=(int)96" ! rtpjitterbuffer latency=20 drop-on-latency=true ! rtph264depay ! decodebin ! videoconvert ! autovideosink sync=false
```

#### Publisher

```
gst-launch-1.0 -v videotestsrc ! video/x-raw,width=640,height=480,framerate=20/1 ! videoscale ! videoconvert ! x264enc tune=zerolatency bitrate=500 speed-preset=superfast ! rtph264pay ! udpsink host=127.0.0.1 port=5000
```
