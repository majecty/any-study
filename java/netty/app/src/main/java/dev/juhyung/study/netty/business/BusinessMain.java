package dev.juhyung.study.netty.business;

import dev.juhyung.study.netty.business.db.CountDB;
import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.Channel;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.logging.LogLevel;
import io.netty.handler.logging.LoggingHandler;
import io.netty.util.concurrent.DefaultEventExecutorGroup;
import io.netty.util.concurrent.EventExecutorGroup;
import java.sql.SQLException;

public class BusinessMain {
  static final int PORT = 8080;

  public static void main(String[] args) throws InterruptedException, SQLException {
    EventLoopGroup bossGroup = new NioEventLoopGroup(1);
    EventLoopGroup workerGroup = new NioEventLoopGroup(1);
    EventExecutorGroup businessGroup = new DefaultEventExecutorGroup(2);

    final var countDB = CountDB.openDB();
    try {
      ServerBootstrap b = new ServerBootstrap();
      b.option(ChannelOption.SO_BACKLOG, 1024);
      b.group(bossGroup, workerGroup)
          .channel(NioServerSocketChannel.class)
          .handler(new LoggingHandler(LogLevel.INFO))
          .childHandler(new BusinessInitializer(businessGroup, countDB));

      Channel ch = b.bind(PORT).sync().channel();

      System.out.println("listen on " + PORT);

      ch.closeFuture().sync();
    } finally {
      bossGroup.shutdownGracefully();
      workerGroup.shutdownGracefully();
    }
  }
}
