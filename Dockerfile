FROM golang:1.24.3-bookworm AS builder

WORKDIR /app
COPY app .

COPY lib/voicevox_core.h .
COPY lib/libvoicevox_core.so /usr/lib
COPY lib/libvoicevox_onnxruntime.so.1.17.3 /usr/lib

RUN go mod tidy
RUN go build -o main .



FROM public.ecr.aws/lambda/provided:al2023

COPY lib/voicevox_core.h .
COPY lib/libvoicevox_core.so .
COPY lib/libvoicevox_onnxruntime.so.1.17.3 .
COPY lib/open_jtalk_dic_utf_8-1.11 open_jtalk_dic_utf_8-1.11
COPY lib/vvms/0.vvm vvms/

COPY --from=builder /app/main ./main
ENTRYPOINT [ "./main" ]




#FROM golang:1.24.3-bookworm AS builder
#
#WORKDIR /app
#COPY app .
#COPY go.sum .
#COPY main.go .
#COPY voicevox_core.go .
#COPY voicevox_wrapper.go .
#
#COPY voicevox_core.h .
#COPY libvoicevox_core.so .
#COPY libvoicevox_core.so /usr/lib
#
#COPY libvoicevox_onnxruntime.so.1.17.3 .
#COPY libvoicevox_onnxruntime.so.1.17.3 /usr/lib
#
#COPY open_jtalk_dic_utf_8-1.11/ open_jtalk_dic_utf_8-1.11/
#
#COPY vvms/0.vvm vvms/0.vvm
#
#RUN go mod tidy
#
#RUN go build -o main .
#
#
#
#FROM public.ecr.aws/lambda/provided:al2023
#
#COPY voicevox_core.h .
#COPY libvoicevox_core.so .
#COPY libvoicevox_core.so /usr/lib
#
#COPY libvoicevox_onnxruntime.so.1.17.3 .
#COPY libvoicevox_onnxruntime.so.1.17.3 /usr/lib
#COPY open_jtalk_dic_utf_8-1.11/ open_jtalk_dic_utf_8-1.11/
#COPY vvms/0.vvm vvms/0.vvm
#
#COPY --from=builder /app/main ./main
#ENTRYPOINT [ "./main" ]
#
