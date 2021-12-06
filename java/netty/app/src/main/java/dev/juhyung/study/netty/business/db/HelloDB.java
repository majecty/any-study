package dev.juhyung.study.netty.business.db;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class HelloDB {
  public static void main(String[] args) {
    Connection connection = null;
    try {
      connection = DriverManager.getConnection("jdbc:sqlite:borre.db");
      final var statement = connection.createStatement();
      statement.setQueryTimeout(30);

      statement.executeUpdate("drop table if exists counts");
      statement.executeUpdate("create table counts (id integer, count integer)");
      statement.executeUpdate("insert into counts values(0, 0)");
      ResultSet rs = statement.executeQuery("select * from counts");
      while (rs.next()) {
        System.out.println("id " + rs.getInt("id"));
        System.out.println("count " + rs.getInt("count"));
      }
    } catch (SQLException e) {
      e.printStackTrace();
    } finally {
      try {
        if (connection != null) {
          connection.close();
        }
      } catch (SQLException e) {
        System.err.println(e.getMessage());
      }
    }
  }
}
