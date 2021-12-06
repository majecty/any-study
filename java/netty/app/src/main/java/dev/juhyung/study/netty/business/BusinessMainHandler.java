package dev.juhyung.study.netty.business;

import static io.netty.handler.codec.http.HttpHeaderValues.CLOSE;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpHeaderNames;
import io.netty.handler.codec.http.HttpHeaderValues;
import io.netty.handler.codec.http.HttpObject;
import io.netty.handler.codec.http.HttpRequest;
import io.netty.handler.codec.http.HttpResponseStatus;

public class BusinessMainHandler extends SimpleChannelInboundHandler<HttpObject> {
  private static final byte[] CONTENT = {'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'};

  @Override
  public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
    ctx.flush();
  }

  @Override
  protected void channelRead0(ChannelHandlerContext ctx, HttpObject msg) throws Exception {
    if (msg instanceof HttpRequest) {
      HttpRequest req = (HttpRequest) msg;
      System.out.println(req.method().toString());
      System.out.println(req.uri());
      System.out.println(Thread.currentThread().getName());
      Thread.sleep(5000);

      FullHttpResponse response =
          new DefaultFullHttpResponse(
              req.protocolVersion(), HttpResponseStatus.OK, Unpooled.wrappedBuffer(CONTENT));
      response
          .headers()
          .set(HttpHeaderNames.CONTENT_TYPE, HttpHeaderValues.TEXT_PLAIN)
          .setInt(HttpHeaderNames.CONTENT_LENGTH, response.content().readableBytes());

      response.headers().set(HttpHeaderNames.CONNECTION, CLOSE);

      ChannelFuture f = ctx.write(response);
      f.addListener(ChannelFutureListener.CLOSE);
    }
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
    cause.printStackTrace();
    ctx.close();
  }
}
