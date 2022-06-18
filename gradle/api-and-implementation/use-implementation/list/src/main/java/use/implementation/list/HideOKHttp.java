package use.implementation.list;

import java.io.IOException;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Request.Builder;
import okhttp3.Response;

public class HideOKHttp {
  String doSomething() throws IOException {
    final OkHttpClient okHttpClient = new OkHttpClient();
    final Request request = new Builder()
        .url("https://cat-fact.herokuapp.com/facts")
        .build();
    try (Response response = okHttpClient.newCall(request).execute()) {
      return response.body().string();
    }
  }

  // 아래 코드를 다른 모듈에서 사용하면 컴파일 에러 난다.
  // 컴파일 에러를 해결하려면 okhttp3를 사용할 때 `api`를 사용해야 한다.
  public Response exposeOkHTTP3Type() throws IOException {
    final OkHttpClient okHttpClient = new OkHttpClient();
    final Request request = new Builder()
        .url("https://cat-fact.herokuapp.com/facts")
        .build();
    return okHttpClient.newCall(request).execute();
  }

  public static void main(String[] args) throws IOException {
    System.out.println(new HideOKHttp().doSomething());
  }
}
