// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 14 Jul 2022 13:22:26 BST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ndigo

/*
#cgo linux LDFLAGS: -L/usr/local/lib -lndi
#cgo darwin LDFLAGS: -Wl,-rpath,/Library/NDI\ SDK\ for\ Apple/lib/macOS -L/Library/NDI\ SDK\ for\ Apple/lib/macOS -lndi
#cgo windows LDFLAGS: -LC:\\Program Files\\NDI\NDI 5 Runtime\\v5 -lProcessing.NDI.Lib.x64
#include <stdlib.h>
#include "include/Processing.NDI.Lib.h"
#include "cgo_helpers.h"
*/
import "C"

// SourceType as declared in include/Processing.NDI.structs.h:207
type SourceType struct {
	PNdiName       string
	ref667670a2    *C.NDIlib_source_t
	allocs667670a2 interface{}
}

// VideoFrameV2Type as declared in include/Processing.NDI.structs.h:259
type VideoFrameV2Type struct {
	Xres               int32
	Yres               int32
	FourCC             FourCCVideoType
	FrameRateN         int32
	FrameRateD         int32
	PictureAspectRatio float32
	FrameFormatType    FrameFormatType
	Timecode           int64
	PData              *byte
	PMetadata          string
	Timestamp          int64
	refffd8086c        *C.NDIlib_video_frame_v2_t
	allocsffd8086c     interface{}
}

// AudioFrameV2Type as declared in include/Processing.NDI.structs.h:295
type AudioFrameV2Type struct {
	SampleRate           int32
	NoChannels           int32
	NoSamples            int32
	Timecode             int64
	PData                *float32
	ChannelStrideInBytes int32
	PMetadata            string
	Timestamp            int64
	ref7f2b41d7          *C.NDIlib_audio_frame_v2_t
	allocs7f2b41d7       interface{}
}

// AudioFrameV3Type as declared in include/Processing.NDI.structs.h:341
type AudioFrameV3Type struct {
	SampleRate     int32
	NoChannels     int32
	NoSamples      int32
	Timecode       int64
	FourCC         FourCCAudioType
	PData          *byte
	PMetadata      string
	Timestamp      int64
	ref7ee92be0    *C.NDIlib_audio_frame_v3_t
	allocs7ee92be0 interface{}
}

// MetadataFrameType as declared in include/Processing.NDI.structs.h:359
type MetadataFrameType struct {
	Length         int32
	Timecode       int64
	PData          *byte
	ref612bc0fe    *C.NDIlib_metadata_frame_t
	allocs612bc0fe interface{}
}

// TallyType as declared in include/Processing.NDI.structs.h:373
type TallyType struct {
	OnProgram      bool
	OnPreview      bool
	refc51e93b6    *C.NDIlib_tally_t
	allocsc51e93b6 interface{}
}

// FindInstanceType as declared in include/Processing.NDI.Find.h:33
type FindInstanceType C.NDIlib_find_instance_t

// FindCreateType as declared in include/Processing.NDI.Find.h:56
type FindCreateType struct {
	ShowLocalSources bool
	PGroups          string
	PExtraIps        string
	ref298166dd      *C.NDIlib_find_create_t
	allocs298166dd   interface{}
}

// RecvInstanceType as declared in include/Processing.NDI.Recv.h:33
type RecvInstanceType C.NDIlib_recv_instance_t

// RecvCreateV3Type as declared in include/Processing.NDI.Recv.h:133
type RecvCreateV3Type struct {
	SourceToConnectTo SourceType
	ColorFormat       RecvColorFormat
	Bandwidth         RecvBandwidth
	AllowVideoFields  bool
	PNdiRecvName      string
	ref2125df04       *C.NDIlib_recv_create_v3_t
	allocs2125df04    interface{}
}

// RecvPerformanceType as declared in include/Processing.NDI.Recv.h:152
type RecvPerformanceType struct {
	VideoFrames    int64
	AudioFrames    int64
	MetadataFrames int64
	ref53902e91    *C.NDIlib_recv_performance_t
	allocs53902e91 interface{}
}

// RecvQueueType as declared in include/Processing.NDI.Recv.h:169
type RecvQueueType struct {
	VideoFrames    int32
	AudioFrames    int32
	MetadataFrames int32
	ref83c703c1    *C.NDIlib_recv_queue_t
	allocs83c703c1 interface{}
}

// RecvRecordingTimeType as declared in include/Processing.NDI.Recv.ex.h:202
type RecvRecordingTimeType struct {
	NoFrames       int64
	StartTime      int64
	LastTime       int64
	ref9e97433b    *C.NDIlib_recv_recording_time_t
	allocs9e97433b interface{}
}

// SendInstanceType as declared in include/Processing.NDI.Send.h:33
type SendInstanceType C.NDIlib_send_instance_t

// SendCreateType as declared in include/Processing.NDI.Send.h:54
type SendCreateType struct {
	PNdiName       string
	PGroups        string
	ClockVideo     bool
	ClockAudio     bool
	ref85a33e95    *C.NDIlib_send_create_t
	allocs85a33e95 interface{}
}

// RoutingInstanceType as declared in include/Processing.NDI.Routing.h:33
type RoutingInstanceType C.NDIlib_routing_instance_t

// RoutingCreateType as declared in include/Processing.NDI.Routing.h:47
type RoutingCreateType struct {
	PNdiName      string
	PGroups       string
	reffc210a9    *C.NDIlib_routing_create_t
	allocsfc210a9 interface{}
}

// AudioFrameInterleaved16sType as declared in include/Processing.NDI.utilities.h:69
type AudioFrameInterleaved16sType struct {
	SampleRate     int32
	NoChannels     int32
	NoSamples      int32
	Timecode       int64
	ReferenceLevel int32
	PData          *int16
	refc5c2e401    *C.NDIlib_audio_frame_interleaved_16s_t
	allocsc5c2e401 interface{}
}

// AudioFrameInterleaved32sType as declared in include/Processing.NDI.utilities.h:101
type AudioFrameInterleaved32sType struct {
	SampleRate     int32
	NoChannels     int32
	NoSamples      int32
	Timecode       int64
	ReferenceLevel int32
	PData          *int32
	ref30602036    *C.NDIlib_audio_frame_interleaved_32s_t
	allocs30602036 interface{}
}

// AudioFrameInterleaved32fType as declared in include/Processing.NDI.utilities.h:125
type AudioFrameInterleaved32fType struct {
	SampleRate     int32
	NoChannels     int32
	NoSamples      int32
	Timecode       int64
	PData          *float32
	ref2a8d41ad    *C.NDIlib_audio_frame_interleaved_32f_t
	allocs2a8d41ad interface{}
}

// VideoFrameType as declared in include/Processing.NDI.deprecated.h:65
type VideoFrameType struct {
	Xres               int32
	Yres               int32
	FourCC             FourCCVideoType
	FrameRateN         int32
	FrameRateD         int32
	PictureAspectRatio float32
	FrameFormatType    FrameFormatType
	Timecode           int64
	PData              *byte
	LineStrideInBytes  int32
	refccebe7b6        *C.NDIlib_video_frame_t
	allocsccebe7b6     interface{}
}

// AudioFrameType as declared in include/Processing.NDI.deprecated.h:93
type AudioFrameType struct {
	SampleRate           int32
	NoChannels           int32
	NoSamples            int32
	Timecode             int64
	PData                *float32
	ChannelStrideInBytes int32
	ref16b107c7          *C.NDIlib_audio_frame_t
	allocs16b107c7       interface{}
}

// RecvCreateType as declared in include/Processing.NDI.deprecated.h:134
type RecvCreateType struct {
	SourceToConnectTo SourceType
	ColorFormat       RecvColorFormat
	Bandwidth         RecvBandwidth
	AllowVideoFields  bool
	refb0fafbc2       *C.NDIlib_recv_create_t
	allocsb0fafbc2    interface{}
}

// FramesyncInstanceType as declared in include/Processing.NDI.FrameSync.h:78
type FramesyncInstanceType C.NDIlib_framesync_instance_t

// V5 as declared in include/Processing.NDI.DynamicLoad.h:579
type V5 struct {
	refbe10a6cd    *C.NDIlib_v5
	allocsbe10a6cd interface{}
}

// V45 as declared in include/Processing.NDI.DynamicLoad.h:581
type V45 struct {
	ref6a6aecbe    *C.NDIlib_v4_5
	allocs6a6aecbe interface{}
}

// V4 as declared in include/Processing.NDI.DynamicLoad.h:582
type V4 struct {
	refc917965b    *C.NDIlib_v4
	allocsc917965b interface{}
}

// V3 as declared in include/Processing.NDI.DynamicLoad.h:583
type V3 struct {
	ref577303f8    *C.NDIlib_v3
	allocs577303f8 interface{}
}

// V2 as declared in include/Processing.NDI.DynamicLoad.h:584
type V2 struct {
	ref2074336e    *C.NDIlib_v2
	allocs2074336e interface{}
}
