package main

/*
#cgo LDFLAGS: -L. -lvoicevox_core
#include "voicevox_core.h"
*/
import "C"

// VoicevoxCore is a function group that wraps the C API
type VoicevoxCore struct{}

// const char *voicevox_get_onnxruntime_lib_versioned_filename(void);
func (r *VoicevoxCore) voicevoxGetOnnxruntimeLibVersionedFilename() *C.char {
	return C.voicevox_get_onnxruntime_lib_versioned_filename()
}

// const char *voicevox_get_onnxruntime_lib_unversioned_filename(void);
func (r *VoicevoxCore) voicevoxGetOnnxruntimeLibUnversionedFilename() *C.char {
	return C.voicevox_get_onnxruntime_lib_unversioned_filename()
}

// const struct VoicevoxOnnxruntime *voicevox_onnxruntime_get(void);
func (r *VoicevoxCore) voicevoxOnnxruntimeGet() *C.VoicevoxOnnxruntime {
	return C.voicevox_onnxruntime_get()
}

// VoicevoxResultCode voicevox_onnxruntime_load_once(struct VoicevoxLoadOnnxruntimeOptions options,
//
//	const struct VoicevoxOnnxruntime **out_onnxruntime);
func (r *VoicevoxCore) voicevoxOnnxruntimeLoadOnce(options C.VoicevoxLoadOnnxruntimeOptions, out_onnxruntime **C.VoicevoxOnnxruntime) C.VoicevoxResultCode {
	return C.voicevox_onnxruntime_load_once(options, out_onnxruntime)
}

// なんか動かん
// VoicevoxResultCode voicevox_onnxruntime_init_once(const struct VoicevoxOnnxruntime **out_onnxruntime);
// func (r *VoicevoxCore) voicevoxOnnxruntimeInitOnce(out_onnxruntime **C.VoicevoxOnnxruntime) C.VoicevoxResultCode {
// 	return C.voicevox_onnxruntime_init_once(out_onnxruntime)
// }

// VoicevoxResultCode voicevox_open_jtalk_rc_new(const char *open_jtalk_dic_dir,
//
//	struct OpenJtalkRc **out_open_jtalk);
func (r *VoicevoxCore) voicevoxOpenJtalkRcNew(open_jtalk_dic_dir *C.char, out_open_jtalk **C.OpenJtalkRc) C.VoicevoxResultCode {
	return C.voicevox_open_jtalk_rc_new(open_jtalk_dic_dir, out_open_jtalk)
}

// VoicevoxResultCode voicevox_open_jtalk_rc_use_user_dict(const struct OpenJtalkRc *open_jtalk,
//
//	const struct VoicevoxUserDict *user_dict);
func (r *VoicevoxCore) voicevoxOpenJtalkRcUseUserDict(open_jtalk *C.OpenJtalkRc, user_dict *C.VoicevoxUserDict) C.VoicevoxResultCode {
	return C.voicevox_open_jtalk_rc_use_user_dict(open_jtalk, user_dict)
}

// VoicevoxResultCode voicevox_open_jtalk_rc_analyze(const struct OpenJtalkRc *open_jtalk,
//
//	const char *text,
//	char **output_accent_phrases_json);
func (r *VoicevoxCore) voicevoxOpenJtalkRcAnalyze(open_jtalk *C.OpenJtalkRc, text *C.char, output_accent_phrases_json **C.char) C.VoicevoxResultCode {
	return C.voicevox_open_jtalk_rc_analyze(open_jtalk, text, output_accent_phrases_json)
}

// void voicevox_open_jtalk_rc_delete(struct OpenJtalkRc *open_jtalk);
// C.void という型がないのでこうした。
func (r *VoicevoxCore) voicevoxOpenJtalkRcDelete(open_jtalk *C.OpenJtalkRc) {
	C.voicevox_open_jtalk_rc_delete(open_jtalk)
}

// struct VoicevoxInitializeOptions voicevox_make_default_initialize_options(void);
// C.void という型がないのでこうした。
func (r *VoicevoxCore) voicevoxMakeDefaultInitializeOptions() C.VoicevoxInitializeOptions {
	return C.voicevox_make_default_initialize_options()
}

// const char *voicevox_get_version(void);
func (r *VoicevoxCore) voicevoxGetVersion() *C.char {
	return C.voicevox_get_version()
}

// VoicevoxResultCode voicevox_audio_query_create_from_accent_phrases(const char *accent_phrases_json,
//
//	char **output_audio_query_json);
func (r *VoicevoxCore) voicevox_audio_query_create_from_accent_phrases(accent_phrases_json *C.char, output_audio_query_json **C.char) C.VoicevoxResultCode {
	return C.voicevox_audio_query_create_from_accent_phrases(accent_phrases_json, output_audio_query_json)
}

// VoicevoxResultCode voicevox_voice_model_file_open(const char *path,
//
//	struct VoicevoxVoiceModelFile **out_model);
func (r *VoicevoxCore) voicevoxVoiceModelFileOpen(path *C.char, out_model **C.VoicevoxVoiceModelFile) C.VoicevoxResultCode {
	return C.voicevox_voice_model_file_open(path, out_model)
}

// --- ここから ---

// struct VoicevoxLoadOnnxruntimeOptions voicevox_make_default_load_onnxruntime_options(void);
func (r *VoicevoxCore) voicevoxMakeDefaultLoadOnnxruntimeOptions() C.VoicevoxLoadOnnxruntimeOptions {
	return C.voicevox_make_default_load_onnxruntime_options()
}

// VoicevoxResultCode voicevox_synthesizer_new(const struct VoicevoxOnnxruntime *onnxruntime,
//
//	const struct OpenJtalkRc *open_jtalk,
//	struct VoicevoxInitializeOptions options,
//	struct VoicevoxSynthesizer **out_synthesizer);
func (r *VoicevoxCore) voicevoxSynthesizerNew(onnxruntime *C.VoicevoxOnnxruntime, open_jtalk *C.OpenJtalkRc, options C.VoicevoxInitializeOptions, out_synthesizer **C.VoicevoxSynthesizer) C.VoicevoxResultCode {
	return C.voicevox_synthesizer_new(onnxruntime, open_jtalk, options, out_synthesizer)
}

// VoicevoxResultCode voicevox_synthesizer_load_voice_model(const struct VoicevoxSynthesizer *synthesizer,
//
//	const struct VoicevoxVoiceModelFile *model);
func (r *VoicevoxCore) voicevoxSynthesizerLoadVoiceModel(synthesizer *C.VoicevoxSynthesizer, model *C.VoicevoxVoiceModelFile) C.VoicevoxResultCode {
	return C.voicevox_synthesizer_load_voice_model(synthesizer, model)
}

// struct VoicevoxTtsOptions voicevox_make_default_tts_options(void);
func (r *VoicevoxCore) voicevoxMakeDefaultTtsOptions() C.VoicevoxTtsOptions {
	return C.voicevox_make_default_tts_options()
}

// VoicevoxResultCode voicevox_synthesizer_tts(const struct VoicevoxSynthesizer *synthesizer,
//
//	const char *text,
//	VoicevoxStyleId style_id,
//	struct VoicevoxTtsOptions options,
//	uintptr_t *output_wav_length,
//	uint8_t **output_wav);
func (r *VoicevoxCore) voicevoxSynthesizerTts(synthesizer *C.VoicevoxSynthesizer, text *C.char, style_id C.VoicevoxStyleId, options C.VoicevoxTtsOptions, output_wav_length *C.ulong, output_wav **C.uchar) C.VoicevoxResultCode {
	return C.voicevox_synthesizer_tts(synthesizer, text, style_id, options, output_wav_length, output_wav)
}

// --- ここまで ---

// func (r *VoicevoxCore) voicevoxAudioQuery(
// 	text *C.char,
// 	speakerID C.uint,
// 	options C.VoicevoxAudioQueryOptions,
// 	outputAudioQueryJson **C.char,
// ) C.VoicevoxResultCode {
// 	return C.voicevox_audio_query(text, speakerID, options, outputAudioQueryJson)
// }

// func (r *VoicevoxCore) voicevoxMakeDefaultInitializeOptions() C.VoicevoxInitializeOptions {
// 	return C.voicevox_make_default_initialize_options()
// }

// // func (r *VoicevoxCore) voicevoxInitialize(options C.VoicevoxInitializeOptions) C.int {
// // 	return C.voicevox_initialize(options)
// // }

// func (r *VoicevoxCore) voicevoxGetVersion() *C.char {
// 	return C.voicevox_get_version()
// }

// // func (r *VoicevoxCore) voicevoxLoadModel(speakerID C.uint) C.int {
// // 	return C.voicevox_load_model(speakerID)
// // }

// // func (r *VoicevoxCore) voicevoxIsGpuMode() C.bool {
// // 	return C.voicevox_is_gpu_mode()
// // }

// // func (r *VoicevoxCore) voicevoxIsModelLoaded(speakerID C.uint) C.bool {
// // 	return C.voicevox_is_model_loaded(speakerID)
// // }

// // func (r *VoicevoxCore) voicevoxFinalize() C.void {
// // 	return C.voicevox_finalize()
// // }

// // func (r *VoicevoxCore) voicevoxGetMetasJson() *C.char {
// // 	return C.voicevox_get_metas_json()
// // }

// // func (r *VoicevoxCore) voicevoxGetSupportedDevicesJson() *C.char {
// // 	return C.voicevox_get_supported_devices_json()
// // }

// // func (r *VoicevoxCore) voicevoxPredictDuration(
// // 	length C.ulong,
// // 	phonemeVector *C.int64_t,
// // 	speakerID C.uint,
// // 	outputPredictDurationDataLength *C.ulong,
// // 	outputPredictDurationData **C.float,
// // ) C.int {
// // 	return C.voicevox_predict_duration(
// // 		length,
// // 		phonemeVector,
// // 		speakerID,
// // 		outputPredictDurationDataLength,
// // 		outputPredictDurationData,
// // 	)
// // }

// // func (r *VoicevoxCore) voicevoxPredictDurationDataFree(predictDurationData *C.float) C.void {
// // 	return C.voicevox_predict_duration_data_free(predictDurationData)
// // }

// // func (r *VoicevoxCore) voicevoxPredictIntonation(
// // 	length C.ulong,
// // 	vowel_phoneme_vector *C.int64_t,
// // 	consonantPhonemeVector *C.int64_t,
// // 	startAccentVector *C.int64_t,
// // 	endAccentVector *C.int64_t,
// // 	startAccentPhraseVector *C.int64_t,
// // 	endAccentPhraseVector *C.int64_t,
// // 	speakerID C.uint,
// // 	outputPredictIntonationDataLength *C.ulong,
// // 	outputPredictIntonationData **C.float,
// // ) C.int {
// // 	return C.voicevox_predict_intonation(
// // 		length,
// // 		vowel_phoneme_vector,
// // 		consonantPhonemeVector,
// // 		startAccentVector,
// // 		endAccentVector,
// // 		startAccentPhraseVector,
// // 		endAccentPhraseVector,
// // 		speakerID,
// // 		outputPredictIntonationDataLength,
// // 		outputPredictIntonationData,
// // 	)
// // }

// // func (r *VoicevoxCore) voicevoxPredictIntonationDataFree(predictIntonationData *C.float) C.void {
// // 	return C.voicevox_predict_intonation_data_free(predictIntonationData)
// // }

// // func (r *VoicevoxCore) voicevoxDecode(
// // 	length C.ulong,
// // 	phonemeSize C.ulong,
// // 	f0 *C.float,
// // 	phonemeVector *C.float,
// // 	speakerID C.uint,
// // 	outputDecodeDataLength *C.ulong,
// // 	outputDecodeData **C.float,
// // ) C.int {
// // 	return C.voicevox_decode(
// // 		length,
// // 		phonemeSize,
// // 		f0,
// // 		phonemeVector,
// // 		speakerID,
// // 		outputDecodeDataLength,
// // 		outputDecodeData)
// // }

// // func (r *VoicevoxCore) voicevoxDecodeDataFree(decodeData *C.float) C.void {
// // 	return C.voicevox_decode_data_free(decodeData)
// // }

// // func (r *VoicevoxCore) voicevoxMakeDefaultAudioQueryOptions() C.VoicevoxAudioQueryOptions {
// // 	return C.voicevox_make_default_audio_query_options()
// // }

// // func (r *VoicevoxCore) voicevoxAudioQuery(
// // 	text *C.char,
// // 	speakerID C.uint,
// // 	options C.VoicevoxAudioQueryOptions,
// // 	outputAudioQueryJson **C.char,
// // ) C.VoicevoxResultCode {
// // 	return C.voicevox_audio_query(text, speakerID, options, outputAudioQueryJson)
// // }

// // func (r *VoicevoxCore) voicevoxMakeDefaultSynthesisOptions() C.VoicevoxSynthesisOptions {
// // 	return C.voicevox_make_default_synthesis_options()
// // }

// // func (r *VoicevoxCore) voicevoxSynthesis(
// // 	audioQueryJson *C.char,
// // 	speakerID C.uint,
// // 	options C.VoicevoxSynthesisOptions,
// // 	outputWavLength *C.ulong,
// // 	outputWav **C.uchar,
// // ) C.int {
// // 	return C.voicevox_synthesis(audioQueryJson, speakerID, options, outputWavLength, outputWav)
// // }

// // func (r *VoicevoxCore) voicevoxMakeDefaultTtsOptions() C.VoicevoxTtsOptions {
// // 	return C.voicevox_make_default_tts_options()
// // }

// // func (r *VoicevoxCore) voicevoxTts(
// // 	text *C.char,
// // 	speakerID C.uint,
// // 	options C.VoicevoxTtsOptions,
// // 	outputWavLength *C.ulong,
// // 	outputWav **C.uchar,
// // ) C.int {
// // 	return C.voicevox_tts(
// // 		text,
// // 		speakerID,
// // 		options,
// // 		outputWavLength,
// // 		outputWav,
// // 	)
// // }

// // func (r *VoicevoxCore) voicevoxAudioQueryJsonFree(audioQueryJson *C.char) C.void {
// // 	return C.voicevox_audio_query_json_free(audioQueryJson)
// // }

// func (r *VoicevoxCore) voicevoxWavFree(wav *C.uchar) C.void {
// 	return C.voicevox_wav_free(wav)
// }

// // func (r *VoicevoxCore) voicevoxErrorResultToMessage(resultCode C.VoicevoxResultCode) *C.char {
// // 	return C.voicevox_error_result_to_message(resultCode)
// // }

// //////////////////////

// // struct VoicevoxLoadOnnxruntimeOptions voicevox_make_default_load_onnxruntime_options(void);

// func (r *VoicevoxCore) voicevoxWavFree(wav *C.uchar) C.VoicevoxLoadOnnxruntimeOptions {
// 	return C.voicevox_make_default_load_onnxruntime_options(C.void)
// }
