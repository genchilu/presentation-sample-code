public class ThreadNum {
    public static void main(String []args) throws InterruptedException {
        for (int i=0;i<10;i++){
            new Thread(()->{
                try {
                    Thread.sleep(100000000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }).start();
        }

    }
}
