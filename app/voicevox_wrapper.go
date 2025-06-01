package main

/*
#cgo LDFLAGS: -L. -lvoicevox_core
#include "voicevox_core.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type VoicevoxWrapper struct{}

func (v *VoicevoxWrapper) generate(text string, openJTalkDicDir string, onnxruntimePath string, voiceModelPath string, styleId int) ([]byte, error) {

	core := VoicevoxCore{}

	// Generate VoicevoxInitializeOptions
	var initializeOptions C.VoicevoxInitializeOptions = core.voicevoxMakeDefaultInitializeOptions()
	fmt.Println("Generate VoicevoxInitializeOptions")
	fmt.Printf("Acceleration Mode: %v\n", initializeOptions.acceleration_mode)
	fmt.Printf("Cpu Num Threads: %v\n", initializeOptions.cpu_num_threads)

	// Generate VoicevoxLoadOnnxruntimeOptions
	var loadOnnxruntimeOptions C.VoicevoxLoadOnnxruntimeOptions = core.voicevoxMakeDefaultLoadOnnxruntimeOptions()
	fmt.Println("Generate VoicevoxLoadOnnxruntimeOptions")
	fmt.Printf("Onnxruntime Filename: %v\n", C.GoString(loadOnnxruntimeOptions.filename))

	// Update OnnxruntimeLibVersionedFilename
	loadOnnxruntimeOptions.filename = C.CString(onnxruntimePath)

	// Load Onnxruntime
	var onnxruntime *C.VoicevoxOnnxruntime
	var onnxruntimeLoadOnceResultCode C.VoicevoxResultCode = core.voicevoxOnnxruntimeLoadOnce(loadOnnxruntimeOptions, &onnxruntime)
	if onnxruntimeLoadOnceResultCode == 0 {
		fmt.Println("Load Onnxruntime Success")
	} else {
		fmt.Println("Load Onnxruntime Failed")
		return nil, errors.New("Onnxruntimeの読み込みに失敗しました。")
	}

	// Load Open JDK
	var openJtalk *C.OpenJtalkRc
	var openJtalkRcNewResultCode C.VoicevoxResultCode = core.voicevoxOpenJtalkRcNew(C.CString(openJTalkDicDir), &openJtalk)
	if openJtalkRcNewResultCode == 0 {
		fmt.Println("Load OpenJTalk Success")
	} else {
		fmt.Println("Load OpenJTalk Failed")
		return nil, errors.New("OpenJDKの読み込みに失敗しました。")
	}

	// Make Synthesizer
	var synthesizer *C.VoicevoxSynthesizer
	var synthesizerNewResultCode C.VoicevoxResultCode = core.voicevoxSynthesizerNew(onnxruntime, openJtalk, initializeOptions, &synthesizer)
	if synthesizerNewResultCode == 0 {
		fmt.Println("Make Synthesizer Success")
	} else {
		fmt.Println("Make Synthesizer Failed")
		return nil, errors.New("Synthesizerの作成に失敗しました。")
	}

	// Open Voice Model
	var model *C.VoicevoxVoiceModelFile
	var voiceModelFileOpenResultCode C.VoicevoxResultCode = core.voicevoxVoiceModelFileOpen(C.CString(voiceModelPath), &model)
	if voiceModelFileOpenResultCode == 0 {
		fmt.Println("Open Voice Model Success")
	} else {
		fmt.Println("Open Voice Model Failed")
		return nil, errors.New("Voice Modelの読み込みに失敗しました。")
	}

	// Synthesizer Load Voice Model
	var synthesizerLoadVoiceModelResultCode C.VoicevoxResultCode = core.voicevoxSynthesizerLoadVoiceModel(synthesizer, model)
	if synthesizerLoadVoiceModelResultCode == 0 {
		fmt.Println("Synthesizer Load Voice Model Success")
	} else {
		fmt.Println("Synthesizer Load Voice Model Failed")
		return nil, errors.New("SynthesizerのVoice Modelの読み込みに失敗しました。")
	}

	// Generate
	// StyleID対応表: https://github.com/VOICEVOX/voicevox_vvm/blob/main/README.md#%E9%9F%B3%E5%A3%B0%E3%83%A2%E3%83%87%E3%83%ABvvm%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%A8%E5%A3%B0%E3%82%AD%E3%83%A3%E3%83%A9%E3%82%AF%E3%82%BF%E3%83%BC%E3%82%B9%E3%82%BF%E3%82%A4%E3%83%AB%E5%90%8D%E3%81%A8%E3%82%B9%E3%82%BF%E3%82%A4%E3%83%AB-id-%E3%81%AE%E5%AF%BE%E5%BF%9C%E8%A1%A8
	var voicevoxTtsOptions C.VoicevoxTtsOptions = core.voicevoxMakeDefaultTtsOptions()
	var outputWavLength C.ulong = 0
	var outputWav *C.uchar
	var synthesizerTtsResultCode C.VoicevoxResultCode = core.voicevoxSynthesizerTts(synthesizer, C.CString(text), C.VoicevoxStyleId(styleId), voicevoxTtsOptions, &outputWavLength, &outputWav)
	if synthesizerTtsResultCode == 0 {
		fmt.Println("Synthesize Success")
	} else {
		fmt.Println("Synthesize Failed")
		return nil, errors.New("音声の合成に失敗しました。")
	}

	// Output
	bytes := C.GoBytes(unsafe.Pointer(outputWav), C.int(outputWavLength))
	return bytes, nil
}
