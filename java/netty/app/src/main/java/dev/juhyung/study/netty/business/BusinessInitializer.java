package dev.juhyung.study.netty.business;

import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.http.HttpServerCodec;
import io.netty.handler.codec.http.HttpServerExpectContinueHandler;
import io.netty.util.concurrent.EventExecutorGroup;

public class BusinessInitializer extends ChannelInitializer<SocketChannel> {

  private final EventExecutorGroup businessGroup;

  public BusinessInitializer(EventExecutorGroup businessGroup) {
   this.businessGroup = businessGroup;
  }

  @Override
  protected void initChannel(SocketChannel ch) throws Exception {
    ChannelPipeline p = ch.pipeline();
    p.addLast(new HttpServerCodec());
    p.addLast(new HttpServerExpectContinueHandler());
    p.addLast(businessGroup, new BusinessMainHandler());
  }
}
