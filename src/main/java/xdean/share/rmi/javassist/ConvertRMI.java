package xdean.share.rmi.javassist;

import static xdean.jex.util.cache.CacheUtil.cache;
import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.rmi.Naming;
import java.rmi.Remote;
import java.rmi.RemoteException;
import java.rmi.registry.LocateRegistry;
import java.rmi.server.UnicastRemoteObject;

import javassist.ClassPool;
import javassist.CtClass;
import javassist.CtMethod;
import javassist.Modifier;
import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodInterceptor;

public class ConvertRMI {

  public static void main(String[] args) throws Exception {
//     startServer();
    startClient();
  }

  public static void startClient() throws Exception {
    String stringURL = "rmi://127.0.0.1/" + MyService.class.getName();
    toRemoteInterface(MyService.class);// define the Remote interface in client classloader
    MyService service = fromRemote(Naming.lookup(stringURL), MyService.class);
    String said = service.sayHello("Dean");
    System.out.println("Client heard: " + said);
    service.throwIt();
  }

  public static void startServer() throws Exception {
    LocateRegistry.createRegistry(1099);
    Remote remote = toRemote(new MyServiceImpl(), MyService.class);
    Naming.rebind(MyService.class.getName(), remote);
    System.out.println(remote);
    System.out.println(remote.getClass());
    System.out.println("Server started!");
  }

  @SuppressWarnings("unchecked")
  public static <T> T fromRemote(Remote remote, Class<T> inter) throws Exception {
    Enhancer e = new Enhancer();
    e.setInterfaces(new Class[] { inter });
    e.setCallback((MethodInterceptor) (obj, method, args, proxy) -> {
      Method remoteMethod = remote.getClass().getMethod(method.getName(), method.getParameterTypes());
      try {
        return remoteMethod.invoke(remote, args);
      } catch (InvocationTargetException ex) {
        Throwable targetException = ex.getTargetException();
        while (targetException instanceof RemoteException) {
          targetException = targetException.getCause();
        }
        throw targetException;
      }
    });
    return (T) e.create();
  }

  public static <T> Remote toRemote(T local, Class<T> inter) throws Exception {
    Enhancer e = new Enhancer();
    e.setSuperclass(UnicastRemoteObject.class);
    e.setInterfaces(new Class[] { toRemoteInterface(inter) });
    e.setCallback((MethodInterceptor) (obj, method, args, proxy) -> {
      Method targetMethod = local.getClass().getMethod(method.getName(), method.getParameterTypes());
      try {
        return targetMethod.invoke(local, args);
      } catch (InvocationTargetException ex) {
        Throwable targetException = ex.getTargetException();
        throw new RemoteException(targetException.getMessage(), targetException);
      }
    });
    return (Remote) e.create();
  }

  @SuppressWarnings("unchecked")
  public static Class<? extends Remote> toRemoteInterface(Class<?> inter) throws Exception {
    return cache("toRemote", inter, () -> uncheck(() -> {
      ClassPool pool = ClassPool.getDefault();
      CtClass cc = pool.getAndRename(inter.getName(), inter.getName() + "$RemoteVersion");
      cc.setModifiers(Modifier.PUBLIC | cc.getModifiers());
      cc.addInterface(pool.get(Remote.class.getName()));
      for (CtMethod cm : cc.getMethods()) {
        cm.setExceptionTypes(new CtClass[] { pool.getCtClass(RemoteException.class.getName()) });
      }
      cc.writeFile();
      return cc.toClass();
    }));
  }
}

class ServiceException extends Exception {
  public ServiceException(String msg) {
    super(msg);
  }
}

interface MyService {
  String sayHello(Object who) throws ServiceException;

  void throwIt() throws ServiceException;
}

class MyServiceImpl implements MyService {

  @Override
  public String sayHello(Object who) throws ServiceException {
    String hello = who.toString();
    System.out.println("Server said: " + hello);
    return "Hello! " + hello;
  }

  @Override
  public void throwIt() throws ServiceException {
    throw new ServiceException("throw in server");
  }

}