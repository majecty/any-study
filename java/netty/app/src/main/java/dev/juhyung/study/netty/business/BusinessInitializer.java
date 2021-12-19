package dev.juhyung.study.netty.business;

import dev.juhyung.study.netty.business.db.CountDB;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.http.HttpServerCodec;
import io.netty.handler.codec.http.HttpServerExpectContinueHandler;
import io.netty.util.concurrent.EventExecutorGroup;

public class BusinessInitializer extends ChannelInitializer<SocketChannel> {

  private final EventExecutorGroup businessGroup;
  private final CountDB countDB;

  public BusinessInitializer(EventExecutorGroup businessGroup, CountDB countDB) {
   this.businessGroup = businessGroup;
   this.countDB = countDB;
  }

  @Override
  protected void initChannel(SocketChannel ch) throws Exception {
    ChannelPipeline p = ch.pipeline();
    p.addLast(new HttpServerCodec());
    p.addLast(new HttpServerExpectContinueHandler());
    p.addLast(businessGroup, new BusinessMainHandler(countDB));
  }
}
