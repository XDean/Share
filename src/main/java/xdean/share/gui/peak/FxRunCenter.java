package xdean.share.gui.peak;

import com.google.common.base.Stopwatch;
import javafx.application.Platform;

import java.util.Deque;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;

/**
 * Utility to schedule tasks to UI thread more fluently.
 * <p>
 * # Feature 1, Peak Clipping
 * <p>
 * If schedule to UI thread directly, all tasks will be called in one UI cycle. It may cause UI block
 * there. To avoid this defect, you should use this utility.
 * <p>
 * FxRunCenter will limit the tasks to UI thread. If tasks have take more than {@link #FRAME_MILLIS} time, it will exit
 * tasks run and schedule remaining tasks to next UI cycle.
 * <p>
 * # Feature 2. De-duplication
 * <p>
 * There may same task do several times. (Usually like relayout, redraw or validate). You can use id to mark tasks to
 * avoid duplicate do same task.
 * <p>
 * # Sample
 *
 * <pre>
 * <code>
 * builder()
 * .display(your-display) // optional
 * .id(your-id) // optional
 * .wait(true) // optional
 * .run(your-task)
 * </code>
 * </pre>
 *
 * @author XDean
 */
public abstract class FxRunCenter {

    private static final long FRAME_MILLIS = 50;
    private static final SwtContext INSTANCE = new SwtContext();

    public static class Config {
        boolean wait = false;
        Object id;

        /**
         * Set block wait. If true, current thread will block until the task done.
         * <p>
         * DO NOT WAIT ON UI THREAD.
         */
        public Config wait(boolean b) {
            wait = b;
            return this;
        }

        /**
         * The task id.
         */
        public Config id(Object o) {
            id = o;
            return this;
        }

        /**
         * Run the task
         */
        public void run(Runnable r) {
            if (wait) {
                r = new WaitRunnable(r);
            }
            INSTANCE.schedule(id, r);
            if (wait) {
                ((WaitRunnable) r).await();
            }
        }
    }

    private static class WaitRunnable implements Runnable {
        final Runnable task;
        final CountDownLatch c = new CountDownLatch(1);

        public WaitRunnable(Runnable task) {
            this.task = task;
        }

        public void await() {
            try {
                c.await();
            } catch (InterruptedException e) {
                // no op, because UI thread never interrupted
            }
        }

        public void done() {
            c.countDown();
        }

        @Override
        public void run() {
            task.run();
            c.countDown();
        }
    }

    public static Config builder() {
        return new Config();
    }

    public static void runLater(Runnable r) {
        builder().run(r);
    }

    public static void runLaterAndWait(Runnable r) {
        builder().wait(true).run(r);
    }


    private static class SwtContext {
        final Deque<Runnable> tasks = new LinkedBlockingDeque<>();
        final Deque<Runnable> advTasks = new LinkedBlockingDeque<>();
        final Map<Object, Runnable> taskMap = new ConcurrentHashMap<>();
        final AtomicBoolean scheduled = new AtomicBoolean(false);
        boolean running;

        void schedule(Runnable r) {
            schedule(null, r);
        }

        void schedule(Object id, Runnable r) {
            Objects.requireNonNull(r, "Task can't be null");
            if (id != null) {
                Runnable origin = r;
                r = new Runnable() {
                    @Override
                    public void run() {
                        origin.run();
                        taskMap.remove(id, this);
                    }
                };
                Runnable old = taskMap.put(id, r);
                if (old != null) {
                    tasks.remove(old);
                    advTasks.remove(old);
                    if (old instanceof WaitRunnable) {
                        ((WaitRunnable) old).done();
                    }
                }
            }
            if (Platform.isFxApplicationThread() && running) {
                advTasks.addFirst(r);
            } else {
                tasks.addLast(r);
            }
            scheduleThis();
        }

        void scheduleThis() {
            if (scheduled.compareAndSet(false, true)) {
                Platform.runLater(this::run);
            }
        }

        void run() {
            Stopwatch sw = Stopwatch.createStarted();
            advTasks.forEach(tasks::addFirst);
            advTasks.clear();
            running = true;
            while (!tasks.isEmpty()) {
                Runnable r = tasks.pollFirst();
                r.run();
                advTasks.forEach(tasks::addFirst);
                advTasks.clear();
                if (sw.elapsed(TimeUnit.MILLISECONDS) > FRAME_MILLIS) {
                    break;
                }
            }
            running = false;
            if (scheduled.compareAndSet(true, false)) {
                if (!tasks.isEmpty()) {
                    scheduleThis();
                }
            }
        }
    }
}
