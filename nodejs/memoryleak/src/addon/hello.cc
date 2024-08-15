// hello.cc
#include <node.h>

namespace demo {

using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

void Method(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  args.GetReturnValue().Set(String::NewFromUtf8(
      isolate, "world").ToLocalChecked());
}

void MethodLeakJS(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  args.GetReturnValue().Set(String::NewFromUtf8(
      isolate, "world").ToLocalChecked());
  // Memory leak
  for (int i = 0; i < 1000000; i++) {
    String::NewFromUtf8(isolate, "world").ToLocalChecked();
  }
}

void MethodLeakMalloc(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  args.GetReturnValue().Set(String::NewFromUtf8(
      isolate, "world").ToLocalChecked());
  // Memory leak
  for (int i = 0; i < 1000000; i++) {
    malloc(1024);
  }
}

void Initialize(Local<Object> exports) {
  NODE_SET_METHOD(exports, "hello", Method);
  NODE_SET_METHOD(exports, "helloLeakJS", MethodLeakJS);
  NODE_SET_METHOD(exports, "helloLeakMalloc", MethodLeakMalloc);
}

NODE_MODULE(NODE_GYP_MODULE_NAME, Initialize)

}  // namespace demo