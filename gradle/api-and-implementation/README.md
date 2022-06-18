# api와 implementation

[라이브러리의 build.gradle](./use-implementation/list/build.gradle.kts#L10) 에서 okhttp3 라이브러리를
implementation로 가져다 쓰는지, api로 가져다 쓰는지에 따라서 컴파일이 실패하고 성공합니다.

[App.java](./use-implementation/app/src/main/java/use/implementation/app/App.java#L23) 에서
exposeOkHTTP3Type을 쓰는지 안쓰는지에 따라 컴파일 결과가 달라집니다.

* api를 쓰면 항상 컴파일이 성공합니다.
* implementation을 쓰고 main 함수에서 exposeOkHTTP3Type 함수를 사용하면 컴파일 에러가 납니다.
* implementation을 쓰고 main 함수에서 exposeOkHTTP3Type 함수를 사용하지 않읓면 컴파일 에러가 나지 않습니다.
