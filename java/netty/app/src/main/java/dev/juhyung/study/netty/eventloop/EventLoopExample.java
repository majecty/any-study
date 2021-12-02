package dev.juhyung.study.netty.eventloop;

import io.netty.channel.DefaultEventLoop;
import java.util.concurrent.TimeUnit;

public class EventLoopExample {

  public static void main(String[] args) throws InterruptedException {
    final var eventLoop = new DefaultEventLoop();
    eventLoop.schedule(() -> {
      System.out.println("in eventloop");
      eventLoop.shutdownGracefully();
    }, 100, TimeUnit.MILLISECONDS);

    eventLoop.awaitTermination(1, TimeUnit.SECONDS);
  }
}
