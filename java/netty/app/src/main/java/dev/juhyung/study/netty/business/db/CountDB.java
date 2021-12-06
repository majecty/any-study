package dev.juhyung.study.netty.business.db;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class CountDB {
  private final Connection connection;

  private CountDB(Connection connection) {
    this.connection = connection;
  }

  public static CountDB openDB() throws SQLException {
    Connection connection = null;
    connection = DriverManager.getConnection("jdbc:sqlite:borre.db");
    try (final var statement = connection.createStatement()) {
      statement.setQueryTimeout(30);

      statement.executeUpdate("drop table if exists counts");
      statement.executeUpdate("create table counts (id integer, count integer)");
      statement.executeUpdate("insert into counts values(0, 0)");
      return new CountDB(connection);
    }
  }

  public void countUp() throws SQLException {
    try (final var statement = connection.createStatement()) {
      statement.setQueryTimeout(30);
      statement.executeUpdate("update counts set count=count+1 where id=0");
    }
  }

  public int get() throws SQLException {
    try (final var statement = connection.createStatement()) {
      statement.setQueryTimeout(30);
      final var resultSet = statement.executeQuery("select count from counts where id=0");
      return resultSet.getInt("count");
    }
  }
}
