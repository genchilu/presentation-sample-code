import java.util.Random;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class ThreadOverhead {
    public static void doSomething() {
        for (int j = 0; j < 1000; j++) {
            Random random = new Random();
            random.nextInt();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        int threadNum = Integer.parseInt(args[0]);

        ExecutorService executorService = Executors.newFixedThreadPool(threadNum);

        for (int j = 0; j < 200000; j++) {
            executorService.execute(new Thread(() -> {
                doSomething();
            }));
        }
        executorService.shutdown();
        executorService.awaitTermination(Long.MAX_VALUE, TimeUnit.NANOSECONDS);
    }
}