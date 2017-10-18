package xdean.share.mybatis.plugin;

import java.util.Properties;

import org.apache.ibatis.executor.Executor;
import org.apache.ibatis.mapping.MappedStatement;
import org.apache.ibatis.plugin.Interceptor;
import org.apache.ibatis.plugin.Intercepts;
import org.apache.ibatis.plugin.Invocation;
import org.apache.ibatis.plugin.Plugin;
import org.apache.ibatis.plugin.Signature;
import org.apache.ibatis.session.ResultHandler;
import org.apache.ibatis.session.RowBounds;

/*<pre>

<plugins>
    <plugin interceptor="xdean.share.mybatis.plugin.MyBatisPluginDemo">
        <property name="someProperty" value="100"/>
    </plugin>
</plugins>

</pre>
 */
@Intercepts({
    @Signature(
        type = Executor.class,
        method = "query",
        args = { MappedStatement.class, Object.class, RowBounds.class, ResultHandler.class }),
    @Signature(
        type = Executor.class,
        method = "update",
        args = { MappedStatement.class, Object.class })
})
public class MyBatisPluginDemo implements Interceptor {
  @Override
  public Object intercept(Invocation invocation) throws Throwable {
    Object[] args = invocation.getArgs();
    MappedStatement statement = (MappedStatement) args[0];
    Object param = args[1];
    System.out.println("sql: " + statement.getBoundSql(param).getSql());
    return invocation.proceed();
  }

  @Override
  public Object plugin(Object target) {
    return Plugin.wrap(target, this);
  }

  @Override
  public void setProperties(Properties properties) {
  }
}