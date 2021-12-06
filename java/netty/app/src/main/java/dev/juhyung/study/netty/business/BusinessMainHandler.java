package dev.juhyung.study.netty.business;

import static io.netty.handler.codec.http.HttpHeaderValues.CLOSE;
import static io.netty.handler.codec.http.HttpHeaderValues.KEEP_ALIVE;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpRequest;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpHeaderNames;
import io.netty.handler.codec.http.HttpHeaderValues;
import io.netty.handler.codec.http.HttpObject;
import io.netty.handler.codec.http.HttpRequest;
import io.netty.handler.codec.http.HttpResponseStatus;
import io.netty.handler.codec.http.HttpUtil;

public class BusinessMainHandler extends SimpleChannelInboundHandler<HttpObject> {
  private static final byte[] CONTENT = { 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd' };

  @Override
  public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
    ctx.flush();
  }

  @Override
  protected void channelRead0(ChannelHandlerContext ctx, HttpObject msg) throws Exception {
    if (msg instanceof HttpRequest) {
      HttpRequest req = (HttpRequest) msg;

      boolean keepAlive = HttpUtil.isKeepAlive(req);
      FullHttpResponse response = new DefaultFullHttpResponse(
          req.protocolVersion(),
          HttpResponseStatus.OK,
          Unpooled.wrappedBuffer(CONTENT)
      );
      response.headers()
          .set(HttpHeaderNames.CONTENT_TYPE, HttpHeaderValues.TEXT_PLAIN)
          .setInt(HttpHeaderNames.CONTENT_LENGTH, response.content().readableBytes());

      if (keepAlive) {
        if (!req.protocolVersion().isKeepAliveDefault()) {
          response.headers().set(HttpHeaderNames.CONNECTION, KEEP_ALIVE);
        }
      } else {
        response.headers().set(HttpHeaderNames.CONNECTION, CLOSE);
      }

      ChannelFuture f = ctx.write(response);

      if (!keepAlive) {
        f.addListener(ChannelFutureListener.CLOSE);
      }
    }
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
    cause.printStackTrace();
    ctx.close();
  }
}
